package lobby

type StateDto struct {
	Code    string      `json:"code"`
	Host    string      `json:"host"`
	Members []MemberDto `json:"members"`
}
