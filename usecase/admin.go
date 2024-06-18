package usecase

import (
	"errors"
	"strconv"

	_entities "github.com/thanapatjitmung/entities"
	_repository "github.com/thanapatjitmung/repository"
)

type (
	AdminUsecase interface {
		GetAllData() ([]*_entities.UserRes, error)
		GetByIdForAdmin(int) (*_entities.UserRes, error)
		UpdateUserForAdmin(int, *_entities.UserRes) error
		DeleteUserForAdmin(int) error
	}

	adminUsecaseImpl struct {
		userRepo _repository.AuthRepo
	}
)

func NewAdminUsecaseImpl(userRepo _repository.AuthRepo) AdminUsecase {
	return &adminUsecaseImpl{userRepo: userRepo}
}

func (a *adminUsecaseImpl) GetAllData() ([]*_entities.UserRes, error) {
	var dataResponse []*_entities.UserRes
	data, err := a.userRepo.GetAllData()
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		userResponse := &_entities.UserRes{
			ID:       d.ID,
			Username: d.Username,
			Role:     d.Role,
		}
		dataResponse = append(dataResponse, userResponse)
	}

	return dataResponse, nil
}

func (a *adminUsecaseImpl) GetByIdForAdmin(id int) (*_entities.UserRes, error) {
	data, err := a.userRepo.GetByIdForAdmin(id)
	if err != nil {
		return nil, err
	}
	var dataResponse = &_entities.UserRes{
		ID:       data.ID,
		Username: data.Username,
		Role:     data.Role,
	}
	// fmt.Println("============== GetById Admin UseCase ==============")
	// fmt.Println(dataResponse)
	return dataResponse, nil
}

func (a *adminUsecaseImpl) UpdateUserForAdmin(id int, user *_entities.UserRes) error {
	check := false
	data, err := a.userRepo.GetAllData()
	if err != nil {
		return err
	}
	for _, checkUsername := range data {
		if user.Username == checkUsername.Username && id != checkUsername.ID {
			return errors.New("username already exists")
		}
	}
	for i, d := range data {
		if id == d.ID {
			data[i].Username = user.Username
			data[i].Role = user.Role
			check = true
			break
		}
	}
	if check {
		var clientData, adminData [][]string
		for _, d := range data {
			record := []string{strconv.Itoa(d.ID), d.Username, d.Password, d.Role}
			if d.Role == "client" {
				clientData = append(clientData, record)
			} else if d.Role == "admin" {
				adminData = append(adminData, record)
			}
		}
		err = a.userRepo.UpdateClient(clientData)
		if err != nil {
			return err
		}
		err = a.userRepo.UpdateAdmin(adminData)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("user not found")
}

func (a *adminUsecaseImpl) DeleteUserForAdmin(id int) error {
	var updatedData []*_entities.User
	check := false
	data, err := a.userRepo.GetAllData()
	if err != nil {
		return err
	}
	for _, d := range data {
		if id != d.ID {
			updatedData = append(updatedData, d)
		} else {
			check = true
		}
	}
	if check {
		var clientData, adminData [][]string

		for _, d := range updatedData {
			record := []string{strconv.Itoa(d.ID), d.Username, d.Password, d.Role}
			if d.Role == "client" {
				clientData = append(clientData, record)
			} else if d.Role == "admin" {
				adminData = append(adminData, record)
			}
		}
		err = a.userRepo.UpdateClient(clientData)
		if err != nil {
			return err
		}
		err = a.userRepo.UpdateAdmin(adminData)
		if err != nil {
			return err
		}
		return nil

	}

	return errors.New("user not found")
}
