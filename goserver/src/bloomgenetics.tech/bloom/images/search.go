package images

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func GetImage(i entity.Image) (entity.Image, error) {
	const qBase = "SELECT id,image,ftype,fsize FROM img WHERE id = $1"
	o := entity.Image{}
	rows, err := util.Database.Query(qBase, i.ID)
	defer rows.Close()
	if err != nil {
		util.PrintDebug(err)
		return o, err
	}
	for rows.Next() {
		err = rows.Scan(&o.ID, &o.Data, &o.Type, &o.Size)
		if err != nil {
			util.PrintError("Unable to get Image")
			util.PrintDebug(err)
		}
	}
	return o, nil
}
