package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ActionType int

const (
	actionGuardChange ActionType = 0
	actionAsleep      ActionType = 1
	actionAwoke       ActionType = 2
	actionOther       ActionType = 3
)

type GuardData struct {
	Action ActionType
	Guard  string
	Minute int
}

func main() {

	filename := "sorted.txt"
	lines, err := getInput(filename)
	if err != 0 {
		fmt.Println("there was an error reading the file")
		return
	}

	shiftData := make(map[string][]int)
	// guardData["#10"][0] = 1

	// elem, ok := guardData["#10"]
	// if ok == false {
	// 	fmt.Println("The guard doesn't exist yet.")
	// } else {
	// 	fmt.Println(elem)
	// }

	// [GuardID][array of ints (0-59 for each minute)]
	var currentGuard string
	// var currentState ActionType
	var lastTime int

	for _, line := range lines {
		guardData := parseLine(line)

		switch guardData.Action {

		case actionGuardChange:
			// fmt.Println("Changing guards...")
			// fmt.Println("Guard", guardData.Guard, "is on duty.")
			currentGuard = guardData.Guard
			// currentGuard

		case actionAsleep:
			// fmt.Println("Guard is asleep")
			lastTime = guardData.Minute

		case actionAwoke:
			// fmt.Println("Guard woke up.")
			times, ok := shiftData[currentGuard]

			if !ok {
				shiftData[currentGuard] = make([]int, 60)
				times = shiftData[currentGuard]
			}
			//keep track of the minutes asleep from lastTime to guardData.Minute which is the current time
			for lastTime < guardData.Minute {
				times[lastTime]++
				lastTime++
			}
		}
		// currentState = guardData.Action
		//fmt.Println(currentGuard, currentState, guardData.Minute)
	}

	// find which guard slept the most minutes
	biggestSleeper, mostMinutesSlept := getBiggestSleeper(shiftData)

	worstGuardData := shiftData[biggestSleeper]

	most, _ := getMinuteSleptMost(worstGuardData)
	fmt.Printf("Guard %s slept the most during minute %d.\n", biggestSleeper, most)
	// fmt.Printf("He slept %d times that minute.\n", mostMinutesSlept)
	fmt.Printf("He slept a total of %d minutes.\n", mostMinutesSlept)

	// find the guard that slept the most asleep on the same minute
	singleMostMinute := 0
	singleGuardID := ""
	mostSlept := 0

	for guardID, minutes := range shiftData {
		min, amountSlept := getMinuteSleptMost(minutes)
		if amountSlept > mostSlept {
			mostSlept = amountSlept
			singleMostMinute = min
			singleGuardID = guardID
		}
	}
	fmt.Printf("Guard %s was asleep during minute %d %d times.\n", singleGuardID, singleMostMinute, mostSlept)
	// fmt.Println(shiftData)
}

func getBiggestSleeper(shiftData map[string][]int) (string, int) {
	most := 0
	biggestSleeper := ""

	for guardID, minutes := range shiftData {
		minutesSlept := addUpMinutes(minutes)
		if minutesSlept > most {
			most = minutesSlept
			biggestSleeper = guardID
		}
	}
	return biggestSleeper, most
}

func addUpMinutes(sleepData []int) int {
	sum := 0

	for _, minutes := range sleepData {
		sum += minutes
	}
	return sum

}

func getMinuteSleptMost(minutesSlept []int) (int, int) {
	mostMinutes := 0
	most := 0
	mostAmountSlept := 0

	for min, amountSlept := range minutesSlept {
		if amountSlept > mostMinutes {
			mostMinutes = amountSlept
			mostAmountSlept = amountSlept
			most = min
		}
	}
	return most, mostAmountSlept
}
func getInput(fileName string) ([]string, int) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading the file")
		return nil, 1
	}
	stringData := string(data)
	stringData = strings.Trim(stringData, "\n ")
	lines := strings.Split(stringData, "\n")
	return lines, 0
}

func parseLine(line string) GuardData {
	var guard GuardData

	if strings.Contains(line, "Guard") == true {
		guard.Action = actionGuardChange
		data := strings.Split(line, " ")
		guard.Guard = data[3]
	} else if strings.Contains(line, "falls") {
		guard.Action = actionAsleep
	} else if strings.Contains(line, "wakes") {
		guard.Action = actionAwoke
	} else {
		guard.Action = actionOther
	}

	guard.Minute = getTimeFromString(line)

	return guard
}

func getTimeFromString(line string) int {
	s := strings.Split(line, " ")
	s = strings.Split(s[1], ":")
	min, err := strconv.Atoi(strings.Trim(s[1], "]"))
	if err != nil {
		fmt.Println("Error doing the conversion")
		return 0
	}
	return min
}
