# Redovalnica

Aplikacija za upravljanje študentov in ocen, napisana v programskem jeziku Go.

## Opis

Redovalnica omogoča:
- Dodajanje ocen študentom
- Izpis redovalnice z vsemi študenti in ocenami
- Izračun in izpis končnega uspeha študentov s komentarjem

## Namestitev

```bash
go install github.com/nikkastigar/redovalnica/cmd/redovalnica@latest
```

## Uporaba

### Osnovna uporaba

```bash
redovalnica
```

### Uporaba s stikali

```bash
redovalnica --stOcen=5 --minOcena=1 --maxOcena=10 --mejaPozitivna=5.5
```

### Stikala

- `--stOcen` - Najmanjše število ocen potrebnih za pozitivno oceno (privzeto: 6)
- `--minOcena` - Najmanjša možna ocena (privzeto: 1)
- `--maxOcena` - Največja možna ocena (privzeto: 10)
- `--mejaPozitivna` - Mejna povprečna ocena za pozitivno oceno (privzeto: 6.0)

### Pomoč

```bash
redovalnica --help
```

## Paket redovalnica

Paket lahko uvozite v svoj Go projekt:

```go
import "github.com/nikkastigar/redovalnica/redovalnica"
```

### Izvoženi tipi

```go
type Student struct {
    Ime     string
    Priimek string
    Ocene   []int
}
```

### Izvožene funkcije

#### DodajOceno

```go
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena, maxOcena int)
```

Doda oceno študentu z dano vpisno številko. Preveri veljavnost ocene in obstoj študenta.

#### IzpisVsehOcen

```go
func IzpisVsehOcen(studenti map[string]Student)
```

Izpiše redovalnico z vsemi študenti in njihovimi ocenami, urejeno po vpisni številki.

#### IzpisiKoncniUspeh

```go
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int, mejaPozitivna float64)
```

Izpiše končni uspeh vseh študentov s povprečno oceno in komentarjem.

## Primer uporabe v kodi

```go
package main

import (
    "github.com/nikkastigar/redovalnica/redovalnica"
)

func main() {
    studenti := make(map[string]redovalnica.Student)
    
    studenti["12345678"] = redovalnica.Student{
        Ime:     "Janez",
        Priimek: "Novak",
        Ocene:   []int{6, 8, 9},
    }
    
    redovalnica.DodajOceno(studenti, "12345678", 10, 0, 10)
    redovalnica.IzpisVsehOcen(studenti)
    redovalnica.IzpisiKoncniUspeh(studenti, 6, 6.0)
}
```

## Lokalna gradnja

```bash
git clone https://github.com/nikkastigar/redovalnica.git
cd redovalnica
go build ./cmd/redovalnica
./redovalnica
```

## Dokumentacija

Celotna dokumentacija je na voljo na: https://pkg.go.dev/github.com/nikkastigar/redovalnica

## Licenca

Projekt je namenjen za akademske namene.
