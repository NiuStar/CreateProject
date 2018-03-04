package Method

import (
	"CreateProject/Variable"
	"CreateProject/Params"
	"go/format"
	"github.com/NiuStar/utils"
	"CreateProject/Process"
	"fmt"
)

func CreateMethod(path ,projectName string,value map[string]interface{}) {


	var method_value string = "package " + value["name"].(string) + "\n" + "import(\n"


	var imports map[string]string = make(map[string]string)


	var variable_value string = `
/**************************函数内成员变量************************/
`
	variable := value["variable"].(map[string]interface{})
	for name,info := range variable {
		variable_value += Variable.GetVariable(info.(map[string]interface{}),name)
	}





	var params_value string= `

/**************************客户端传值************************/
`
	if value["params"] != nil {
		params := value["params"].([]interface{})

		var get_params string
		for _,v_o := range params {
			get_params_str ,imports_l := Params.GetParams(v_o.(map[string]interface{}))
			get_params += get_params_str

			for k,v := range imports_l {
				imports[k] = v
			}

		}
		params_value += get_params
	}


	var processes_value string = `

/**************************工作流************************/
`

	if value["process"] != nil {
		processes := value["process"].([]interface{})

		var get_processes string
		for _,v_o := range processes {
			v := v_o.(map[string]interface{})
			/*if strings.EqualFold(v["type"].(string) , "operator") {
				get_processes += v["value"].(string) + "\n"
			} else if strings.EqualFold(v["type"].(string) , "fmt") {
				get_processes += v["value"].(string) + "\n"
			}*/

			p , ok ,imports_l := Process.Process(v,imports)

			if !ok {
				fmt.Println("错误的转义：",v)
			} else {
				get_processes += p + "\n"
				for k,v := range imports_l {
					imports[k] = v
				}
			}

		}
		processes_value += get_processes
	}


	processes_value += `

/**************************返回结果************************/
`


	var imports_value string = `
/**************************包含包文件************************/
`


	imports_value += "\"github.com/NiuStar/server/gin\"\n"
	imports_value += "\"net/http\"\n"
	imports_value += "\"encoding/json\"\n"

	for k_o,v_o := range imports {
			imports_value += v_o + " \"" + k_o + "\"\n"
	}

	imports_value += ")\n" +
		"func " + value["name"].(string) + "(c *gin.Context) {\n"


	imports_value += `
/**************************最后结束时候的回调，需要返回客户端数据，result是结果************************/
	var result map[string]interface{} = make(map[string]interface{})
	defer func() {
		body,_ := json.Marshal(result)
		c.String(http.StatusOK,string(body))
	}()
`

	method_value += imports_value

	method_value += variable_value

	method_value += params_value
	method_value += processes_value


	var reslut string = value["result"].(string)

	method_value += "	result[\"result\"] = " + reslut + "\n"

	method_value += "	return\n}"

	fmt.Println("method_value:",method_value)

	data,err := format.Source([]byte(method_value))

	if err != nil {
		panic(err)
	}

	utils.WriteToFile(path + "/" + projectName + "/" + value["name"].(string) + "/" + value["name"].(string) + ".go",string(data))
}
