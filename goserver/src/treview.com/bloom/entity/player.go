package entity

import (
	"database/sql"
	"nullandvoidgaming.com/projectCloud/item"
	"nullandvoidgaming.com/projectCloud/util"
)

type Equipment struct {
	LeftThumb   item.Ring
	LeftIndex   item.Ring
	LeftMiddle  item.Ring
	LeftRing    item.Ring
	LeftPinky   item.Ring
	RightThumb  item.Ring
	RightIndex  item.Ring
	RightMiddle item.Ring
	RightRing   item.Ring
	RightPinky  item.Ring
}

type Player struct {
	ID           int64
	Name         string
	God          string
	Affinity     string
	Equipment    Equipment
	Intelligence int32
	Strength     int32
	Wisdom       int32
	Agility      int32
	Life         int32
	Vitality     int32
}


func SearchPlayer(p Player, database *sql.DB) ([]Player, error) {
        var output []Player
        rows, err := database.Query("SELECT id, name, intelligence, strength, wisdom,"+
                "agility, life FROM player WHERE id LIKE $1", p.ID)
        if err != nil {
                util.PrintError(err.Error())
                return output, err 
        }
        defer rows.Close()
        for rows.Next() {
                var player Player
                err := rows.Scan(&player.ID, &player.Name, &player.Intelligence,
                        &player.Strength, &player.Wisdom, &player.Agility, &player.Life)
                if err != nil {
                        util.PrintError(output)
                } else {
                        if cap(output) != len(output) {
                                last := len(output)
                                output = output[0 : last+1]
                                output[last] = player
                        } else {
                                break
                        }
                }
        }
        return output, nil 
}
