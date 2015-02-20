package prismatic

import "log"

type UrlTopic struct {
	Topic []Topic `json:"topics"`
}

// Tag URL with interests.
//
// Prismatic API: https://github.com/Prismatic/interest-graph#tag-url-with-interests.
func (s *TopicService) TagUrl(url string) (UrlTopic, *Response, error) {
	urlTopic := new(UrlTopic)
	if url == "" {
		log.Fatalln("url is required")
	}

	data := url.Values{}
	data.Set("url", url)

	req, err := s.client.NewRequest("POST", "/url/topic", data)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return UrlTopic{}, nil, err
	}

	resp, err := s.client.Do(req, topics)
	if err != nil {
		return UrlTopic{}, resp, err
	}

	return *urlTopic, resp, err
}
