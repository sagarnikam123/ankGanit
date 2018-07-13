package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type paddhat struct {
	Name      string
	Example   string
	PaddhatNo int
}

func main() {
	methods := []paddhat{
		{Name: "Addition", Example: "5 + 5", PaddhatNo: 1},
		{Name: "Substration", Example: "30 - 10", PaddhatNo: 2},
		{Name: "Multiplication", Example: "5 x 5", PaddhatNo: 3},
		{Name: "Division", Example: "40/8", PaddhatNo: 4},
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
{{ "Method No:" | faint }}	{{ .PaddhatNo }}`,
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

}
