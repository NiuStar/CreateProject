package CreateProject

import (
	"github.com/NiuStar/utils"
	"strings"
	"go/format"
)

/*
func main() {


	body := utils.ReadFile("project.json")
	//fmt.Println("body:",body)

	var list map[string]interface{}
	err := json.Unmarshal([]byte(body),&list)
	if err != nil {
		panic(err)
	}

	method := list["method"].([]interface{})

	Process.InitProcess()
	for _,value_o := range method {

		Method.CreateMethod(list["name"].(string),value_o.(map[string]interface{}))

	}
	CreateMain(list["name"].(string),method)
}*/

func CreateMain(path ,userName,projectName string,methods []interface{}) {

	var method_value string = "package main\n" + "import(\n"
	method_value += "\"github.com/NiuStar/server\"\n"

	var methods_str string
	for _,method_l := range methods {
		method := method_l.(map[string]interface{})
		if strings.EqualFold(method["type"].(string) , "POST") {
			methods_str += "ser.POST(\"" + method["name"].(string) + "\"," + method["name"].(string) + "." + method["name"].(string) + ")\n"
		} else if strings.EqualFold(method["type"].(string) , "GET") {
			methods_str += "ser.GET(\"" + method["name"].(string) + "\"," + method["name"].(string) + "." + method["name"].(string) + ")\n"
		}
		method_value += "\"" + userName + "/" + projectName + "/" + method["name"].(string) + "\"\n"
	}

	method_value +=  ")\n func main() { \n ser := server.Default() //服务器初始化\n"

	method_value +=  methods_str

	method_value += "	ser.RunServer() // for a hard coded port\n}"

	data,err := format.Source([]byte(method_value))

	if err != nil {
		panic(err)
	}

	utils.WriteToFile(path + "/" + projectName + "/maintest.go",string(data))
}

