package prismatic

import (
	"fmt"
	"log"
)

type TextTopic struct {
	Topic []Topic `json:"topics"`
}

// Tag Text with Interests.
//
// https://github.com/Prismatic/interest-graph#tag-text-with-interests.
func (s *TopicService) TagText(title, body string) (*TextTopic, *Response, error) {
	topics := new(TextTopic)
	if body == "" {
		log.Fatalln("The body of the text to tag. Must be at least 140 characters.")
	}
	u := fmt.Sprintf("/text/topic/%v/%v", title, body)
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	return *topics, resp, err

}
