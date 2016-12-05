package candidate

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
	"strconv"
)

type CandidateQuery struct {
	ID        sql.NullInt64
	CrossID   sql.NullInt64
	ProjectID sql.NullInt64
}

func GetTraits(e entity.Candidate) (entity.Candidate, error) {
	const qBase = "SELECT t.id,t.project_id,t.name,t.pool,tt.id,tt.name,tt.weight FROM specimen_trait st JOIN trait t ON st.trait_id = t.id JOIN trait_t tt ON t.class = tt.id WHERE st.specimen_id = $1ORDER BY t.pool"
	rows, err := util.Database.Query(qBase, e.ID)
	if err != nil {
		util.PrintError("Unable to get Candidates traits")
		util.PrintDebug(err)
		return e, err
	}
	defer rows.Close()
	for rows.Next() {
		t := entity.Trait{}
		err = rows.Scan(&t.ID, &t.Project_ID, &t.Name, &t.Pool, &t.Type_ID, &t.Type, &t.Weight)
		if err != nil {
			util.PrintError("Unable to read Candidate Traits")
			util.PrintDebug(err)
		}
		e.Traits = append(e.Traits, t)
	}
	return e, nil
}

func GetAll(pid int64) ([]entity.Candidate, error) {
	const qBase = "SELECT s.id,c.id,c.project_id,s.note,s.img_id FROM crosses c JOIN specimen s ON c.id = s.cross_id WHERE c.project_id = $1"
	out := make([]entity.Candidate, 0)
	rows, err := util.Database.Query(qBase, pid)
	defer rows.Close()
	for rows.Next() {
		e := entity.Candidate{}
		var iid sql.NullInt64
		err = rows.Scan(&e.ID, &e.CrossID, &e.ProjectID, &e.Note, &iid)
		if err != nil {
			util.PrintError("Unable to read project")
			util.PrintDebug(err)
		}
		if iid.Valid {
			e.ImageID = iid.Int64
		}
		e, _ = GetTraits(e)
		out = append(out, e)
	}
	return out, nil

}

func GetParents(c entity.Cross) ([]entity.Candidate, error) {
	q1 := CandidateQuery{}
	q1.ID.Int64 = c.Parent1ID
	q1.ID.Valid = true
	q2 := CandidateQuery{}
	q2.ID.Int64 = c.Parent2ID
	q2.ID.Valid = true
	out1, _ := SearchCandidates(q1)
	out2, _ := SearchCandidates(q2)
	return append(out1, out2...), nil
}

func SearchCandidates(q CandidateQuery) ([]entity.Candidate, error) {
	const qBase = "SELECT s.id,c.id,c.project_id,s.note,s.img_id FROM crosses c JOIN specimen s ON c.id = s.cross_id"
	queryVars := make([]interface{}, 0)
	out := make([]entity.Candidate, 0)
	query := " WHERE "
	endQuery := qBase
	if q.ID.Valid {
		queryVars = append(queryVars, q.ID.Int64)
		query = query + "s.id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if q.CrossID.Valid {
		if len(queryVars) > 0 {
			query += "AND "
		}
		queryVars = append(queryVars, q.CrossID.Int64)
		query = query + "s.cross_id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if q.ProjectID.Valid {
		if len(queryVars) > 0 {
			query += "AND "
		}
		queryVars = append(queryVars, q.ProjectID.Int64)
		query = query + "c.project_id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if len(queryVars) > 0 {
		endQuery = qBase + query
	}
	rows, err := util.Database.Query(endQuery, queryVars...)
	if err != nil {
		util.PrintError("Candidate Query Error")
		util.PrintDebug(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		e := entity.Candidate{}
		var iid sql.NullInt64
		err = rows.Scan(&e.ID, &e.CrossID, &e.ProjectID, &e.Note, &iid)
		if err != nil {
			util.PrintError("Unable to read project")
			util.PrintDebug(err)
		}
		if iid.Valid {
			e.ImageID = iid.Int64
		}
		e, _ = GetTraits(e)
		out = append(out, e)
	}
	return out, nil
}
