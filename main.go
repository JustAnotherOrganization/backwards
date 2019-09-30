/*‮*/ package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	/*‮*/ os.Args = os.Args[1:]
	/*‮*/ data := []byte(strings.Join(os.Args, " "))
	/*‮*/ clipboard.WriteAll(fmt.Sprintf("‮%s", reverse(data)))
}

func reverse(data []byte) (r string) {
	/*‮*/ for left, right := 0, len(data)-1; left < right; left, right = left+1, right-1 {
		/*‮*/ data[left], data[right] = data[right], data[left]
	}
	/*‮*/ return string(data)
}
