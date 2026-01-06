// https://leetcode.com/problems/complete-prime-number/description/

func completePrime(num int) bool {
    if !isPrime(num) {
        return false
    }

    for trimmed := trimLeft(num); trimmed > 0; trimmed = trimLeft(trimmed) {
        if !isPrime(trimmed) {
            return false
        }
    }

    for trimmed := trimRight(num); trimmed > 0; trimmed = trimRight(trimmed) {
        if !isPrime(trimmed) {
            return false
        }
    }

    return true
}

func trimLeft(num int) int {
    if num < 10 {
        return 0
    }

    power := int(math.Pow(10, math.Floor(math.Log10(float64(num)))))
    return num % power
}

func trimRight(num int) int {
    return num / 10
}

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }

    for i := 2; i*i <= num; i++ {
        if num % i == 0 {
            return false
        }
    }

    return true
}
