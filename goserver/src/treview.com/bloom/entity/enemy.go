package entity


import (
	"strconv"
	"database/sql"
	"nullandvoidgaming.com/projectCloud/util"
)

type Enemy struct {
	ID           int64
	Name         string
	God          string
	Affinity     string
	Intelligence int32
	Strength     int32
	Wisdom       int32
	Agility      int32
	Life         int32
	Vitality     int32

	current_hp int32
}

func SearchEnemy(e Enemy, database *sql.DB) ([]Enemy, error) {
	var output []Enemy
	a := make([]interface{}, 0)
	query := "SELECT e.id, e.name " +
		",strength,agility,vitality " +
		",intelligence,wisdom,life " +
		"FROM enemy e,god g,god a WHERE " +
		"g.id = e.god_id AND " +
		"a.id = e.element_id "
	where := ""
	var count int64 = 0
	if e.Name != "" {
		count++
		a = append(a, e.Name)
		where = where + " AND e.name LIKE $" + strconv.FormatInt(count, 10)
	}
	if e.ID != 0 {
		count++
		a = append(a, e.ID)
		where = "AND e.id = $" + strconv.FormatInt(count, 10)
	}
	query = query + where
	rows, err := database.Query(query, a...)
	if err != nil {
		util.PrintError(err.Error())
		return output, err
	}
	for rows.Next() {
		var enemy Enemy
		err := rows.Scan(&enemy.ID, &enemy.Name,
			&enemy.Strength, &enemy.Agility,
			&enemy.Vitality, &enemy.Intelligence,
			&enemy.Wisdom, &enemy.Life)
		if err != nil {
			util.PrintError(err.Error())
			continue
		}
		output = append(output, enemy)
	}
	return output, nil
}
