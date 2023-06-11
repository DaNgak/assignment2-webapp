package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {

	getUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if getUser.Email == "" {
		return nil, errors.New("user not found")
	}

	if getUser.Password != user.Password {
		return nil, errors.New("wrong email or password")
	}

	*user = getUser

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.Claims{
		UserID: getUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokens.SignedString(model.JwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil // TODO: replace this
}

func (s *userService) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	return s.userRepo.GetUserTaskCategory()
}