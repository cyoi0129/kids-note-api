package services

import (
	"fmt"
	"kids-note-api/models"
)

func FetchSchoolList() ([]models.School, error) {
	var schools []models.School
	rows, err := models.DB.Query("SELECT id, prefecture, city, type, name FROM \"kids_schools\"")

	if err != nil {
		fmt.Println(err)
		return schools, err
	}

	for rows.Next() {
		var school models.School
		rows.Scan(&school.Id, &school.Prefecture, &school.City, &school.Type, &school.Name)
		schools = append(schools, school)
	}

	return schools, nil
}
