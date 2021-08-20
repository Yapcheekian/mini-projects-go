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
	lerr    error
}

func newParser() *parser {
	return &parser{
		sum: make(map[string]result),
	}
}

func parse(p *parser, line string) (parsed result) {
	if p.lerr != nil {
		return
	}

	p.lines++

	fields := strings.Fields(line)

	if len(fields) < 2 {
		p.lerr = fmt.Errorf("wrong input: %d (lines #%d)", len(fields), p.lines)
		return
	}

	var err error

	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])

	if parsed.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (lines #%d)", fields[1], p.lines)
	}
	return
}

func update(p *parser, res result) {
	if p.lerr != nil {
		return
	}

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

func err(p *parser) error {
	return p.lerr
}
