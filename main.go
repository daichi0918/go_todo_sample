package main

import (
	"fmt"
	"go_sample_todo/app/controllers"
	"go_sample_todo/app/models"
)

func main() {


	fmt.Println(models.Db)

	controllers.StartMainServer()
}
