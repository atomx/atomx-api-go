// +build ignore

package main

import (
	"os"
	"strings"
	"text/template"
)

type names struct {
	Name        string
	NameUpper   string
	Plural      string
	PluralUpper string
}

var code = template.Must(template.New("code").Parse(`package atomx

import (
	"strconv"
	"strings"
)

func (this {{.NameUpper}}) path() string {
	if this.ID > 0 {
		return "{{.Name}}/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "{{.Name}}"
	}
}

type {{.Name}}Response struct {
	Success  bool      "json:\"success\""
	Error    string    "json:\"error\""
	{{.NameUpper}} *{{.NameUpper}} "json:\"{{.Name}}\""
}

func (this {{.Name}}Response) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *{{.NameUpper}}) response() response {
	return &{{.Name}}Response{
		{{.NameUpper}}: this,
	}
}

type {{.PluralUpper}}List struct {
	List
	{{.PluralUpper}} []{{.NameUpper}} "json:\"{{.Plural}}\""
}

func (this {{.PluralUpper}}List) path() string {
	return "{{.Plural}}?" + this.str()
}

type {{.PluralUpper}} []{{.NameUpper}}

func (this {{.PluralUpper}}) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this {{.PluralUpper}}) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *{{.PluralUpper}}) Add(y {{.NameUpper}}) {
	*this = append(*this, y)
}

func (this *{{.PluralUpper}}) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type {{.NameUpper}}Relation struct {
	{{.NameUpper}}
}

func (this *{{.NameUpper}}Relation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

`))

func title(name string) string {
	return strings.Replace(strings.Title(strings.Replace(name, "_", " ", -1)), " ", "", -1)
}

func main() {
	for name, plural := range map[string]string{
		"advertiser":          "advertisers",
		"creative":            "creatives",
		"creative_attribute":  "creative_attributes",
		"creative_ban_reason": "creative_ban_reasons",
		"domain":              "domains",
		"domain_attribute":    "domain_attributes",
		"network":             "networks",
		"placement":           "placements",
		"publisher":           "publishers",
		"site":                "sites",
		"size":                "sizes",
		"user":                "users",
	} {
		if f, err := os.Create(name + "_gen.go"); err != nil {
			panic(err)
		} else if err := code.Execute(f, names{
			Name:        name,
			NameUpper:   title(name),
			Plural:      plural,
			PluralUpper: title(plural),
		}); err != nil {
			panic(err)
		} else if err := f.Close(); err != nil {
			panic(err)
		}
	}
}
