package users

import "context"

type UserService struct {
    repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) error {
    return s.repo.Create(ctx, User{Name: name, Email: email})
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]User, error) {
    return s.repo.List(ctx)
}

