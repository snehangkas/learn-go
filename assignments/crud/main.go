package main

import (
	"fmt"
	"learn-go/assignments/domain"
	"learn-go/assignments/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (controller CustomerController) Add(c domain.Customer) {
	err := controller.store.Create(c)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (controller CustomerController) Modify(c domain.Customer) {
	err := controller.store.Update(c.ID, c)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Customer is updated")

}

func (controller CustomerController) Remove(id string) {
	err := controller.store.Delete(id)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
}

func (controller CustomerController) Fetch(id string) domain.Customer {
	customer, err := controller.store.GetById(id)
	if err != nil {
		fmt.Println("Error : ", err)
		return domain.Customer{}
	}
	return customer
}

func (controller CustomerController) FetchAll() []domain.Customer {
	customers, err := controller.store.GetAll()
	if err != nil {
		fmt.Println("Error : ", err)
		return nil
	}
	return customers
}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(), // Inject the dependency
	}
	controller.Add(domain.Customer{
		ID:   "001",
		Name: "Sam",
	})
	controller.Add(domain.Customer{
		ID:   "002",
		Name: "John",
	})
	controller.Add(domain.Customer{
		ID:   "003",
		Name: "Jerry",
	})
	controller.Modify(domain.Customer{
		ID:   "002",
		Name: "Tom",
	})
	fmt.Println("Current map store contains", controller.FetchAll())
	fmt.Println("Updated recode: ", controller.Fetch("001"))
	fmt.Println("Adding a new record")
	controller.Add(domain.Customer{
		ID:   "004",
		Name: "Harry",
	})
	fmt.Println("Updated map store", controller.FetchAll())
	controller.Remove("004")
	fmt.Println("Final map store", controller.FetchAll())
}
