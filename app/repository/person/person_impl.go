package person

import (
	"fmt"
	"swag-tutor/app/adapter/database"
	model "swag-tutor/app/models/sql"
)

type personrepository struct {
	db *database.MySQL
}

func NewPersonRepository(db *database.MySQL) PersonRepository {
	return &personrepository{
		db: db,
	}
}


func (pr personrepository) GetAllPersons() (result *[]model.Person, err error) {
	err = pr.db.DB.Find(&result).Error
	if len(*result) == 0 {
		err = fmt.Errorf("%s", "Data Not Found")
		return result, err
	}
	if err != nil {
		return result, err
	}
	return result, nil
}

func (pr personrepository) GetOnePerson(id int) (result *model.Person, err error) {
	err = pr.db.DB.First(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (pr personrepository) CreatePerson(person *model.Person) (result *model.Person, err error) {
	err = pr.db.DB.Create(person).Error
	if err != nil {
		fmt.Println(err.Error(), "ini error di create")
		return result, err
	}
	fmt.Println(person)
	result = person
	return result, nil
}

func (pr personrepository) UpdatePerson(person *model.Person, id int) (result string, err error) {

	db := pr.db.DB.Where("id = ?", id).Updates(person)
	if db.Error != nil {
		return "", err
	}
	if db.RowsAffected < 1 {
		return "", fmt.Errorf("%s", "Data not Found")
	}

	return "Sukses to Update", nil
}

func (pr personrepository) DeletePerson(id int) (result string, err error) {
	db := pr.db.DB.Delete(&model.Person{}, id)
	if db.Error != nil {
		return "", err
	}
	if db.RowsAffected < 1 {
		return "", fmt.Errorf("%s", "Data not Found")
	}
	return "Success to Delete", nil
}