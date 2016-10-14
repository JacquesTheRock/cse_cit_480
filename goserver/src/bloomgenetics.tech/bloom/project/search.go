package project

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
	"errors"
	"strconv"
)

func GetAllProjects() ([]entity.Project, error) {
	const qBase = "SELECT id,name,description,visibility FROM project"
	output := []entity.Project{}
	rows, err := util.Database.Query(qBase)
	var p entity.Project
	defer rows.Close()
	for rows.Next() {
		p = entity.Project{}
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

func SearchProjects(p entity.Project) ([]entity.Project, error) {
	const qBase = "SELECT id,name,description,visibility FROM project"
	queryVars := make([]interface{}, 0)
	out := make([]entity.Project, 0)
	query := " WHERE "
	endQuery := qBase
	if p.ID != 0 {
		queryVars = append(queryVars, p.ID)
		query = query + "id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if p.Name != "" {
		queryVars = append(queryVars, p.Name)
		query = query + "name LIKE $" + strconv.Itoa(len(queryVars)) + " "
	}
	if p.Description != "" {
		queryVars = append(queryVars, p.Description)
		query = query + "description LIKE $" + strconv.Itoa(len(queryVars)) + " "
	}
	if p.Visibility != false {
		queryVars = append(queryVars, p.Visibility)
		query = query + "visibility = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if len(queryVars) > 0 {
		endQuery = qBase + query
	}
	rows, err := util.Database.Query(endQuery, queryVars...)
	if err != nil {
		util.PrintError("Query Error")
		util.PrintError(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		e := entity.Project{}
		var desc sql.NullString
		err = rows.Scan(&e.ID, &e.Name, &desc, &e.Visibility)
		if err != nil {
			util.PrintError("Unable to read project")
			util.PrintError(err)
		}
		if desc.Valid {
			e.Description = desc.String
		}
		out = append(out, e)
	}
	return out, nil
}

func GetProject(p entity.Project) (entity.Project, error) {
	res, err := SearchProjects(p)
	if err != nil {
		return entity.Project{}, err
	}
	if len(res) > 1 {
		return entity.Project{}, errors.New("Too many results")
	}
	if len(res) == 0 {
		return entity.Project{}, errors.New("No results")
	}
	return res[0], nil
}
