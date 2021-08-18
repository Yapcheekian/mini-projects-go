package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int
		domains []string
		total   int
		lines   int
	)

	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())

		if len(fields) < 2 {
			fmt.Printf("wrong input: %d (lines #%d)\n", len(fields), lines)
			os.Exit(1)
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (lines #%d)\n", fields[1], lines)
			os.Exit(1)
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits
		sum[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)

	for _, domain := range domains {
		fmt.Printf("%-30s %10d\n", domain, sum[domain])
	}
	fmt.Printf("%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("Err", err)
	}
}
