package services

import (
	"fmt"
	"kids-note-api/models"

	"github.com/lib/pq"
)

func FetchItemById(item_id int) (models.Item, error) {
	var item models.Item
	err := models.DB.QueryRow("SELECT id, name, detail, type, image, kid, family FROM \"kids_items\" WHERE id = $1", item_id).Scan(&item.Id, &item.Name, &item.Detail, &item.Type, &item.Image, &item.Kid, &item.Family)

	if err != nil {
		fmt.Println(err)
		return item, err
	}
	return item, nil
}

func FetchItemListByFamilyId(family_id int) ([]models.Item, error) {
	var items []models.Item
	rows, err := models.DB.Query("SELECT id, name, detail, type, image, kid, family FROM \"kids_items\" WHERE family = $1", family_id)

	if err != nil {
		fmt.Println(err)
		return items, err
	}

	for rows.Next() {
		var item models.Item
		rows.Scan(&item.Id, &item.Name, &item.Detail, &item.Type, &item.Image, &item.Kid, &item.Family)
		items = append(items, item)
	}

	return items, nil
}

func FetchItemListByIds(item_ids []int) ([]models.Item, error) {
	var items []models.Item
	rows, err := models.DB.Query("SELECT id, name, detail, type, image, kid, family FROM \"kids_items\" WHERE id in $1", pq.Array(item_ids))

	if err != nil {
		fmt.Println(err)
		return items, err
	}

	for rows.Next() {
		var item models.Item
		rows.Scan(&item.Id, &item.Name, &item.Detail, &item.Type, &item.Image, &item.Kid, &item.Family)
		items = append(items, item)
	}

	return items, nil
}

func CreateItem(input models.Item) (models.Item, error) {
	item := models.Item{
		Name:   input.Name,
		Detail: input.Detail,
		Type:   input.Type,
		Image:  input.Image,
		Kid:    input.Kid,
		Family: input.Family,
	}
	err := models.DB.QueryRow("INSERT INTO kids_items(name, detail, type, image, kid, family) VALUES($1,$2,$3,$4,$5,$6) RETURNING id", item.Name, item.Detail, item.Type, item.Image, item.Kid, item.Family).Scan(&item.Id)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	return item, nil
}

func UpdateItem(input models.Item) (models.Item, error) {
	item := models.Item{
		Id:     input.Id,
		Name:   input.Name,
		Detail: input.Detail,
		Type:   input.Type,
		Image:  input.Image,
		Kid:    input.Kid,
		Family: input.Family,
	}
	_, err := models.DB.Query("UPDATE \"kids_items\" SET name = $1, detail = $2, type = $3, image = $4, kid = $5, family = $6 WHERE id = $7", item.Name, item.Detail, item.Type, item.Image, item.Kid, item.Family, item.Id)
	if err != nil {
		return item, err
	}
	return item, nil
}
