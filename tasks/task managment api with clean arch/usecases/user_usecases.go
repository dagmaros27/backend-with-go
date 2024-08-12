package usecases

import (
	"context"
	//"errors"
	//"net/http"
	"task_managment_api/domain"
	"task_managment_api/infrastructure"
	"time"

	//"go.mongodb.org/mongo-driver/mongo"
)

type userUsecase struct {
	userRepository domain.UserRepository
	ctxTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, ctxTimeout time.Duration) domain.UserUsecase {
	return &userUsecase{userRepository: userRepository, ctxTimeout: ctxTimeout}
}


func (uc *userUsecase)RegisterUser(c context.Context, user domain.User) domain.CustomError{
	_ ,err := uc.userRepository.GetUserByUsername(c, user.Username)

	if err.ErrCode != 0 {
		return err
	}
	


	count, err := uc.userRepository.GetUserCount(c)

	if err.ErrCode != 0 {
		return err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}	
	
	return uc.userRepository.CreateUser(c, user)
}


func (uc *userUsecase)AuthenticateUser(c context.Context, username, password string) (string, domain.CustomError){
	
	user, err := uc.userRepository.GetUserByUsername(c, username)

	if err.ErrCode != 0 {
		return "", err
	}

	err = infrastructure.VerifyPassword(user, password)

	if err.ErrCode != 0 { 
		return "", err
	}

	return infrastructure.GenerateUserToken(user)
}


func (uc *userUsecase)PromoteUser(c context.Context, username string) domain.CustomError{
	user, err := uc.userRepository.GetUserByUsername(c, username)
	if err.ErrCode != 0 {
		return err
	}
	user.Role = "admin"
	return uc.userRepository.UpdateUser(c, user)
}

