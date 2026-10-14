package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	departmentsCount, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for range departmentsCount {
		if !scanner.Scan() {
			return
		}

		employeesCount, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		minAllowedTemp := 15
		maxAllowedTemp := 30

		const expectedPartsCount = 2

		for range employeesCount {
			if !scanner.Scan() {
				return
			}

			inputParts := strings.Fields(scanner.Text())
			if len(inputParts) != expectedPartsCount {
				continue
			}

			operator := inputParts[0]
			desiredTemp, _ := strconv.Atoi(inputParts[1])

			if operator == ">=" {
				if desiredTemp > minAllowedTemp {
					minAllowedTemp = desiredTemp
				}
			} else if operator == "<=" {
				if desiredTemp < maxAllowedTemp {
					maxAllowedTemp = desiredTemp
				}
			}

			if minAllowedTemp <= maxAllowedTemp {
				fmt.Println(minAllowedTemp)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
