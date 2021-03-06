package slow

import (
	"strings"
)

func Event(in <-chan string, out chan<- string) {
	var buffer string
	var isHeader bool
	var isQuery bool

	for line := range in {
		e := ""
		l := len(line)

		if isQuery == false && strings.HasPrefix(line, "# ") {
			isHeader = true
		}

		if isHeader == true && l >= 6 {
			buffer += line + "\n"

			s := string(line[0:6])
			s = strings.ToUpper(s)

			if s == "SELECT" || s == "INSERT" || s == "UPDATE" || s == "DELETE" {
				isQuery = true
			}
		}

		if l > 1 {
			e = string(line[l-1:])
		} else {
			e = string(line)
		}

		if isQuery == true && e == ";" {
			out <- strings.TrimRight(buffer, "\n")

			buffer = line + "\n"
			isHeader = false
			isQuery = false
		}
	}
}

func Properties(event string) map[string]string {
	property := map[string]string{}
	whiteSpaceStart := 0
	whiteSpaceEnd := 0
	startQuery := 0

	p := []rune(event)
	l := len(p)

	for x := 0; x < l; x++ {
		// Register first White Space:
		if p[x] == ' ' {
			whiteSpaceStart = x
		}

		// Start second loop to find next property:
		if p[x] == ':' && p[x+1] == ' ' {
			for y := x + 1; y < l; y++ {
				// Stop when is finished header and start SQL:
				if p[y] == '\n' && p[y+1] != '#' {
					whiteSpaceEnd = y
					break
				}

				// Remove header comments:
				if p[y] == '#' || p[y] == '\n' || p[y] == '\r' {
					p[y] = ' '
				}

				// Replace unnecessary symbols:
				if p[y] == '@' {
					p[y] = '_'
				}

				// Register last White Space:
				if p[y] == ' ' {
					whiteSpaceEnd = y
					continue
				}

				// Stop when find next property:
				if p[y] == ':' && p[y+1] == ' ' {
					break
				}
			}

			key := string(p[whiteSpaceStart:x])
			key = strings.TrimSpace(key)
			key = strings.ToLower(key)
			value := strings.TrimSpace(string(p[x+1 : whiteSpaceEnd]))

			property[key] = value
		}

		// Find timestamp value:
		if (x+24) <= l && string(p[x:x+14]) == "SET timestamp=" {
			property["timestamp"] = string(p[x+14 : x+24])
			startQuery = x + 25
		}
	}
	// Find query:
	property["query"] = string(p[startQuery:l])
	property["query"] = strings.Trim(property["query"], "\n")

	return property
}
