package controllers

import (
	"net/http"

	"kids-note-api/models"
	"kids-note-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchFamilyById(c *gin.Context) {
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
	family, family_err := services.FetchFamilyById(id)
	members, member_err := services.FetchFamilyMembers(id)
	familyResponse := models.FamilyResponse{
		Id:      family.Id,
		Name:    family.Name,
		Members: members,
	}
	if err != nil || family_err != nil || member_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": familyResponse})
	}
}

func CreateFamily(c *gin.Context) {
	var input models.Family
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 4, "Data": err.Error()})
		return
	}
	family, db_error := services.CreateFamily(input)
	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": family})
}

func UpdateFamily(c *gin.Context) {
	var input models.Family
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 1, "Data": err.Error()})
		return
	}
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckFamilyPermission(tokenString, int(input.Id))
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	family, db_error := services.UpdateFamily(input)

	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": family})
}
