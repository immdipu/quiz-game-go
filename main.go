package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type user struct{
	firstName string
	lastName string
	username string
}

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

	problems := parseLine(lines)

	correct := 0;
	for index, value := range problems{
		fmt.Printf("Problem #%d: %s\n", index+1, value.question)
		var answer string
		fmt.Scanf("%s\n",&answer)
		if answer == value.answer{
			fmt.Printf("Correct answer\n")
			correct++;
		}
	}

	fmt.Printf("You have answered %d answers\n", correct)

}

func parseLine(lines [][]string) []problem{
	res := make([]problem,len((lines)))
	for index, values := range lines{
		res[index]= problem{
			question: values[0],
			answer: values[1],
		}
	}
	return res;
}

type problem struct{
	question string
	answer string
}