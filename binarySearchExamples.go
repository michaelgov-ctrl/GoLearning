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
