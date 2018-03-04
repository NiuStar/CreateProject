/*
配置文件里的工作流
*/
package Process

import (
	"CreateProject/InterfaceFormat"
	"fmt"
)

func operator(list map[string]interface{},operator string,imports_list map[string]string) (result string,ok bool,imports map[string]string) {

	ok = false

	if list["left"] == nil || list["right"] == nil {
		fmt.Println(`list["left"] == nil || list["right"] == nil`)
		fmt.Println(`list:`,list,operator)
		return
	}

	imports = make(map[string]string)
	switch list["left"].(type) {
	case string:{
		var variable bool = true
		if list["variable"] != nil {
			variable = list["variable"].(bool)
		}

		if variable {
			result = list["left"].(string)
		} else {
			result = `"` + list["left"].(string) + `"`
		}
	}
	case map[string]interface{}:{
		var ok1 bool = false
		var imports_l map[string]string
		result,ok1 ,imports_l = Process(list["left"].(map[string]interface{}),imports_list)
		if !ok1 {
			fmt.Println(`operator(list map[string]interface{},operator string) switch list["left"].(type) {`)
			fmt.Println(list["right"])
			return
		}
		for k,v := range imports_l {
			imports[k] = v
			imports_list[k] = v
		}
	}
	case []interface{}:{
		for index,value := range list["left"].([]interface{}) {


			var r string
			r,ok1 ,imports_l := Process(value.(map[string]interface{}),imports_list)

			if !ok1 {
				fmt.Println(`operator(list map[string]interface{},operator string) switch list["left"].(type) 2`)
				fmt.Println(list["right"])
				return
			}

			result += r
			if index != len(list["left"].([]interface{})) - 1 {
				result += " , "
			}
			for k,v := range imports_l {
				imports[k] = v
				imports_list[k] = v
			}

		}
	}
	default:
		result = InterfaceFormat.Interface2StringValue(list["left"])
	}

	result += operator

	r , ok ,imports_l := Process(list["right"].(map[string]interface{}),imports_list)

	if !ok {
		fmt.Println(`operator(list map[string]interface{},operator string) (result string,ok bool)`)
		fmt.Println(list["right"])
		return
	}
	for k,v := range imports_l {
		imports[k] = v
		imports_list[k] = v
	}
	result += r

	ok = true
	return
}
