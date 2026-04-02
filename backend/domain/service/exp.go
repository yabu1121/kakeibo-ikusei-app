package service

import "github.com/kakebon/backend/domain/model"

func CalcExp(char *model.Character, exp int) {
	char.CurrentExp += exp
	for char.CurrentExp >= char.ExpToNextLevel {
		char.CurrentExp -= char.ExpToNextLevel
		char.CurrentLevel++
		char.ExpToNextLevel = char.CurrentLevel * 100
		if char.CurrentLevel >= 99 {
			char.CurrentLevel = 99
			char.ExpToNextLevel = 9900
		}
		if char.CurrentLevel >= 2 {
			char.ImageURL = "https://example.com/images/level2_bird.png"
		}
	}
}