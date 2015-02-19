package prismatic

import "fmt"

// Result represents a single value returned from a topic search.
type Result struct {
	Topic string `json:"topic,omitempty"`
	Id    int    `json:"id,omitempty"`
}

// Result represents the returned collection from a topic search.
type ResultResponse struct {
	Result []Result `json:"results"`
}

// SearchForInterest.
//
// Prismatic API docs: https://github.com/Prismatic/interest-graph#search-for-an-interest
func (s *TopicService) SearchForInterest(param string) (ResultResponse, *Response, error) {
	result := new(ResultResponse)
	u := fmt.Sprintf("/topic/search?search-query=%v", param)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return ResultResponse{}, nil, err
	}

	resp, err := s.client.Do(req, result)

	if err != nil {
		return ResultResponse{}, resp, err
	}

	return *result, resp, err

}
