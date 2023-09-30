package main

import (
	"fmt"
	"os"
	"video_filter/internal/app"
)

func main() {
	fmt.Println("Initializing application...")

	a := app.New()

	fmt.Println("Finish initialize application...")
	fmt.Println("Starting application...")

	var err error

	if len(os.Args) > 1 {
		if len(os.Args) > 2 {
			fmt.Println("Error: too many arguments...")
			os.Exit(-1)
		} else {
			err = a.Run(os.Args[1])
		}
	} else {
		err = a.Run("")
	}

	if err != nil {
		os.Exit(-1)
	}

	fmt.Println("Ending application...")
	os.Exit(0)
}
