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
	const qBase = "UPDATE project SET name = $2, description = $3, visibility = $4, location = $5, species = $6, ptype = $7 WHERE id = $1"
	output := entity.Project{}
	util.PrintDebug(p)
	visible := 0
	if p.Visibility {
		visible = 1
	}
	_, err := util.Database.Exec(qBase, p.ID, p.Name, p.Description, visible, p.Location, p.Species, p.Type)
	if err != nil {
		util.PrintError("Unable to update Project")
		return output, err
	}
	return p, nil
}
