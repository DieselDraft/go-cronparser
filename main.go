package cronparser

import (
	"regexp"
	"strings"
)

type Cron struct {
	rules []Rule
}

type Rule struct {
	min        string
	hour       string
	dayOfMonth string
	month      string
	dayOfWeek  string

	command string
}

func Parce(str string) (r Cron) {
	var rules = make([]Rule, 0)

	lines := strings.Split(str, "\n")

	for _, line := range lines {
		line = filterComments(line)

		if len(line) == 0 {
			continue
		}

		rule, isRule := readRule(line)

		if isRule {
			rules = append(rules, rule)
		}
	}

	r.rules = rules

	return
}

func filterComments(line string) string {
	p := strings.SplitN(line, "#", 2)

	if len(p) > 0 {
		return p[0]
	}

	return line
}

func readRule(line string) (Rule, bool) {
	reg, _ := regexp.Compile("\\s*(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(.*)")

	r := reg.FindStringSubmatch(line)

	if len(r) != 7 {
		return Rule{}, false
	}

	rule := Rule{
		min:        r[1],
		hour:       r[2],
		dayOfMonth: r[3],
		month:      r[4],
		dayOfWeek:  r[5],
		command:    r[6],
	}

	return rule, true
}
