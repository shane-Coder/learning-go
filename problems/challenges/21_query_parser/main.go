package main

import (
	"errors"
	"fmt"
	"strings"
)

func parseQuery(queryString string) (map[string]string, error) {
	strPairs := strings.Split(queryString, "&")
	var result = make(map[string]string)
	for _, part := range strPairs {
		pairs := strings.SplitN(part, "=", 2)
		if len(pairs) != 2 {
			return nil, fmt.Errorf("malformed query parameter: %s", part)
		}
		key := pairs[0]
		value := pairs[1]
		if key == "" {
			return nil, errors.New("malformed query parameter: key cannot be empty")
		}
		result[key] = value

	}
	return result, nil
}

func main() {
	query1 := "name=shivam&city=gurgaon&active=true"
	params1, err1 := parseQuery(query1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Parsed Query 1:", params1)
	}

	query2 := "id=123&user=alice&malformed_param"
	params2, err2 := parseQuery(query2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Parsed Query 2:", params2)
	}
}
