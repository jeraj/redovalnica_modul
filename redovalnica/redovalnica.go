//Package redovalnica vsebuje funkcije za delo z ocenami in izpis končnega stanja uspeha študentov
package redovalnica

import "fmt"

//Student predstavlja enega posameznega študenta in njegove ocene
type Student struct {
    Ime     string
    Priimek string
    Ocene   []int
}

//Redovalnica hrani podatke o ocenah študentov
var Redovalnica = make(map[string]Student) //NUJNO treba pisat z veliko!!

//IzpisVsehOcen izpiše ocene vseh študentov
func IzpisVsehOcen(){
	fmt.Println("REDOVALNICA:")
	for kljuc,s := range Redovalnica {
		fmt.Printf("%s - %s %s: %v\n", kljuc, s.Ime, s.Priimek, s.Ocene)
	}
}

//DodajOceno je funkcija, ki študentu z vpisno številko <x> doda (celoštevilsko) oceno 
func DodajOceno(vpisnaStevilka string, ocena int, minOcena int, maxOcena int){
	if ocena < minOcena || ocena > maxOcena{
		fmt.Println("Neveljavna ocena!")
		return
	}

	if s, ok := Redovalnica[vpisnaStevilka]; ok { //s je kopija structa, tega specificnega studenta
		s.Ocene = append(s.Ocene, ocena) //dodamo oceno tej kopiji
		Redovalnica[vpisnaStevilka] = s //jo shranimo nazaj, nadomesti original
		fmt.Println("Ocena uspesno vpisana.")
	} else {
		fmt.Println("Student s tako vpisno stevilko ne obstaja!")
		return
	}
}

//povprecje je skrita funkcija, računa povprečje ocen študenta z vpisno številko <x>
func povprecje(vpisnaStevilka string)float64{
	//da ostane skrita funkcija, se mora zaceti ime funkcije z malo zacetnico 
	//ostale funkcije ki morajo biti vidne izven paketa pa imajo veliko zacetnico
	if s, ok := Redovalnica[vpisnaStevilka]; ok {
		if len(s.Ocene) == 0{
			return 0
		}

		vsota := 0
		for _, ocena := range s.Ocene{
			vsota +=ocena
		}

		var povprecje float64 = float64(vsota) / float64(len(s.Ocene))

		if povprecje < 6 {
			return 0.0
		}

		return povprecje

	} else { //student ne obstaja
		return -1
	}
}

//IzpisiKoncniUspeh je funkcija, ki izpiše končni uspeh vseh študentov v redovalnici
func IzpisiKoncniUspeh(stOcen int){
	for vpisna,s := range Redovalnica {
		if len(s.Ocene) < stOcen{
			fmt.Printf(" -> Negativno ocenjen (premalo ocen)\n")
		}
		var povp float64 = povprecje(vpisna)
		if povp == 0{
			if len(s.Ocene) == 0{
				continue
			}
			var vsota float64 = 0
			for _,ocena := range s.Ocene{
				vsota += float64(ocena)
			}
			povp = vsota / float64(len(s.Ocene))
		}

		fmt.Printf("%s %s: povprena ocena %.1f", s.Ime, s.Priimek, povp)
		if povp >= 9{
			fmt.Printf(" -> Odlicen student!\n")
		}else if povp >= 6 && povp < 9{
			fmt.Printf(" -> Povprecen student\n")
		}else{
			fmt.Printf(" -> Neuspesen student\n")
		}
	}
}