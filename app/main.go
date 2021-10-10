package main //required package

//Turned off Go modules by using: $env -w  GO111MODULE=off
//Installed Git on local machine
//Installed GCC for Go on loacal machine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../timer" //Package import - local package
)

type Brewer struct { //creating a new struct type named brewer
	Name         string
	Yield        int64 //For easier conversions
	Grind        string
	RatioMedium  string
	RatioStrong  string
	BrewTime     float32
	Instructions string
}

var Frenchpress = Brewer{ //frenchpress is of the brewer type
	Name:         "French Press",
	Yield:        8,
	Grind:        "Course",
	RatioMedium:  "1:12",
	RatioStrong:  "1:10",
	BrewTime:     4,
	Instructions: "\n•Add the ground coffee to the beaker.\n•Add boiling water in the desired ratio.\n•Stir briefly and let brew for 4 minutes.\n•Slowly press filter down and then serve.",
}
var Conedripper = Brewer{
	Name:         "Pour-over Cone Dripper",
	Yield:        2,
	Grind:        "Medium",
	RatioMedium:  "1:17",
	RatioStrong:  "1:15",
	BrewTime:     3.5,
	Instructions: "\n•Put filter in cone and place on top of coffee mug or glass jug.\n•Measure coffee and water ratio. Add the coffee grounds to the filter.\n•Boil water and then let it sit 2-3 minutes after boiling.\n•If the water is too hot it will result in bitter coffee.\n•Pour the hot water into the cone until coffee becomes saturated.\n•Let bloom for 30 seconds.\n•Continue to pour water, in a circular motion, until cone is filled.•\nThe slower you pour the hot water, the better.",
}
var Chemex = Brewer{
	Name:         "Chemex",
	Yield:        8,
	Grind:        "Medium-Course",
	RatioMedium:  "1:15",
	RatioStrong:  "1:14",
	BrewTime:     3.5,
	Instructions: "\n•Put filter in chemex.\n•Measure coffee and water ratio. Add the coffee grounds to the filter.\n•Boil water and then let it sit 1-2 minutes after boiling.\n•Pour the hot water into the cone until coffee becomes wet and saturated.\n•Let bloom for 30 seconds.\n•Continue to pour water, in a circular motion, until cone is filled.•\nThe slower you pour the hot water, the better.",
}
var Aero = Brewer{
	Name:         "Aeropress",
	Yield:        1,
	Grind:        "Fine",
	RatioMedium:  "1:16",
	RatioStrong:  "1:15",
	BrewTime:     1.5,
	Instructions: "\n•Measure coffee and water ratio.\n•Put filter in the plastic cup.\n•Boil water and then let it sit 1-2 minutes after boiling.\n•Wet filter and cup with hot water.\n•Assemble the aero and add the coffee grounds to the aeropress. The numbers on the aero should be upside down when facing it.\n•Pour some hot water on the grounds gently to wet them evenly.\n•Start the timer.\n•Let bloom for 30 seconds.\n•Fill the rest of the chamber with the remaining water. After 1 minute stir gently.\n•Fasten the cap and begin applying downward pressure into mug.\n•Your coffee is fully brewed once it begins to make a hissing sound.",
}
var Mokasmall = Brewer{
	Name:         "200ml Moka Pot",
	Yield:        3,
	Grind:        "Fine",
	RatioMedium:  "1:7",
	RatioStrong:  "1:6",
	BrewTime:     3.5,
	Instructions: "\n•Boil water and then let it sit 1-2 minutes after boiling.\n•Measure coffee and water ratio with scale.\n•Fill bottom chamber with hot water. Do not fill past the pressure seal.\n•Insert the filter cup and add the ground coffee. Do not compact the coffee grounds.\n•Tap the pot gently to spread coffee grounds.\n•Screw on the top chamber and place on gas stove with medium heat.\n•Start timer and wait for boil.\n•Once coffee starts rapidly filling the top chamber, remove from heat.\n•Wait for coffee to stop filling the top chamber.\n•Serve as espresso shots or ad water to dilute for Americano.",
}
var Moka = Brewer{
	Name:         "300ml Moka Pot",
	Yield:        6,
	Grind:        "Fine",
	RatioMedium:  "1:7",
	RatioStrong:  "1:6",
	BrewTime:     4,
	Instructions: "\n•Boil water and then let it sit 1-2 minutes after boiling.\n•Measure coffee and water ratio with scale.\n•Fill bottom chamber with hot water. Do not fill past the pressure seal.\n•Insert the filter cup and add the ground coffee. Do not compact the coffee grounds.\n•Tap the pot gently to spread coffee grounds.\n•Screw on the top chamber and place on gas stove with medium heat.\n•Start timer and wait for boil.\n•Once coffee starts rapidly filling the top chamber, remove from heat.\n•Wait for coffee to stop filling the top chamber.\n•Serve as espresso shots or ad water to dilute for Americano.",
}
var Siphon = Brewer{
	Name:         "Siphon",
	Yield:        3,
	Grind:        "Medium",
	RatioMedium:  "1:18",
	RatioStrong:  "1:16",
	BrewTime:     4,
	Instructions: "\n•Attach the filter and hook clip to bottom of funnel.\n•Measure coffee and water ratios with scale.\n•Fill bottom chamber with the water.\n•Place bottom chamber on the heat source (butane burner) and turn to maximum.\n•Start timer.\n•Place the funnel loosely on the bottom chamber\n•Start timer and wait for boil.\n•Attach the funnel to the bottom chamber securely and tight.\n•Wait for top chamber to fill with water, then add coffee grounds.\n•Brew for the next (~90s) and stir every 45s.\n•Turn down heat and wait for drawdown.",
}
var Numbers = []*Brewer{ //Using a pointer for the array
	&Aero,
	&Chemex,
	&Conedripper,
	&Frenchpress,
	&Moka,
	&Mokasmall,
	&Siphon}

