package main

import (
	"fmt"
	"os"
)

type Troll struct {
	Id                      int
	Nom                     string
	IdGuilde                int
	Race                    raceTroll
	Niveau                  int
	NbKillsTrolls           int // suicides non compris
	NbKillsMonstres         int
	ClassementKillsTrolls   int
	ClassementKillsMonstres int
	Tag                     tag
	NbKillsTK               int // attention : ceci n'a de sens que statistique, une part importante des kills est forcément mal tagguée
	NbKillsATK              int
	NbKilledByATK           int
	NbKillsTKRécents        int
	NbKillsRécents          int
}

func NewTroll(id int) *Troll {
	t := new(Troll)
	t.Id = id
	return t
}

// bâtit une classification textuelle qualitative pour un affichage dans l'extension Chrall
func (troll *Troll) ChrallClassifHtml() string {
	if troll.NbKillsTrolls == 0 {
		if troll.NbKillsMonstres == 0 {
			return "<b>NK</b>"
		}
		return "pur <b>MK</b>"
	}
	if troll.Tag == mk {
		if troll.NbKillsTrolls < 2 && troll.NbKillsMonstres > 200 {
			return "pur <b>MK</b>"
		}
		return "<b>MK</b>"
	}
	if troll.Tag == tk {
		s := ""
		if troll.NbKillsTKRécents == 0 && troll.NbKillsRécents > 10 {
			s = "Ancien "
		}
		if troll.NbKillsTrolls <= 5 {
			if troll.NbKillsMonstres > 30*troll.NbKillsTrolls {
				return s + "<b>TK</b> occasionnel"
			}
			return s + "<b>TK</b> probable"
		}
		if troll.NbKillsTrolls > 30 && troll.NbKillsMonstres < 40 {
			return s + "pur <b>TK</b>"
		}
		if troll.NbKillsTrolls >= 3*troll.NbKillsTK && troll.NbKillsTK < 20 {
			return s + "<b>TK</b> probable"
		}
		return s + "<b>TK</b>"
	}
	if troll.Tag == atk {
		if troll.NbKillsMonstres > 10*troll.NbKillsTrolls {
			return "<b>MK</b> et <b>ATK</b>"
		}
		return "<b>ATK</b>"
	}
	return "indéterminé"
}

// imprime un tableau lisible des principales caractéristiques des trolls
func PrintTrolls(trolls []*Troll, max int) {
	fmt.Printf("| %5s | %7s | %24s | %15s | %17s | %19s | %21s | %7s | %8s | %9s | %30s |\n", "#", "Num", "Nom", "Kills de trolls", "Kills de monstres", "Class. kills trolls", "Class. kills monstres", "Classif", "Kills TK", "Kills ATK", "Classif HTML")
	i := 0
	for _, troll := range trolls {
		if troll != nil {
			i++
			fmt.Printf("| %5d | %7d | %24s | %15d | %17d | %19d | %21d | %7s | %8d | %9d | %30s |\n", i, troll.Id, troll.Nom, troll.NbKillsTrolls, troll.NbKillsMonstres, troll.ClassementKillsTrolls, troll.ClassementKillsMonstres, troll.Tag.string(), troll.NbKillsTK, troll.NbKillsATK, troll.ChrallClassifHtml())
			if i == max {
				break
			}
		}
	}
}

// écrit un fichier csv des trolls
func WriteTrolls(w *os.File, trolls []*Troll, includeHeader bool) { // je ne sais pas pourquoi je ne peux pas définir w comme un *io.Writer
	if includeHeader {
		fmt.Fprintf(w, "%s;%s;%s;%s;%s;%s;%s;%s;%s;%s\n", "ID Troll", "Kills de trolls", "Kills de monstres", "Classement kills trolls", "Classement kills monstres", "Classif", "Kills TK", "Kills ATK", "Classif HTML", "Nom", "Race", "Niveau", "ID Guilde")
	}
	for _, troll := range trolls {
		if troll != nil {
			fmt.Fprintf(w, "%d;%d;%d;%d;%d;%s;%d;%d;%s;%s;%s;%d;%d\n", troll.Id, troll.NbKillsTrolls, troll.NbKillsMonstres, troll.ClassementKillsTrolls, troll.ClassementKillsMonstres, troll.Tag.string(), troll.NbKillsTK, troll.NbKillsATK, troll.ChrallClassifHtml(), troll.Nom, troll.Race.string(), troll.Niveau, troll.IdGuilde)
		}
	}
}
