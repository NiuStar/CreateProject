/*
配置文件里的工作流
*/
package Process

import (
	"fmt"
)

func FMT(list map[string]interface{},operator string,imports_list map[string]string) (result string,ok bool,imports map[string]string) {

	ok = false
	if list["value"] == nil  {
		fmt.Println(`fmt(list map[string]interface{},operator string)`)
		fmt.Println(`list:`,list,operator)
		return
	}

	imports = make(map[string]string)
	r_head := ""
	switch list["value"].(type) {
	case []interface{}:{
		for index,value := range list["value"].([]interface{}) {

			switch value.(type) {
			case map[string]interface{}:
				{
					r,ok,imports_l := Process(value.(map[string]interface{}),imports_list)
					fmt.Println("value.(map[string]interface{}:",value.(map[string]interface{}))
					fmt.Println("r:",r)
					if !ok {
						fmt.Println(`FMT(list map[string]interface{},operator string) Process(value.(map[string]interface{}))`)
						fmt.Println(value)
						continue
					}
					for k,v := range imports_l {
						imports[k] = v
						imports_list[k] = v
					}
					r_head += r
				}

			default:
				r_head += getStringValue(value)
			}

			if index != len(list["value"].([]interface{})) - 1 {
				r_head += " , "
			}


		}
	}

	default:
		return
	}

	imports["fmt"] = ""
	imports_list["fmt"] = ""
	result = "fmt.Println("

	result += r_head


	result += ")"

	ok = true

	return
}