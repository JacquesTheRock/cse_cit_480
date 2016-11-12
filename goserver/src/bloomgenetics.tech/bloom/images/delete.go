package images

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func DeleteImage(i entity.Image) error {
	const qBase = "DELETE FROM img WHERE id = $1"
	rows, err := util.Database.Query(qBase, i.ID)
	defer rows.Close()
	if err != nil {
		util.PrintDebug(err)
		return err
	}
	return nil
}
