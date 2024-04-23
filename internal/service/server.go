package service

type Service struct {
	AnimeService
	UserService
}

var Server = new(Service)
