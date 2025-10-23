package services

import (
	"fmt"
	"kids-note-api/models"
)

func FetchFamilyById(family_id int) (models.Family, error) {
	var family models.Family
	err := models.DB.QueryRow("SELECT id, name FROM \"kids_families\" WHERE id = $1", family_id).Scan(&family.Id, &family.Name)

	if err != nil {
		fmt.Println(err)
		return family, err
	}
	return family, nil
}

func CreateFamily(input models.Family) (models.Family, error) {
	family := models.Family{
		Name: input.Name,
	}
	err := models.DB.QueryRow("INSERT INTO kids_families(name) VALUES($1) RETURNING id", family.Name).Scan(&family.Id)
	if err != nil {
		fmt.Println(err)
		return family, err
	}
	return family, nil
}

func UpdateFamily(input models.Family) (models.Family, error) {
	family := models.Family{
		Id:   input.Id,
		Name: input.Name,
	}
	_, err := models.DB.Query("UPDATE \"kids_families\" SET name = $1 WHERE id = $2", family.Name, family.Id)
	if err != nil {
		return family, err
	}
	return family, nil
}
