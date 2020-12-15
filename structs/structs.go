package structs

import "github.com/jinzhu/gorm"

// Person is ...
type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}
