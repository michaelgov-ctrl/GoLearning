// https://leetcode.com/problems/simplified-fractions/description/

func simplifiedFractions(n int) []string {
    var res []string
    
    for denominator := 1; denominator <= n; denominator++ {
        for numerator := 1; numerator < denominator; numerator++ {
            if numerator % denominator != 0 && gcd(numerator, denominator) == 1 {
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
