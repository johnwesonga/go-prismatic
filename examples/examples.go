package main

import (
	"fmt"

	"github.com/johnwesonga/go-prismatic/prismatic"
)

func main() {
	apiToken := "MTQyNDE4NDE4MTQxNQ.cHJvZA.am9obndlc29uZ2FAZ21haWwuY29t.ap9Wvx-Jd_kPA8g4ErfJD5UNBhA"
	client := prismatic.NewClient(nil, apiToken)
	results, _, err := client.Topics.SearchForInterest("Kimende")
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}

	if len(results.Result) == 0 {
		fmt.Println("No result returned")
	}

	for _, v := range results.Result {
		fmt.Printf("id:%v topic:%v\n", v.Id, v.Topic)
	}

	t, _, err := client.Topics.SearchForRelatedTopic(12011)
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}

	fmt.Println(t.Topic)

}
