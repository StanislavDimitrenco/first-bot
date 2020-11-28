package main

//structure for request get me
type GetMe struct {
	Ok     bool        `json:"ok"`
	Result GetMeResult `json:"result"`
}

type GetMeResult struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	UserName  string `json:"user_name"`
}
