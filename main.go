package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	service := NewService(TextAnalyzer{})
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter text (empty line to analyze, 'exit' to quit):")

		var lines []string

		for scanner.Scan() {
			line := scanner.Text()

			if len(lines) == 0 && line == "exit" {
				fmt.Println("Goodbye!")
				return
			}

			if line == "" {
				break
			}

			lines = append(lines, line)
		}

		if len(lines) == 0 {
			continue
		}

		input := strings.Join(lines, "\n")

		task, err := service.CreateTask(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if err := service.ProcessTask(task.ID); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %+v\n", *task.Result)
		fmt.Printf("State: %s\n\n", task.State)
	}
}
