package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/jeraj/redovalnica_modul/redovalnica"
)

func main() {
	cmd := &cli.Command{ //command line interface
		Name:  "redovalnica",
		Usage: "Program za dodajanje ocen in izpis končnega uspeha študentov",
		Flags: []cli.Flag{ //dodajanje stikal
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "najmanjše število ocen potrebnih za pozitivno oceno", //opis
				Value: 3, //default vrednost
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "najmanjša možna ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "največja možna ocena",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			//preberemo vrednosti stikal
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			//zazenemo glavni del programo
			return runRedovalnica(stOcen, minOcena, maxOcena)
		},
	}

	//zazenemo CLI aplikacijo
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func runRedovalnica(stOcen, minOcena, maxOcena int) error {
	//simulacija, da funkcije delajo
	fmt.Println("REDOVALNICA\n")
	fmt.Printf("Parametri podani v stikalih: stOcen=%d, minOcena=%d, maxOcena=%d\n\n", stOcen, minOcena, maxOcena)

	//inicializacija primerov študentov
	redovalnica.Redovalnica["632101300"] = redovalnica.Student{"Marko", "Novak", []int{6, 6, 7}}
	redovalnica.Redovalnica["632101302"] = redovalnica.Student{"Marija", "Horvat", []int{10, 10, 8}}
	redovalnica.Redovalnica["632101304"] = redovalnica.Student{"Maja", "Kos", []int{4, 3, 5}}

	//uporaba funkcij iz paketa redovalnica
	redovalnica.DodajOceno("632101302", 10, minOcena, maxOcena)
	fmt.Println()
	redovalnica.IzpisVsehOcen()
	fmt.Println()
	redovalnica.IzpisiKoncniUspeh(stOcen)


	return nil //isto kot return null
}