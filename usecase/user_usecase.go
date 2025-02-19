package usecase

import (
	"errors"

	"github.com/Ayyasy123/todo-list-api/models"
	"github.com/Ayyasy123/todo-list-api/repository"
	"github.com/Ayyasy123/todo-list-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(req *models.RegisterReq) (*models.UserRes, error)
	LoginUser(req *models.LoginReq) (*models.UserRes, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) RegisterUser(req *models.RegisterReq) (*models.UserRes, error) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &models.UserRes{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *userUsecase) LoginUser(req *models.LoginReq) (*models.UserRes, error) {
	user, err := u.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	// Generate token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.UserRes{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token, // Tambahkan token ke response
	}, nil
}
