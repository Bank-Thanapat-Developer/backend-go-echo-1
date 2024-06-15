package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	_entities "github.com/thanapatjitmung/entities"
)

type (
	AuthRepo interface {
		Save(user *_entities.User) error
		GetAllClients() ([]*_entities.User, error)
		GetAllAdmins() ([]*_entities.User, error)
		GetAllData() ([]*_entities.User, error)
		GetByIdForAdmin(int) (*_entities.User, error)
		GetByIdForUser(id int) (*_entities.User, error)
		UpdateAdmin([][]string) error
		UpdateClient([][]string) error
		DeleteAdmin([][]string) error
		DeleteClient([][]string) error
	}

	userRepoImpl struct {
		clientFilePath string
		adminFilePath  string
	}
)

func NewUserRepoImpl(clientFilePath, adminFilePath string) AuthRepo {
	_, _ = os.OpenFile(clientFilePath, os.O_RDWR|os.O_CREATE, 0644)
	_, _ = os.OpenFile(adminFilePath, os.O_RDWR|os.O_CREATE, 0644)
	return &userRepoImpl{
		clientFilePath: clientFilePath,
		adminFilePath:  adminFilePath,
	}
}

func (a *userRepoImpl) Save(user *_entities.User) error {
	var filePath string
	if user.Role == "client" {
		filePath = a.clientFilePath
	} else if user.Role == "admin" {
		filePath = a.adminFilePath
	} else {
		return errors.New("invalid user role")
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{strconv.Itoa(user.ID), user.Username, user.Password, user.Role}
	err = writer.Write(record)
	if err != nil {
		return err
	}

	return nil
}

func (a *userRepoImpl) GetAllClients() ([]*_entities.User, error) {
	clientFile, err := os.Open("client.csv")
	if err != nil {
		return nil, err
	}
	defer clientFile.Close()
	var data []*_entities.User

	reader := csv.NewReader(clientFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		d := &_entities.User{
			ID:       id,
			Username: record[1],
			Password: record[2],
			Role:     record[3],
		}
		data = append(data, d)
	}

	return data, nil
}

func (a *userRepoImpl) GetAllAdmins() ([]*_entities.User, error) {
	adminFile, err := os.Open("admin.csv")
	if err != nil {
		return nil, err
	}
	defer adminFile.Close()
	var data []*_entities.User

	reader := csv.NewReader(adminFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		d := &_entities.User{
			ID:       id,
			Username: record[1],
			Password: record[2],
			Role:     record[3],
		}
		data = append(data, d)
	}
	return data, nil
}

func (a *userRepoImpl) GetAllData() ([]*_entities.User, error) {
	var data []*_entities.User

	dataClients, err := a.GetAllClients()
	if err != nil {
		return nil, err
	}

	data = append(data, dataClients...)

	dataAdmin, err := a.GetAllAdmins()
	if err != nil {
		return nil, err
	}

	data = append(data, dataAdmin...)
	fmt.Println(len(data))

	return data, nil
}

func (a *userRepoImpl) GetByIdForAdmin(id int) (*_entities.User, error) {
	data, err := a.GetAllData()
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		if id == d.ID {
			user := &_entities.User{
				ID:       d.ID,
				Username: d.Username,
				Password: d.Password,
				Role:     d.Role,
			}
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (a *userRepoImpl) GetByIdForUser(id int) (*_entities.User, error) {
	data, err := a.GetAllClients()
	if err != nil {
		return nil, err
	}
	fmt.Println(id)

	for _, d := range data {
		fmt.Println(d.ID)
		if id == d.ID {
			user := &_entities.User{
				ID:       d.ID,
				Username: d.Username,
				Password: d.Password,
				Role:     d.Role,
			}
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
func (a *userRepoImpl) UpdateAdmin(adminData [][]string) error {
	adminFile, err := os.Create("admin.csv")
	if err != nil {
		return err
	}
	defer adminFile.Close()
	adminWriter := csv.NewWriter(adminFile)
	defer adminWriter.Flush()
	err = adminWriter.WriteAll(adminData)
	if err != nil {
		return err
	}
	return nil
}

func (a *userRepoImpl) UpdateClient(clientData [][]string) error {
	clientFile, err := os.Create("client.csv")
	if err != nil {
		return err
	}
	clientWriter := csv.NewWriter(clientFile)
	defer clientWriter.Flush()
	err = clientWriter.WriteAll(clientData)
	if err != nil {
		return err
	}
	return nil
}

func (a *userRepoImpl) DeleteClient(clientData [][]string) error {
	clientFile, err := os.Create("client.csv")
	if err != nil {
		return err
	}
	clientWriter := csv.NewWriter(clientFile)
	defer clientWriter.Flush()
	err = clientWriter.WriteAll(clientData)
	if err != nil {
		return err
	}
	return nil
}

func (a *userRepoImpl) DeleteAdmin(adminData [][]string) error {
	adminFile, err := os.Create("admin.csv")
	if err != nil {
		return err
	}
	adminWriter := csv.NewWriter(adminFile)
	defer adminWriter.Flush()
	err = adminWriter.WriteAll(adminData)
	if err != nil {
		return err
	}
	return nil
}
