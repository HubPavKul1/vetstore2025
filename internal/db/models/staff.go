package models

import (
	"github.com/HubPavKul1/vetstore2025/internal/utils"
	"gorm.io/gorm"
)


type Department struct {
	gorm.Model
	Name string
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
	
	return e.LastName + " " + e.FirstName + " " + e.Patronymic
}

func(e Employee)GetShortName()string {
	firstNameLetter := utils.FirstLetterFromString(e.FirstName)
	firstPatrLetter := utils.FirstLetterFromString(e.Patronymic)
	return e.LastName + " " + firstNameLetter + ". " + firstPatrLetter + "."
}