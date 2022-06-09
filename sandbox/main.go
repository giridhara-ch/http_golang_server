package main

import (
	"fmt"
	"giridhara-ch/http_golang_server/internal/database"
	"log"
	"path/filepath"
)

func main() {
	dbPath, err := filepath.Abs("../http_golang_server/internal/database/db.json")
	if err != nil {
		log.Fatal(err)
	}
	c := database.NewClient(dbPath)
	err = c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB ensured")
	user, err := c.CreateUser("ravikiran@domain.com", "password", "Ravi Kiran", 30)
	if err != nil {
		// log.Panic(err)
		fmt.Println(err)
	}
	// fmt.Println("user created", user)
	fmt.Println("user created", user.Name, user.Email)

	updatedUser, err := c.UpdateUser("ravikiran@domain.com", "password", "Ravi Kiran", 29)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("user updated", updatedUser)
	fmt.Println("user updated", updatedUser.Name, updatedUser.Email)

	getUser, err := c.GetUser("ravikiran@domain.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("got user", getUser)
	// fmt.Println(c.GetUser("s@domain.com"))

	_, err = c.GetUser("ravikiran@domain.com")
	if err != nil {
		log.Fatal("shouldn't be able to get deleted user")
	}
	fmt.Println("Deletion confirmed")

	post, err := c.CreatePost("ravikiran@domain.com", "simple post text")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created post", post)

	post2, err := c.CreatePost("ravikiran@domain.com", "simple post text for 2nd time")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created post", post2)

	posts, err := c.GetPosts("ravikiran@domain.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeletePost(post.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted first post")

	posts, err = c.GetPosts("ravikiran@domain.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeletePost(post2.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted second post")

	posts, err = c.GetPosts("ravikiran@domain.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeleteUser("ravikiran@domain.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted User")

}
