package treview

import (
	"bloomgenetics.tech/bloom/cross"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
)

func GenerateForest(project_id int64) ([]entity.TreeNode, error) {
	out := make([]entity.TreeNode, 0)
	q := cross.CrossQuery{}
	q.ProjectID.Int64 = project_id
	q.ProjectID.Valid = true
	all, err := cross.SearchCrosses(q)
	if err != nil {
		util.PrintInfo("Issue Finding it")
		util.PrintDebug(err)
		return out, err
	}
	nodes := make(map[int64]*entity.TreeNode)
	for _, el := range all {
		e := entity.TreeNode{}
		e.Self = el
		nodes[el.ID] = &e
	}
	for _, el := range all {
		n := nodes[el.ID]
		parIDs := getParentIDs(n.Self)
		for _, p := range parIDs {
			if nodes[p] != nil {
				n.Parents = append(n.Parents, nodes[p])
			}
		}
		childIDs := getChildrenIDs(n.Self)
		for _, c := range childIDs {
			if nodes[c] != nil {
				n.Children = append(n.Children, nodes[c])
			}
		}
	}
	roots := make(map[int64]*entity.TreeNode)
	used := make(map[int64]bool)
	for _, el := range nodes {
		for len(el.Parents) > 0 {
			if used[el.Self.ID] == true {
				break //Cycled Back on Self
			}
			used[el.Self.ID] = true
			next := el.Parents[0]
			el.Parents = nil //el.Parents[1:]
			el = next
		}
		if !used[el.Self.ID] {
			roots[el.Self.ID] = el
		}
	}
	for _, r := range roots {
		out = append(out, *r)
	}
	return out, nil
}

func getChildrenIDs(q entity.Cross) []int64 {
	const qBase = "SELECT DISTINCT cp.cross_id AS child FROM cross_parent cp, specimen s WHERE cp.specimen_id = s.id AND s.cross_id = $1"
	out := make([]int64, 0)
	rows, err := util.Database.Query(qBase, q.ID)
	if err != nil {
		util.PrintInfo("Unable to query children for cross")
		util.PrintDebug(err)
		return out
	}
	defer rows.Close()
	for rows.Next() {
		var id sql.NullInt64
		err = rows.Scan(&id)
		if err != nil {
			util.PrintDebug(err)
		}
		if id.Valid {
			out = append(out, id.Int64)
		}
	}
	return out
}

func getParentIDs(q entity.Cross) []int64 {
	const qBase = "SELECT DISTINCT s.cross_id AS parent FROM cross_parent cp, specimen s WHERE cp.specimen_id = s.id AND cp.cross_id = $1"
	out := make([]int64, 0)
	rows, err := util.Database.Query(qBase, q.ID)
	if err != nil {
		util.PrintInfo("Unable to query parents for cross")
		util.PrintDebug(err)
		return out
	}
	defer rows.Close()
	for rows.Next() {
		var id sql.NullInt64
		err = rows.Scan(&id)
		if err != nil {
			util.PrintDebug(err)
		}
		if id.Valid {
			out = append(out, id.Int64)
		}
	}
	return out
}

func Generate(project_id int64, cross_id int64) (entity.TreeNode, error) {
	out := entity.TreeNode{}
	q := cross.CrossQuery{}
	q.ID.Int64 = cross_id
	q.ProjectID.Int64 = project_id
	q.ID.Valid = true
	q.ProjectID.Valid = true
	center, err := cross.GetCross(q)
	if err != nil {
		util.PrintInfo("Issue Finding Cross")
		util.PrintDebug(err)
		return out, err
	}
	parIDs := getParentIDs(center)
	for _, pID := range parIDs {
		qP := cross.CrossQuery{}
		qP.ID.Int64 = pID
		qP.ID.Valid = true
		qP.ProjectID.Int64 = project_id
		qP.ProjectID.Valid = true
		p, err := cross.GetCross(qP)
		if err != nil || p.ID == 0 {
			continue
		}
		parent := entity.TreeNode{}
		parent.Self = p
		out.Parents = append(out.Parents, &parent)
	}
	childIDs := getChildrenIDs(center)
	for _, cID := range childIDs {
		qC := cross.CrossQuery{}
		qC.ID.Int64 = cID
		qC.ID.Valid = true
		qC.ProjectID.Int64 = project_id
		qC.ProjectID.Valid = true
		c, err := cross.GetCross(qC)
		if err != nil || c.ID == 0 {
			continue
		}
		child := entity.TreeNode{}
		child.Self = c
		out.Children = append(out.Children, &child)
	}

	out.Self = center
	return out, nil
}
