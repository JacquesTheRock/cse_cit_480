package project

import (
	"bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func NewProject(uid string, p entity.Project) (entity.Project, error) {
	const qBase = "INSERT INTO project(name,description,visibility,location,species,ptype) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
	output := entity.Project{}
	visible := 0
	if p.Visibility {
		visible = 1
	}
	err := util.Database.QueryRow(qBase, p.Name, p.Description, visible, p.Location, p.Species, p.Type).Scan(&output.ID)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("New Project Method error")
		return output, err
	}
	q := QueryProject{}
	q.ID.Valid = true
	q.ID.Int64 = output.ID
	pArray, _ := SearchProjects(q)
	if len(pArray) != 1 {
		return output, nil
	}
	output = pArray[0]
	err = auth.SetRole(auth.Role{UserID: uid, ProjectID: output.ID, RoleID: 5})
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("Unable to associate Project ")
		util.PrintDebug(uid + " to " + string(output.ID))
		return output, err
	}
	return output, nil
}

//TODO: Expand fields to all other fields
func UpdateProject(p entity.Project) (entity.Project, error) {
	const qBase = "UPDATE project SET name = $1, description = $2, visibility = $3 WHERE id = $4"
	output := entity.Project{}
	visible := 0
	if p.Visibility {
		visible = 1
	}
	_, err := util.Database.Exec(qBase, p.Name, p.Description, visible, p.ID)
	if err != nil {
		util.PrintError("Unable to update Project")
		return output, err
	}
	return p, nil
}
