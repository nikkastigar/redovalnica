# Primeri uporabe

## 1. Zagon aplikacije s privzetimi nastavitvami

```bash
go run cmd/redovalnica/main.go
```

Izhod:
```
Nastavitve: minOcena=1, maxOcena=10, stOcen=6, mejaPozitivna=6.0

=== ZAČETNO STANJE ===

REDOVALNICA:
11111111 - Ana Zupan: [9 10 10 9 10 9 10]
12345678 - Janez Novak: [6 8 9]
65432123 - Luka Kovač: [5 5 6]
87654321 - Maja Horvat: [10 9 8 7 9 10]

=== DODAJANJE OCEN ===
Dodana ocena 10 študentu 12345678
Dodana ocena 8 študentu 87654321
Študent s vpisno številko 00000000 ne obstaja.
Ocena 11 ni v ustreznem območju (1-10).

REDOVALNICA:
11111111 - Ana Zupan: [9 10 10 9 10 9 10]
12345678 - Janez Novak: [6 8 9 10]
65432123 - Luka Kovač: [5 5 6]
87654321 - Maja Horvat: [10 9 8 7 9 10 8]

KONČNI USPEH:
Ana Zupan: povprečna ocena 9.6 -> Odličen študent!
Janez Novak: premalo ocen (4/6)
Luka Kovač: premalo ocen (3/6)
Maja Horvat: povprečna ocena 8.7 -> Povprečen študent
```

## 2. Zagon z nižjimi zahtevami (3 ocene)

```bash
go run cmd/redovalnica/main.go --stOcen=3
```

Izhod:
```
Nastavitve: minOcena=1, maxOcena=10, stOcen=3, mejaPozitivna=6.0

...

KONČNI USPEH:
Ana Zupan: povprečna ocena 9.6 -> Odličen študent!
Janez Novak: povprečna ocena 8.2 -> Povprečen študent
Luka Kovač: povprečna ocena 5.3 -> Neuspešen študent
Maja Horvat: povprečna ocena 8.7 -> Povprečen študent
```

## 3. Zagon z drugačnimi mejami ocen

```bash
go run cmd/redovalnica/main.go --minOcena=5 --maxOcena=10
```

## 4. Zagon z nižjo mejo za pozitivno oceno

```bash
go run cmd/redovalnica/main.go --stOcen=3 --mejaPozitivna=5.0
```

Izhod:
```
KONČNI USPEH:
Ana Zupan: povprečna ocena 9.6 -> Odličen študent!
Janez Novak: povprečna ocena 8.2 -> Povprečen študent
Luka Kovač: povprečna ocena 5.3 -> Povprečen študent
Maja Horvat: povprečna ocena 8.7 -> Povprečen študent
```

## 5. Kombinirane nastavitve

```bash
go run cmd/redovalnica/main.go --stOcen=4 --minOcena=1 --maxOcena=5 --mejaPozitivna=3.0
```

## 6. Pomoč

```bash
go run cmd/redovalnica/main.go --help
```

Izhod:
```
NAME:
   redovalnica - Upravljanje študentov in ocen

USAGE:
   redovalnica [global options]

GLOBAL OPTIONS:
   --stOcen int           Najmanjše število ocen potrebnih za pozitivno oceno (default: 6) [%ST_OCEN%]
   --minOcena int         Najmanjša možna ocena (default: 1) [%MIN_OCENA%]
   --maxOcena int         Največja možna ocena (default: 10) [%MAX_OCENA%]
   --mejaPozitivna float  Mejna povprečna ocena za pozitivno oceno (default: 6) [%MEJA_POZITIVNA%]
   --help, -h             show help
```

## 7. Prevajanje v izvršljivo datoteko

```bash
go build -o redovalnica cmd/redovalnica/main.go
./redovalnica --stOcen=3 --mejaPozitivna=5.0
```

## 8. Namestitev

```bash
go install github.com/nikkastigar/redovalnica/cmd/redovalnica@latest
redovalnica --help
```

## 9. Uporaba kot knjižnica v drugem projektu

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
        Ocene:   []int{6, 8, 9, 7, 8, 10},
    }
    
    redovalnica.DodajOceno(studenti, "12345678", 9, 1, 10)
    redovalnica.IzpisVsehOcen(studenti)
    redovalnica.IzpisiKoncniUspeh(studenti, 6, 6.0)
}
```

## 10. Nastavljanje preko okoljskih spremenljivk

```bash
export ST_OCEN=4
export MIN_OCENA=1
export MAX_OCENA=10
export MEJA_POZITIVNA=5.5

go run cmd/redovalnica/main.go
```
