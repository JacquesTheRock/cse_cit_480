package auth

import (
	"bloomgenetics.tech/bloom/util"
	"database/sql"
)

type Role struct {
	UserID    string `json:"user_id"`
	Name      string `json:"role_name"`
	ProjectID int64  `json:"project_id"`
	RoleID    int64  `json:"role_id"`
}

func SetRole(r Role) error {
	const qBase = "INSERT INTO roles(user_id,project_id,role_id) VALUES ($1,$2,$3)"
	_, err := util.Database.Exec(qBase, r.UserID, r.ProjectID, r.RoleID)
	if err != nil {
		util.PrintError("Unable to Attach user to role " + r.UserID)
		util.PrintDebug(err)
		return err
	}
	return nil
}

func UpdateRole(r Role) error {
	const qBase = "UPDATE roles SET role_id = $1 WHERE user_id = $2 AND project_id = $3"
	_, err := util.Database.Exec(qBase, r.RoleID, r.UserID, r.ProjectID)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("Unable to update role for: " + r.UserID)
		return err
	}
	return nil
}

func DeleteRole(r Role) error {
	const qBase = "DELETE FROM roles WHERE user_id = $1 AND project_id = $2 AND role_id <> 5"
	_, err := util.Database.Exec(qBase, r.UserID, r.ProjectID)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("Unable to update role for: " + r.UserID)
		return err
	}
	return nil
}

func GetRole(uid string, pid int64) Role {
	out := Role{}
	out.UserID = uid
	out.ProjectID = pid
	out.RoleID = -1
	out.Name = "Not Found"
	const qBase = "SELECT role_id, name FROM roles r JOIN role_t rt ON r.role_id = rt.id WHERE user_id = $1 AND project_id = $2"
	rows, err := util.Database.Query(qBase, uid, pid)
	if err != nil {
		util.PrintDebug(err)
		util.PrintError("Unable to get Role for: " + uid)
		return out
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
			out.Name = name.String
		}
		out.RoleID = id.Int64
	}
	return out
}

func GetProjectRoles(pid int64) []Role {
	var out []Role
	const qBase = "SELECT r.user_id,r.project_id,r.role_id,rt.name FROm roles r JOIN role_t rt ON r.role_id = rt.id WHERE project_id = $1"
	rows, err := util.Database.Query(qBase, pid)
	if err != nil {
		util.PrintDebug(err)
		return out
	}
	defer rows.Close()
	for rows.Next() {
		r := Role{}
		err = rows.Scan(&r.UserID, &r.ProjectID, &r.RoleID, &r.Name)
		if err != nil {
			util.PrintDebug(err)
			continue
		}
		out = append(out, r)
	}
	return out
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
