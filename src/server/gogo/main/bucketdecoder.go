package main

/*
Décodeur du seau versé dans le Puit de chrall.

Ce seau est une chaine qui peut contenir des CDM, et peut-être un jour autre chose
*/

import (
	"strings"
	"fmt"
)

type BucketDecoder struct {
	Cdm []*CDM // les CDM trouvées dans l'input	
	nbResults int
}

func (bd *BucketDecoder) addCdm(cdm *CDM) {
	l := len(bd.Cdm)
	if l==cap(bd.Cdm) {
		newSlice := make([]*CDM, (l+1)*3/2)
		copy(newSlice, bd.Cdm)
		bd.Cdm = newSlice
    }
    bd.Cdm = bd.Cdm[0:l+1]
	bd.Cdm[l]=cdm
	bd.nbResults++
}


func (bd *BucketDecoder) Decode(input string) {
	lines := strings.Split(input, "\n", -1)
	
	var currentCdm *CDM
	var numLineAtCdmStart int
	for numLine, line := range lines {
		if (len(line)<4) {
			continue
		}
		if cdm := NewCdm(line); cdm!=nil {
			currentCdm = cdm
			bd.addCdm(currentCdm)
			numLineAtCdmStart = numLine		
		} else if currentCdm!=nil && numLine-numLineAtCdmStart<40 { // 40 : nombre de lignes maximal d'une cdm (à déterminer)
			name, char := AnalyseLineAsCdmChar(line)
			if name != "" {
				currentCdm.Chars[name] = char
			} else {
				fmt.Println("Ligne pas comprise : " + line);
			}
		}
	}
	
	fmt.Println("Résultats :")
	for _, cdm := range bd.Cdm {
		if cdm==nil {
			fmt.Println("cdm is nil !")
		} else {
			cdm.Print()
		}
	}
}