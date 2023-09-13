package tickets

import (
	"context"
)

type service struct {
	repository RepositoryTicket
}

func NewServiceTicket(repository RepositoryTicket) ServiceTicket {
	return &service{
		repository: repository,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	totalTickets, err := s.repository.GetAll(ctx)

	if err != nil {
		return 0, err
	}

	destTickets, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	return float64(len(destTickets)) / float64(len(totalTickets)), nil
}
