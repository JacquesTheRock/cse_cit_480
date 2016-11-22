package user

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"strconv"
)

func SearchUsers(u entity.User) ([]entity.User, error) {
	const qBase = "SELECT id,email,name,location,growzone,season,specialty, about FROM users"
	op := "OR "
	queryVars := make([]interface{}, 0)
	out := make([]entity.User, 0)
	query := " WHERE "
	endQuery := qBase
	if u.ID != "" {
		queryVars = append(queryVars, u.ID)
		query = query + "id LIKE $" + strconv.Itoa(len(queryVars)) + " "
	}
	if u.DisplayName != "" {
		if len(queryVars) > 0 {
			query += op
		}
		queryVars = append(queryVars, u.DisplayName)
		query = query + "name LIKE $" + strconv.Itoa(len(queryVars)) + " "
	}
	if u.Location != "" {
		if len(queryVars) > 0 {
			query += "OR "
		}
		queryVars = append(queryVars, u.Location)
		query = query + "location LIKE $" + strconv.Itoa(len(queryVars)) + " "
	}
	if u.Growzone != "" {
		if len(queryVars) > 0 {
			query += op
		}
		queryVars = append(queryVars, u.DisplayName)
		query = query + "growzone = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if u.Season != "" {
		if len(queryVars) > 0 {
			query += op
		}
		queryVars = append(queryVars, u.DisplayName)
		query = query + "season = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if u.Specialty != "" {
		if len(queryVars) > 0 {
			query += op
		}
		queryVars = append(queryVars, u.DisplayName)
		query = query + "specialty = $" + strconv.Itoa(len(queryVars)) + " "
	}

	if len(queryVars) > 0 {
		endQuery = qBase + query
	}
	util.PrintDebug(endQuery)
	rows, err := util.Database.Query(endQuery, queryVars...)
	if err != nil {
		util.PrintDebug(err)
		return out, err
	}
	defer rows.Close()
	for rows.Next() {
		e := entity.User{}
		err = rows.Scan(&e.ID, &e.Email, &e.DisplayName, &e.Location, &e.Growzone, &e.Season, &e.Specialty, &e.About)
		if err != nil {
			util.PrintError("Unable to read user")
			util.PrintDebug(err)
		}
		out = append(out, e)
	}

	return out, nil
}

func GetUser(u entity.User) (entity.User, error) {
	const qBase = "SELECT id,email,name,location,growzone,season,specialty,about FROM users WHERE id = $1"
	out := entity.User{}
	rows, err := util.Database.Query(qBase, u.ID)
	defer rows.Close()
	for rows.Next() {
		e := entity.User{}
		err = rows.Scan(&e.ID, &e.Email, &e.DisplayName, &e.Location, &e.Growzone, &e.Season, &e.Specialty, &e.About)
		if err != nil {
			util.PrintError("Unable to read user")
			util.PrintDebug(err)
		}
		out = e
	}

	return out, nil
}
