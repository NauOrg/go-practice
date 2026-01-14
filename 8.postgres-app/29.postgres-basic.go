package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect() *gorm.DB {

	dsn := "host=localhost user=root password=example dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect postgres database:", err)
	}

	return db
}

type Post struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	UserID uint
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {

	db := connect()

	db.AutoMigrate(&User{}, &Post{})

	//creating
	createdUser1 := User{Name: "Naushad1", Email: "naushad1@example.com"}
	db.Create(&createdUser1)
	fmt.Println(createdUser1.ID)

	createdUser2 := User{Name: "Naushad2", Email: "naushad2@example.com"}
	db.Create(&createdUser2)
	fmt.Println(createdUser2.ID)

	//fetching one
	var fetchUser User
	db.First(&fetchUser, 1)

	//updating
	db.Model(&createdUser1).Updates(User{Name: "Naushad Ansari"})
	db.Model(&createdUser2).Update("Email", "Naushadansari@example.com")

	//fetch all
	var fetchUsersbefore []User
	db.Find(&fetchUsersbefore)
	fmt.Println("fetchUsersbefore", fetchUsersbefore)

	//delete one
	db.Delete(fetchUsersbefore[0])

	//fetch all after delete
	var fetchUsersAfter []User
	db.Find(&fetchUsersAfter)
	fmt.Println("fetchUsersAfter", fetchUsersAfter)

	//post creating
	postUser1 := User{Name: "User1", Email: "user1@example.com"}
	db.Create(&postUser1)
	postUser2 := User{Name: "User2", Email: "user2@example.com"}
	db.Create(&postUser2)

	//for postUser1
	post1 := Post{
		Title:  "post1 Go",
		UserID: postUser1.ID,
	}
	db.Create(&post1)
	post2 := Post{
		Title:  "post2 Python",
		UserID: postUser1.ID,
	}
	db.Create(&post2)

	//for postUser2
	post3 := Post{
		Title:  "Hello",
		UserID: postUser2.ID,
	}
	db.Create(&post3)

	//preload
	var preloadUser User
	db.Preload("Posts").First(&preloadUser, postUser1.ID)
	fmt.Println("Preload:", preloadUser)

	//conditional preload
	var conditionalPreloadUser User
	db.Preload("Posts", "title LIKE ?", "%Go%").First(&conditionalPreloadUser, postUser1.ID)
	fmt.Println("conditional Preload:", conditionalPreloadUser)

	//column preload
	var columnPreloadUser User
	db.Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "user_id")
	}).First(&columnPreloadUser, postUser1.ID)
	fmt.Println("column Preload:", columnPreloadUser)

	//join
	var joinUsers []User
	db.Joins("JOIN posts ON posts.user_id = users.id").Preload("Posts").Find(&joinUsers)
	fmt.Println("join:", joinUsers)
}
