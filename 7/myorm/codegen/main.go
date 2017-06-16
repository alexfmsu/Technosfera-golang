package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
)

type Counter struct {
	n int
}

func (c *Counter) Set(n int) int {
	c.n = n

	return n
}

func (c *Counter) Inc() int {
	c.n++

	return c.n
}

type Column struct {
	Name string
	Type string
	Tag  string

	ColumnName string

	Is_Primary_key bool
	Is_Username    bool
	Is_Null        bool
}

type ORM struct {
	StructName string

	TableName string

	Primary_key string
	Username    string

	Columns []Column
}

// TEMPLATE BEGINS-------------------------------------------------------------------------------------------------------------------------
var packages = `
import (
    "database/sql"
)
`

// ----------------------------------------------------------------------------------------------------------------------------------------
var SetDB = `
func SetDB(_db *sql.DB){
    db = *_db
}
`

// ----------------------------------------------------------------------------------------------------------------------------------------
var select_fields = `{{$exist := Set 0}}{{range $_, $field := .Columns}}{{if and $field.ColumnName (ne $field.Is_Primary_key true)}}{{$exist := Inc}}{{if ne $exist 1}}, {{end}}{{$field.ColumnName}}{{end}}{{end}}`
var select_values = `{{$exist := Set 0}}{{range $_, $field := .Columns}}{{if and $field.ColumnName (ne $field.Is_Primary_key true)}}{{$exist := Inc}}{{if ne $exist 1}}, {{end}}` + scan_values + `{{end}}{{end}}`
var scan_values = `{{if ne $field.Is_Null true}}&{{$tmp}}.{{$field.Name}}{{else}}&null_{{ToLower $field.Name}}{{end}}`
var SELECT = `{{range $_, $v := .Columns}}{{if eq $v.Is_Null true}}var null_{{ToLower $v.Name}} sql.Null{{Title $v.Type}}{{end}}{{end}}{{$tb := .TableName}}{{$pr := .Primary_key}}
    
    err := db.QueryRow("SELECT ` + select_fields + ` FROM {{$tb}} WHERE {{$pr}}=?", PK).Scan(` + select_values + `)
`

var FindByPK = `func ({{$tmp}}* {{.StructName}}) FindByPK(PK uint) error{
    ` + SELECT + `
    if err != nil{
        return err
    }
    {{range $_, $v := .Columns}}{{if eq $v.Is_Null true}}
    if null_{{ToLower $v.Name}}.Valid{
        {{$tmp}}.{{$v.Name}} = null_{{ToLower $v.Name}}.{{Title $v.Type}}
    }{{end}}{{end}}

    return nil
}
`

// ----------------------------------------------------------------------------------------------------------------------------------------
var if_field_creates = `{{if and (ne $field.ColumnName "") (ne $field.Is_Primary_key true)}}`

var create_prepare_values = `{{$exist := Set 0}}{{range $_, $field := .Columns}}{{if and (ne $field.ColumnName "") (ne $field.Is_Primary_key true)}}{{$exist := Inc}}{{if ne $exist 1}}, {{end}}?{{end}}{{end}}`
var create_exec_fields = `{{$exist := Set 0}}{{range $_, $field := .Columns}}` + if_field_updates + `{{$exist := Inc}}{{if ne $exist 1}}, {{end}}` + create_exec_values + `{{end}}{{end}}`
var create_exec_values = `&{{$tmp}}.{{.Name}}`

var create_prepare_fields = `{{$exist := Set 0}}{{range $_, $field := .Columns}}` + if_field_creates + `{{$exist := Inc}}{{if ne $exist 1}}, {{end}}{{$field.ColumnName}}{{end}}{{end}}`

var Create = `
func ({{$tmp}}* {{.StructName}}) Create() error{
    stmt, err := db.Prepare("INSERT INTO {{.TableName}}(` + create_prepare_fields + `) VALUES(` + create_prepare_values + `)")
    
    if err != nil{
        return err
    }
    {{range $_, $v := .Columns}}{{if and $v.ColumnName (eq $v.Is_Null true)}}
    var null_{{ToLower $v.Name}} sql.Null{{Title $v.Type}}
    
    null_{{ToLower $v.Name}}.{{Title $v.Type}} = {{$tmp}}.{{$v.Name}}
    ` +
	if_ne_zero + `
        null_{{ToLower $v.Name}}.Valid = true
    }else{
        null_{{ToLower $v.Name}}.Valid = false
    }
    {{end}}{{end}}
    
    row, err := stmt.Exec(` + create_exec_fields + `)

    if err != nil{
        return err
    }

    last_id, err := row.LastInsertId()
    
    if err != nil{
        return err
    }

    u.ID = uint(last_id)

    return nil
}
`

// ----------------------------------------------------------------------------------------------------------------------------------------
var if_field_updates = `{{if and (ne $field.Is_Primary_key true) ($field.ColumnName)}}`
var update_prepare_fields = `{{$exist := Set 0}}{{range $_, $field := .Columns}}` + if_field_creates + `{{$exist := Inc}}{{if ne $exist 1}}, {{end}}{{$field.ColumnName}}=?{{end}}{{end}}`
var update_exec_values = `{{if eq $field.Is_Null true}}&null_{{ToLower .Name}}{{else}}&{{$tmp}}.{{.Name}}{{end}}`
var update_exec_fields = `{{$exist := Set 0}}{{range $_, $field := .Columns}}` + if_field_updates + `{{$exist := Inc}}{{if ne $exist 1}}, {{end}}` + update_exec_values + `{{end}}{{end}}`

