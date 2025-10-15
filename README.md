# mensaapi
An unofficial golang API for the Studierendenwerk Ulms Mensa meal plans.

# Available locations
    Universität Ulm
    - Mensa in der Uni Ulm
    - Cafeteria Helmholtzstraße
    - Cafeteria West 

    Technische Hochschule Ulm
    - THU Ulm Mensa
    - THU Ulm Außenstelle Oberer Eselsberg Essensausgabe

    Hochschule Aalen
    - Mensa
    - Cafeteria in der Hochschule

    Pädagogische Hochschule Schwäbisch Gmünd
    - Mensaria PH

    Hochschule für Gestaltung Schwäbisch Gmünd
    - Hochschule für Gestaltung Schwäbisch Gmünd

    HBC. Biberach
    - HBC. Biberach

    Duale Hochschule Heidenheim
    - Duale Hochschule Heidenheim

# Example
``` golang
package main

import (
	"fmt"
	"time"

	"github.com/tiemingo/mensaapi"
)

func main() {
	// day, err := mensaapi.GetDayByLocation(mensaapi.LOC_MENSA_UNI_ULM, mensaapi.ConvertStringToDate("2025-10-30"))
	day, err := mensaapi.GetDayByLocation(mensaapi.LOC_MENSA_UNI_ULM, time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("---" + mensaapi.ConvertDateToString(day.Date) + "---")
	for category, items := range day.Categories {
		fmt.Println("---" + category + "---")
		for _, item := range items {
			fmt.Print("  ")
			if len(item.Additives.Tips) > 0 {
				fmt.Print("(")
				for _, v := range item.Additives.Tips {
					fmt.Print(mensaapi.GetTipDescription(v), ", ")
				}
				fmt.Print(") ")
			}
			fmt.Println(item.Name)
			if item.Info != "" {
				fmt.Println("	" + item.Info)
			}
			fmt.Println("	", item.Prices.PriceStudent, "€")
		}
	}
}
```