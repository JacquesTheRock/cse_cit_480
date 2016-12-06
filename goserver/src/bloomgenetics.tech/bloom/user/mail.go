package user

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
)

func GetMails(m entity.Mail) ([]entity.Mail, error) {
	const qBase = "SELECT id,prev,dest,src,subject,message,to_char(arrival,'YYYY MM DD HH12:MI:SS') FROM mail WHERE dest = $1"
	out := make([]entity.Mail, 0)
	rows, err := util.Database.Query(qBase, m.Dest)
	if err != nil {
		util.PrintDebug(err)
	}
	defer rows.Close()
	for rows.Next() {
		var Prev sql.NullInt64
		e := entity.Mail{}
		err = rows.Scan(&e.ID, &Prev, &e.Dest, &e.Src, &e.Subject, &e.Message, &e.Date)
		if err != nil {
			util.PrintError("Unable to read Mail")
			util.PrintDebug(err)
		}
		if Prev.Valid {
			e.Prev = Prev.Int64
		}
		out = append(out, e)

	}

	return out, nil
}

func GetMailByID(m entity.Mail) ([]entity.Mail, error) {
	const qBase = "SELECT id,prev,dest,src,subject,message,to_char(arrival,'YYYY MM DD HH12:MI:SS') FROM mail WHERE dest = $1 AND id = $2"
	out := make([]entity.Mail, 0)
	rows, err := util.Database.Query(qBase, m.Dest, m.ID)
	if err != nil {
		util.PrintDebug(err)
	}
	defer rows.Close()
	for rows.Next() {
		var Prev sql.NullInt64
		e := entity.Mail{}
		err = rows.Scan(&e.ID, &Prev, &e.Dest, &e.Src, &e.Subject, &e.Message, &e.Date)
		if err != nil {
			util.PrintError("Unable to read Mail")
			util.PrintDebug(err)
		}
		if Prev.Valid {
			e.Prev = Prev.Int64
		}
		out = append(out, e)

	}

	return out, nil
}

func PostMail(m entity.Mail) (entity.Mail, error) {
	const qBase = "INSERT INTO mail(dest,src,subject,message) VALUES ($1,$2,$3,$4)"
	_, err := util.Database.Exec(qBase, m.Dest, m.Src, m.Subject, m.Message)
	if err != nil {
		util.PrintDebug(err)
		return entity.Mail{}, err
	}
	m = entity.Mail{Dest: m.Dest, Src: m.Src}
	return m, nil
}

func ReplyMail(m entity.Mail) (entity.Mail, error) {
	const qBase = "INSERT INTO mail(dest,src,subject,message,prev) VALUES ($1,$2,$3,$4,$5)"
	_, err := util.Database.Exec(qBase, m.Dest, m.Src, m.Subject, m.Message, m.Prev)
	if err != nil {
		util.PrintDebug(err)
		return entity.Mail{}, err
	}
	m = entity.Mail{Dest: m.Dest, Src: m.Src, Message: "Message Replied successfully"}
	return m, nil
}

func DeleteMail(m entity.Mail) (entity.Mail, error) {
	const qBase = "DELETE FROM mail WHERE src = $1 AND id = $2"
	_, err := util.Database.Exec(qBase, m.Src, m.ID)
	if err != nil {
		util.PrintDebug(err)
		return entity.Mail{}, nil
	}
	m = entity.Mail{ID: m.ID, Src: m.Src}
	return m, nil

}
