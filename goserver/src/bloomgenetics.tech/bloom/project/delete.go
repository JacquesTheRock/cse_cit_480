package project

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func DeleteProject(p entity.Project) (entity.Project, error) {
	const qBase = "DELETE FROM project WHERE id = $1"
	output := entity.Project{}
	_, err := util.Database.Exec(qBase, p.ID)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("Delete Project Method error")
		return output, err
	}
	return output, nil
}
