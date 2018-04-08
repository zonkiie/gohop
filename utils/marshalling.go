/****** MIT License **********
Copyright (c) 2017 Datzer Rainer
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
***************************/

package utils

import (
	"encoding/json"
	"encoding/xml"
	"gopkg.in/yaml.v2"
	"fmt"
	"strings"
)

func JsonMarshal(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return "Error"
	}
	return string(b[:])
}

func XmlMarshal(data interface{}) string {
	b, err := xml.Marshal(data)
	if err != nil {
		return "Error"
	}
	return string(b[:])
}

func YamlMarshal(data interface{}) string {
	b, err := yaml.Marshal(data)
	if err != nil {
		return "Error"
	}
	return string(b[:])
}

func DumpMarshal(data interface{}) string {
	return fmt.Sprintf("Entry: %v\n", data)
}

func Marshal(data interface{}, format string) string {
	switch strings.ToLower(format) {
		case "xml":
			return XmlMarshal(data)
		case "json":
			return JsonMarshal(data)
		case "yaml":
			return YamlMarshal(data)
		case "dump":
			return DumpMarshal(data)
		default:
			panic("Format \"" + format + "\" not known")
	}
}
