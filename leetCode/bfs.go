
// https://leetcode.com/problems/keys-and-rooms/solutions/4773560/841-keys-and-rooms-using-go-bfs-search/

func canVisitAllRooms(rooms [][]int) bool {
    visited := make([]bool, len(rooms))
    visited[0] = true
    stack := make([]int, 0)
    stack = append(stack, 0)

    for len(stack) != 0 {
        pop := stack[0]
        stack = stack[1:]

        for _, key := range rooms[pop] {
            if !visited[key] {
                visited[key] = true
                stack = append(stack, key)
            }
        }
    }

    for _, visit := range visited {
        if visit == false {
            return false
        }
    }

    return true
}
