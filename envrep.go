package envrep

import (
	"fmt"
	"os"
	"strings"
)

func Replace(txt string) string {
	return env(txt)
}

func env(txt string) string {
	for _, v := range os.Environ() {
		e := strings.SplitN(v, "=", 2)
		txt = strings.Replace(txt, fmt.Sprintf("<env:%s>", e[0]), e[1], -1)
	}

	return txt
}
