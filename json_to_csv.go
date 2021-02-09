package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/fatih/stopwatch"
	"io/ioutil"
	"os"
)

// SET CSV Structure
// `json:"example"` for the corresponding json field
type Ticket struct {
	Titel        string `json:"title"`
	Beschreibung string `json:"description"`
}

func main() {

	s := stopwatch.Start(0)

	jsonDataFromFile, err := ioutil.ReadFile("./dump3.json")
	if err != nil {
		fmt.Println(err)
	}

	var jsonData []Ticket
	err = json.Unmarshal([]byte(jsonDataFromFile), &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	csvFile, err := os.Create("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, usance := range jsonData {
		var row []string
		row = append(row, usance.Titel)
		row = append(row, usance.Beschreibung)
		writer.Write(row)
	}

	d := s.ElapsedTime()
	fmt.Printf("Elapsed Time: %s", d)
	writer.Flush()
}
