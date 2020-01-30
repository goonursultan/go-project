package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *gorm.DB  {
	db, err := gorm.Open("mysql", "root:@(localhost)/kaspilab?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		panic(err.Error())
	}
	return db
}

type Pagination struct {
	Limit int64 `json:"limit"`
	Page int64 `json:"page"`
	Data interface{} `json:"data"`
}

func validatePagination(r *http.Request) (bool,string,int64,int64) {
	type PathVariables struct {
		Limit int64 `validate:"required,min=1"`
		Page int64 `validate:"required,min=1"`
	}
	limit, _ :=strconv.ParseInt(r.FormValue("limit"),10, 64)
	page,_:=strconv.ParseInt(r.FormValue("page"),10,64)

	pathVars :=PathVariables{Limit:limit,Page:page}
	validate := validator.New()
	err := validate.Struct(pathVars)
	if err != nil {
		return false, err.Error(),limit,page
	}
	return true,"",limit,page
}

func users(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("connect -> users method")
	valid,message,limit,page :=validatePagination(r)
	
	if valid {
		type Direction struct {
			Id int64 `gorm:"primary_key",json:"id"`
			Direction string `json:"direction"`
		}
		type Result struct {
			Id string `gorm:"primary_key",json:"id"`
			Fio  string `json:"fio"`
			University  string `json:"university"`
			Course  string `json:"course"`
			MobilePhone  string `json:"mobile_phone"`
			Email  string `json:"email"`
			CreatedAt  *time.Time `json:"created_at"`
			UpdatedAt  *time.Time `json:"updated_at"`
			Directions [] Direction `gorm:"many2many:user_direction",json:"directions"`
		}

		var results []Result
		//var directions []Direction
		DB().
			Table("users as u").
			Select("u.*").
			Scan(&results)
		DB().Model(&results)
		fmt.Println()

		data:=Pagination{Limit:limit,Page:page,Data:results}
		respondJSON(w, http.StatusOK, data)
	}else {
		respondError(w,500,message)
	}
}

//роутинг + точка входа
func main()  {
	router:=mux.NewRouter()
	router.HandleFunc("/users",users).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}


func respondJSON(w http.ResponseWriter, status int, res interface{}) {
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}