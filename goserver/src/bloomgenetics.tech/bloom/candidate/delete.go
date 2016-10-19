package candidate

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
)

func DeleteTraits(e entity.Candidate, trans *sql.Tx) (entity.Candidate, error) {
	var t2 *sql.Tx
	if trans == nil {
		t2, _ = util.Database.Begin()
	} else {
		t2 = trans
	}
	const qBase = "DELETE FROM specimen_trait WHERE specimen_id = $1"
	_, err := t2.Exec(qBase, e.ID)
	if err != nil {
		util.PrintError("Unable to delete Candidates traits")
		util.PrintDebug(err)
		if trans == nil {
			t2.Rollback()
		}
		return e, err
	}
	if trans == nil {
		t2.Commit()
	}
	return e, nil
}

func DeleteCandidate(e entity.Candidate) (entity.Candidate, error) {
	const qBase = "DELETE FROM specimen WHERE id = $1 AND cross_id = $2"
	_, err := util.Database.Exec(qBase, e.ID, e.CrossID)
	return e, err
}
