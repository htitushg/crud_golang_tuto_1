package config

import (
	"text/template"
	"time"
)

type User struct {
	Id        string
	Nom       string
	Prenom    string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

var Arr User

type ListeUtilisateurs struct {
	Nombre int
	Name   string
	Users  []User
}
type Utilisateur struct {
	Name string
	User User
}

type Config struct {
	TemplateCache map[string]*template.Template
	Port          string
}

var AppConfig Config

type Car struct {
	Id           int       `json:"id"`
	Manufacturer string    `json:"manufacturer"`
	Design       string    `json:"design"`
	Style        string    `json:"style"`
	Doors        uint8     `json:"doors"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Cars []Car

/*
Pour les besoins d'inscription dans une base de données MySql:

Conversion de date : string vers []Byte
DatedeCreation = []byte(user.CreatedAt)
DatedeMaj = []byte(time.Now().Format("2006-01-02"))

Conversion de date: []Byte vers string
Arr.UpdatedAt = string(DatedeMaj)
Arr.CreatedAt = string(DatedeCreation)
*/
