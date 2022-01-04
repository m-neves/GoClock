package service

import (
	"errors"
	"time"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/repository"
)

type EntryServiceInterface interface {
	Create(entry *data.Entry) error
	FindById(id, userId int) (*data.Entry, error)
	SumByPeriod(start, end time.Time, userId int) (*data.Entries, error)
}

type entryService struct {
	er repository.EntryRepositoryInterface
}

func NewEntryService(er repository.EntryRepositoryInterface) EntryServiceInterface {
	return &entryService{er}
}

func (es *entryService) Create(entry *data.Entry) error {

	if !entry.IsValid() {
		return errors.New("invalid entry type")
	}

	return es.er.Create(entry)
}

func (es *entryService) FindById(id, userId int) (*data.Entry, error) {
	return es.er.FindById(id, userId)
}

func (es *entryService) SumByPeriod(start, end time.Time, userId int) (*data.Entries, error) {
	return es.er.SumByPeriod(start, end, userId)
}
