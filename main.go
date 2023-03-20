package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Characters struct {
	Fullname string
	Family   string
	Title    string
	Image    string
	ImageUrl string
}

type Continents struct {
	Name string
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Utilise un template HTML pour afficher la page d'accueil
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Charac(w http.ResponseWriter, r *http.Request) {
	// Récupère les données depuis une API
	resp, err := http.Get("https://thronesapi.com/api/v2/Characters")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Convertit les données JSON en structures de données Go
	var chars []Characters
	err = json.NewDecoder(resp.Body).Decode(&chars)
	if err != nil {
		log.Fatal(err)
	}

	// Utilise un template HTML pour afficher les données
	tmpl := template.Must(template.ParseFiles("chara.html"))
	err = tmpl.Execute(w, chars)
	if err != nil {
		log.Fatal(err)
	}
}

func Conti(w http.ResponseWriter, r *http.Request) {
	// Récupère les données depuis une API
	resp, err := http.Get("https://thronesapi.com/api/v2/Continents")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Convertit les données JSON en structures de données Go
	var chars []Continents
	err = json.NewDecoder(resp.Body).Decode(&chars)
	if err != nil {
		log.Fatal(err)
	}

	// Utilise un template HTML pour afficher les données
	tmpl := template.Must(template.ParseFiles("conti.html"))
	err = tmpl.Execute(w, chars)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/Chara/", Charac)
	http.HandleFunc("/Conti/", Conti)

	fmt.Println("(http://localhost:8080/) - Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
