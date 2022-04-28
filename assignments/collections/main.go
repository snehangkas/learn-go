package main

import (
	"fmt"
)

var data map[string]string

func init() {
	data = make(map[string]string)
}

func main() {
	add("frontend", "react")
	add("backend", "go")
	add("infrastructure", "k8s")
	add("db", "cockroachdb")
	fmt.Println(getAll())
	updateItem("db", "cloud spanner")
	fmt.Println("The new db value is", getById("db"))
	deleteItem("frontend")
	fmt.Println(getAll())
}

func add(k, v string) {
	data[k] = v
}

func updateItem(k, v string) {
	data[k] = v
}
func getById(k string) string {
	return data[k]
}
func getAll() map[string]string {
	return data
}
func deleteItem(k string) {
	delete(data, k)
}
