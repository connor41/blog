package routes

import (
	"fmt"
	"net/http"

	"github.com/martini-contrib/render"

	"../db/users"
	"../utils"
)

func GetRegisterHandler(rnd render.Render) {
	rnd.HTML(200, "register", nil)
}

func PostRegisterHandler(rnd render.Render, r *http.Request) {
	id := ""
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	perm := ""

	fmt.Println(email)
	fmt.Println(username)
	fmt.Println(password)

	userTable := users.UsersTable{id, email, username, password, perm}
	fmt.Println("createuser")
	id = utils.GenerateNameId(username)
	perm = "user"
	userTable.Id = id
	userTable.Permission = perm
	err := usersTables.Insert(userTable)
	if err != nil {
		fmt.Println("такой юзер есть")
	}

	rnd.Redirect("/")
}