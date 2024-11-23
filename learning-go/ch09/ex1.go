package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var MySentinel = errors.New("Invalid Id")

type MyError struct {
	FieldName string
}

func (me MyError) Error() string {
	return me.FieldName
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)

		var myError MyError
		if err != nil {
			if errors.Is(err, MySentinel) {
				fmt.Printf("This is a sentinel error")
				continue
			} else if errors.As(err, &myError) {
				fmt.Printf("record %d: %+v error: empty field %s\n", count, emp, myError.FieldName)
				continue
			} else {
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
				continue
			}
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var validID = regexp.MustCompile(`\w{4}-\d{3}`)

func ValidateEmployee(e Employee) error {
	var err []error
	if len(e.ID) == 0 {
		err = append(err, MyError{"ID"})
	}
	if !validID.MatchString(e.ID) {
		err = append(err, MySentinel)
		// return errors.New("invalid ID")
	}
	if len(e.FirstName) == 0 {
		err = append(err, MyError{"FirstName"})
	}
	if len(e.LastName) == 0 {
		err = append(err, MyError{"LastName"})
	}
	if len(e.Title) == 0 {
		err = append(err, MyError{"Title"})
	}

	if len(err) != 0 {
		return errors.Join(err...)
	}
	return nil
}