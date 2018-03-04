/*
配置文件里的工作流
*/
package Process

import (
	"fmt"
)

func Append(list map[string]interface{},operator string,imports_list map[string]string) (result string,ok bool,imports map[string]string) {

	ok = false
	if list["left"] == nil || list["right"] == nil  {
		fmt.Println(`fmt(list map[string]interface{},operator string)`)
		fmt.Println(`list:`,list,operator)
		return
	}

	imports = make(map[string]string)
	r_head := ""
	switch list["left"].(type) {
	case string:{
		var variable bool = true
		if list["variable"] != nil {
			variable = list["variable"].(bool)
		}

		if variable {
			r_head = list["left"].(string)
		} else {
			return
		}
	}
	case map[string]interface{}:{
		r ,ok1,imports_l := Process(list["left"].(map[string]interface{}),imports_list)
		if !ok1 {
			fmt.Println(`append(list map[string]interface{},operator string) Process(list["left"].(map[string]interface{}))`)
			fmt.Println(`list:`,list,operator)
			return
		}
		for k,v := range imports_l {
			imports[k] = v
			imports_list[k] = v
		}

		r_head = r
	}
	default:
		return
	}



	result = "append("

	result += r_head + " , "

	r ,ok1,imports_l := Process(list["right"].(map[string]interface{}),imports_list)
	if !ok1 {
		fmt.Println(`append(list map[string]interface{},operator string) Process(list["left"].(map[string]interface{}))`)
		fmt.Println(`list:`,list,operator)
		return
	}

	result += r

	for k,v := range imports_l {
		imports[k] = v
		imports_list[k] = v
	}

	result += ")"

	ok = true

	return
}