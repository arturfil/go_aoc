package day1

import (
	"fmt"
	"log"
	"sort"

	"github.com/arturfil/go_aoc/helpers"
)

func Day1() {
	totDistance := 0

	sort.Slice(ids1, func(i, j int) bool {
		return ids1[i] < ids1[j]
	})

	sort.Slice(ids2, func(i, j int) bool {
		return ids2[i] < ids2[j]
	})

	if len(ids1) != len(ids2) {
		log.Println("IDs length are not the same")
	}

	for i := range ids1 {
		distance := helpers.Abs(ids1[i] - ids2[i])
		totDistance += distance
	}

	fmt.Println(totDistance)
}

func Day1PartTwo() {
    locIds := map[int]int {}
    totScore := 0

    // put all the keys in the map
    for _, id := range ids1 {
        if _, ok := locIds[id]; ok {
            fmt.Println("Key already there")
        } else {
            locIds[id] = 0 
        }
    }

    for _, id := range ids2 {
        if _, ok := locIds[id]; ok {
            locIds[id] += 1
        }
    }

    for k, val := range locIds {
        totScore += (k * val)
    }

    fmt.Println("Tot Score:", totScore)
}


