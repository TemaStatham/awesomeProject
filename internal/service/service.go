package service

import "log/slog"

type Service struct {
	log *slog.Logger
}

type GoodSaver interface {
	SaveGood()
}

type GoodUpdater interface {
	ChangeDescription()
	RedistributePriorities()
}

type GoodRemover interface {
	Remove()
}

type GoodGetter interface {
	GetList()
}

func NewService(_l *slog.Logger) *Service {
	return &Service{
		log: _l,
	}
}

func (s *Service) Create() {

}

func (s *Service) Update() {

}

func (s *Service) Reprioritize() {

}

func (s *Service) Remove() {

}

func (s *Service) GetList() {

}
