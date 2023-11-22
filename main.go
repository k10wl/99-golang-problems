package main

import (
	"bytes"
	"flag"
	"os"
	"text/template"
	"unicode"
)

func main() {
	createProblem()
}

func createProblem() {
	problem := flag.String("problem", "", "name of problem")
	flag.Parse()
	if problem == nil {
		panic("problem name was not specified")
	}
	funcMap := template.FuncMap{
		"capitalize": capitalize,
	}
	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates.txt"))
	generateProblemFunction(t, *problem)
	generateProblemTest(t, *problem)
}

func generateProblemFunction(t *template.Template, problem string) {
	var buf bytes.Buffer
	t.ExecuteTemplate(&buf, "function", map[string]string{
		"Name": problem,
	})
	generate(problem, ".go", buf.String())
}

func generateProblemTest(t *template.Template, problem string) {
	var buf bytes.Buffer
	t.ExecuteTemplate(&buf, "test", map[string]string{
		"Name": problem,
	})
	generate(problem, "_test.go", buf.String())
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func generate(name string, prefix string, content string) {
	f, err := os.Create(name + prefix)
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}