func reader() string { // function for reading in input from user
	reader := bufio.NewReader(os.Stdin) //getting input from user using reader
	input, err := reader.ReadString('\n')
	if err != nil { //checks if there was an error with the input
		fmt.Println("An error occured while reading input. Please try again", err)
	}
	input = strings.Replace(input, "\r\n", "", -1) //Windows version of strings.Replace
	return input
}
func menuMain() {

	for { //infinite loop. Like a while loop.
		var input string

		fmt.Println("\n☕ Welcome to Brewer ☕")
		fmt.Println("\nHere you can learn how to get the best taste out of your manual coffee maker!")
		fmt.Println("Please ENTER the most appropriate option:")
		fmt.Println("1: Brewer settings\n2: Customise Brewer settings\n3: Timer\n4: Exit")
		fmt.Print("\nInput-> ")

		input = reader()

		switch input { //switch for menu selection
		case "1":
			menuOne()
		case "2":
			menuTwo(0)
		case "3":
			menuThree()
		case "4":
			os.Exit(0)
		default:
			fmt.Println("\nInvalid selection.")
		}
	}
}

func menuOne() {
	//declaring common variables
	var input string
	var selection int64
	//defining different brewer types

	fmt.Println("\nPlease ENTER the most appropriate option:")
	fmt.Printf("1: %v\n2: %v\n3: %v\n4: %v\n5: %v\n6: %v\n7: %v\n\n8: Go back\n",
		Aero.Name, Chemex.Name, Conedripper.Name, Frenchpress.Name,
		Moka.Name, Mokasmall.Name, Siphon.Name)

	fmt.Print("Input-> ")
	input = reader()
	selection, _ = strconv.ParseInt(input, 9, 12) //convert string to int

	for { //input validation
		if selection < 1 || selection > 8 {
			fmt.Print("Invalid input. Try again: ")
			input = reader()
			selection, _ = strconv.ParseInt(input, 9, 12)
			continue //start at the top of loop
		} else if selection == 8 {
			menuMain()
			break
		} else {
			break
		}
	}
	selection--
	fmt.Printf("\nSelected Brewer: %v\nYield (cups): %v\nGrind setting: %v\nCoffee to Water ratio (Medium): %v\nCoffee to water ratio (Strong): %v\nBrew time: %v minutes\n",
		Numbers[selection].Name, Numbers[selection].Yield, Numbers[selection].Grind, Numbers[selection].RatioMedium,
		Numbers[selection].RatioStrong, Numbers[selection].BrewTime)

	fmt.Println("\nFor instructions on how to use this brewer, enter HELP\nTo edit brewer settings, enter EDIT.\nTo continue, enter DONE")
	fmt.Print("\nInput-> ")

	input = reader()
	input = strings.ToUpper(input)

	if strings.Compare("HELP", input) == 0 {
		fmt.Printf("\nInstructions: %v\n", Numbers[selection].Instructions) //shows instructions
		fmt.Println("\nPlease select:\n1. Main Menu\n2. Brewer selection screen")
		fmt.Print("\nInput-> ")

		for { //checking for input. Loops until valid input is provided.
			input = reader()
			if strings.Compare("1", input) == 0 {
				menuMain()
			} else if strings.Compare("2", input) == 0 {
				menuOne()
			} else {
				fmt.Print("Invalid selection. Input: ")
				continue
			}
		}
	} else if strings.Compare("EDIT", input) == 0 {
		selection++
		menuTwo(selection)

	} else if strings.Compare("DONE", input) == 0 {
		menuMain()
	} else {
		fmt.Println("Invalid input. Returning to selection screen.")
		menuOne()
	}
}
func menuTwo(i int64) { //Brewer settings edit menu
	var input string
	fmt.Println("⚙️  Brewer editor screen ⚙️\nHere you can edit or adjust the Brewers' settings to your liking.")
	if i == 0 { //Default i = 0 if coming from the home screen. Still need to select the brewer to edit.
		fmt.Println("\nPlease ENTER the most appropriate option:")
		for x, item := range Numbers { //print names of Brewers in loop
			fmt.Printf("%v. %v\n", x+1, item.Name)
		}
		fmt.Print("Input-> ")
		input = reader()
		selection, _ := strconv.ParseInt(input, 8, 12) //convert string to int

		for { //basic input validation
			if selection < 1 || selection > 8 {
				fmt.Print("Invalid input. Try again: ")
				continue //start at the top of loop
			} else if selection == 8 {
				menuMain()
				break
			} else {
				break
			}
		}
		menuTwo(selection) //Brewer selected (i>0), next part of function can now be executed

	} else {
		fmt.Println("\nWhat property would you like to edit?\nPlease note: some values may not be edited.")
		fmt.Println("\n1. The yield\n2. The optimal grind setting (Fine/Medium/Course)\n3. The ratio for a medium brew (1:?)\n4. The ratio for a strong brew (1:?)\n\n5. Go back")
		fmt.Print("Input-> ")
		input = reader()
		selection, _ := strconv.ParseInt(input, 6, 12) //convert string to int

		for { //input validation
			if selection < 1 || selection > 5 {
				fmt.Print("Invalid input. Try again: ")
				input = reader()
				selection, _ = strconv.ParseInt(input, 6, 12)
				continue //start at the top of loop
			} else if selection == 5 {
				menuMain()
				break
			} else {
				break
			}
		}
		fmt.Println("What should the new setting be?")
		fmt.Print("Input-> ")
		input = reader()
		if selection == 1 {
			var yield, _ = strconv.ParseInt(input, 12, 64)
			Numbers[i-1].Yield = yield //selection -1 to specify the correct position in array
		} else if selection == 2 {
			Numbers[i-1].Grind = input
		} else if selection == 3 {
			Numbers[i-1].RatioMedium = input
		} else if selection == 4 {
			Numbers[i-1].RatioStrong = input
		} else {
			fmt.Println("Something went wrong. Returning to the main menu")
			menuMain()
		}
	}
}

func menuThree() { // Timer menu.

	fmt.Println("\n⏰ Timer screen ⏰\nPlease enter the required number of minutes.")
	fmt.Print("\nInput-> ")
	var minutes string = reader()
	timer.Time(minutes) //call the Time func in timer package

	fmt.Println("\nPlease select:\n1. New timer\n2. Main Menu")
	fmt.Print("\nInput-> ")
	var input string = reader()
	for {
		if strings.Compare("1", input) == 0 { //input checking
			menuThree()
		} else if strings.Compare("2", input) == 0 {
			menuMain()
		} else {
			fmt.Print("Invalid selection. Try again: ")
			continue
		}
	}
}

func main() {

	menuMain()

}
