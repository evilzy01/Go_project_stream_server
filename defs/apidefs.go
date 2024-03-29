package defs

// Request
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type SignedUp struct {
	Success   bool   `json:"success`
	SessionId string `json:session_id`
}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DispalyCtime string
}

type Comments struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string
	TTL      int64
}
