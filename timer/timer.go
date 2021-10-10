package timer

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"    //GUI window
	"fyne.io/fyne/v2/widget" //GUI window
)

func Time(minutes string) {
	x := true

	a := app.New()
	w := a.NewWindow("Berwer Timer")

	fmt.Printf("\nThe timer for %v minutes has started. Please wait.\n...", minutes)
	selection, _ := strconv.ParseFloat(minutes, 32) //convert string to float for half minutes
	selection = selection * 60                      //total amount of seconds for timer

	calcDelay := time.Duration(selection) * time.Second // How many seconds need to elapse

	var endTime <-chan time.Time // signal for when timer us up

	for x { //loop unitil endTime
		// start the timer
		if endTime == nil {
			endTime = time.After(calcDelay)
		}
		select {
		case <-endTime: //channel
			fmt.Println("Done!. Close the pop-up window to continue.")
			w.SetContent(widget.NewLabel("DONE! The count down is over."))
			w.ShowAndRun()
			x = false

		default:
			time.Sleep(5000 * time.Millisecond) // Sleep for 5000 milliseconds
			fmt.Print(".")
		}
	}
}
