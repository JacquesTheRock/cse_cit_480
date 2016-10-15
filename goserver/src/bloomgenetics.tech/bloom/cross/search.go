package cross

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
)

type CrossQuery struct {
	ID        sql.NullInt64
	ProjectID sql.NullInt64
}

func SearchCrosses(q CrossQuery) ([]entity.Cross, error) {
	const qBase = "SELECT id,project_id,name FROM crosses"
	queryVars := make([]interface{}, 0)
	out := make([]entity.Cross, 0)
	query := " WHERE "
	endQuery := qBase
	if q.ID.Valid {
		queryVars = append(queryVars, q.ID.Int64)
		query = query + "id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if q.ProjectID.Valid {
		queryVars = append(queryVars, q.ProjectID.Int64)
		query = query + "project_id = $" + strconv.Itoa(len(queryVars)) + " "
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
		e := entity.Cross{}
		var name sql.NullString
		err = rows.Scan(&e.ID, &e.ProjectID, &name)
		if err != nil {
			util.PrintError("Unable to read Cross")
			util.PrintError(err)
		}
		if name.Valid {
			e.Name = desc.String
		}
		out = append(out, e)
	}
	return out, nil
}
