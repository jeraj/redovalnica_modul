package redovalnica

import "fmt"

type Student struct {
    ime     string
    priimek string
    ocene   []int
}

var Redovalnica = make(map[string]Student)

func izpisRedovalnice(studenti map[string]Student){
	fmt.Println("REDOVALNICA:")
	for kljuc,s := range studenti {
		fmt.Printf("%s - %s %s: %v\n", kljuc, s.ime, s.priimek, s.ocene)
	}
}

func dodajOceno(redovalnica map[string]Student, vpisnaStevilka string, ocena int){
	if ocena < 0 || ocena > 10{
		fmt.Println("Neveljavna ocena!")
		return
	}

	if s, ok := redovalnica[vpisnaStevilka]; ok { //s je kopija structa, tega specificnega studenta
		s.ocene = append(s.ocene, ocena) //dodamo oceno tej kopiji
		redovalnica[vpisnaStevilka] = s //jo shranimo nazaj, nadomesti original
		fmt.Println("Ocena uspesno vpisana.")
	} else {
		fmt.Println("Student s tako vpisno stevilko ne obstaja!")
		return
	}
}

func povprecje(redovalnica map[string]Student, vpisnaStevilka string)float64{
	if s, ok := redovalnica[vpisnaStevilka]; ok {
		if len(s.ocene) == 0{
			return 0
		}

		vsota := 0
		for _, ocena := range s.ocene{
			vsota +=ocena
		}

		var povprecje float64 = float64(vsota) / float64(len(s.ocene))

		if povprecje < 6 {
			return 0.0
		}

		return povprecje

	} else { //student ne obstaja
		return -1
	}
}

func izpisiKoncniUspeh(studenti map[string]Student){
	for vpisna,s := range studenti {
		var povp float64 = povprecje(studenti, vpisna)
		if povp == 0{
			if len(s.ocene) == 0{
				continue
			}
			var vsota float64 = 0
			for _,ocena := range s.ocene{
				vsota += float64(ocena)
			}
			povp = vsota / float64(len(s.ocene))
		}

		fmt.Printf("%s %s: povprena ocena %.1f", s.ime, s.priimek, povp)
		if povp >= 9{
			fmt.Printf(" -> Odlicen student!\n")
		}else if povp >= 6 && povp < 9{
			fmt.Printf(" -> Povprecen student\n")
		}else{
			fmt.Printf(" -> Neuspesen student\n")
		}
	}
}