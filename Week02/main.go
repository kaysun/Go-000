package main

import "github.com/kaysun/Go-000/Week02/service"

func main() {
	userService := service.UserService{}
	userService.DemandUser(1)
}