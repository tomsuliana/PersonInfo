package dto

import (
	"database/sql"
)

type DBGetPerson struct {
	ID         uint
	Name       string
	Surname    string
	Patronymic sql.NullString
	Age        uint
	Gender     string
	Nation     string
}

type Person struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        uint   `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}

type Age struct {
	Age uint `json:"age"`
}

type Gender struct {
	Gender string `json:"gender"`
}

type CountryId struct {
	CountryId string `json:"country_id"`
}

type Nation struct {
	Nation []*CountryId `json:"country"`
}

type RespID struct {
	ID uint `json:"id"`
}

func ToPerson(person *DBGetPerson) *Person {
	return &Person{
		ID:         person.ID,
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: transformSQLStringToString(person.Patronymic),
		Age:        person.Age,
		Gender:     person.Gender,
		Nation:     person.Nation,
	}
}

func ToDBGetPerson(person *Person) *DBGetPerson {
	return &DBGetPerson{
		ID:         person.ID,
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: *transformStringToSQLString(person.Patronymic),
		Age:        person.Age,
		Gender:     person.Gender,
		Nation:     person.Nation,
	}
}

func transformStringToSQLString(str string) *sql.NullString {
	if str != "" {
		return &sql.NullString{String: str, Valid: true}
	}
	return &sql.NullString{Valid: false}
}

func transformSQLStringToString(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return ""
}
