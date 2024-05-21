package users

type (
	IService interface {
		GetUsers(id uint) (*UserResp, error)
	}

	Service struct {
		repo Repo
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
