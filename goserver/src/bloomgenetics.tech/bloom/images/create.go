package images

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func CreateImage(i entity.Image) (entity.Image, error) {
	const qBase = "INSERT INTO img(image,ftype,fsize) VALUES ($1,$2,$3) RETURNING id"
	o := entity.Image{}
	util.PrintDebug("Inserting: " + i.Data)
	err := util.Database.QueryRow(qBase, i.Data, i.Type, i.Size).Scan(&i.ID)
	if err != nil {
		util.PrintDebug(err)
		return o, err
	}
	return i, nil
}
