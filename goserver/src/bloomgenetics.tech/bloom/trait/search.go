package trait

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"strconv"
)

func SearchTraits(t entity.Trait) ([]entity.Trait, error) {
	const qBase = "SELECT a.id,a.project_id,a.name,a.pool,b.weight,b.name,b.id FROM trait a, trait_t b WHERE a.class = b.id"
	queryVars := make([]interface{}, 0)
	out := make([]entity.Trait, 0)
	query := " "
	endQuery := qBase
	if t.ID != 0 {
		queryVars = append(queryVars, t.ID)
		query = query + "AND a.id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if t.Project_ID != 0 {
		queryVars = append(queryVars, t.Project_ID)
		query = query + "AND a.project_id = $" + strconv.Itoa(len(queryVars)) + " "
	}
	if t.Name != "" {
		queryVars = append(queryVars, t.Name)
		query = query + "AND a.name LIKE $" + strconv.Itoa(len(queryVars)) + " "
	}
	if len(queryVars) > 0 {
		endQuery = qBase + query
	}
	util.PrintInfo(endQuery)
	rows, err := util.Database.Query(endQuery, queryVars...)
	defer rows.Close()
	for rows.Next() {
		e := entity.Trait{}
		err = rows.Scan(&e.ID, &e.Project_ID, &e.Name, &e.Pool, &e.Weight, &e.Type, &e.Type_ID)
		if err != nil {
			util.PrintError("Unable to read trait")
			util.PrintDebug(err)
		}
		out = append(out, e)
	}
	return out, nil
}

func GetTrait(t entity.Trait) (entity.Trait, error) {
	out := entity.Trait{}
	tArray, _ := SearchTraits(t)
	if len(tArray) == 1 {
		return tArray[0], nil
	}
	return out, nil
}
