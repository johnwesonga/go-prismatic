package prismatic

type UrlTopic struct {
	Topic []Topic `json:"topics"`
}

// Tag URL with interests.
//
// Prismatic API: https://github.com/Prismatic/interest-graph#tag-url-with-interests.
