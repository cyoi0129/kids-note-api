package controllers

import (
	"fmt"
	"net/http"

	"kids-note-api/models"
	"kids-note-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IDを指定してユーザー情報を取得
func FetchUserInfo(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		id = 0
	}
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckUserPermission(tokenString, id)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	user, err := services.FetchUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": user})
	}
}

// 新たなユーザーの追加
func CreateUser(c *gin.Context) {
	var input models.NewUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 4, "Data": err.Error()})
		return
	}
	// メールトークンの有効性チェック
	isTokenValid := services.CheckMailToken(input.Token)
	if !isTokenValid {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": 4, "Data": "Token Error"})
		return
	}
	// 家族パラメータない場合は、家族を新規作成
	if input.Family > 0 {
	} else {
		family, err := services.CreateFamily(models.Family{Name: ""})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
			return
		}
		input.Family = int(family.Id)
	}
	// ユーザー新規作成
	user, err := services.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	// 新トークン発行プロセス
	user_token, err := services.CreateToken(int(user.Id), user.Family)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}

	c.Header("Authorization", user_token)
	userResponse := models.UserResponse{
		Token: user_token,
		Info: models.User{
			Id:     user.Id,
			Name:   user.Name,
			Email:  user.Email,
			Family: user.Family,
		},
	}

	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": userResponse})
}

// 既存ユーザーの更新
func UpdateUser(c *gin.Context) {
	user_param := c.Param("id")
	id, err := strconv.Atoi(user_param)
	if err != nil {
		id = 0
	}
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 1, "Data": err.Error()})
		return
	}
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckUserPermission(tokenString, id)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	user, db_error := services.UpdateUser(id, input)
	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": user})
	}
}

// ログインAPIがPOSTされたときのプロセス
func LoginHandler(c *gin.Context) {
	var inputUser models.LoginUser
	// リクエストからユーザー情報を取得
	if err := c.BindJSON(&inputUser); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Status": 4, "Data": "Login Error"})
		return
	}
	loginResult, user := services.CheckUserVaildation(inputUser.Email, inputUser.Password)
	// ユーザー情報の検証
	if !loginResult {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": 4, "Data": "Login Error"})
		return
	}
	// トークンの発行（ヘッダー・ペイロード）
	token, err := services.CreateToken(int(user.Id), user.Family)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	// ヘッダーにトークンをセット
	c.Header("Authorization", token)
	userResponse := models.UserResponse{
		Info:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": userResponse})
}

// メールAPIのトークンからの認証チェック
func MailAuthMiddleware(c *gin.Context) {
	// Authorizationヘッダーからトークンを取得
	tokenString := c.GetHeader("Authorization")
	isTokenValid := services.CheckMailToken(tokenString)
	if isTokenValid {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": 4, "Data": "Auth Error"})
		c.Abort()
		return
	}
	c.Next()
}

// APIのトークンからの認証チェック
func AuthMiddleware(c *gin.Context) {
	// Authorizationヘッダーからトークンを取得
	tokenString := c.GetHeader("Authorization")
	_, err := services.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": 4, "Data": "Auth Error"})
		c.Abort()
		return
	}
	c.Next()
}
