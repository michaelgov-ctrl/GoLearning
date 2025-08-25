// https://leetcode.com/problems/trionic-array-i/description/

func isTrionic(nums []int) bool {
	if nums[1] <= nums[0] {
		return false
	}
	state := 1
	for i := 2; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return false
		}

        switch state {
        case 1:
            if nums[i-1] < nums[i] {
				continue
			}
			state = 2
        case 2:
            if nums[i-1] > nums[i] {
				continue
			}
			state = 3
        default:
            if nums[i-1] < nums[i] {
				continue
			}
			return false
        }

	}
    
	return state == 3
}

