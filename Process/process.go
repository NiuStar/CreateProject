package Process

import (
	"CreateProject/InterfaceFormat"
	"fmt"
)

type Operator struct {
	method ProcessMethod
	operator string
}

var order map[string]*Operator = make(map[string]*Operator)

type ProcessMethod func(map[string]interface{},string,map[string]string) (string,bool,map[string]string)


func InitProcess() {
	AddProcess("operator:=",operator," := ")
	AddProcess("operator=",operator," = ")
	AddProcess("operator+",operator," + ")
	AddProcess("operator-",operator," - ")
	AddProcess("fmt",FMT,"")
	AddProcess("mapvalue",mapValue,"")
	AddProcess("append",Append,"")
	AddProcess("plugin",Plugin,"")
	AddProcess("CodeBlock",CodeBlock,"")

}


func AddProcess(name string ,method ProcessMethod,operator string) {
	order[name] = &Operator{method:method,operator:operator}
}

func Process(list map[string]interface{},imports_list map[string]string) (result string,ok bool,imports map[string]string) {

	ok = false
	if list["type"] != nil {
		if order[list["type"].(string)] != nil {
			return order[list["type"].(string)].method(list,order[list["type"].(string)].operator,imports_list)
		}

	} else {
		if list["left"] == nil {
			return
		}
		if list["variable"] != nil && list["variable"].(bool) {
			switch list["left"].(type) {
				case string:{
					return list["left"].(string),true,nil
				}
				case map[string]interface{}:{
					return Process(list["left"].(map[string]interface{}),imports_list)
				}

			}
		} else if list["variable"] == nil {

			switch list["left"].(type) {
			case map[string]interface{}:{
				 return Process(list["left"].(map[string]interface{}),imports_list)
			}
			default:
				return InterfaceFormat.Interface2StringValue(list["left"]),true,nil
			}


		} else {
			return getStringValue(list["left"]),true,nil
			//return `"` + list["left"].(string) + `"`
		}
	}

	fmt.Println("Process OVER")
	return
}

func getStringValue(value interface{}) string {
	switch value.(type) {
		case string:{
			return `"` + value.(string) + `"`
		}
	default:
		return InterfaceFormat.Interface2StringValue(value)
	}
}