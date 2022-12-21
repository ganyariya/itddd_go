package main

import (
	"fmt"

	"github.com/ganyariya/itddd_go/pkg/application/command"
)

func main() {
	fmt.Print("Input your name: ")

	var name string
	fmt.Scan(&name)

	userAppService := Initialize()
	userDto, err := userAppService.Register(command.NewUserRegisterCommand(name))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("登録処理が成功しました")
	fmt.Printf("%+v", userDto)
}
