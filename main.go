package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Sort() {
	sort.Sort(p)
}

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v %v", p.firstName, p.lastName, p.birthDay)
}

func (p People) Less(i, j int) bool {
	if p[i].birthDay == p[j].birthDay {
		return p[i].lastName < p[j].lastName || p[i].firstName < p[j].firstName
	}
	return p[i].birthDay.After(p[j].birthDay)
}

func (p People) Append(person Person) People {
	return append(p, person)
}

func createNewPerson(firstName string, lastName string, birthDay string) (Person, error) {
	const (
		layoutISO = "2006-01-02"
	)
	birthDayTime, err := time.Parse(layoutISO, birthDay)
	if err != nil {
		return Person{}, fmt.Errorf("incorrect birthDay %s. Error: %s", birthDay, err)
	}
	return Person{
		firstName: firstName,
		lastName:  lastName,
		birthDay:  birthDayTime,
	}, nil
}

func main() {
	var peopleList People
	testData := [][]string{
		{"Ivan", "Ivanov", "2005-08-10"},
		{"Ivan", "Ivanov", "2003-08-05"},
		{"Artiom", "Ivanov", "2005-08-10"},
	}
	for _, data := range testData {
		if p, err := createNewPerson(data[0], data[1], data[2]); err == nil {
			peopleList = peopleList.Append(p)
		}
	}
	sort.Sort(peopleList)
	fmt.Println(peopleList)
}
