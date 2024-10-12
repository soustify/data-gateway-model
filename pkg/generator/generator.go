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
	implementation.CreateLambdaClient(clientName, preffixName, "client")
}

func RegisterServer(clientName string, class interface{}, packages []string) {
	implementation := GetDefinition(class)
	implementation.CreateServer(Name(clientName), "server", packages)
}

func RegisterAdapter(clientName string, class interface{}, packages []string) {
	implementation := GetDefinition(class)
	implementation.CreateAdapter(Name(clientName), "adapter", packages)
}

func (m *Implementation) CreateLambdaClient(clientName, preffixName, _package string) {
	implName := fmt.Sprintf("lambda%s", clientName)
	content := fmt.Sprintf("package %s\n\n", _package)
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
		content += method.GetLambdaClient(implName, preffixName)
	}

	createFile("pkg/client", fmt.Sprintf("%s.go", m.Interface.Name.ToSnakeCase()), content)
}

func (m *Implementation) CreateServer(clientName Name, _package string, packages []string) {
	content := fmt.Sprintf("package %s\n\n", _package)
	content += "import (\n"
	content += "	\"context\"\n"
	content += "	\"errors\"\n"
	content += "    \"github.com/soustify/data-gateway-model/pkg/pb\"\n"
	if packages != nil {
		for _, s := range packages {
			content += fmt.Sprintf("	\"%s\"\n", s)
		}
	}
	content += ")\n\n"
	content += fmt.Sprintf("var %s pb.%s\n\n", clientName, m.Interface.Name)

	content += "type (\n"
	content += fmt.Sprintf("	%s struct {\n", clientName.ToLowerFirst())
	content += fmt.Sprintf("		pb.Unimplemented%s\n", m.Interface.Name)
	content += fmt.Sprintf("		//dao repository.%sRepository\n", clientName)
	content += "		db  func() *gorm.DB\n"
	content += "	}\n"
	content += ")\n\n"

	content += "func init() {\n"
	content += fmt.Sprintf("	%s = &%s{\n", clientName, clientName.ToLowerFirst())
	content += fmt.Sprintf("		//dao: repository.%sDao,\n", clientName)
	content += "		db: func() *gorm.DB {\n"
	content += "			return gorm_dao.GetInstance(\"default\")\n"
	content += "		},\n"
	content += "	}\n"
	content += "}\n\n"

	for _, method := range m.Methods {
		content += method.GetEmptyServer(clientName.ToLowerFirst())
	}
	createFile("pkg/server", fmt.Sprintf("%s.go", m.Interface.Name.ToSnakeCase()), content)
}

func (m *Implementation) CreateAdapter(clientName Name, _package string, packages []string) {
	content := fmt.Sprintf("package %s\n\n", _package)
	content += "import (\n"
	content += "	\"context\"\n"
	content += "    \"github.com/soustify/data-gateway-model/pkg/pb\"\n"
	content += "    \"github.com/soustify/data-gateway-model/pkg/types\"\n"
	content += "    \"github.com/soustify/data-gateway/pkg/server\"\n"

	if packages != nil {
		for _, s := range packages {
			content += fmt.Sprintf("	\"%s\"\n", s)
		}
	}
	content += ")\n\n"
	content += fmt.Sprintf("var %s %sAdapter = &%s{}\n\n", clientName, clientName, clientName.ToLowerFirst())

	content += "type (\n"

	content += fmt.Sprintf("\t%sAdapter interface {", clientName)
	for _, method := range m.Methods {
		content += method.GetAdapter(clientName.ToLowerFirst(), clientName, true)
	}
	content += "\n}\n"
	content += fmt.Sprintf("	%s struct {\n", clientName.ToLowerFirst())
	content += "	}\n"
	content += ")\n\n"

	for _, method := range m.Methods {
		content += method.GetAdapter(clientName.ToLowerFirst(), clientName, false)
	}
	createFile("pkg/adapter", fmt.Sprintf("%s.go", m.Interface.Name.ToSnakeCase()), content)
}

// createFile cria ou sobrescreve um arquivo na pasta "pkg/client"
func createFile(dir, fileName string, content string) error {
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

func (m *Method) GetLambdaClient(structName, lambdaVarName string) string {
	content := "\n"
	if m != nil {
		content += fmt.Sprintf("func (c *%s) %s(ctx %s, in %s, opts ... %s) (%s, %s){\n",
			structName, m.Name, m.Arguments[0], m.Arguments[1], strings.Replace(m.Arguments[2], "[]", "", 1), m.Returns[0], m.Returns[1])
		content += fmt.Sprintf("\tvar result %s\n", m.Returns[0])
		content += fmt.Sprintf("\terr := invokeAndUnmarshal(ctx, \"%s_%s\", in, &result)\n", lambdaVarName, m.Name.ToSnakeCase())
		content += "\treturn result, err\n"
		content += "}\n"
	}
	return content
}

func (m *Method) GetEmptyServer(structName string) string {
	content := "\n"
	if m != nil {
		if len(m.Arguments) == 0 {
			return ""
		}
		content += fmt.Sprintf("func (c *%s) %s(ctx %s, in %s) (%s, %s){\n",
			structName, m.Name, m.Arguments[0], m.Arguments[1], m.Returns[0], m.Returns[1])
		content += "\treturn nil, errors.New(\"Not Implemented Server\")\n"
		content += "}\n"
	}
	return content
}

func (m *Method) GetAdapter(structName string, clientName Name, signature bool) string {
	content := "\n"
	if m != nil {
		if len(m.Arguments) == 0 {
			return ""
		}
		if signature {
			content += fmt.Sprintf("\t\t%s(ctx context.Context, parameter *types.LambdaParameter[%s]) (%s, %s)", m.Name, m.Arguments[1], m.Returns[0], m.Returns[1])
		} else {
			content += fmt.Sprintf("func (c *%s) %s(ctx context.Context, parameter *types.LambdaParameter[%s]) (%s, %s){\n", structName, m.Name, m.Arguments[1], m.Returns[0], m.Returns[1])
			content += "\tcontext, err := parameter.GenerateContext(ctx)\n"
			content += "\tif err != nil {\n"
			content += "\t\treturn nil, err\n"
			content += "\t}\n"
			content += fmt.Sprintf("\treturn server.%s.%s(context, *parameter.Content)\n", clientName, m.Name)
			content += "}\n"
		}
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

func (s Name) ToLowerFirst() string {
	if len(s) == 0 {
		return string(s)
	}
	return strings.ToLower(string(s[0])) + string(s[1:])
}
