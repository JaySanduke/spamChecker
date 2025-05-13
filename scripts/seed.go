package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"spamChecker/config"
	"spamChecker/database"
	"spamChecker/models"

	"golang.org/x/crypto/bcrypt"
)

func randomPhone() string {
	prefixes := []string{"9", "8", "7"}
	prefix := prefixes[rand.Intn(len(prefixes))]
	number := ""

	for range 9 {
		number += strconv.Itoa(rand.Intn(10))
	}

	return prefix + number
}

func randomEmail(name string, rnd *rand.Rand) *string {
	if rnd.Intn(2) == 0 {
		email := fmt.Sprintf("%s%d@example.com", name, rnd.Intn(100))
		return &email
	}
	return nil
}

func randomName(rnd *rand.Rand) string {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Helen", "Ivy", "John"}
	return names[rnd.Intn(len(names))]
}

func main() {

	config.LoadConfig("development")

	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	database.Connect()

	var users []models.User

	for range 10 {
		name := randomName(rnd)
		phone := randomPhone()
		email := randomEmail(name, rnd)
		password, _ := bcrypt.GenerateFromPassword([]byte("pass123"), bcrypt.DefaultCost)

		user := models.User{
			Name:     name,
			Phone:    phone,
			Email:    email,
			Password: string(password),
		}
		database.DB.Create(&user)
		users = append(users, user)

		for range 5 + rnd.Intn(6) {
			contact := models.Contact{
				UserID: user.ID,
				Name:   randomName(rnd),
				Phone:  randomPhone(),
			}
			database.DB.Create(&contact)
		}
	}

	for range 30 {
		reporter := users[rnd.Intn(len(users))]
		phone := randomPhone()

		report := models.SpamReport{
			ReporterID:  reporter.ID,
			PhoneNumber: phone,
		}

		database.DB.Create(&report)
	}

	fmt.Println("✅ Seed data created successfully")
}
