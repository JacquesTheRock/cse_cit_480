package entity

import (
	"bloomgenetics.tech/bloom/util"
	"database/sql"
)

type Project struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Visibility  bool   `json:"public"`
}

func NewProject(uid string, n string, d string, p bool) (Project, error) {
	const qBase = "INSERT INTO project(name,description,visibility) VALUES ($1,$2,$3) RETURNING id"
	output := Project{}
	visible := 0
	if p {
		visible = 1
	}
	err := util.Database.QueryRow(qBase, n, d, visible).Scan(&output.ID)
	if err != nil {
		util.PrintError(err)
		util.PrintError("New Project Method error")
		return output, err
	}
	output, _ = GetProject(uid, output.ID)
	const rBase = "INSERT INTO roles VALUES($1,$2,1)"
	_, err = util.Database.Exec(rBase, uid, output.ID)
	if err != nil {
		util.PrintError(err)
		util.PrintError("Unable to associate Project")
		return output, err
	}
	return output, nil
}

func GetProject(uid string, id int64) (Project, error) {
	const qBase = "SELECT id,name,description,visibility FROM project WHERE id = $1"
	output := Project{}
	rows, err := util.Database.Query(qBase, id)
	defer rows.Close()
	for rows.Next() {
		var desc sql.NullString
		err = rows.Scan(&output.ID, &output.Name, &desc, &output.Visibility)
		if desc.Valid {
			output.Description = desc.String
		}
		if err != nil {
			util.PrintError(err)
			return output, err
		}
	}
	return output, nil
}

func GetAllProjects() ([]Project, error) {
	const qBase = "SELECT id,name,description,visibility FROM project"
	output := []Project{}
	rows, err := util.Database.Query(qBase)
	var p Project
	defer rows.Close()
	for rows.Next() {
		p = Project{}
		var desc sql.NullString
		err = rows.Scan(&p.ID, &p.Name, &desc, &p.Visibility)
		if desc.Valid {
			p.Description = desc.String
		}
		if err != nil {
			util.PrintError(err)
			return output, err
		}
		output = append(output, p)
	}
	return output, nil
}
