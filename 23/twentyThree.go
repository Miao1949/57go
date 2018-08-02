package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

const CarSilentQuestion = "Is the car silent when you turn the key? "
const BatteryTerminalsCorrodedQuestion = "Are the battery terminals corroded?"
const BatteryTerminalsCorrodedSolution = "Clean terminals and trystarting again."
const BatteryTerminalsNotCorrodedSolution = "Replace cables and try again."

const CarMakesClickingNoiseQuestion = "Does the car make a clicking noise? "
const CarMakesClickingNoiseSolution = "Replace the battery."

const CarCranksButDoesNotStartQuestion = "Does the car crank up but fail to start? "
const CarCranksButDoesNotStartSolution = "Check spark plug connections."

const EngineStartsAndThenDiesQuestion = "Does the engine start and then die? "
const CarHasFuelInjectionQuestion = "Does your car have fuel injection? "
const CarDoesNotHaveFuelInjectionSolution = "Check to ensure the choke is opening and closing."
const CarHasFuelInjectionSollution =  "Get it in for service."


func getYesOrNowInput(msg string) (answer bool) {
	yesOrNoOptions :=  make(map[string]bool)
	yesOrNoOptions["y"] = true
	yesOrNoOptions["Y"] = true
	yesOrNoOptions["n"] = true
	yesOrNoOptions["N"] = true
	retString := getRestrictedStringInput(msg, yesOrNoOptions)
	answer = retString == "y" || retString == "Y"
	return
}

func getRestrictedStringInput(msg string, allowedResponses map[string]bool) (retString string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) > 0 && allowedResponses[input]{
			retString = input
			done = true
		}
	}
	return
}


func main() {
	carSilent := getYesOrNowInput(CarSilentQuestion)
	if carSilent {
		batteryTerminalsCorroded := getYesOrNowInput(BatteryTerminalsCorrodedQuestion)
		if batteryTerminalsCorroded {
			fmt.Println(BatteryTerminalsCorrodedSolution)
		} else {
			fmt.Println(BatteryTerminalsNotCorrodedSolution)
		}
	} else {
		carMakesClickingNoise := getYesOrNowInput(CarMakesClickingNoiseQuestion)
		if carMakesClickingNoise {
			fmt.Println(CarMakesClickingNoiseSolution)
		} else {
			carCranksButDoesNotStart := getYesOrNowInput(CarCranksButDoesNotStartQuestion)
			if carCranksButDoesNotStart {
				fmt.Println(CarCranksButDoesNotStartSolution)
			} else {
				engineStartsButThenDies := getYesOrNowInput(EngineStartsAndThenDiesQuestion)
				if engineStartsButThenDies {
					carHasFuelInjection := getYesOrNowInput(CarHasFuelInjectionQuestion)
					if carHasFuelInjection {
						fmt.Println(CarHasFuelInjectionSollution)
					} else {
						fmt.Println(CarDoesNotHaveFuelInjectionSolution)
					}
				} else {
					fmt.Println("Soory, does not know what to do!")
				}
			}
		}
	}
}
