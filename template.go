package thebrief

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"strings"
)

// PageValues holds the value for the template
type PageValues struct {
	Title       string
	ScriptFiles []string
	StyleFiles  []string
	Content     interface{}
}

type assetmap map[string]string

// name looks for a template named "templates/<name>.tmpl"
func renderWithPartials(wr io.Writer, name string, title string, content interface{}) error {
	pv := &PageValues{}
	pv.Title = title
	pv.Content = content

	// Get js and css files from assets.json
	assetJSON, err := ioutil.ReadFile("./www/assets.json")
	if err != nil {
		return err
	}

	var hashedAssets assetmap
	err = json.Unmarshal(assetJSON, &hashedAssets)
	if err != nil {
		return err
	}

	for _, v := range hashedAssets {
		if strings.HasSuffix(v, "js") {
			pv.ScriptFiles = append(pv.ScriptFiles, v)
		} else if strings.HasSuffix(v, "css") {
			pv.StyleFiles = append(pv.StyleFiles, v)
		}
	}

	// Load allTemplates and template with name
	var allTemplates []string
	files, err := ioutil.ReadDir("./templates/partials")
	if err != nil {
		return err
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".tmpl") {
			allTemplates = append(allTemplates, "./templates/partials/"+filename)
		}
	}
	allTemplates = append(allTemplates, "./templates/"+name+".tmpl")

	templates, err := template.ParseFiles(allTemplates...)
	if err != nil {
		return err
	}

	// Lookup template `name`
	page := templates.Lookup(name)
	if page == nil {
		return errors.New("template not found with name: " + name)
	}

	// Execute
	return page.ExecuteTemplate(wr, name, pv)
}
