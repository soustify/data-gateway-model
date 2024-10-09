package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

type (
	Implementation struct {
		Interface Interface
		Methods   []Method
	}

	Interface struct {
		Name    Name
		Package string
	}

	Method struct {
		Name      Name
		Arguments Arguments
		Returns   Returns
	}

	Arguments []string
	Returns   []string
	Name      string
)

func RegisterLambdaClient(clientName, preffixName string, class interface{}) {
	implementation := GetDefinition(class)
	implementation.ContentFile(clientName, preffixName)
}

func (m *Implementation) ContentFile(clientName, preffixName string) {
	implName := fmt.Sprintf("lambda%s", clientName)
	content := "package client\n\n"
	content += "import (\n"
	content += "	\"context\"\n"
	content += "    \"github.com/soustify/data-gateway-model/pkg/pb\"\n"
	content += "	\"google.golang.org/grpc\"\n"
	content += ")\n\n"
	content += "var (\n"
	content += fmt.Sprintf("\t%s pb.%s = &%s{}\n", clientName, m.Interface.Name, implName)
	content += ")\n\n"
	content += fmt.Sprintf("type %s struct{}\n", implName)

	for _, method := range m.Methods {
		content += method.GetSignature(implName, preffixName)
	}

	createFile(fmt.Sprintf("%s.go", m.Interface.Name.ToSnakeCase()), content)
}

// createFile cria ou sobrescreve um arquivo na pasta "pkg/client"
func createFile(fileName string, content string) error {
	dir := "pkg/client"
	filePath := filepath.Join(dir, fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("erro ao criar diretório: %v", err)
		}
	}
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("erro ao criar ou sobrescrever arquivo: %v", err)
	}
	fmt.Printf("Arquivo %s criado/atualizado com sucesso na pasta pkg/client\n", fileName)
	return nil
}

func (m *Method) GetSignature(structName, lambdaVarName string) string {
	content := "\n"
	if m != nil {
		fmt.Sprintf("func (c *%s) %s(ctx %s, in %s, opts ... %s){",
			"%s", m.Name, m.Arguments[0], m.Arguments[1], m.Arguments[2])
		content += fmt.Sprintf("func (c *%s) %s(ctx %s, in %s, opts ... %s) (%s, %s){\n",
			structName, m.Name, m.Arguments[0], m.Arguments[1], strings.Replace(m.Arguments[2], "[]", "", 1), m.Returns[0], m.Returns[1])
		content += fmt.Sprintf("\tvar result %s\n", m.Returns[0])
		content += fmt.Sprintf("\terr := invokeAndUnmarshal(ctx, \"%s_%s\", in, &result)\n", lambdaVarName, m.Name.ToSnakeCase())
		content += "\treturn result, err\n"
		content += "}\n"
	}
	return content
}

func GetDefinition(class interface{}) Implementation {
	typ := reflect.TypeOf(class).Elem() // Obtém o tipo subjacente da interface
	var impl Implementation
	if typ.Kind() == reflect.Interface {
		impl.Interface = Interface{
			Name:    Name(typ.Name()),
			Package: typ.PkgPath(),
		}
		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)
			methodObj := Method{
				Name:      Name(method.Name),
				Arguments: getMethodArguments(method),
				Returns:   getMethodReturns(method),
			}
			impl.Methods = append(impl.Methods, methodObj)
		}
	}
	return impl
}

func getMethodArguments(method reflect.Method) []string {
	var args []string
	for j := 0; j < method.Type.NumIn(); j++ {
		param := method.Type.In(j)
		args = append(args, param.String())
	}
	return args
}

func getMethodReturns(method reflect.Method) []string {
	var returns []string
	for k := 0; k < method.Type.NumOut(); k++ {
		returnType := method.Type.Out(k)
		returns = append(returns, returnType.String())
	}
	return returns
}

func (s Name) ToSnakeCase() string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(string(s), "${1}_${2}")
	return strings.ToLower(snake)
}
