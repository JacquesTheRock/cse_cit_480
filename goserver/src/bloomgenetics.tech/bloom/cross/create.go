package cross

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func AssignParents(e entity.Cross) (entity.Cross, error) {
	const qBase = "INSERT INTO cross_parent(cross_id,specimen_id) VALUES ($1,$2)"
	var err error
	if e.ID == 0 {
		return e, nil
	}
	DeleteAllParents(e)
	if e.Parent1ID == 0 {
		e.Parent1ID = e.Parent2ID
		e.Parent2ID = 0
	}
	if e.Parent1ID != 0 {
		_, err = util.Database.Exec(qBase, e.ID, e.Parent1ID)
		if err != nil {
			util.PrintError("Insert Error")
			util.PrintDebug(err)
			e.Parent1ID = 0
		}
		if e.Parent2ID != 0 {
			_, err = util.Database.Exec(qBase, e.ID, e.Parent2ID)
			if err != nil {
				util.PrintError("Insert Error")
				util.PrintDebug(err)
				e.Parent2ID = 0
			}

		}
	}
	return e, err
}

func CreateCross(e entity.Cross) (entity.Cross, error) {
	const qBase = "INSERT INTO crosses(project_id,name) VALUES ($1,$2) RETURNING id"
	err := util.Database.QueryRow(qBase, e.ProjectID, e.Name).Scan(&e.ID)
	if err != nil {
		util.PrintError("Insert Error")
		util.PrintDebug(err)
		return entity.Cross{}, err
	}
	e, err = AssignParents(e)
	return e, err
}

func UpdateCross(e entity.Cross) (entity.Cross, error) {
	const qBase = "UPDATE crosses SET name = $1 WHERE id = $2 AND project_id = $3"
	_, err := util.Database.Exec(qBase, e.Name, e.ID, e.ProjectID)
	if err != nil {
		util.PrintError("Unable to Update cross")
		util.PrintDebug(err)
		return entity.Cross{}, err
	}
	e, err = AssignParents(e)
	return e, nil
}
