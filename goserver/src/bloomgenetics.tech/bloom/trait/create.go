package trait

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func NewTrait(t entity.Trait) (entity.Trait, error) {
	const qBase = "INSERT INTO trait(project_id,name,class,pool) VALUES($1,$2,$3,$4) RETURNING id"
	err := util.Database.QueryRow(qBase, t.Project_ID, t.Name, t.Type_ID, t.Pool).Scan(&t.ID)
	if err != nil {
		util.PrintError("Unable to add trait to project")
		return t, err
	}
	var output entity.Trait
	output, err = GetTrait(t)
	return output, err
}

func UpdateTrait(t entity.Trait) (entity.Trait, error) {
	const qBase = "UPDATE trait SET name = $1, class = $2, pool = $3 WHERE project_id = $4 AND id = $5"
	_, err := util.Database.Exec(qBase, t.Name, t.Type_ID, t.Pool, t.Project_ID, t.ID)
	if err != nil {
		util.PrintError("Unable to update trait")
		util.PrintError(err)
		return t, err
	}
	var output entity.Trait
	output, err = GetTrait(t)
	return output, err
}

func DeleteTrait(t entity.Trait) (entity.Trait, error) {
	const qBase = "DELETE FROM trait WHERE project_id = $1 AND id = $2"
	c, _ := GetTrait(t)
	_, err := util.Database.Exec(qBase, t.Project_ID, t.ID)
	if err != nil {
		util.PrintError("Unable to delete trait")
		util.PrintError(err)
		return t, err
	}
	return c, nil
}
