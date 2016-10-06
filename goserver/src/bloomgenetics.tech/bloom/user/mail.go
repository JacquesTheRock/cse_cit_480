package user

import (
	"database/sql"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func GetMails(m entity.Mail) ([]entity.Mail, error) {
	const qBase = "SELECT id,prev,dest,src,message,to_char(arrival,'YYYY MM DD HH12:MI:SS') FROM mail WHERE dest = $1"
	out := make([]entity.Mail, 0)	
	rows, err := util.Database.Query(qBase, m.Dest)
	if err != nil {
		util.PrintError(err)
	}
	defer rows.Close()
	for rows.Next() {
		var Prev sql.NullInt64
		e := entity.Mail{}
		err = rows.Scan(&e.ID,&Prev,&e.Dest, &e.Src, &e.Message, &e.Date)
		if err != nil {
			util.PrintError("Unable to read Mail")
			util.PrintError(err)
		}
		if Prev.Valid {
			e.Prev = Prev.Int64
		}
		out = append(out, e)
		
	}

	return out, nil
}

func GetMailByID(m entity.Mail) ([]entity.Mail, error) {
	const qBase = "SELECT id,prev,dest,src,message,to_char(arrival,'YYYY MM DD HH12:MI:SS') FROM mail WHERE dest = $1 AND id = $2"
	out := make([]entity.Mail, 0)	
	rows, err := util.Database.Query(qBase, m.Dest, m.ID)
	if err != nil {
		util.PrintError(err)
	}
	defer rows.Close()
	for rows.Next() {
		var Prev sql.NullInt64
		e := entity.Mail{}
		err = rows.Scan(&e.ID,&Prev,&e.Dest, &e.Src, &e.Message, &e.Date)
		if err != nil {
			util.PrintError("Unable to read Mail")
			util.PrintError(err)
		}
		if Prev.Valid {
			e.Prev = Prev.Int64
		}
		out = append(out, e)
		
	}

	return out, nil
}
