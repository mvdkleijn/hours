package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Create a map to store the responses
	responses := make(map[string]string)

	// Create a scanner to read input from the user
	scanner := bufio.NewScanner(os.Stdin)

	// Explanation
	fmt.Printf("A standard working day is assumed to be 8 hours.\nA standard week 40 hours.\nWeekends are assumed to be free time.\n\n")
	fmt.Printf("When answering the following questions, assume a full-time, 40 hour per week scenario:\n\n")

	// Ask the four questions and store the responses
	fmt.Print("Number of 8-hour PUBLIC holidays? ")
	scanner.Scan()
	responses["publicHolidaysPerYear"] = scanner.Text()

	fmt.Print("Number of 8-hour PERSONAL holidays? ")
	scanner.Scan()
	responses["personalHolidaysPerYear"] = scanner.Text()

	fmt.Print("Number of extra hours used for sickness, extra hours off, etc? ")
	scanner.Scan()
	responses["extraFreeHours"] = scanner.Text()

	fmt.Print("What is the part-time percentage you want to apply? ")
	scanner.Scan()
	responses["partTimePercentage"] = scanner.Text()

	fmt.Print("How many days of work per week? (4 for a 36 hour, 4x9 week for example) ")
	scanner.Scan()
	responses["personalDaysPerWeek"] = scanner.Text()

	fmt.Print("How many hours of work per day? (9 for a 36 hour, 4x9 week for example) ")
	scanner.Scan()
	responses["personalHoursPerDay"] = scanner.Text()

	fmt.Print("What is the rate for over target hours? € ")
	scanner.Scan()
	responses["overTargetCompensation"] = scanner.Text()

	fmt.Printf("\nThank you for your responses!\n\n")

	if len(os.Args) < 1 {
		fmt.Println("Usage: main.go")
		return
	}

	publicHolidaysPerYear, err := strconv.ParseFloat(responses["publicHolidaysPerYear"], 64)
	if err != nil {
		fmt.Println("Invalid # of public holidays provided.")
		return
	}

	personalHolidaysPerYear, err := strconv.ParseFloat(responses["personalHolidaysPerYear"], 64)
	if err != nil {
		fmt.Println("Invalid # of personal holidays provided.")
		return
	}

	extraFreeHours, err := strconv.ParseFloat(responses["extraFreeHours"], 64)
	if err != nil {
		fmt.Println("Invalid sickdays (in hours) / extra free hours provided.")
		return
	}

	partTimePercentage, err := strconv.ParseFloat(responses["partTimePercentage"], 64)
	if err != nil {
		fmt.Println("Invalid part-time percentage provided.")
		return
	}

	personalHoursPerDay, err := strconv.ParseFloat(responses["personalHoursPerDay"], 64)
	if err != nil {
		fmt.Println("Invalid personalHoursPerDay provided.")
		return
	}

	personalDaysPerWeek, err := strconv.ParseFloat(responses["personalDaysPerWeek"], 64)
	if err != nil {
		fmt.Println("Invalid personalDaysPerWeek provided.")
		return
	}

	overTargetCompensation, err := strconv.ParseFloat(responses["overTargetCompensation"], 64)
	if err != nil {
		fmt.Println("Invalid overTargetCompensation provided.")
		return
	}

	var weeksPerYear, workableDaysPerYear, workableHoursPerYear float64

	// Target hours per year
	targetHoursPerYear := 1600.0

	// Number of weeks in a year
	weeksPerYear = 52

	// Number of workable days in a year (not including holidays)
	workableDaysPerYear = weeksPerYear*5 - publicHolidaysPerYear

	// Number of workable hours in a year
	workableHoursPerYear = workableDaysPerYear * 8

	// Adjust the number of workable hours based on the part-time percentage
	workableHoursPerYear *= (partTimePercentage / 100)

	// Calculate number of personal holiday hours
	personalHolidayHoursPerYear := personalHolidaysPerYear * 8

	// Adjust personal holidays based on part-time percentage
	parttimePersonalHolidaysPerYear := personalHolidaysPerYear * (partTimePercentage / 100)

	//Adjust number of personal holiday hours based on part-time percentage
	personalHolidayHoursPerYear *= (partTimePercentage / 100)

	// Adjust the number of workable hours based on personal holiday hours
	workableHoursPerYear = workableHoursPerYear - personalHolidayHoursPerYear

	// Adjust workable days per year based on personal holidays
	parttimeWorkableDaysPerYear := workableDaysPerYear - parttimePersonalHolidaysPerYear

	// Target hours adjusted for part-time percentage
	parttimeTargetHoursPerYear := targetHoursPerYear * (partTimePercentage / 100)

	// Doublecheck: what are the hours I should book?
	bookableHours := (weeksPerYear * 40) * (partTimePercentage / 100)

	fmt.Printf("Given %.2f days per week with %.2f public holidays & %.2f personal holidays at a %.2f part-time percentage\n(Use 8-hour days.)\n\n", 5.0, publicHolidaysPerYear, personalHolidaysPerYear, partTimePercentage)

	fmt.Printf("In a year there are: \n\n")
	fmt.Printf("Public holidays: %.2f\n", publicHolidaysPerYear)
	fmt.Printf("Personal holidays: %.2f (%.2f weeks at %1.f hrs/day, %1.f days/wk)\n", parttimePersonalHolidaysPerYear, parttimePersonalHolidaysPerYear/personalDaysPerWeek, personalHoursPerDay, personalDaysPerWeek)
	fmt.Printf("Sickdays/Extra holidays: %.2f (%.2f weeks at %1.f hrs/day, %1.f days/wk)\n", extraFreeHours/personalHoursPerDay, (extraFreeHours/personalHoursPerDay)/personalDaysPerWeek, personalHoursPerDay, personalDaysPerWeek)
	fmt.Printf("Workable 8-hour days: %.2f (excl extra free time)\n", parttimeWorkableDaysPerYear)
	fmt.Printf("Workable %.0f-hour days: %.2f (incl extra free time)\n", personalHoursPerDay, (workableHoursPerYear/personalHoursPerDay)-(extraFreeHours/personalHoursPerDay))
	fmt.Printf("Workable hours: %.2f\n", workableHoursPerYear-extraFreeHours)
	fmt.Printf("Bookable hours: %.2f\n", bookableHours)

	fmt.Printf("\nFor your target hours this means:\n\n")
	fmt.Printf("Target hours: %.2f\n", parttimeTargetHoursPerYear)
	fmt.Printf("Target hours per quarter: %.2f\n", parttimeTargetHoursPerYear/personalDaysPerWeek)
	fmt.Printf("Possible over target hours: %.2f\n", workableHoursPerYear-extraFreeHours-parttimeTargetHoursPerYear)
	fmt.Printf("Possible over target earnings: € %.2f\n\n", (workableHoursPerYear-extraFreeHours-parttimeTargetHoursPerYear)*overTargetCompensation)
}
