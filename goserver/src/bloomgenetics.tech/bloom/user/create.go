package user

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func UpdateUser(u entity.User) (entity.User, error) {
	const qBase = "UPDATE users SET name = $1, email = $2, location = $3 WHERE id = $4"
	_, err := util.Database.Exec(qBase, u.DisplayName, u.Email, u.Location, u.ID)
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

func CreateUser(uid string, email string, name string, location string, hash []byte, salt []byte) (entity.User, error) {
	const qBase = "INSERT INTO users(id,email,name,location,hash,salt,algorithm) VALUES ($1,$2,$3,$4,$5,$6,'SHA512')"
	user := entity.User{}
	_, err := util.Database.Exec(qBase, uid, email, name, location, hash, salt)
	if err != nil {
		util.PrintError("createUser Function")
		util.PrintError(err)
		return user, err
	}
	user.ID = uid
	user.Email = email
	user.DisplayName = name
	user.Location = location
	return user, nil
}
