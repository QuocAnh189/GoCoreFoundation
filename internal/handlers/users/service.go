package users

type UserService struct {
	repo IRepository
}

func NewService(repo IRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
