package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}
	// fmt.Printf("%+v", jim)

	// go is pass by so we need to use pointers
	// &variable - give me the memory address of the value and assign to pointer
	// memory address of where exists at
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// *pointer - gives me the value at that address
// * person - type description- it means we're working with a pointer to a person
// * pointerToPerson - this is an operator- it means we want to manipulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
