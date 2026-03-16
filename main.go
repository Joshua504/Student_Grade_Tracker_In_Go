package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type student struct {
	name string
	grade int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	storage := make(map[string]int)

	for {

		fmt.Print("Add Student NAME: ")
		nameInput, _ := reader.ReadString('\n')
		trimName := strings.TrimSpace(nameInput)

		if trimName == "exit"{
			break
		}

		fmt.Print("Add Students GRADES: ")
		gradeInput, _ := reader.ReadString('\n')
		trimGrade := strings.TrimSpace(gradeInput)
		
		if trimGrade == "exit" {
			break
		}

		toDigit, err := strconv.Atoi(trimGrade)

		if err != nil{
			fmt.Println("not a number")
			continue
		}
		storage[trimName] = toDigit

	}
	sliceOfMap  := make([]student, 0, len(storage))

	for i, v := range storage {
		sliceOfMap = append(sliceOfMap, student{i, v})
	}

	sort.Slice(sliceOfMap, func(k, v int) bool{
		return sliceOfMap[k].name < sliceOfMap[v].name
	})

	var builder strings.Builder

	for _, i := range sliceOfMap {
		line := fmt.Sprintf("%s: %d\n", i.name, i.grade)
		builder.WriteString(line)
	}

	fmt.Println("check the result.txt file for the RESULT")
	finalResult := builder.String()

	os.WriteFile("result.txt", []byte(finalResult), 0644)
}