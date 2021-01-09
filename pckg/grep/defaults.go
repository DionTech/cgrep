package grep

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/DionTech/stdoutformat"
)

type Template struct {
	Expression string
}

type Templates map[string]Template

func SaveExpression(name string, expression string) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		stdoutformat.Error(err)
	}

	dir := homeDir + "/cgrep/"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)

		if err != nil {
			stdoutformat.Fatalf("cannot create output dir %s", dir)
		}
	}

	fileName := dir + "templates.json"

	file, err := os.Open(fileName)

	if err != nil {
		file, err = os.Create(fileName)

		if err != nil {
			stdoutformat.Fatalf("cannot create output dir %s", dir)
		}
	}

	defer file.Close()

	templates := make(Templates, 0)

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &templates)

	templates[name] = Template{Expression: expression}

	jsonVar, err := json.Marshal(templates)
	err = ioutil.WriteFile(fileName, jsonVar, os.ModePerm)

	if err != nil {
		stdoutformat.Fatalf("cannot write file: %s", err)
	}

}

func LoadExpression(name string) string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		stdoutformat.Error(err)
	}

	fileName := homeDir + "/cgrep/templates.json"
	file, err := os.Open(fileName)

	if err != nil {
		stdoutformat.Error(err)
	}

	defer file.Close()

	templates := make(Templates, 0)

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &templates)

	template, exists := templates[name]

	if !exists {
		stdoutformat.Error(errors.New(name + " not exists"))
	}

	return template.Expression
}

func loadTemplates() {

}
