package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`

	r := csv.NewReader(strings.NewReader(in))

	r.Read()
	recordOne, _ := r.Read()
	recordTwo, _ := r.Read()
	recordThree, _ := r.Read()

	printRecord(recordOne)
	printRecord(recordTwo)
	printRecord(recordThree)
}

func printRecord(record []string) {
	fmt.Println(record)
}