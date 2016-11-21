package user

import (
	"bloomgenetics.tech/bloom/auth"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func UpdateUser(u entity.User) (entity.User, error) {
	const qBase = "UPDATE users SET name = $2, email = $3, location = $4, growzone = $5, season = $6, specialty = $7, about = $8 WHERE id = $1"
	_, err := util.Database.Exec(qBase, u.ID, u.DisplayName, u.Email, u.Location, u.Growzone, u.Season, u.Specialty, u.About)
	if err != nil {
		util.PrintError("Failure to Update User")
		return entity.User{}, err
	}
	u, err = GetUser(u)
	if err != nil {
		util.PrintError("Failure to find user data")
		return entity.User{}, err
	}
	return u, nil

}

func CreateUser(u entity.User, hash []byte, salt []byte) (entity.User, error) {
	const qBase = "INSERT INTO users(id,email,name,location,growzone,season,specialty,about, hash,salt,algorithm) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,'SHA512')"
	user := entity.User{}
	_, err := util.Database.Exec(qBase, u.ID, u.Email, u.DisplayName, u.Location, u.Growzone, u.Season, u.Specialty, u.About, hash, salt)
	if err != nil {
		util.PrintError("createUser Function")
		util.PrintDebug(err)
		return user, err
	}
	user, _ = GetUser(u)
	auth.SetRole(auth.Role{UserID: user.ID, ProjectID: 0, RoleID: 2})
	return user, nil
}
