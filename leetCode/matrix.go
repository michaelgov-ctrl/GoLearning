// generic matrix related *things*

// https://leetcode.com/problems/flood-fill/description/
type pos struct {
    x, y int
}

func newPos(x, y int) pos {
    return pos{
        x: x,
        y: y,
    }
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    pos, srcColor := newPos(sc, sr), image[sr][sc]
    if srcColor != color {
        fillAdjacent(image, pos, srcColor, color)
    }

    return image
}

func fillAdjacent(image [][]int, pos pos, srcColor, color int) {
    if pos.y < 0 || pos.y >= len(image) || pos.x < 0 || pos.x >= len(image[0]) || image[pos.y][pos.x] != srcColor {
        return
    }

    image[pos.y][pos.x] = color
    
    fillAdjacent(image, newPos(pos.x, pos.y-1), srcColor, color)
    fillAdjacent(image, newPos(pos.x, pos.y+1), srcColor, color)
    fillAdjacent(image, newPos(pos.x-1, pos.y), srcColor, color)
    fillAdjacent(image, newPos(pos.x+1, pos.y), srcColor, color)
}
