package dto

type User struct {
	ID       string   `json:"_id"`
	Index    int      `json:"index"`
	GUID     string   `json:"guid"`
	IsActive bool     `json:"isActive"`
	Balance  string   `json:"balance"`
	Tags     []string `json:"tags"`
	Friends  []Friend `json:"friends"`
}
