package main

import (
	"fmt"

	"github.com/johnwesonga/go-prismatic/prismatic"
)

func main() {
	apiToken := "MTQyNDE4NDE4MTQxNQ.cHJvZA.am9obndlc29uZ2FAZ21haWwuY29t.ap9Wvx-Jd_kPA8g4ErfJD5UNBhA"
	client := prismatic.NewClient(nil, apiToken)
	results, _, err := client.Topics.SearchForInterest("Clojure")
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

	//j := fmt.Sprint(`{"title": "Clojure", "body":%v`, )
	body := `Clojure is a dynamic programming language that targets the Java Virtual Machine 
	(and the CLR, and JavaScript). It is designed to be a general-purpose language, 
	combining the approachability and interactive development of a 
	scripting language with an efficient and robust infrastructure for multithreaded programming.
	Clojure is a compiled language - it compiles directly to JVM bytecode, 
	yet remains completely dynamic. Every feature supported by Clojure is supported at runtime.`

	tagger, _, err := client.Topics.TagText("Clojure", body)
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}

	fmt.Println(tagger)

	r, _, err := client.Topics.TagUrl("http://en.wikipedia.org/wiki/Clojure")
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}

	fmt.Println(r)

}
