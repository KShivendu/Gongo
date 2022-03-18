package core

import (
	"bytes"
	"os"
	ttemplate "text/template"
)

type Template struct {
	filepath string
}

type RequestContext struct {
	Data interface{}
	Crsf string
}

func generateCsrf() string {
	// TODO: Generate CSRF Token
	return "CSRF"
}

func (template *Template) render(context interface{}) Response {
	// Add CSRF token to the interface
	// data map[string]interface{}

	// TODO(fix): CSRF isn't showing up in the rendered template
	rctx := RequestContext{Data: context, Crsf: generateCsrf()}

	// Read the template
	var buffer bytes.Buffer
	// TODO: Don't allow files outside templates directory
	// TODO: Add relative path of gongo instead of putting './'
	dat, err := os.ReadFile("./templates/" + template.filepath)
	check(err)

	// Use the template
	t := ttemplate.Must(ttemplate.New("test").Parse(string(dat)))
	t.Execute(&buffer, rctx)

	// Set content type as text/html and return the rendered html
	return HttpResponse(buffer.String())
}
