package services

import (
	"fmt"
	"kids-note-api/models"
)

func FetchTaskTypeById(id int) (models.TaskType, error) {
	var task_type models.TaskType
	err := models.DB.QueryRow("SELECT id, name, family FROM \"kids_task_types\" WHERE id = $1", id).Scan(&task_type.Id, &task_type.Name, &task_type.Family)

	if err != nil {
		fmt.Println(err)
		return task_type, err
	}
	return task_type, nil

}

func FetchTaskTypeListByFamilyId(family_id int) ([]models.TaskType, error) {
	var task_types []models.TaskType
	rows, err := models.DB.Query("SELECT id, name, family FROM \"kids_task_types\" WHERE family = $1", family_id)
	if err != nil {
		fmt.Println(err)
		return task_types, err
	}
	for rows.Next() {
		var task_type models.TaskType
		rows.Scan(&task_type.Id, &task_type.Name, &task_type.Family)
		task_types = append(task_types, task_type)
	}
	return task_types, nil
}

func CreateTaskType(input models.TaskType) (models.TaskType, error) {
	task_type := models.TaskType{
		Name:   input.Name,
		Family: input.Family,
	}
	err := models.DB.QueryRow("INSERT INTO kids_task_types(name, family) VALUES($1,$2) RETURNING id", task_type.Name, task_type.Family).Scan(&task_type.Id)
	if err != nil {
		fmt.Println(err)
		return task_type, err
	}
	return task_type, nil
}

func UpdateTaskType(input models.TaskType) (models.TaskType, error) {
	task_type := models.TaskType{
		Id:     input.Id,
		Name:   input.Name,
		Family: input.Family,
	}
	_, err := models.DB.Query("UPDATE \"kids_task_types\" SET name = $1, family = $2 WHERE id = $3", task_type.Name, task_type.Family, task_type.Id)
	if err != nil {
		return task_type, err
	}
	return task_type, nil
}
