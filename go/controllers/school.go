package controllers

import (
	"fmt"
	"kids-note-api/models"
	"kids-note-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchSchoolList(c *gin.Context) {
	schools, err := models.GetCache("school") // キャッシュに保存済みのデータを取得
	if err != nil {                           // キャッシュが存在しない場合、DBアクセスへ
		fmt.Println("Access DB")
		schools, err = services.FetchSchoolList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
			return
		}
		models.SetCache("master", schools) // DBから取得したデータをキャッシュへ保存
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": schools})
}
