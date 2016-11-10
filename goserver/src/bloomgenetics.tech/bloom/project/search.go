package project

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
	"errors"
	"strconv"
)

type QueryProject struct {
	ID          sql.NullInt64
	Name        sql.NullString
	Location    sql.NullString
	Type        sql.NullString
	Species     sql.NullString
	Description sql.NullString
}

func GetAllProjects() ([]entity.Project, error) {
	const qBase = "SELECT id,name,description,visibility,location, species, ptype FROM project"
	output := []entity.Project{}
	rows, err := util.Database.Query(qBase)
	var p entity.Project
	defer rows.Close()
	for rows.Next() {
		p = entity.Project{}
		var desc sql.NullString
		err = rows.Scan(&p.ID, &p.Name, &desc, &p.Visibility, &p.Location, &p.Species, &p.Type)
		if desc.Valid {
			p.Description = desc.String
		}
		if err != nil {
			util.PrintDebug(err)
			return output, err
		}
		output = append(output, p)
	}
	return output, nil
}

func SearchProjects(q QueryProject) ([]entity.Project, error) {
	const qBase = "SELECT id,name,description,visibility,location,species,ptype FROM project"
	queryVars := make([]interface{}, 0)
	out := make([]entity.Project, 0)
	query := " WHERE "
	endQuery := qBase
	if q.ID.Valid {
		queryVars = append(queryVars, q.ID.Int64)
		query = query + "id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if q.Name.Valid {
		queryVars = append(queryVars, q.Name.String)
		query = query + "name LIKE $" + strconv.Itoa(len(queryVars)) + " "
	}
	if q.Description.Valid {
		queryVars = append(queryVars, q.Description.String)
		query = query + "description LIKE $" + strconv.Itoa(len(queryVars)) + " "
	}
	if len(queryVars) > 0 {
		endQuery = qBase + query
	}
	rows, err := util.Database.Query(endQuery, queryVars...)
	if err != nil {
		util.PrintError("Query Error")
		util.PrintDebug(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		e := entity.Project{}
		var desc sql.NullString
		err = rows.Scan(&e.ID, &e.Name, &desc, &e.Visibility, &e.Location, &e.Species, &e.Type)
		if err != nil {
			util.PrintError("Unable to read project")
			util.PrintDebug(err)
		}
		if desc.Valid {
			e.Description = desc.String
		}
		out = append(out, e)
	}
	return out, nil
}

func GetProject(q QueryProject) (entity.Project, error) {
	res, err := SearchProjects(q)
	if err != nil {
		return entity.Project{}, err
	}
	if len(res) > 1 {
		util.PrintInfo(res)
		err = errors.New("Too many results")
		util.PrintDebug(err)
		return entity.Project{}, err
	}
	if len(res) == 0 {
		err = errors.New("No results")
		util.PrintDebug(err)
		return entity.Project{}, err
	}
	return res[0], nil
}
