package trait

import (
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/util"
)

func NewTrait(t entity.Trait) (entity.Trait, error) {
	const qBase = "INSERT INTO trait(project_id,name,class) VALUES($1,$2,$3) RETURNING id"
	err := util.Database.QueryRow(qBase, t.Project_ID, t.Name, t.Type_ID).Scan(&t.ID)
	if err != nil {
		util.PrintError("Unable to add trait to project")
		return t, err
	}
	var output entity.Trait
	output, err = GetTrait(t)
	return output, err
}

func UpdateTrait(t entity.Trait) (entity.Trait, error) {
	const qBase = "UPDATE trait SET name = $1, class = $2 WHERE project_id = $3 AND id = $4"
	_, err := util.Database.Exec(qBase, t.Name, t.Type_ID, t.Project_ID, t.ID)
	if err != nil {
		util.PrintError("Unable to update trait")
		util.PrintError(err)
		return t, err
	}
	var output entity.Trait
	output, err = GetTrait(t)
	return output, err
}
