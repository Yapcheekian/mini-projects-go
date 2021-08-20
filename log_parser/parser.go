package main

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}
type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func newParser() parser {
	return parser{
		sum: make(map[string]result),
	}
}

func parse(p *parser, line string) (parsed result, err error) {
	p.lines++

	fields := strings.Fields(line)

	if len(fields) < 2 {
		err = fmt.Errorf("wrong input: %d (lines #%d)", len(fields), p.lines)
		return
	}

	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])

	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (lines #%d)", fields[1], p.lines)
		return
	}
	return
}

func update(p *parser, res result) {
	if _, ok := p.sum[res.domain]; !ok {
		p.domains = append(p.domains, res.domain)
	}
	p.total += res.visits

	r := result{
		domain: res.domain,
		visits: res.visits + p.sum[res.domain].visits,
	}

	p.sum[res.domain] = r
}
