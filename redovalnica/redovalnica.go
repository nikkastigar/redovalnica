// Package redovalnica omogoča upravljanje študentov in njihovih ocen.
// Paket vsebuje funkcije za dodajanje ocen, izpis redovalnice in izračun končnega uspeha.
package redovalnica

import (
	"fmt"
	"sort"
)

// Student predstavlja študenta z imenom, priimkom in seznamom ocen.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno študentu z dano vpisno številko.
// Preveri, ali je ocena v veljavnem območju (med minOcena in maxOcena) in ali študent obstaja.
// Če študent ne obstaja ali je ocena neveljavna, izpiše obvestilo.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena, maxOcena int) {
	if ocena < minOcena || ocena > maxOcena {
		fmt.Printf("Ocena %d ni v ustreznem območju (%d-%d).\n", ocena, minOcena, maxOcena)
		return
	}

	s, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent s vpisno številko", vpisnaStevilka, "ne obstaja.")
		return
	}

	s.Ocene = append(s.Ocene, ocena)
	studenti[vpisnaStevilka] = s
}

// povprecje izračuna povprečno oceno študenta z dano vpisno številko.
// Vrne -1.0, če študent ne obstaja, ali 0.0, če študent nima ocen.
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	s, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}

	n := len(s.Ocene)
	if n == 0 {
		return 0.0
	}

	var sum int
	for _, oc := range s.Ocene {
		sum += oc
	}

	return float64(sum) / float64(n)
}

// IzpisVsehOcen izpiše redovalnico z vsemi študenti in njihovimi ocenami.
// Študenti so izpisani urejeno po vpisni številki.
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("\nREDOVALNICA:")
	keys := make([]string, 0, len(studenti))
	for k := range studenti {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		s := studenti[k]
		fmt.Printf("%s - %s %s: %v\n", k, s.Ime, s.Priimek, s.Ocene)
	}
}

// IzpisiKoncniUspeh izpiše končni uspeh vseh študentov s povprečno oceno in komentarjem.
// Komentar je odvisen od povprečne ocene in minimalnega števila ocen.
// Če študent nima dovolj ocen (manj kot stOcen), povprečje ni upoštevano.
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int, mejaPozitivna float64) {
	fmt.Println("\nKONČNI USPEH:")
	keys := make([]string, 0, len(studenti))
	for k := range studenti {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		s := studenti[k]

		if len(s.Ocene) < stOcen {
			fmt.Printf("%s %s: premalo ocen (%d/%d)\n", s.Ime, s.Priimek, len(s.Ocene), stOcen)
			continue
		}

		avg := povprecje(studenti, k)
		if avg == -1.0 {
			fmt.Printf("%s %s: študent ne obstaja\n", s.Ime, s.Priimek)
			continue
		}

		komentar := ""
		if avg >= 9.0 {
			komentar = "Odličen študent!"
		} else if avg >= mejaPozitivna {
			komentar = "Povprečen študent"
		} else {
			komentar = "Neuspešen študent"
		}

		fmt.Printf("%s %s: povprečna ocena %.1f -> %s\n", s.Ime, s.Priimek, avg, komentar)
	}
}
