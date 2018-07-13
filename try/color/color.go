package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Starting...")

	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")

	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")

	// Create a custom print function for convenience
	err2 := "This is error"
	red2 := color.New(color.FgRed).PrintfFunc()
	red2("Warning")
	red2("Error: %s", err2)

	// Create SprintXxx functions to mix strings with other non-colorized strings:
	yellow := color.New(color.FgYellow).SprintFunc()
	red3 := color.New(color.FgRed).SprintFunc()
	fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red3("error"))

	info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	fmt.Printf("This %s rocks!\n", info("package"))

	// Use helper functions
	fmt.Println("This", color.RedString("warning"), "should be not neglected.")
	fmt.Printf("%v %v\n", color.GreenString("Info:"), "an important message.")

	// Windows supported too! Just don't forget to change the output to color.Output
	fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))

	// Use handy standard colors
	color.Set(color.FgYellow)

	fmt.Println("Existing text will now be in yellow")
	fmt.Printf("This one %s\n", "too")

	color.Unset() // Don't forget to unset

	// You can mix up parameters
	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset() // Use it in your function

	fmt.Println("All text will now be bold magenta.")

	c2 := color.New(color.FgCyan)
	c2.Println("Prints cyan text")

	c2.DisableColor()
	c2.Println("This is printed without any color")

	c2.EnableColor()
	c2.Println("This prints again cyan...")
}
