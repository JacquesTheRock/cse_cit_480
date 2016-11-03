package auth

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
	"database/sql"
)

func VerifyPermissions(auth string) bool {
	uid, _ := ParseAuthorization(auth)
	u := GetLogin(auth)
	return u.ID == uid
}

func GetLogin(auth string) entity.UserLogin {
	u := entity.UserLogin{
		"",
		"Guest",
		"",
	}
	uid, token := ParseAuthorization(auth)
	result, err := searchToken(uid, token)
	if err != nil {
		util.PrintError("Failure to Auth Token for: " + uid)
		return u
	}
	return result
}

func CheckAuth(uid string, pid int, pagename string, method string) (bool, error) {
	const qBase = "select page from roles r join role_perm rp using(role_id) join perm p on p.id = rp.perm_id where r.user_id = $1 and r.project_id = $2 and p.page = $3 and p.action = $4"
	rows, err := util.Database.Query(qBase, uid, pid, pagename, method)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("User ID not found: " + uid)
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		var page sql.NullString
		err = rows.Scan(&page)
		if page.Valid && page.String == pagename {
			return true, nil
		}
		if err != nil {
			util.PrintDebug(err)
		}
	}
	return false, nil
}

func SetRole(u entity.User, pid, rid int) error {
	const qBase = "INSERT INTO roles(user_id,project_id,role_id) VALUES ($1,$2,$3)"
	_, err := util.Database.Exec(qBase, u.ID, pid, rid)
	if err != nil {
		util.PrintError("Unable to Attach user to role" + u.ID)
		util.PrintDebug(err)
		return err
	}
	return nil
}

func GetRole(uid string, pid int64) (int64, string) {
	const qBase = "SELECT role_id, name FROm roles r JOIN role_t rt ON r.role_id = rt.id WHERE user_id = $1 AND project_id = $2"
	rows, err := util.Database.Query(qBase, uid, pid)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("Unable to get Role for: " + uid)
		return -1, "NOT FOUND"
	}
	defer rows.Close()
	for rows.Next() {
		var name sql.NullString
		var id sql.NullInt64
		err = rows.Scan(&id, &name)
		if err != nil {
			util.PrintDebug(err)
			continue
		}
		if name.Valid {
			return id.Int64, name.String
		} else {
			return id.Int64, "Unnamed"
		}
	}
	return -1, "NOT FOUND"
}

func GetRoleID(name string) int64 {
	const qBase = "SELECT id FROM role_t WHERE name = $1"
	rows, err := util.Database.Query(qBase, name)
	if err != nil {
		util.PrintDebug(err)
		return -1
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			util.PrintDebug(err)
			continue
		}
		return id
	}
	return -1
}
