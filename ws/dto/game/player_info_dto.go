package game

type PlayerInfoDto struct {
	Hand       []MemoryCardDto `json:"hand"`
	IsActive   bool            `json:"is_active"`
	IsDeciding bool            `json:"is_deciding"`
	IsDanny    bool            `json:"is_danny"`
}
