func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var m int
	for l <= r {
		m = (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func guessNumber(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		switch resp := guess(mid); resp {
		case -1:
			right = mid
		case 1:
			left = mid + 1
		case 0:
			return mid
		}
	}

	return left
}

//https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
func nextGreatestLetter(letters []byte, target byte) byte {
	for target <= byte('z') {
		target = target + 1
		left := 0
		right := len(letters) - 1
		var middle int
		for left <= right {
			middle = (left + right) / 2
			if letters[middle] == target {
				return letters[middle]
			} else if letters[middle] > target {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return letters[0]
}
