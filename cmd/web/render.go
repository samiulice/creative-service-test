package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type templateData struct {
	StringMap            map[string]string
	IntMap               map[string]int
	FloatMap             map[string]float64
	Data                 map[string]interface{}
	CSRFToken            string
	Flash                string
	Warning              string
	Error                string
	IsAuthenticated      int
	API                  string
	CSSVersion           string
	StripeSecretKey      string
	StripePublishableKey string
	AccountType          string
}

var funcMap = template.FuncMap{
	"titleCase":      titleCase,
	"formatCurrency": formatCurrency,
	"formatDate":     formatDate,
	"userlink":       userlink,
	"sumOfIntegers":  sumOfIntegers,
	"calcPercentage": calcPercentage,
}

// titleCase returns a copy of the string s with all Unicode letters mapped to their Unicode title case
func titleCase(s string) string {
	if len(strings.Split(s, " ")) < 2 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return strings.ToTitle(s)
}

// formatCurrency returns the currency with prefix $
func formatCurrency(n int) string {
	return fmt.Sprintf("$%.2f", float64(n)/100.0)
}

// formatDate returns Date in a specific format
func formatDate(t time.Time, format string) string {
	return t.Format(format)
}

// userlink returns link for user
func userlink(str string) string {
	return str[:len(str)-1]
}

// sumOfIntegers returns the sum of a slice of int numbers
func sumOfIntegers(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// calcPercentage returns the percentage of a to b
func calcPercentage(numbers ...int) string {
	var a, b float32
	if len(numbers) == 2 {
		a, b = float32(numbers[0]), float32(numbers[1])
	} else {
		a, b = float32(numbers[0]), float32(numbers[1])+float32(numbers[2])

	}
	if b <= 0 {
		return "0%"
	}
	return fmt.Sprintf("%.2f%%", (a/b)*100.00)
}

//go:embed templates/*
var templateFS embed.FS

// addDefaultData adds default variables to the templatedata
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	td.API = app.config.api
	if app.Session.Exists(r.Context(), "user_id") {
		td.IsAuthenticated = 1
	} else {
		td.IsAuthenticated = 0
	}
	if app.Session.Exists(r.Context(), "account_type") {
		td.AccountType = app.Session.GetString(r.Context(), "account_type")
	} else {
		td.AccountType = "users"
	}

	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error

	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateInMap := app.temlateCache[templateToRender]

	if templateInMap && app.config.env != "development" {
		t = app.temlateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page)
		if err != nil {
			app.errorLog.Println(err)
			return nil
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, r)
	err = t.Execute(w, td)

	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	return nil
}

func (app *application) parseTemplate(partials []string, page string) (*template.Template, error) {
	var t *template.Template
	var err error

	//Identifying Base Template path
	var baseTemplate string
	if strings.Contains(page, "admin") {
		baseTemplate = "templates/admin.layout.gohtml"
	} else if strings.Contains(page, "employee") {
		baseTemplate = "templates/employee.layout.gohtml"
	} else {
		baseTemplate = "templates/base.layout.gohtml"
	}

	//Identifying Template path that needs to be rendered
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	//build partials
	if len(partials) > 0 {
		for i, v := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partials.gohtml", v)
		}
	}

	//patterns of Templates path that is to be parsed into the single template
	var templates []string
	templates = append(templates, baseTemplate, templateToRender)
	templates = append(templates, partials...)

	t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(funcMap).ParseFS(templateFS, templates...)

	if err != nil {
		return nil, err
	}

	app.temlateCache[templateToRender] = t

	return t, nil
}
