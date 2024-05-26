package users

import (
	"gopkg.in/yaml.v3"
)

//go:generate go run github.com/vektra/mockery/v2 --name IService --case snake

type (
	IService interface {
		GetUsers(id uint) (*UserResp, error)
		ParseFile(data string) error
	}

	Service struct {
		repo Repo
	}

	T struct {
		Name        string `yaml:"name,omitempty"`
		Description string `yaml:"description,omitempty"`
	}
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func InitService(repo Repo) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetUsers(id uint) (*UserResp, error) {
	return s.repo.GetUsers(id)
}

func (s Service) ParseFile(data string) (err error) {
	var content T
	err = yaml.Unmarshal([]byte(data), &content)
	if err != nil {
		return err
	}

	return
}
