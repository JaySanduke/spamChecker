package controllers

import (
	"net/http"
	"strconv"

	"spamChecker/database"
	"spamChecker/models"

	"github.com/gin-gonic/gin"
)

type SpamInput struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func MarkSpam(c *gin.Context) {
	var input SpamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number required"})
		return
	}

	userIDStr, _ := c.Get("userID")
	userID, _ := strconv.Atoi(userIDStr.(string))

	report := models.SpamReport{
		ReporterID:  uint(userID),
		PhoneNumber: input.PhoneNumber,
	}

	if err := database.DB.Create(&report).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not mark as spam"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Marked as spam"})
}
