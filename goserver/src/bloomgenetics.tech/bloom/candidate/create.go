package candidate

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
)

func LinkTraits(e entity.Candidate, trans *sql.Tx) (entity.Candidate, error) {
	//const qBase = "SELECT s.id,c.id,c.project_id FROM crosses c JOIN specimen s ON c.id = s.cross_id"
	var t2 *sql.Tx
	if trans == nil {
		t2, _ = util.Database.Begin()
	} else {
		t2 = trans
	}
	const qBase = "INSERT INTO specimen_trait(trait_id,specimen_id) VALUES ($1,$2)"
	for _, t := range e.Traits {
		_, err := t2.Exec(qBase, t.ID, e.ID)
		if err != nil {
			util.PrintError("Unable to get Candidates traits")
			util.PrintDebug(err)
			if trans == nil {
				t2.Rollback()
			}
			return e, err
		}
	}
	if trans == nil {
		t2.Commit()
	}
	return e, nil
}

func CreateCandidate(e entity.Candidate) (entity.Candidate, error) {
	const qBase = "INSERT INTO specimen(cross_id) VALUES($1) RETURNING id"
	trans, _ := util.Database.Begin()
	err := trans.QueryRow(qBase, e.CrossID).Scan(&e.ID)
	if err != nil {
		util.PrintError("Candidate Query Error")
		util.PrintDebug(err)
		trans.Rollback()
		return e, err
	}
	e, err = LinkTraits(e, trans)
	if err != nil {
		util.PrintError("Failure Linking Traits Error")
		util.PrintDebug(err)
		trans.Rollback()
		return e, err
	}
	trans.Commit()
	return e, nil
}
