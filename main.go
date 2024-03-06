package main

import "crud-server/controllers"

func main() {
	controllers.NewController(":8080")
}