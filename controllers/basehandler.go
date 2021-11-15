package controllers

import "containerization/repository/user"

type Handlers struct {
	Repository user.UserRepository
}

func NewHandler(repository user.UserRepository) *Handlers {
	return &Handlers{
		Repository: repository,
	}
}
