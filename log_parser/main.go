package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	p := newParser()

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		res := parse(p, in.Text())

		update(p, res)
	}

	summarize(p)
	dumpErr([]error{in.Err(), err(p)})
}

func dumpErr(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("Err", err)
		}
	}
}

func summarize(p *parser) {
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(p.domains)

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("%-30s %10d\n", "TOTAL", p.total)
}
