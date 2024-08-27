package usecase

import (
	domain "github.com/legend123213/loan_managment/Domain"
	infrastructure "github.com/legend123213/loan_managment/Infrastructure"
	repositery "github.com/legend123213/loan_managment/Repositery"
)



type UserServiceUsecase struct{
	UserRepo repositery.UserServiceRepo
}

func NewUserServiceUsecase(repo repositery.UserServiceRepo) *UserServiceUsecase{
	return &UserServiceUsecase{
		UserRepo: repo,
	}
}	
func (usecase *UserServiceUsecase) AddUser(user *domain.User) (*domain.User, error){
	user.ID = ""
	user.Password,_=infrastructure.PasswordHasher(user.Password)
	
	return usecase.UserRepo.CreateUser(user)
}
func (usecase *UserServiceUsecase) GetUser(id string) (*domain.User, error){
	return usecase.UserRepo.GetUser(id)
}
func (usecase *UserServiceUsecase) GetUsers() ([]*domain.User, error){
	return usecase.UserRepo.GetUsers()
}

func (usecase *UserServiceUsecase) UpdateUser(user *domain.User) (*domain.User, error){
	return usecase.UserRepo.UpdateUser(user)
}
func (usecase *UserServiceUsecase) DeleteUser(id string) error{
	return usecase.UserRepo.DeleteUser(id)
}
