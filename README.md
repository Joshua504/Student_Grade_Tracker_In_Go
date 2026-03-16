# Student_Grade_Tracker_In_Go

# what i learnt 
- reading the input from the terminal using the **bufio.newreader(os.stin)** and **reader.readstring('\n')**
- removing space from strings using **strings.trimspace(nameinput)**
- how to create an empty map using the **make** keyword **storage := make(map[string]int)**
- how to add to a slice using the **append** method **sliceOfMap = append(sliceOfMap, student{i, v})**
- how to sort a map in go using the **struct** and the **sort.slice** keyword
```
type student struct {
	name string
	grade int
}

sort.Slice(sliceOfMap, func(k, v int) bool{
		return sliceOfMap[k].name < sliceOfMap[v].name
	})
```
- how to concate or build a string using the **strings.Builder**, **builder.WriteString(line)** and **builder.String()** method
- how to write to a file using **os.WriteFile("result.txt", []byte(finalResult), 0644)**
