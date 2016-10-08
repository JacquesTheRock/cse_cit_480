package project

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func NewProject(uid string, p entity.Project) (entity.Project, error) {
	const qBase = "INSERT INTO project(name,description,visibility) VALUES ($1,$2,$3) RETURNING id"
	output := entity.Project{}
	visible := 0
	if p.Visibility {
		visible = 1
	}
	err := util.Database.QueryRow(qBase, p.Name, p.Description, visible).Scan(&output.ID)
	if err != nil {
		util.PrintError(err)
		util.PrintError("New Project Method error")
		return output, err
	}
	pArray, _ := SearchProjects(output)
	if len(pArray) != 1 {
		return output, nil
	}
	output = pArray[0]
	const rBase = "INSERT INTO roles VALUES($1,$2,1)"
	_, err = util.Database.Exec(rBase, uid, output.ID)
	if err != nil {
		util.PrintError(err)
		util.PrintError("Unable to associate Project")
		return output, err
	}
	return output, nil
}
