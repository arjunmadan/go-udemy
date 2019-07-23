package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	arjun := person{
		firstName: "Arjun",
		lastName:  "Madan",
		contact: contactInfo{
			email: "arjunmadan1992@gmail.com",
			zip:   11111,
		},
	}

	arjun.print()
	arjun.updateFirstName("blah")
	arjun.print()

	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)

	name := "bill"
	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
