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

	recordOne, _ := r.Read()
	recordTwo, _ := r.Read()

	fmt.Println(recordOne)
	fmt.Println(recordTwo)
}