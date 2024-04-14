package service

type Repository interface {
}

type Service struct {
	repository Repository
}

func NewService(
	repository Repository,
) (Service, error) {
	return Service{
		repository: repository,
	}, nil
}
