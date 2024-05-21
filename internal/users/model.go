package users

import "fmt"

type (
	IModel interface {
		User(id uint) *UserModel
	}

	Model struct {
	}

	UserModel struct {
		ID   uint
		Name string
	}

	UserSlice []*UserModel
)

func InitModel() Model {
	return Model{}
}

func (m Model) User(id uint) *UserModel {
	return &UserModel{ID: id, Name: fmt.Sprintf("Test%d", id)}
}

func (m Model) Users() UserSlice {
	return UserSlice{
		{ID: 1, Name: "Test1"},
		{ID: 2, Name: "Test2"},
		{ID: 3, Name: "Test3"},
	}
}
