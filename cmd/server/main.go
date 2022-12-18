package main

import (
	"fmt"
	"net/http"

	"github.com/ganyariya/itddd_go/pkg/application/command"
	"github.com/ganyariya/itddd_go/pkg/application/service"
)

var userAppService *service.UserApplicationService

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")

	user, err := userAppService.Register(name)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
	} else {
		fmt.Fprintf(w, "%s registered: %+v", name, user)
	}
}
func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	userDtos, _ := userAppService.GetAll()
	for _, user := range userDtos {
		fmt.Fprintf(w, "user: %+v", user)
	}
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	uid := values.Get("uid")

	user, err := userAppService.Get(uid)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
	} else {
		fmt.Fprintf(w, "%s : %+v", uid, user)
	}
}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	uid := values.Get("uid")

	err := userAppService.Delete(command.NewUserDeleteCommand(uid))
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
	} else {
		fmt.Fprintf(w, "%s deleted", uid)
	}

}

func main() {
	userAppService = Initialize()

	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/getall", GetAllHandler)
	http.HandleFunc("/delete", DeleteHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
