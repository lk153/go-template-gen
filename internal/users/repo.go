package users

type (
	IRepo interface {
		GetUsers(id uint) (*UserResp, error)
	}

	Repo struct {
		model Model
	}
)

type UserResp struct {
	ID   uint
	Name string
}

func InitRepo(model Model) Repo {
	return Repo{model: model}
}

func (repo Repo) GetUsers(id uint) (*UserResp, error) {
	return transform(repo.model.User(id)), nil
}

func transform(model *UserModel) *UserResp {
	return &UserResp{
		ID:   model.ID,
		Name: model.Name,
	}
}
