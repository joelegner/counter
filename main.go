package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
    "github.com/ahmetb/go-cursor"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	counter := 0

    // Hide the cursor
	fmt.Print(cursor.Hide())

	fmt.Println("counter by Joe Legner")
	fmt.Println("(with lots of help from ChatGPT)")
    fmt.Println("Escape: reset to 0")
    fmt.Println("Spacebar or Right Arrow: +1")
    fmt.Println("Backspace or Left Arrow: -1")
    fmt.Println("q or CTRL+C: Quit")

	for {
        if counter < 0 { counter = 0 }
        fmt.Printf("Counter: %d            \r", counter)
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowRight {
			counter++
		} else if key == keyboard.KeySpace {
			counter++
		} else if key == keyboard.KeyBackspace2 {
			counter--
		} else if key == keyboard.KeyArrowLeft {
			counter--
		} else if key == keyboard.KeyEsc {
			counter = 0
		} else if char == 'q' {
	        fmt.Print(cursor.Show())
            break
		} else if key == keyboard.KeyCtrlC {
	        fmt.Print(cursor.Show())
			break
		}

	}

	fmt.Println("\nDone counting.")
	os.Exit(0)
}

