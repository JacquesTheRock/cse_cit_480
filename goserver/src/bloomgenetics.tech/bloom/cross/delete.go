package cross

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func DeleteAllParents(e entity.Cross) error {
	const qBase = "DELETE FROM cross_parent WHERE cross_id = $1"
	_, err := util.Database.Exec(qBase, e.ID)
	if err != nil {
		util.PrintError("Failure to delete parents")
		util.PrintDebug(err)
		return err
	}
	return nil
}

func Delete(e entity.Cross) error {
	DeleteAllParents(e)
	const qBase = "DELETE FROM crosses WHERE id = $1 AND project_id = $2"
	_, err := util.Database.Exec(qBase, e.ID, e.ProjectID)
	if err != nil {
		util.PrintError("Failure to delete Cross")
		util.PrintDebug(err)
		return err
	}
	return nil
}
