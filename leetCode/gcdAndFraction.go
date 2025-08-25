// https://leetcode.com/problems/simplified-fractions/description/

func simplifiedFractions(n int) []string {
    var res []string
    for denominator := 1; denominator <= n; denominator++ {
        for numerator := 1; numerator < denominator; numerator++ {
            if gcd(numerator, denominator) == 1 {
                n, d := strconv.Itoa(numerator), strconv.Itoa(denominator)
                res = append(res, n+"/"+d)
            }
        }
    }

    return res
}

func gcd(a, b int) int {
    if b==0 {
        return a
    }
    return gcd(b,a%b)
}




// https://leetcode.com/problems/gcd-of-odd-and-even-sums/

type Parity int

const (
    Even Parity = iota
    Odd
)

func (p Parity) Value() int {
    if p == Even {
        return 2
    }
    return 1
}

func (p Parity) Sum(n int) int {
    var sum int
    v := p.Value()
    for range n {
        sum += v
        v += 2
    }
    return sum
}

func gcd(a, b int) int {
    if b==0 {
        return a
    }
    return gcd(b,a%b)
}

func gcdOfOddEvenSums(n int) int {
    return gcd(Even.Sum(n), Odd.Sum(n))
}

// or just
// proof: https://leetcode.com/problems/gcd-of-odd-and-even-sums/solutions/7118260/just-simple-math
func gcdOfOddEvenSums(n int) int {
    return n
}
