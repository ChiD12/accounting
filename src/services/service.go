package services

import (
	"example.com/accounting/src/db"
	"example.com/accounting/src/routes/validators"
	"example.com/accounting/src/services/models"
)

type Service struct {
	Database db.DB
}

func MakeService(db db.DB) Service {
	return Service{db}
}

func (this Service) GetUser(email string) (models.User, error) {
	userDao, err := this.Database.GetUser(email)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
			Id:    userDao.Id,
			Email: userDao.Email,
			User:  userDao.User,
			Pass:  userDao.Pass},
		err
}

func (this Service) CreateUser(input validators.CreateUser) (models.User, error) {
	createUser := db.CreateUser{input.Email, input.User, input.Pass}
	userDao, err := this.Database.CreateUser(createUser)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
			Id:    userDao.Id,
			Email: userDao.Email,
			User:  userDao.User,
			Pass:  userDao.Pass},
		err
}
