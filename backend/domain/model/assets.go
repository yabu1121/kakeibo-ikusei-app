package model

var CharacterImages = map[int]string{
    1: "/images/char/level1.jpg",
    2: "/images/char/level2.jpg",
    3: "/images/char/level3.jpg",
    4: "/images/char/level4.jpg",
    5: "/images/char/level5.jpg",
    6: "/images/char/level6.jpg",
    7: "/images/char/level7.jpg",
    8: "/images/char/level8.jpg",
    9: "/images/char/level9.jpg",
    10: "/images/char/level10.jpg",
    11: "/images/char/level11.jpg",
    12: "/images/char/level12.jpg",
    13: "/images/char/level13.jpg",
    14: "/images/char/level14.jpg",
    15: "/images/char/level15.jpg",
    16: "/images/char/level16.jpg",
    17: "/images/char/level17.jpg",
    18: "/images/char/level18.jpg",
    19: "/images/char/level19.jpg",
    20: "/images/char/level20.jpg",
}

func GetImageByLevel(level int) string {
    if img, ok := CharacterImages[level]; ok {
        return img
    }
    return CharacterImages[1]
}