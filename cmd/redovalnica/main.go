package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nikkastigar/redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Upravljanje študentov in ocen",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "stOcen",
				Value:   6,
				Usage:   "Najmanjše število ocen potrebnih za pozitivno oceno",
				Sources: cli.EnvVars("ST_OCEN"),
			},
			&cli.IntFlag{
				Name:    "minOcena",
				Value:   1,
				Usage:   "Najmanjša možna ocena",
				Sources: cli.EnvVars("MIN_OCENA"),
			},
			&cli.IntFlag{
				Name:    "maxOcena",
				Value:   10,
				Usage:   "Največja možna ocena",
				Sources: cli.EnvVars("MAX_OCENA"),
			},
			&cli.Float64Flag{
				Name:    "mejaPozitivna",
				Value:   6.0,
				Usage:   "Mejna povprečna ocena za pozitivno oceno",
				Sources: cli.EnvVars("MEJA_POZITIVNA"),
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			mejaPozitivna := cmd.Float64("mejaPozitivna")

			fmt.Printf("Nastavitve: minOcena=%d, maxOcena=%d, stOcen=%d, mejaPozitivna=%.1f\n\n",
				minOcena, maxOcena, stOcen, mejaPozitivna)

			studenti := make(map[string]redovalnica.Student)

			studenti["12345678"] = redovalnica.Student{
				Ime:     "Janez",
				Priimek: "Novak",
				Ocene:   []int{6, 8, 9},
			}
			studenti["87654321"] = redovalnica.Student{
				Ime:     "Maja",
				Priimek: "Horvat",
				Ocene:   []int{10, 9, 8, 7, 9, 10},
			}
			studenti["65432123"] = redovalnica.Student{
				Ime:     "Luka",
				Priimek: "Kovač",
				Ocene:   []int{5, 5, 6},
			}
			studenti["11111111"] = redovalnica.Student{
				Ime:     "Ana",
				Priimek: "Zupan",
				Ocene:   []int{9, 10, 10, 9, 10, 9, 10},
			}

			fmt.Println("=== ZAČETNO STANJE ===")
			redovalnica.IzpisVsehOcen(studenti)

			fmt.Println("\n=== DODAJANJE OCEN ===")
			redovalnica.DodajOceno(studenti, "12345678", 10, minOcena, maxOcena)
			fmt.Println("Dodana ocena 10 študentu 12345678")

			redovalnica.DodajOceno(studenti, "87654321", 8, minOcena, maxOcena)
			fmt.Println("Dodana ocena 8 študentu 87654321")

			redovalnica.DodajOceno(studenti, "00000000", 7, minOcena, maxOcena)

			redovalnica.DodajOceno(studenti, "12345678", 11, minOcena, maxOcena)

			redovalnica.IzpisVsehOcen(studenti)
			redovalnica.IzpisiKoncniUspeh(studenti, stOcen, mejaPozitivna)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Napaka: %v\n", err)
		os.Exit(1)
	}
}
