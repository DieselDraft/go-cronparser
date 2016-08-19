package cronparser

import "testing"

func TestParce(t *testing.T) {
	var str = `
# Crontab
MAILTO="server@domain.com"
0 1 * * * php ${WEB_ROOT}/cron/test.php
`

	r := Parce(str)

	if len(r.rules) != 1 {
		t.Errorf("Parce(...) -> len(rules) = %d instead of %d", len(r.rules), 1)
	}

}

func TestFilterComments(t *testing.T) {
	list := map[string]string{
		"":                    "",
		"test # test":         "test ",
		"# Crontab":           "",
		"some text # after #": "some text ",
		"blabla":              "blabla",
	}

	for in, out := range list {
		r := filterComments(in)

		if r != out {
			t.Errorf("filterComments('%s') -> '%s' instead of '%s'", in, r, out)
		}

	}
}

func TestIsRule(t *testing.T) {

	list := map[string]struct {
		valid bool
		rule  Rule
	}{
		`MAILTO="server@domain.com"`: {
			valid: false,
			rule:  Rule{},
		},
		`0 1 * * * php ${WEB_ROOT}/cron/test.php`: {
			valid: true,
			rule: Rule{
				min:        "0",
				hour:       "1",
				dayOfMonth: "*",
				month:      "*",
				dayOfWeek:  "*",
				command:    "php ${WEB_ROOT}/cron/test.php",
			},
		},
	}

	for in, out := range list {
		r, is := readRule(in)

		if is != out.valid {
			t.Errorf("readRule('%s') 1 -> %t instead of %t", in, is, out.valid)
		}

		if r != out.rule {
			t.Errorf("readRule('%s') 2 -> %+v instead of %+v", in, r, out.rule)
		}

	}
}
