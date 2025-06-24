package models

import (
	"fmt"

	"github.com/HubPavKul1/vetstore2025/internal/utils"
	"gorm.io/gorm"
)


type Address struct {
	gorm.Model
	Index string `gorm:"uniqueIndex"`
	City string `gorm:"default:'Иваново'"`
	Street string 
	HouseNumber string
	PhoneNumber1 string
	PhoneNumber2 string `gorm:"null"`
	EMail string `gorm:"null"`
	Departments []Department
}

type Department struct {
	gorm.Model
	Name string
	AddressID uint
	Employees []Employee
	MovingsFromStore []MovingFromStore

}

type Position struct {
	gorm.Model
	Name string 
	Employees []Employee
}

type Employee struct {
	gorm.Model
	PositionID uint
	DepartmentID uint
	FirstName string
	Patronymic string
	LastName string
}


func(e Employee)GetFullName()string {
	
	return fmt.Sprintf("%s %s %s", e.LastName, e.FirstName, e.Patronymic)
}

func(e Employee)GetShortName()string {
	firstNameLetter := utils.FirstLetterFromString(e.FirstName)
	firstPatrLetter := utils.FirstLetterFromString(e.Patronymic)
	return fmt.Sprintf("%s %s. %s.",e.LastName, firstNameLetter, firstPatrLetter)
}

func(a Address)GetAddressString()string {
	return fmt.Sprintf("%s, г. %s, ул.%s, д. %s", a.Index, a.City, a.Street, a.HouseNumber)
}