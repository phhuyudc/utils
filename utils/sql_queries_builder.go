package sql_queries_builder

import (
	"os"
	"strings"
	"sync"
)

var (
	queries map[string]string = make(map[string]string)
	l       sync.Mutex
)

func AddQueries(fileName string) {
	l.Lock()
	defer l.Unlock()
	content, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	text := string(content)
	queries_text := strings.Split(text, "\n\n")
	for _, v := range queries_text {
		lines := strings.Split(v, "\n")
		nameText := ""
		queryText := ""
		for _, l := range lines {
			if len(l) < 2 {
				continue
			}
			if l[:2] == "--" {
				// handler description
				if strings.Contains(l, "name") {
					nameRaw := strings.Split(l, ":")
					if len(nameRaw) < 2 {
						continue
					}
					nameText = strings.Split(nameRaw[1], ",")[0]
				}
			} else {
				queryText += "\n" + l
			}
		}
		nameText = strings.Trim(nameText, " ")
		nameText = strings.Trim(nameText, "\n")
		queryText = strings.Trim(queryText, " ")
		queryText = strings.Trim(queryText, "\n")
		queries[nameText] = queryText
	}
}

func GetQuery(queryName string) string {
	l.Lock()
	defer l.Unlock()
	return queries[queryName]
}
