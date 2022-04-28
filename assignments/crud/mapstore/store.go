package mapstore

import (
	"errors"
	"learn-go/assignments/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func (m MapStore) Create(customer domain.Customer) error {
	m.store[customer.ID] = customer
	return nil
}

func (m MapStore) Update(ID string, customer domain.Customer) error {
	if _, ok := m.store[ID]; ok {
		m.store[ID] = customer
		return nil
	}
	return errors.New("customer record cannot be updated")
}

func (m MapStore) Delete(ID string) error {
	if _, ok := m.store[ID]; ok {
		delete(m.store, ID)
		return nil
	}
	return errors.New("customer record cannot be deleted")
}

func (m MapStore) GetById(ID string) (domain.Customer, error) {
	if _, ok := m.store[ID]; ok {
		return m.store[ID], nil
	}
	return domain.Customer{}, errors.New("customer record not found")
}
func (m MapStore) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	for _, customer := range m.store {
		customers = append(customers, customer)
	}

	if customers != nil {
		return customers, nil
	}
	return nil, errors.New("no customer in the map")

}
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}
