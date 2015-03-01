package prismatic

type TopicService struct {
	client *Client
}

type Topic struct {
	Topic string `json:"topic"`
	Id    int    `json:"id"`
	Score int    `json:"score"`
}
