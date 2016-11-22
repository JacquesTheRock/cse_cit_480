package cross

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
	"strconv"
)

type CrossQuery struct {
	ID        sql.NullInt64
	ProjectID sql.NullInt64
	Name      sql.NullString
}

func SearchCrosses(q CrossQuery) ([]entity.Cross, error) {
	const qBase = "SELECT id,project_id,name,description FROM crosses"
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
	if q.Name.Valid {
		queryVars = append(queryVars, q.Name.String)
		query = query + "name LIKE $" + strconv.Itoa(len(queryVars)) + " "

	}
	if len(queryVars) > 0 {
		endQuery = qBase + query
	}
	rows, err := util.Database.Query(endQuery, queryVars...)
	if err != nil {
		util.PrintError("Query Error")
		util.PrintDebug(err)
		return out, err
	}
	defer rows.Close()
	for rows.Next() {
		e := entity.Cross{}
		var name sql.NullString
		err = rows.Scan(&e.ID, &e.ProjectID, &name, &e.Description)
		if err != nil {
			util.PrintError("Unable to read Cross")
			util.PrintDebug(err)
		}
		if name.Valid {
			e.Name = name.String
		}
		out = append(out, e)
	}
	return out, nil
}

func GetCross(q CrossQuery) (entity.Cross, error) {
	const qBase = "SELECT c.id,c.project_id,c.name,c.description,p.specimen_id FROM crosses c LEFT JOIN cross_parent p ON c.id = p.cross_id WHERE c.id = $1 AND c.project_id = $2"
	out := entity.Cross{}
	rows, err := util.Database.Query(qBase, q.ID.Int64, q.ProjectID.Int64)
	if err != nil {
		util.PrintError("Query Error")
		util.PrintDebug(err)
		return out, err
	}
	defer rows.Close()
	for rows.Next() {
		var name sql.NullString
		var parent sql.NullInt64
		err = rows.Scan(&out.ID, &out.ProjectID, &name, &out.Description, &parent)
		if err != nil {
			util.PrintError("Unable to read Cross")
			util.PrintDebug(err)
		}
		if name.Valid {
			out.Name = name.String
		}
		if parent.Valid {
			if out.Parent1ID != 0 {
				out.Parent2ID = parent.Int64
			} else {
				out.Parent1ID = parent.Int64
			}
		}
	}
	return out, nil

}
