package usecase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	personRep "server/server/internal/Person/repository"
	"server/server/internal/domain/dto"
)

type PersonUsecaseI interface {
	GetPersons() ([]*dto.Person, error)
	GetPersonsByAge(age uint) ([]*dto.Person, error)
	GetPersonsByGender(gender string) ([]*dto.Person, error)
	GetPersonsByNation(nation string) ([]*dto.Person, error)
	GetPersonsWithLimit(limit uint) ([]*dto.Person, error)
	DeletePerson(id uint) error
	UpdatePerson(newPerson *dto.Person) error
	CreatePerson(newPerson *dto.Person) (uint, error)
}

type PersonUsecase struct {
	personRepo personRep.PersonRepositoryI
}

func NewPersonUsecase(personRepI personRep.PersonRepositoryI) *PersonUsecase {
	return &PersonUsecase{
		personRepo: personRepI,
	}
}

func (per PersonUsecase) GetPersons() ([]*dto.Person, error) {
	dbpers, err := per.personRepo.GetPersons()
	if err != nil {
		return nil, err
	}
	persons := []*dto.Person{}
	for _, dbper := range dbpers {
		person := dto.ToPerson(dbper)
		persons = append(persons, person)
	}

	return persons, nil
}

func (per PersonUsecase) GetPersonsByAge(age uint) ([]*dto.Person, error) {
	dbpers, err := per.personRepo.GetPersonsByAge(age)
	if err != nil {
		return nil, err
	}
	persons := []*dto.Person{}
	for _, dbper := range dbpers {
		person := dto.ToPerson(dbper)
		persons = append(persons, person)
	}

	return persons, nil
}

func (per PersonUsecase) GetPersonsByGender(gender string) ([]*dto.Person, error) {
	dbpers, err := per.personRepo.GetPersonsByGender(gender)
	if err != nil {
		return nil, err
	}
	persons := []*dto.Person{}
	for _, dbper := range dbpers {
		person := dto.ToPerson(dbper)
		persons = append(persons, person)
	}

	return persons, nil
}

func (per PersonUsecase) GetPersonsByNation(nation string) ([]*dto.Person, error) {
	dbpers, err := per.personRepo.GetPersonsByNation(nation)
	if err != nil {
		return nil, err
	}
	persons := []*dto.Person{}
	for _, dbper := range dbpers {
		person := dto.ToPerson(dbper)
		persons = append(persons, person)
	}

	return persons, nil
}

func (per PersonUsecase) GetPersonsWithLimit(limit uint) ([]*dto.Person, error) {
	dbpers, err := per.personRepo.GetPersonsWithLimit(limit)
	if err != nil {
		return nil, err
	}
	persons := []*dto.Person{}
	for _, dbper := range dbpers {
		person := dto.ToPerson(dbper)
		persons = append(persons, person)
	}

	return persons, nil
}

func (per PersonUsecase) DeletePerson(id uint) error {
	err := per.personRepo.DeletePerson(id)
	if err != nil {
		return err
	}
	return nil
}

func (per PersonUsecase) UpdatePerson(newPerson *dto.Person) error {
	pers, err := per.personRepo.GetPersonById(newPerson.ID)
	if err != nil {
		return err
	}

	if pers != nil {
		person := dto.ToPerson(pers)
		if newPerson.Name != "" {
			person.Name = newPerson.Name
		}

		if newPerson.Surname != "" {
			person.Surname = newPerson.Surname
		}

		if newPerson.Patronymic != "" {
			person.Patronymic = newPerson.Patronymic
		}

		if newPerson.Age != 0 {
			person.Age = newPerson.Age
		}

		if newPerson.Gender != "" {
			person.Gender = newPerson.Gender
		}

		if newPerson.Nation != "" {
			person.Nation = newPerson.Nation
		}

		return per.personRepo.UpdatePerson(dto.ToDBGetPerson(person))
	}

	return dto.ErrNotFound
}

func (per PersonUsecase) CreatePerson(newPerson *dto.Person) (uint, error) {
	person := dto.ToDBGetPerson(newPerson)

	ageresp, err := http.Get("https://api.agify.io/?name=" + person.Name)
	if err != nil {
		return 0, err
	}

	agebody, err := ioutil.ReadAll(ageresp.Body)
	if err != nil {
		return 0, err
	}

	ageObject := &dto.Age{}
	err = json.Unmarshal(agebody, &ageObject)

	person.Age = ageObject.Age

	genderresp, err := http.Get("https://api.genderize.io/?name=" + person.Name)
	if err != nil {
		return 0, err
	}

	genderbody, err := ioutil.ReadAll(genderresp.Body)
	if err != nil {
		return 0, err
	}

	genderObject := &dto.Gender{}
	err = json.Unmarshal(genderbody, &genderObject)

	person.Gender = genderObject.Gender

	nationresp, err := http.Get("https://api.nationalize.io/?name=" + person.Name)
	if err != nil {
		return 0, err
	}

	nationbody, err := ioutil.ReadAll(nationresp.Body)
	if err != nil {
		return 0, err
	}

	nationObject := &dto.Nation{}
	err = json.Unmarshal(nationbody, &nationObject)

	person.Nation = nationObject.Nation[0].CountryId

	personid, err := per.personRepo.CreatePerson(person)
	if err != nil {
		return 0, err
	}

	return personid, nil

}
