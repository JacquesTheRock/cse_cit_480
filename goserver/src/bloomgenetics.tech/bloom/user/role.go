package user

import (
	"bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
	"strconv"
)

type QueryProjectRole struct {
	UID string
	PID sql.NullInt64
	RID sql.NullInt64
}

func SearchProjects(p QueryProjectRole) ([]auth.Role, error) {
	const qBase = " SELECT user_id,project_id,role_id FROM roles"
	queryVars := make([]interface{}, 0)
	out := make([]auth.Role, 0)
	query := " WHERE "
	endQuery := qBase
	if p.UID != "" {
		queryVars = append(queryVars, p.UID)
		query = query + "user_id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if p.PID.Valid {
		if len(queryVars) > 0 {
			query += "AND "
		}
		queryVars = append(queryVars, p.PID.Int64)
		query = query + "project_id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if p.RID.Valid {
		if len(queryVars) > 0 {
			query += "AND "
		}
		queryVars = append(queryVars, p.RID.Int64)
		query = query + "role_id = $" + strconv.Itoa(len(queryVars)) + " "
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
		e := auth.Role{}
		err = rows.Scan(&e.UserID, &e.ProjectID, &e.RoleID)
		if err != nil {
			util.PrintError("Unable to read project")
			util.PrintDebug(err)
		}
		out = append(out, e)
	}
	return out, nil
}
