package controllers

import (
	"net/http"

	"kids-note-api/models"
	"kids-note-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchKidById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		id = 0
	}
	kid, err := services.FetchKidById(id)
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckFamilyPermission(tokenString, kid.Family)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": kid})
	}
}

func FetchKidListByFamilyId(c *gin.Context) {
	user_param := c.Param("id")
	id, err := strconv.Atoi(user_param)
	if err != nil {
		id = 0
	}
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckFamilyPermission(tokenString, id)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	kid, err := services.FetchKidListByFamilyId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": kid})
	}
}

func CreateKid(c *gin.Context) {
	var input models.Kid
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 4, "Data": err.Error()})
		return
	}
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckFamilyPermission(tokenString, input.Family)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	kid, db_error := services.CreateKid(input)
	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": kid})
}

func UpdateKid(c *gin.Context) {
	var input models.Kid
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 4, "Data": err.Error()})
		return
	}
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckFamilyPermission(tokenString, input.Family)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	kid, db_error := services.UpdateKid(input)
	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": kid})
}
