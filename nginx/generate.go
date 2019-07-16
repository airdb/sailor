package nginx

import (
	"fmt"
	"reflect"
)

const SpecialName = "Key"

func GenServerConfig(config *ServerConfig) {
	cstr := "server	{\n"
	e := reflect.ValueOf(config).Elem()

	for i := 0; i < e.NumField(); i++ {
		switch e.Field(i).Kind() {
		case reflect.Struct:
			for j := 0; j < e.Field(i).NumField(); j++ {
				if e.Field(i).Field(j).String() != "" {
					cstr += "\t" + underscore(e.Field(i).Type().Field(j).Name) + "\t" + e.Field(i).Field(j).String() + ";\n"
				}
			}
			/*
				e1 := reflect.Value(e.Field(i))
				for j := 0; j < e1.NumField(); j++ {
					if e1.Field(j).String() != "" {
						cstr += "\t" + underscore(e1.Type().Field(j).Name) + "\t" + e1.Field(j).String() + ";\n"
					}
				}
			*/
		case reflect.Slice:
			cstr += "\n"
			for j := 0; j < e.Field(i).Len(); j++ {
				for k := 0; k < e.Field(i).Index(j).NumField(); k++ {
					if e.Field(i).Index(j).Field(k).Kind() == reflect.String {
						if e.Field(i).Index(j).Field(k).String() == "" {
							continue
						}

						if e.Field(i).Index(j).Type().Field(k).Name == SpecialName {
							cstr += "\tlocaltion\t" + e.Field(i).Index(j).Field(k).String() + " {\n"
							continue
						} else {
							cstr += "\t\t" + underscore(e.Field(i).Index(j).Type().Field(k).Name)
							cstr += "\t" + e.Field(i).Index(j).Field(k).String() + ";\n"
						}
					}
					if e.Field(i).Index(j).Field(k).Kind() == reflect.Slice {
						// sliceHandler(e.Field(i).Index(j).Type().Field(k).Name, e.Field(i).Index(j).Field(k))
						cstr += sliceHandler(e.Field(i).Index(j).Type().Field(k).Name, e.Field(i).Index(j).Field(k))
					} else if e.Field(i).Index(j).Field(k).Kind() == reflect.Struct {
						cstr += structHandler(e.Field(i).Index(j).Field(k))
					}
				}
				cstr += "\t}\n\n"
			}
		case reflect.String:
			if e.Field(i).String() != "" {
				cstr += "\t" + underscore(e.Type().Field(i).Name) + "\t" + e.Field(i).String() + ";\n"
			}
		default:
			fmt.Println("Invaild reflect kind.")
		}
	}

	cstr += "}"
	fmt.Println(cstr)
}

func GenUpstreamConfig(config *UpstreamConfig) {
	cstr := "upstream\t" + config.Key + " {\n"

	for _, server := range config.Servers {
		e := reflect.ValueOf(server)
		for i := 0; i < e.NumField(); i++ {
			if e.Type().Field(i).Name == SpecialName {
				cstr += "upstream\t" + e.Field(i).String() + " {\n"
				continue
			}
			cstr += stringHandler(e.Type().Field(i).Name, e.Field(i).String())
		}
	}

	for _, upstream := range config.Upstreams {
		e := reflect.ValueOf(upstream)
		for i := 0; i < e.NumField(); i++ {
			if e.Type().Field(i).Name == SpecialName {
				cstr += "upstream\t" + e.Field(i).String() + " {\n"
				continue
			}

			// some struct have uint, so need use fmt to format.
			value := fmt.Sprint(e.Field(i))
			cstr += stringHandler(e.Type().Field(i).Name, value)
		}
	}

	cstr += "}\n"
	fmt.Println(cstr)

}

func sliceHandler(key string, e reflect.Value) (cstr string) {
	if key == "Locations" || key == "server" || key == "upstream" {
		fmt.Printf("%s is a special name\n", key)
	}

	for i := 0; i < e.Len(); i++ {
		if e.Index(i).Kind().String() == "struct" {
			cstr += structHandler(e.Index(i))
		} else {
			fmt.Println("Invaild config struct")
		}
	}
	return
}

func structHandler(e reflect.Value) (cstr string) {
	for i := 0; i < e.NumField(); i++ {
		if e.Field(i).Kind().String() == "string" {
			cstr += stringHandler(e.Type().Field(i).Name, e.Field(i).String())
		} else if e.Field(i).Kind().String() == "slice" {
			if e.Type().Field(i).Name == "Proxy_set_headers" {
				cstr += sliceHandler("proxy_set_header", e.Field(i))
			}
		}
	}
	return
}

func stringHandler(key, value string) (cstr string) {
	if value != "" {
		if key == "Balancer_by_lua_block" {
			cstr += "\n"
			cstr += "\t\t" + underscore(key) + "\t{\n"
			cstr += "\t\t\t" + value + "\n"
			cstr += "\t\t}\n"
		} else {
			// cstr += "\t\t" + underscore(key) + "\t" + value + ";\n"
			cstr += "\t\t" + underscore(key) + "\t" + value + ";\n"
		}
	}
	return
}
