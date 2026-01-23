package main

import (
	"fmt"
	"go_sample_todo/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.PassWord = "password"
	fmt.Println(u);

	u.CreateUser()



	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

		// user, _ := models.GetUser(2)
	// user.CreateTodo("First todo")
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	// fmt.Println(todos)

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

		t, _ := models.GetTodo(3)
		t.DeleteTodo()
	// 	t.Content = "Updated todo content"
	// fmt.Println(t)
}
