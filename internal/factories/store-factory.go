package factories

import (
	random "github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
)

type Address struct {
	City    string
	State   string
	Country string
	Street  string
}

type Store struct {
	Id        uuid.UUID
	StoreName string
	Address   Address // Todo: turn this into just Address instead of Address Address (see if works)
}

type Customer struct {
	Id uuid.UUID
    First string
    Last string
    Address
    Store
}

type Franchise struct {
	Stores    []Store
	Customers []Customer
}

func (f *Franchise) addStore(s Store) {
	f.Stores = append(f.Stores, s)
}

func CreateFranchise() Franchise {
	franchise := &Franchise{}

	for i := 0; i < 10; i++ {
		address := Address{
			random.City(),
			random.State(random.Large),
			"United States",
			random.Street(),
		}

		store := Store{uuid.New(), random.Letters(10), address}

		franchise.addStore(store)
	}

	for i := 0; i < random

	return *franchise
}
