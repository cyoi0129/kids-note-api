package services

import (
	"fmt"
	"kids-note-api/models"

	"github.com/lib/pq"
)

func FetchTaskById(task_id int) (models.Task, error) {
	var task models.Task
	var type_list, item_list []string
	err := models.DB.QueryRow("SELECT id, name, detail, types, status, update, due, items, kid, userId, family FROM \"kids_tasks\" WHERE id = $1", task_id).Scan(&task.Id, &task.Name, &task.Detail, pq.Array(&type_list), &task.Status, &task.Update, &task.Due, pq.Array(&item_list), &task.Kid, &task.UserId, &task.Family)
	if err != nil {
		fmt.Println(err)
		return task, err
	}
	task.Types = convert2Int(type_list)
	task.Items = convert2Int(item_list)
	return task, nil
}

func FetchTaskListByFamilyId(family_id int) ([]models.Task, error) {
	var tasks []models.Task
	rows, err := models.DB.Query("SELECT id, name, detail, types, status, update, due, items, kid, userId, family FROM \"kids_tasks\" WHERE family = $1", family_id)
	if err != nil {
		fmt.Println(err)
		return tasks, err
	}
	for rows.Next() {
		var task models.Task
		var type_list, item_list []string
		rows.Scan(&task.Id, &task.Name, &task.Detail, pq.Array(&type_list), &task.Status, &task.Update, &task.Due, pq.Array(&item_list), &task.Kid, &task.UserId, &task.Family)
		task.Types = convert2Int(type_list)
		task.Items = convert2Int(item_list)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CreateTask(input models.Task) (models.Task, error) {
	task := models.Task{
		Name:   input.Name,
		Detail: input.Detail,
		Types:  input.Types,
		Status: input.Status,
		Update: input.Update,
		Due:    input.Due,
		Items:  input.Items,
		Kid:    input.Kid,
		UserId: input.UserId,
		Family: input.Family,
	}
	err := models.DB.QueryRow("INSERT INTO kids_tasks(name, detail, types, status, update, due, items, kid, userId, family) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id", task.Name, task.Detail, pq.Array(task.Types), task.Status, task.Update, task.Due, pq.Array(task.Items), task.Kid, task.UserId, task.Family).Scan(&task.Id)
	if err != nil {
		fmt.Println(err)
		return task, err
	}
	return task, nil
}

func UpdateTask(input models.Task) (models.Task, error) {
	task := models.Task{
		Id:     input.Id,
		Name:   input.Name,
		Detail: input.Detail,
		Types:  input.Types,
		Status: input.Status,
		Update: input.Update,
		Due:    input.Due,
		Items:  input.Items,
		Kid:    input.Kid,
		UserId: input.UserId,
		Family: input.Family,
	}
	_, err := models.DB.Query("UPDATE \"kids_tasks\" SET name = $1, detail = $2, types = $3, status = $4, update = $5, due = $6, items = $7, kid = $8, userId = $9, family = $10 WHERE id = $11", task.Name, task.Detail, pq.Array(task.Types), task.Status, task.Update, task.Due, pq.Array(task.Items), task.Kid, task.UserId, task.Family, task.Id)
	if err != nil {
		return task, err
	}
	return task, nil
}

func SetTaskDone(ids []int) error {
	_, err := models.DB.Query("UPDATE \"kids_tasks\" SET status = 'DONE' WHERE id = ANY($1)", pq.Array(ids))
	return err
}
