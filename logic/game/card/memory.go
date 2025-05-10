package card

import (
	gameDto "danny-backend/ws/dto/game"
)

type Memory struct {
	Image string
	Id    int
}

func (mem *Memory) ToDto() gameDto.MemoryCardDto {
	return gameDto.MemoryCardDto{
		Id:    mem.Id,
		Image: mem.Image,
	}
}
