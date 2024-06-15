package usecase

import (
	"errors"
	"strconv"

	_entities "github.com/thanapatjitmung/entities"
	_repository "github.com/thanapatjitmung/repository"
)

type (
	ClientUsecase interface {
		GetAllDataClient() ([]*_entities.UserRes, error)
		GetProfile(int) (*_entities.UserRes, error)
		UpdateProfile(int, *_entities.UserRes) error
	}
	clientUsecaseImpl struct {
		repo _repository.AuthRepo
	}
)

func NewClientUsecaseImpl(repo _repository.AuthRepo) ClientUsecase {
	return &clientUsecaseImpl{
		repo: repo,
	}
}

func (c *clientUsecaseImpl) GetAllDataClient() ([]*_entities.UserRes, error) {
	var dataResponse []*_entities.UserRes
	data, err := c.repo.GetAllClients()
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

func (c *clientUsecaseImpl) GetProfile(id int) (*_entities.UserRes, error) {
	data, err := c.repo.GetByIdForUser(id)
	if err != nil {
		return nil, err
	}
	var dataResponse = &_entities.UserRes{
		ID:       data.ID,
		Username: data.Username,
		Role:     data.Role,
	}
	return dataResponse, nil
}

func (c *clientUsecaseImpl) UpdateProfile(id int, user *_entities.UserRes) error {
	check := false
	allData, err := c.repo.GetAllData()
	if err != nil {
		return err
	}
	for _, checkUsername := range allData {
		if user.Username == checkUsername.Username {
			return errors.New("username already exists")
		}
	}
	data, err := c.repo.GetAllClients()
	if err != nil {
		return err
	}
	for i, d := range data {
		if id == d.ID {
			data[i].Username = user.Username
			check = true
			break
		}
	}
	if check {
		var clientData [][]string
		for _, d := range data {
			record := []string{strconv.Itoa(d.ID), d.Username, d.Password, d.Role}
			clientData = append(clientData, record)
		}
		err = c.repo.UpdateClient(clientData)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("user not found")
}
