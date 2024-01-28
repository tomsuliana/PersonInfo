package repository

import (
	"server/server/internal/domain/dto"
)

type PersonRepositoryI interface {
	GetPersons() ([]*dto.DBGetPerson, error)
	GetPersonById(id uint) (*dto.DBGetPerson, error)
	GetPersonsByAge(age uint) ([]*dto.DBGetPerson, error)
	GetPersonsByGender(gender string) ([]*dto.DBGetPerson, error)
	GetPersonsByNation(nation string) ([]*dto.DBGetPerson, error)
	GetPersonsWithLimit(limit uint) ([]*dto.DBGetPerson, error)
	DeletePerson(id uint) error
	UpdatePerson(person *dto.DBGetPerson) error
	CreatePerson(person *dto.DBGetPerson) (uint, error)
}
