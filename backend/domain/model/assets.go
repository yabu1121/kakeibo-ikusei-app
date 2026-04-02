package model

var CharacterImages = map[int]string{
    1: "https://example.com/egg.png",
    2: "https://example.com/chick.png",
    3: "https://example.com/bird.png",
}

func GetImageByLevel(level int) string {
    if img, ok := CharacterImages[level]; ok {
        return img
    }
    return CharacterImages[1]
}