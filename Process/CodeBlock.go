/*
配置文件里的工作流，代码块，用大括号括起来的，然后里面是工作流
*/
package Process

import (
	"fmt"
)

func CodeBlock(list map[string]interface{},operator string,imports_list map[string]string) (result string,ok bool,imports map[string]string) {

	ok = false
	result = "{"
	if list["process"] != nil {
		processes := list["process"].([]interface{})

		var get_processes string
		for _,v_o := range processes {
			v := v_o.(map[string]interface{})

			p , ok ,imports_l := Process(v,imports)

			if !ok {
				fmt.Println("错误的转义：",v)
			} else {
				get_processes += p + "\n"
				for k,v := range imports_l {
					imports[k] = v
				}
			}

		}
		result += get_processes
	}

	result += "}"
	ok = true
	return
}