var if_ne_zero = `if u.{{$v.Name}} != {{if eq $v.Type "int32"}}int(0){{end}}{{if eq $v.Type "float"}}0{{end}}{{if eq $v.Type "string"}}""{{end}}{{if eq $v.Type "bool"}}false{{end}}{`

var Update = `
func ({{$tmp}}* {{.StructName}}) Update() error{
    stmt, err := db.Prepare("UPDATE {{.TableName}} SET ` + update_prepare_fields + ` WHERE {{.Primary_key}} = ?")
    
    if err != nil{
        return err
    }
    {{range $_, $v := .Columns}}{{if and $v.ColumnName (eq $v.Is_Null true)}}
    var null_{{ToLower $v.Name}} sql.Null{{Title $v.Type}}
    
    null_{{ToLower $v.Name}}.{{Title $v.Type}} = {{$tmp}}.{{$v.Name}}
    ` +
	if_ne_zero + `
        null_{{ToLower $v.Name}}.Valid = true
    }else{
        null_{{ToLower $v.Name}}.Valid = false
    }
    {{end}}{{end}}

    _, err = stmt.Exec(` + update_exec_fields + `, &{{$tmp}}.{{.Primary_key}})

    if err != nil{
        return err
    }

    return nil
}
`

// ----------------------------------------------------------------------------------------------------------------------------------------
var tag = `{{if ne $field.Tag ""}}` + "`" + `{{$field.Tag}}` + "`" + `{{end}}`

var orm_struct_field = ` {{range $field := .Columns}}
    {{$field.Name}} {{$field.Type}} ` + tag + `{{end}}`

var orm_struct = `type {{.StructName}} struct{` + orm_struct_field + `}
`

// ----------------------------------------------------------------------------------------------------------------------------------------
var tmpl = `{{$tmp := FirstCharLower .StructName}}` +
	`package {{ToLower .StructName}}
` + packages + `
var db sql.DB
` + orm_struct + SetDB + FindByPK + Create + Update

// TEMPLATE ENDS---------------------------------------------------------------------------------------------------------------------------

func main() {
	if len(os.Args) < 2 {
		log.Println("Error: err")

		return
	}

	filename := os.Args[1]

	// dir := path.Dir(filename)
	dir := path.Dir("../user/" + filename)

	myorm_data, err := parse(filename)

	if err != nil {
		log.Println("Cannot build ORM:", err)
	}

	var counter Counter

	funcMap := template.FuncMap{
		"ToLower":        strings.ToLower,
		"Title":          strings.Title,
		"FirstCharLower": func(s string) string { return strings.ToLower(s)[0:1] },
		"Set":            counter.Set,
		"Inc":            counter.Inc,
	}

	t := template.Must(template.New("ORM").Funcs(funcMap).Parse(tmpl))

	for k, v := range myorm_data {
		dest := dir + "/" + strings.ToLower(k) + "_myorm.go"
		// dest := "../user" + "/" + strings.ToLower(k) + "_myorm.go"

		file, err := os.Create(dest)

		if err != nil {
			log.Println(err)

			continue
		}

		t.Execute(file, v)

		file.Close()

		break
	}
}

func parse(filename string) (map[string]ORM, error) {
	fset := token.NewFileSet()

	var orm map[string]ORM
	orm = make(map[string]ORM, 0)

	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)

		return orm, err
	}

	tag_tb := regexp.MustCompile(`myorm:(.+)`)

	tag_1 := regexp.MustCompile(`myorm:\"(.+)\"`)
	tag_2 := regexp.MustCompile(`(.+):(.+)`)

	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.GenDecl:
			comment := decl.Doc.Text()

			group := tag_tb.FindStringSubmatch(comment)

			if len(group) != 2 {
				continue
			}

			table_name := group[1]

			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					name := spec.Name.String()

					columns := make([]Column, 0)

					fields := spec.Type.(*ast.StructType).Fields.List

					var primary_key string = ""
					var username string = ""

					for _, id := range fields {
						tag := id.Tag

						column := Column{
							Tag:        "",
							Name:       fmt.Sprintf("%s", id.Names[0]),
							ColumnName: fmt.Sprintf("%s", id.Names[0]),
							Type:       fmt.Sprintf("%s", id.Type),

							Is_Primary_key: false,
							Is_Username:    false,
						}

						if tag != nil {
							substr := tag_1.FindStringSubmatch(tag.Value)

							if len(substr) != 2 {
								return orm, errors.New("parse error")
							} else {
								for _, sub := range substr[1:] {
									if sub == "primary_key" {
										primary_key = fmt.Sprintf("%s", id.Names[0])
										column.Is_Primary_key = true
									}

									if sub == "-" {
										column.ColumnName = ""
									}

									if sub == "null" {
										column.Is_Null = true
									}
								}

								substr2 := tag_2.FindStringSubmatch(substr[1])
								// fmt.Println(substr2)

								if len(substr2) == 3 {
									if substr2[1] == "column" {
										username = fmt.Sprintf("%s", id.Names[0])
										column.Is_Username = true
									}

									column.ColumnName = substr2[2]
								}

								column.Tag = fmt.Sprintf("%s", substr[0])
							}
						}

						columns = append(columns, column)
					}

					orm[name] = ORM{
						StructName: strings.Title(name),

						TableName: table_name,

						Primary_key: primary_key,
						Username:    username,

						Columns: columns,
					}
				}
			}
		}
	}

	return orm, nil
}

// //go:generate ./codegen
// fmt
// Tag
