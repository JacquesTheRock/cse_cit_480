package user

import (
	"bloomgenetics.tech/bloom/util"
	"database/sql"
	"strconv"
)

type QueryProjectRole struct {
	UID string
	PID sql.NullInt64
	RID sql.NullInt64
}

type ProjectRole struct {
	UID string `json:"uid"`
	PID int64  `json:"pid"`
	RID int64  `json:"rid"`
}

func SearchProjects(p QueryProjectRole) ([]ProjectRole, error) {
	const qBase = " SELECT user_id,project_id,role_id FROM roles"
	queryVars := make([]interface{}, 0)
	out := make([]ProjectRole, 0)
	query := " WHERE "
	endQuery := qBase
	if p.UID != "" {
		queryVars = append(queryVars, p.UID)
		query = query + "user_id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if p.PID.Valid {
		queryVars = append(queryVars, p.PID.Int64)
		query = query + "project_id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if p.RID.Valid {
		queryVars = append(queryVars, p.RID.Int64)
		query = query + "role_id = $" + strconv.Itoa(len(queryVars)) + " "
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
		e := ProjectRole{}
		err = rows.Scan(&e.UID, &e.PID, &e.RID)
		if err != nil {
			util.PrintError("Unable to read project")
			util.PrintError(err)
		}
		out = append(out, e)
	}
	return out, nil
}
