package envrep

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var format = "{{envrep.%s}}"

func Replace(txt string) string {
	return clean(value(txt))
}

func value(txt string) string {
	for _, v := range os.Environ() {
		e := strings.SplitN(v, "=", 2)
		txt = strings.Replace(txt, fmt.Sprintf(format, e[0]), e[1], -1)
	}

	return txt
}

func clean(txt string) string {
	return regexp.MustCompile(fmt.Sprintf(format, "*")).ReplaceAllString(txt, "")
}
