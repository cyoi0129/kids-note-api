package controllers

import (
	"net/http"

	"kids-note-api/models"
	"kids-note-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchTaskTypeById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		id = 0
	}
	task_type, err := services.FetchTaskTypeById(id)
	tokenString := c.GetHeader("Authorization")
	hasPermission := services.CheckFamilyPermission(tokenString, task_type.Family)
	if !hasPermission {
		c.JSON(http.StatusForbidden, gin.H{"Status": 4, "Data": "No Permission"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": task_type})
	}
}

func FetchTaskTypeListByFamilyId(c *gin.Context) {
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
	task_types, err := services.FetchTaskTypeListByFamilyId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": task_types})
	}
}

func CreateTaskType(c *gin.Context) {
	var input models.TaskType
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
	task_type, db_error := services.CreateTaskType(input)

	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": task_type})
}

func UpdateTaskType(c *gin.Context) {
	var input models.TaskType
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
	task_type, db_error := services.UpdateTaskType(input)
	if db_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": task_type})
}
