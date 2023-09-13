package tickets

import (
	"context"
	"errors"
	"testing"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/stretchr/testify/assert"
)

var cxt = context.Background()

var tickets = []domain.Ticket{
	{
		Id:      "1",
		Name:    "Tait Mc Caughan",
		Email:   "tmc0@scribd.com",
		Country: "Finland",
		Time:    "17:11",
		Price:   785.00,
	},
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

var ticketsByDestination = []domain.Ticket{
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

type stubRepo struct {
	db *DbMock
}

type DbMock struct {
	db  []domain.Ticket
	spy bool
	err error
}

func NewRepositoryTest(dbM *DbMock) RepositoryTicket {
	return &stubRepo{dbM}
}

func (r *stubRepo) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}

	return r.db.db, nil
}

func (r *stubRepo) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	var tkts []domain.Ticket

	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}

	for _, t := range r.db.db {
		if t.Country == destination {
			tkts = append(tkts, t)
		}
	}

	return tkts, nil
}

func TestGetTicketTotalByDestination(t *testing.T) {
	t.Run("GetTotalTicketsOK", func(t *testing.T) {
		dbMock := &DbMock{
			db:  tickets,
			spy: false,
			err: nil,
		}

		repo := NewRepositoryTest(dbMock)
		service := NewServiceTicket(repo)

		tkts, err := service.GetTotalTickets(cxt, "China")

		assert.Nil(t, err)
		assert.True(t, dbMock.spy)
		assert.Equal(t, len(ticketsByDestination), tkts)
	})

	t.Run("GetTotalTicketsError", func(t *testing.T) {
		dbMock := &DbMock{
			db:  []domain.Ticket{},
			spy: false,
			err: errors.New("Empty list"),
		}

		repo := NewRepositoryTest(dbMock)
		service := NewServiceTicket(repo)

		_, err := service.GetTotalTickets(cxt, "China")

		assert.NotNil(t, err)
		assert.True(t, dbMock.spy)
	})
}

func TestGetTicketAverageByDestination(t *testing.T) {
	t.Run("AverageDestinationOK", func(t *testing.T) {
		dbMock := &DbMock{
			db:  tickets,
			spy: false,
			err: nil,
		}

		repo := NewRepositoryTest(dbMock)
		service := NewServiceTicket(repo)

		avr, err := service.AverageDestination(cxt, "China")

		assert.Nil(t, err)
		assert.NotNil(t, avr)
		assert.True(t, dbMock.spy)
	})

	t.Run("AverageDestinationError", func(t *testing.T) {
		dbMock := &DbMock{
			db:  tickets,
			spy: false,
			err: errors.New("Empty list"),
		}

		repo := NewRepositoryTest(dbMock)
		service := NewServiceTicket(repo)

		_, err := service.AverageDestination(cxt, "China")

		assert.NotNil(t, err)
		assert.True(t, dbMock.spy)
	})
}
