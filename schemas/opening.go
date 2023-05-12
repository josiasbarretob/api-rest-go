package schemas

import "gorm.io/gorm"

type Opening struct{
	gorm.Model
	Role string //Vaga
	Company string
	Location string
	Remote bool
	Link string
	Salary float64
}