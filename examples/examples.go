package main

import (
	"fmt"

	"github.com/johnwesonga/go-prismatic/prismatic"
)

func main() {
	apiToken := "api-token"
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
