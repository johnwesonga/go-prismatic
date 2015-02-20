package prismatic

import (
	"log"
	"net/url"
)

type TextTopic struct {
	Topic []Topic `json:"topics"`
}

// Tag Text with Interests.
//
// Prismatic API docs: https://github.com/Prismatic/interest-graph#tag-text-with-interests.
func (s *TopicService) TagText(title, body string) (TextTopic, *Response, error) {
	topics := new(TextTopic)
	if body == "" {
		log.Fatalln("The body of the text to tag. Must be at least 140 characters.")
	}

	data := url.Values{}
	data.Set("title", title)
	data.Set("body", body)

	req, err := s.client.NewRequest("POST", "/text/topic", data)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return TextTopic{}, nil, err
	}

	resp, err := s.client.Do(req, topics)
	if err != nil {
		return TextTopic{}, resp, err
	}

	return *topics, resp, err

}
