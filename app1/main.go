package main

import (
	"app1/db"
	"fmt"
)

func main() {
	data := db.Open()
	fmt.Println(data)
	casaDani := data.NewObra("casa Dani")
	casaDani.AddIngresos(db.Ingresos{Money: 100})
	casaDani.AddIngresos(db.Ingresos{Money: 2000})
	fmt.Println(casaDani.GetIngresos())
	casaDani.UpDateIngresos(1, db.Ingresos{Name: "casa Dani", Money: 2300})
	fmt.Println(casaDani.GetIngresos())
}
