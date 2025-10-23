package services

import (
	"fmt"
	"kids-note-api/models"
)

// 登録済みのメールアドレスかをチェック
func CheckExistUserEmail(email string) int {
	var count int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM \"kids_users\" WHERE email = $1", email).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

// 家族IDに紐付くユーザーを取得
func FetchFamilyMembers(family_id int) ([]models.FamilyMember, error) {
	var members []models.FamilyMember
	rows, err := models.DB.Query("SELECT id, name FROM \"kids_users\" WHERE family = $1", family_id)
	if err != nil {
		fmt.Println(err)
		return members, err
	}
	for rows.Next() {
		var member models.FamilyMember
		rows.Scan(&member.Id, &member.Name)
		members = append(members, member)
	}
	return members, nil
}

func FetchUserById(user_id int) (models.User, error) {
	var user models.User
	err := models.DB.QueryRow("SELECT id, email, name, gender, family FROM \"kids_users\" WHERE id = $1", user_id).Scan(&user.Id, &user.Email, &user.Name, &user.Gender, &user.Family)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func CreateUser(input models.NewUser) (models.User, error) {
	user := models.User{
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
		Gender:   input.Gender,
		Family:   input.Family,
	}
	err := models.DB.QueryRow("INSERT INTO kids_users(email, password, name, gender, family) VALUES($1,$2,$3,$4,$5) RETURNING id", user.Email, user.Password, user.Name, user.Gender, user.Family).Scan(&user.Id)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func UpdateUser(user_id int, input models.User) (models.User, error) {
	user := models.User{
		Id:       input.Id,
		Name:     input.Name,
		Gender:   input.Gender,
		Password: input.Password,
	}
	_, err := models.DB.Query("UPDATE \"kids_users\" SET name = $1, gender = $2, password = $3 WHERE id = $4", user.Name, user.Gender, user.Password, user_id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// ユーザーログイン検証
func CheckUserVaildation(email string, password string) (bool, models.User) {
	var user models.User
	err := models.DB.QueryRow("SELECT id, email, password, name, gender, family FROM \"kids_users\" WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Gender, &user.Family)
	if err != nil || user.Password != password {
		fmt.Println(err)
		return false, user
	}
	return true, user
}
