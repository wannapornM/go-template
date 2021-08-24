package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Menu struct {
	Id          int
	Name        string
	Description string
	Price       int
}

var menus = []Menu{
	{
		Id:          1,
		Name:        "BB Corn",
		Description: "Giant breed of rare corn that was eaten by Gourmet Nobility as a snack long ago",
		Price:       4000,
	},
	{
		Id:          2,
		Name:        "Century Soup",
		Description: "A soup cooked with hundreds or even thousands of ingredients",
		Price:       10000,
	},
	{
		Id:          3,
		Name:        "Jewel Meat",
		Description: "Incandescent lamp-like radiance that dulls jewels and lights up a night sky",
		Price:       8000,
	},
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/404.html")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/hello.html")
}

func AllMenusHandler(w http.ResponseWriter, r *http.Request) {
	menuTemplate, err := template.ParseFiles("views/menus/index.html")
	if err != nil {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	menuTemplate.Execute(w, menus)
}

func MenusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuIndex, _ := strconv.ParseInt(params["id"], 0, 64)
	menuIndex -= 1

	menuTemplate, err := template.ParseFiles("views/menus/show.html")
	if err != nil || isOutOfRange(menuIndex) {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	menuTemplate.Execute(w, menus[menuIndex])
}

func isOutOfRange(index int64) bool {
	return (index < 0 || index >= int64(len(menus)))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/home", HomeHandler)
	router.HandleFunc("/menus", AllMenusHandler)
	router.HandleFunc("/menus/{id:[0-9]+}", MenusHandler)
	// Please help implement one of following routing
	// router.HandleFunc("/products", AllProductsHandler)
	// router.HandleFunc("/products/{id:[0-9]+}", ProductsHandler)
	// or
	// router.HandleFunc("/books", AllBooksHandler)
	// router.HandleFunc("/books/{id:[0-9]+}", BooksHandler)
	// or
	// router.HandleFunc("/songs", AllSongsHandler)
	// router.HandleFunc("/songs/{id:[0-9]+}", SongsHandler)
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	http.Handle("/", router)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.ListenAndServe(":8000", nil)
}
