package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "package main\n\nimport (\n\t\"fmt\"\n\t\"strings\"\n)\n\nfunc main() {\n\tx := \"%s\"\n\tr := strings.NewReplacer(\"\\n\", \"\\\\n\", \"\\t\", \"\\\\t\", \"\\\"\", \"\\\\\\\"\", \"\\\\\", \"\\\\\\\\\")\n\tfmt.Printf(x, r.Replace(x))\n\treturn\n}\n"
	r := strings.NewReplacer("\n", "\\n", "\t", "\\t", "\"", "\\\"", "\\", "\\\\")
	fmt.Printf(x, r.Replace(x))
	return
}
