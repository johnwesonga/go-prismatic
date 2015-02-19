package prismatic

import "fmt"

type Topic struct {
	Topic string `json:"topic"`
	Id    int    `json:"id"`
	score int    `json:"id"`
}

type TopicResponse struct {
	Topic []Topic `json:"topics"`
}

func (s *TopicService) SearchForRelatedTopic(param int) (TopicResponse, *Response, error) {
	result := new(TopicResponse)
	u := fmt.Sprintf("/topic/topic?id=%v", param)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return TopicResponse{}, nil, err
	}

	resp, err := s.client.Do(req, result)

	if err != nil {
		return TopicResponse{}, resp, err
	}

	return *result, resp, err

}
