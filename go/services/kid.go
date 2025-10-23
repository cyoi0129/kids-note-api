package services

import (
	"fmt"
	"kids-note-api/models"
)

func FetchKidListByFamilyId(family_id int) ([]models.Kid, error) {
	var kids []models.Kid
	rows, err := models.DB.Query("SELECT id, name, birth, gender, family, school FROM \"kids_kids\" WHERE family = $1", family_id)
	if err != nil {
		fmt.Println(err)
		return kids, err
	}
	for rows.Next() {
		var kid models.Kid
		rows.Scan(&kid.Id, &kid.Name, &kid.Birth, &kid.Gender, &kid.Family, &kid.School)
		kids = append(kids, kid)
	}
	return kids, nil
}

func FetchKidById(kid_id int) (models.Kid, error) {
	var kid models.Kid
	err := models.DB.QueryRow("SELECT id, name, birth, gender, family, school FROM \"kids_kids\" WHERE id = $1", kid_id).Scan(&kid.Id, &kid.Name, &kid.Birth, &kid.Gender, &kid.Family, &kid.School)

	if err != nil {
		fmt.Println(err)
		return kid, err
	}
	return kid, nil
}

func CreateKid(input models.Kid) (models.Kid, error) {
	kid := models.Kid{
		Name:   input.Name,
		Birth:  input.Birth,
		Gender: input.Gender,
		Family: input.Family,
		School: input.School,
	}
	err := models.DB.QueryRow("INSERT INTO kids_kids(name, birth, gender, family, school) VALUES($1,$2,$3,$4,$5) RETURNING id", kid.Name, kid.Birth, kid.Gender, kid.Family, kid.School).Scan(&kid.Id)
	if err != nil {
		fmt.Println(err)
		return kid, err
	}
	return kid, nil
}

func UpdateKid(input models.Kid) (models.Kid, error) {
	kid := models.Kid{
		Id:     input.Id,
		Name:   input.Name,
		Birth:  input.Birth,
		Gender: input.Gender,
		Family: input.Family,
		School: input.School,
	}
	_, err := models.DB.Query("UPDATE \"kids_kids\" SET name = $1, birth = $2, gender = $3, family = $4, school = $5 WHERE id = $6", kid.Name, kid.Birth, kid.Gender, kid.Family, kid.School, kid.Id)
	if err != nil {
		return kid, err
	}
	return kid, nil
}
