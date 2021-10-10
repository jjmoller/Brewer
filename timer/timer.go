package timer

import (
	"fmt"
	"strconv"
	"time"

	//"fyne.io/fyne/v2/app"          //GUI window
	//"fyne.io/fyne/v2/widget"       //GUI window
	"github.com/inancgumus/screen" //For clearing the terminal
)

func Time(minutes string) {
	screen.Clear()
	var X bool = true
	type clock [11]string

	clock1 := clock{
		"\n /________ /|\n",
		"|   XII   | |\n",
		"|    |    | |\n",
		"|IX  * III| |\n",
		"|    |    | |\n",
		"| ___VI___| |\n",
		"|    \\    | |\n",
		"|     \\   | |\n",
		"|      \\  | |\n",
		"|      ( )| |\n",
		"|_________|/ \n",
	}
	clock2 := clock{
		"\n /________ /|\n",
		"|   XII   | |\n",
		"|    |    | |\n",
		"|IX  * III| |\n",
		"|    |    | |\n",
		"|____VI___| |\n",
		"|    /    | |\n",
		"|   /     | |\n",
		"|  /      | |\n",
		"|( )      | |\n",
		"|_________|/\n",
	}

	//a := app.New()
	//w := a.NewWindow("Berwer Timer")
	fmt.Println("")
	fmt.Printf("\nThe timer for %v minutes has started. Please wait.\n...", minutes)
	selection, _ := strconv.ParseFloat(minutes, 32) //convert string to float for half minutes
	selection = selection * 60                      //total amount of seconds for timer

	calcDelay := time.Duration(selection) * time.Second // How many seconds need to elapse

	var endTime <-chan time.Time // signal for when timer us up
	for X {                      //loop unitil endTime
		// start the timer
		if endTime == nil {
			endTime = time.After(calcDelay)
		}
		select {
		case <-endTime: //channel
			screen.Clear()
			fmt.Println("\nDone! ")
			//w.SetContent(widget.NewLabel("DONE! The count down is over."))
			//w.ShowAndRun()
			X = false

		default:
			screen.MoveTopLeft()        //clear the terminal
			time.Sleep(time.Second * 1) // Sleep for 5000 milliseconds{
			if time.Now().Second()%2 == 0 {
				next := clock2
				fmt.Print(next, "  ")
			} else {
				next := clock1
				fmt.Print(next, "  ")
			}

		}
	}
}
