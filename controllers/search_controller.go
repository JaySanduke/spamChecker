package controllers

import (
	"net/http"

	"spamChecker/database"
	"spamChecker/models"

	"github.com/gin-gonic/gin"
)

func getSpamCount(phone string) int64 {
	var count int64
	database.DB.Model(&models.SpamReport{}).Where("phone_number = ?", phone).Count(&count)
	return count
}

type SearchResult struct {
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Email     *string `json:"email,omitempty"`
	SpamCount int64   `json:"spam_count"`
}

func Search(c *gin.Context) {
	name := c.Query("name")
	phone := c.Query("phone")
	// userIDStr, _ := c.Get("userID")
	// userID, _ := strconv.Atoi(userIDStr.(string))

	var results []SearchResult

	if name != "" {
		var users []models.User
		database.DB.Where("name ILIKE ?", name+"%").Find(&users)

		var partial []models.User
		database.DB.Where("name ILIKE ?", "%"+name+"%").
			Where("name NOT ILIKE ?", name+"%").
			Find(&partial)

		users = append(users, partial...)

		for _, u := range users {
			results = append(results, SearchResult{
				Name:      u.Name,
				Phone:     u.Phone,
				SpamCount: getSpamCount(u.Phone),
			})
		}

		c.JSON(http.StatusOK, results)
		return
	}

	if phone != "" {
		var user models.User
		if err := database.DB.Where("phone = ?", phone).First(&user).Error; err == nil {
			result := SearchResult{
				Name:      user.Name,
				Phone:     user.Phone,
				SpamCount: getSpamCount(user.Phone),
			}

			var contact models.Contact
			if err := database.DB.Where("user_id = ? AND phone = ?", user.ID, phone).First(&contact).Error; err == nil {
				result.Email = user.Email
			}

			c.JSON(http.StatusOK, []SearchResult{result})
			return
		}

		var contacts []models.Contact
		database.DB.Where("phone = ?", phone).Find(&contacts)

		for _, cEntry := range contacts {
			results = append(results, SearchResult{
				Name:      cEntry.Name,
				Phone:     cEntry.Phone,
				SpamCount: getSpamCount(cEntry.Phone),
			})
		}

		c.JSON(http.StatusOK, results)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Provide either name or phone to search"})
}

func GetProfile(c *gin.Context) {
	phone := c.Param("phone")
	// userIDStr, _ := c.Get("userID")
	// viewerID, _ := strconv.Atoi(userIDStr.(string))

	var targetUser models.User
	if err := database.DB.Where("phone = ?", phone).First(&targetUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	profile := map[string]interface{}{
		"name":       targetUser.Name,
		"phone":      targetUser.Phone,
		"spam_count": getSpamCount(targetUser.Phone),
	}

	var contact models.Contact
	err := database.DB.Where("user_id = ? AND phone = ?", targetUser.ID, phone).First(&contact).Error

	if err == nil && targetUser.Email != nil {
		profile["email"] = targetUser.Email
	}

	c.JSON(http.StatusOK, profile)
}
