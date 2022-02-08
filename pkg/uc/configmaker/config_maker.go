package configmaker

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type GenerateConfigOptions struct {
	OutputPath   string
	TemplatePath string
	TemplateData map[string]interface{}
}

type GenerateConfigOutput struct {
	OutputFilePath string
}

func GenerateConfig(opts GenerateConfigOptions) (*GenerateConfigOutput, error) {

	templateBytes, err := ioutil.ReadFile(opts.TemplatePath)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New("config").Parse(string(templateBytes))
	if err != nil {
		return nil, err
	}

	//outputFile, err := ioutil.TempFile("", opts.OutputPath)
	//if err != nil {
	//	return nil, err
	//}

	directory := filepath.Dir(opts.OutputPath)

	//pwd, err := os.Getwd()
	//if err != nil {
	//	return nil, err
	//}
	//absoultePath := filepath.Join(pwd, directory)
	err = os.MkdirAll(directory, 0777)
	if err != nil {
		return nil, err
	}

	outputFile, err := os.Create(opts.OutputPath)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(outputFile, opts.TemplateData)
	if err != nil {
		return nil, err
	}

	return &GenerateConfigOutput{
		OutputFilePath: outputFile.Name(),
	}, nil
}
