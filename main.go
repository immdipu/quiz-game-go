package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main(){

	fmt.Println("This is a quiz game ")
	var CSVFileName  = flag.String("csv","problem.csv", "A CSV file in the format of question, answer")
	flag.Parse();
	file, err := os.Open(*CSVFileName)
	if err !=nil{
		fmt.Printf("Failed to open csv file %s\n", *CSVFileName);
		os.Exit(1)
	}
	r := csv.NewReader(file);
	lines, err := r.ReadAll()

	if err !=nil{
		fmt.Printf("Something went wrong, couldn't open the file: %s", *CSVFileName)
		os.Exit(1)
	}
	fmt.Println(lines)


}