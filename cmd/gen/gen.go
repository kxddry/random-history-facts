package main

import (
	"bufio"
	"fmt"
	"github.com/kxddry/random-history-facts/internal/lib/factmatcher"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var out []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		fm := factmatcher.FactMatcher{}
		str := fmt.Sprintf(`INSERT INTO facts (fact, normalized_fact) VALUES ('%s', '%s');`, line, fm.Normalize(line))
		out = append(out, str)
	}
	for _, line := range out {
		fmt.Println(line)
	}
}
