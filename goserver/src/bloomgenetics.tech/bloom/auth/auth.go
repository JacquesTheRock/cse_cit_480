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
