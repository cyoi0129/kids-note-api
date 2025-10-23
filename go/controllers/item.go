package controllers

import (
	"net/http"

	"kids-note-api/models"
	"kids-note-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchItemById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		id = 0
	}
	item, err := services.FetchItemById(id)
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckFamilyPermission(tokenString, item.Family)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": item})
	}
}

func FetchItemListByFamilyId(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		id = 0
	}
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckFamilyPermission(tokenString, id)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	item, err := services.FetchItemListByFamilyId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": item})
	}
}

func CreateItem(c *gin.Context) {
	var input models.Item
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
	item, db_error := services.CreateItem(input)

	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": item})
}

func UpdateItem(c *gin.Context) {
	var input models.Item
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
	item, db_error := services.UpdateItem(input)

	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": item})
}
