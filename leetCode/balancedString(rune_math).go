// https://leetcode.com/problems/check-balanced-string/description/

func isBalanced(num string) bool {
    var balance = 0
    for idx, val := range num {     
        if idx % 2 == 0 {
            balance += int(val - '0')
        } else {
            balance -= int(val - '0')
        }
    }

    return balance == 0
}
