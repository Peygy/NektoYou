package routes

import (
	"html/template"
	"log"
	"net/http"

	"web/pkg/data"
	"web/pkg/models"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func StartPage(w http.ResponseWriter, r *http.Request){
	users := []models.User{}

	db,err := data.OpenDb()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from " + viper.GetString("data.table"))
	if err != nil{
		log.Println(err)
	}

	for rows.Next(){
		user := models.User{}
		err = rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil{
			log.Println(err)
			continue
		}
		users = append(users, user)
	}

	tmpl,_ := template.ParseFiles("templates/main.html")
	tmpl.Execute(w, users)
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		login := r.FormValue("userlogin")
		password := r.FormValue("userpassword")

		db,err := data.OpenDb()
		if err != nil{
			log.Fatal(err)
		}
		defer db.Close()

		_,err = db.Exec("insert into " + viper.GetString("data.table") + " (login, password) values (?,?)", login, password)
		
		if err != nil{
			log.Fatal(err)
		}

		http.Redirect(w,r,"/", 301)
	}else{
		http.ServeFile(w,r,"templates/add.html")
	}
}


func Edit_Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := models.User{}

	db,err := data.OpenDb()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("select * from " + viper.GetString("data.table") + " where id = ?", id)
	err = row.Scan(&user.Id, &user.Login, &user.Password)

	if err != nil{
		log.Println(err)
		http.Error(w,http.StatusText(404),http.StatusNotFound)		
	}else{
		tmpl, _ := template.ParseFiles("templates/edit.html")
		tmpl.Execute(w, user)
	}
}


func Edit_Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	login := r.FormValue("userlogin")
	password := r.FormValue("userpassword")

	db,err := data.OpenDb()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	_,err = db.Exec("update " + viper.GetString("data.table") + " set login = ?, password = ? where id = ?", login, password, id)
		
	if err != nil{
		log.Fatal(err)
	}

	http.Redirect(w,r,"/", 301)
}


func RemoveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db,err := data.OpenDb()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	_,err = db.Exec("delete from " + viper.GetString("data.table") + " where id=?", id)

	if err != nil{
		log.Fatal(err)
	}

	http.Redirect(w,r,"/", 301)
}
