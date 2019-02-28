package main

import "fmt"
import "flag"
import "math/rand"
import "time"

func randomGenerator() *rand.Rand {
	seed := rand.NewSource(time.Now().UnixNano())
  return rand.New(seed)
}

func getArguments() (int, bool, bool) {
	iterations := flag.Int(
		"iterations",
		10,
		"How many times you want the game to be simulated?")
	wantsToSwitch := flag.Bool(
		"switch",
		true,
		"Do you want the suggested strategy to be applied?")
	verbose := flag.Bool(
		"verbose",
		true,
		"Do you want a full description of each iteration?")
	flag.Parse()
	return *iterations, *wantsToSwitch, *verbose
}

func randomDoorsContent() [3]string {
	content := [3]string{"goat", "goat", "goat"}
	content[randomGenerator().Intn(3)] = "car"
	return content
}

func playerChoosesDoor() int {
	return randomGenerator().Intn(3) + 1
}

func montyOpensDoor(doorsContent [3]string, playerChoice int) int {
	door := 1
	for i := 0; i < 3; i++ {
		if doorsContent[i] == "goat" && i + 1 != playerChoice {
			door = i + 1
		}
	}
	return door
}

func changeDoor(playerChoice, openDoor int) int {
	door := 1
	for i := 1; i <= 3; i++ {
		if i != playerChoice && i != openDoor {
			door = i
		}
	}
	return door
}

func simulateGame(wantsToSwitch, verbose bool) bool {
	doorsContent := randomDoorsContent()
	choice := playerChoosesDoor()
	if verbose {
		fmt.Printf("\n")
		fmt.Printf("\nPlayer selects door %d.", choice)
	}
	openDoor := montyOpensDoor(doorsContent, choice)
	if verbose {
		fmt.Printf("\nMonty opens door %d (which contains a goat).", openDoor)
		fmt.Printf("\nHe lets the player change its choice.")
	}
	if wantsToSwitch {
		choice = changeDoor(choice, openDoor)
		if verbose {
			fmt.Printf("\nPlayer switches to door number %d.", choice)
		}
	} else {
		if verbose {
			fmt.Printf("\nPlayer sticks with his initial choice.")
		}
	}
	prize := doorsContent[choice - 1]
	if verbose {
		fmt.Printf("\nPlayer won a %s.", prize)
		fmt.Printf("\n\n")
	}
	return prize == "car"
}

func printStats(iterations, numWon int) {
	fmt.Printf("\nThe game was simulated %d times", iterations)
	fmt.Printf("\nNumber of victories: %d", numWon)
	fmt.Printf("\nNumber of losses: %d", iterations - numWon)
	p := float64(numWon) / float64(iterations) * 100
	fmt.Printf("\nPercentage of victories: %f%%", p)
	fmt.Printf("\n\n")
}

func main() {
	iterations, wantsToSwitch, verbose := getArguments()
	numWon := 0
	for i := 0; i < iterations; i++ {
		if verbose {
			fmt.Printf("\n")
			fmt.Printf("\nSimulation number %d:", i + 1)
			fmt.Printf("\n")
			fmt.Printf("\n")
		}
		won := simulateGame(wantsToSwitch, verbose)
		if won {
			numWon++
		}
	}
	printStats(iterations, numWon)
}