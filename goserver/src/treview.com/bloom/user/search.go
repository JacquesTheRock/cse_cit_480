package user

import (
	"treview.com/bloom/entity"
	"treview.com/bloom/util"
)

func SearchUsers(u entity.User) ([]entity.User, error) {
	const qBase = "SELECT id,email,name,location FROM users"
	queryVars := make([]interface{}, 0)
	out := make([]entity.User, 0)
	query := " WHERE "
	endQuery := qBase
	if u.ID != "" {
		queryVars = append(queryVars, u.ID)
		query = query + "id LIKE $" + string(len(queryVars)) + " "
	} else {
		if u.DisplayName != "" {
			queryVars = append(queryVars, u.DisplayName)
			query = query + "name LIKE $" + string(len(queryVars)) + " "
		}
		if len(queryVars) > 0 {
			endQuery = qBase + query
		}
	}
	rows, err := util.Database.Query(endQuery, queryVars...)
	defer rows.Close()
	for rows.Next() {
		e := entity.User{}
		err = rows.Scan(&e.ID, &e.Email, &e.DisplayName, &e.Location)
		if err != nil {
			util.PrintError("Unable to read user")
			util.PrintError(err)
		}
		out = append(out, e)
	}

	return out, nil
}

func GetUser(u entity.User) (entity.User, error) {
	const qBase = "SELECT id,email,name,location FROM users WHERE id = $1"
	out := entity.User{}
	rows, err := util.Database.Query(qBase, u.ID)
	defer rows.Close()
	for rows.Next() {
		e := entity.User{}
		err = rows.Scan(&e.ID, &e.Email, &e.DisplayName, &e.Location)
		if err != nil {
			util.PrintError("Unable to read user")
			util.PrintError(err)
		}
		out = e
	}

	return out, nil
}
