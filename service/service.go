package service

type Repository interface {
}

type Service struct {
	repository Repository
}

func NewService() (Service, error) {
	return Service{}, nil
}
