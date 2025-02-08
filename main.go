package main

import (
	admin_models "goblog/admin/models"
	"goblog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	/*admin_models.Post{
		Title:       "Go ile Web Programlama",
		Slug:        "go-ile-web-programlama",
		Description: "Descriptionne",
	}.Add()
	/*post := admin_models.Post{}.Get(1)
	fmt.Println(post.Title)*/
	//fmt.Println(admin_models.Post{}.GetAll())
	//post := admin_models.Post{}.Get(1)
	//post.Updates(admin_models.Post{Title: "Python ile web", Description: "test"})

	http.ListenAndServe(":8080", config.Routes())

}
