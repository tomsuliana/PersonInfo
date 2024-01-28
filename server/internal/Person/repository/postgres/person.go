package repository

import (
	"database/sql"
	//"server/internal/domain/dto"
	"server/server/internal/domain/dto"
)

//PersonRepo struct
type PersonRepo struct {
	DB *sql.DB
}

//NewPersonRepo creates new object of Person repo
func NewPersonRepo(db *sql.DB) *PersonRepo {
	return &PersonRepo{
		DB: db,
	}
}

//GetPersons gets info about people
func (repo *PersonRepo) GetPersons() ([]*dto.DBGetPerson, error) {
	rows, err := repo.DB.Query(`SELECT id, name, surname, patronymic, age, gender, nation 
								FROM person`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var Persons = []*dto.DBGetPerson{}
	for rows.Next() {
		person := &dto.DBGetPerson{}
		err = rows.Scan(
			&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nation,
		)
		if err != nil {
			return nil, err
		}
		Persons = append(Persons, person)
	}
	return Persons, nil
}

func (repo *PersonRepo) GetPersonsByAge(age uint) ([]*dto.DBGetPerson, error) {
	rows, err := repo.DB.Query(`SELECT id, name, surname, patronymic, age, gender, nation FROM person WHERE age = $1`, age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var Persons = []*dto.DBGetPerson{}
	for rows.Next() {
		person := &dto.DBGetPerson{}
		err = rows.Scan(
			&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nation,
		)
		if err != nil {
			return nil, err
		}
		Persons = append(Persons, person)
	}
	return Persons, nil
}

func (repo *PersonRepo) GetPersonsByGender(gender string) ([]*dto.DBGetPerson, error) {
	rows, err := repo.DB.Query(`SELECT id, name, surname, patronymic, age, gender, nation FROM person WHERE gender = $1`, gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var Persons = []*dto.DBGetPerson{}
	for rows.Next() {
		person := &dto.DBGetPerson{}
		err = rows.Scan(
			&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nation,
		)
		if err != nil {
			return nil, err
		}
		Persons = append(Persons, person)
	}
	return Persons, nil
}

func (repo *PersonRepo) GetPersonsByNation(nation string) ([]*dto.DBGetPerson, error) {
	rows, err := repo.DB.Query(`SELECT id, name, surname, patronymic, age, gender, nation FROM person WHERE nation = $1`, nation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var Persons = []*dto.DBGetPerson{}
	for rows.Next() {
		person := &dto.DBGetPerson{}
		err = rows.Scan(
			&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nation,
		)
		if err != nil {
			return nil, err
		}
		Persons = append(Persons, person)
	}
	return Persons, nil
}

func (repo *PersonRepo) GetPersonsWithLimit(limit uint) ([]*dto.DBGetPerson, error) {
	rows, err := repo.DB.Query(`SELECT id, name, surname, patronymic, age, gender, nation FROM person LIMIT $1`, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var Persons = []*dto.DBGetPerson{}
	for rows.Next() {
		person := &dto.DBGetPerson{}
		err = rows.Scan(
			&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nation,
		)
		if err != nil {
			return nil, err
		}
		Persons = append(Persons, person)
	}
	return Persons, nil
}

func (repo *PersonRepo) DeletePerson(id uint) error {
	deletePerson := `DELETE FROM person WHERE id = $1`
	_, err := repo.DB.Exec(deletePerson, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PersonRepo) UpdatePerson(person *dto.DBGetPerson) error {
	updatePerson := `UPDATE person 
				   SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nation = $6
				   WHERE id = $7`
	_, err := repo.DB.Exec(updatePerson, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nation, person.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PersonRepo) GetPersonById(id uint) (*dto.DBGetPerson, error) {
	person := &dto.DBGetPerson{}
	row := repo.DB.QueryRow(`SELECT id, name, surname, patronymic, age, gender, nation FROM person WHERE id = $1`, id)
	err := row.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return person, nil
}

func (repo *PersonRepo) CreatePerson(person *dto.DBGetPerson) (uint, error) {
	insertPerson := `INSERT INTO person (name, surname, patronymic, age, gender, nation) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var ID uint
	err := repo.DB.QueryRow(insertPerson, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nation).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, nil
}
