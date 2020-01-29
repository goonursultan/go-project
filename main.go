package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"time"
)

func users(w http.ResponseWriter,r *http.Request)  {
	type User struct {
		Id int64 `json:"id"`
		Fio string `json:"fio"`
		University string `json:"university"`
		Course string `json:"course"`
		MobilePhone string `json:"mobile_phone"`
		Email string `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	db, err := gorm.Open("mysql", "root:@(localhost)/kaspilab?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		panic(err.Error())
	}
	defer db.Close()

	var users [] User

	db.Table("users").Select("*").Limit(2).Scan(&users)
	_ = json.NewEncoder(w).Encode(&users)
}
func main()  {
	router:=mux.NewRouter()
	router.HandleFunc("/users",users).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}