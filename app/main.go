package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type paddhat struct {
	Name    string
	Example string
	Sign    string
}

func main() {
	methods := []paddhat{
		{Name: "Addition", Example: "5 + 5", Sign: "+"},
		{Name: "Subtraction", Example: "30 - 10", Sign: "-"},
		{Name: "Multiplication", Example: "5 x 5", Sign: "x"},
		{Name: "Division", Example: "40/8", Sign: "/"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .Example | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Example | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- AnkGanit ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Example:" | faint }}	{{ .Example }}
{{ "Sign:" | faint }}	{{ .Sign }}`,
	}

	searcher := func(input string, index int) bool {
		paddhat := methods[index]
		name := strings.Replace(strings.ToLower(paddhat.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Which method to go ",
		Items:     methods,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	// which method selected
	selectedIndex, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Selection failed %v\n", err)
		return
	}

	//fmt.Println(objectSelected)
	fmt.Printf("You are going for %d: %s\n", selectedIndex+1, methods[selectedIndex].Name)

phirsePrompt:
	aakdeModeLabel, uttar := Aakdemode(methods[selectedIndex].Sign)
	whatIs := PerformPrompt(aakdeModeLabel, uttar)
	if whatIs {
		goto phirsePrompt
	} else {
		fmt.Print(color.YellowString("\nBad Luck:"))
		fmt.Printf(" %s%d", aakdeModeLabel, uttar)
		fmt.Println(color.RedString("\nGame Over..."))
	}

} // main

// Aakdemode -
func Aakdemode(sign string) (aakdeMode string, ans int) {
	first := rand.Intn(10)
	firstString := strconv.Itoa(first)
	second := rand.Intn(10)
	secondString := strconv.Itoa(second)

	switch sign {
	case "+":
		return firstString + " + " + secondString + " = ", first + second
	case "-":
		return firstString + " - " + secondString + " = ", first - second
	case "x":
		return firstString + " x " + secondString + " = ", first * second
	case "/":
		return firstString + " / " + secondString + " = ", first / second
	}
	return "", 0
}

// PerformPrompt -
func PerformPrompt(aakdeModeLabel string, uttar int) bool {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    aakdeModeLabel,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	i, _ := strconv.Atoi(result)
	if uttar == i {
		fmt.Printf("%v %v %d\n\n", color.GreenString("Correct:"), aakdeModeLabel, uttar)
		return true
	}
	return false
}
