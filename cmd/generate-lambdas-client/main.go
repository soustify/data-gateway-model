package main

import (
	"fmt"
	pb "github.com/soustify/data-gateway-model/pkg/pb/proto"
	"os"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

// StructTemplate defines the template for the generated struct and methods.
const StructTemplate = `
package {{.PackageName}}

import (
	"context"
	"google.golang.org/grpc"
	"github.com/soustify/data-gateway-model/pkg/pb/proto"
)

var {{.ClientName}} pb.{{.InterfaceName}} = &{{.StructName}}{}

// {{.StructName}} implements the {{.InterfaceName}} interface
type {{.StructName}} struct{}

{{range .Methods}}
// {{.Name}} implements {{$.InterfaceName}}.{{.Name}}
func (c *{{$.StructName}}) {{.Name}} {{.Signature}} {
	panic("implement me")
}
{{end}}
`

// Method defines the structure for a method in the generated struct.
type Method struct {
	Name      string
	Signature string
}

// TemplateData defines the structure for the template data.
type TemplateData struct {
	PackageName   string
	StructName    string
	InterfaceName string
	ClientName    string
	Methods       []Method
}

// Função para converter camelCase para snake_case
func camelToSnake(s string) string {
	re := regexp.MustCompile("([a-z])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}

// CreateImplementation dynamically generates a Go file implementing the provided interface.
func CreateImplementation(ifaceType interface{}, packageName, structName, clientName string) {
	iface := reflect.TypeOf(ifaceType).Elem()

	if iface.Kind() != reflect.Interface {
		panic("Expected an interface type")
	}

	fmt.Printf("Generating implementation for interface: %s\n", iface.Name())

	methods := make([]Method, 0, iface.NumMethod())
	for i := 0; i < iface.NumMethod(); i++ {
		method := iface.Method(i)
		signature := generateMethodSignature(method)
		methods = append(methods, Method{
			Name:      method.Name,
			Signature: signature,
		})
	}

	data := TemplateData{
		PackageName:   packageName,
		StructName:    structName,
		InterfaceName: iface.Name(),
		ClientName:    clientName,
		Methods:       methods,
	}

	// Generate the file content using the template.
	file := camelToSnake(structName)
	generateFile(data, fmt.Sprintf("generated/lambda/%s.go", file))
}

// generateMethodSignature generates the method signature for a given method.
func generateMethodSignature(method reflect.Method) string {
	// Collect parameter types
	var in, out string
	for j := 0; j < method.Type.NumIn(); j++ { // Skipping receiver (index 0)
		_type := method.Type.In(j).String()

		if j == method.Type.NumIn()-1 && strings.HasPrefix(_type, "[]") {
			in += fmt.Sprintf("param%d %s", j, strings.Replace(_type, "[]", "...", 1))
		} else {
			in += fmt.Sprintf("param%d %s", j, _type)
		}
		if j < method.Type.NumIn()-1 {
			in += ", "
		}
	}

	for j := 0; j < method.Type.NumOut(); j++ {
		out += method.Type.Out(j).String()
		if j < method.Type.NumOut()-1 {
			out += ", "
		}
	}

	result := fmt.Sprintf("(%s) (%s)", in, out)
	return result
}

// generateFile generates the Go file based on the template and the provided data.
func generateFile(data TemplateData, outputPath string) {
	tmpl, err := template.New("struct").Parse(StructTemplate)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated file: %s\n", outputPath)
}

func main() {
	CreateImplementation((*pb.UsersServiceClient)(nil), "lambda", "usersServiceClient", "UsersClient")
	CreateImplementation((*pb.TransactionsServiceClient)(nil), "lambda", "transactionsServiceClient", "TransactionsClient")
}
