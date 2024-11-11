package _01_TwoSum

// Bu O(n^2) dir bunu O(n) yapmaya çalışacağız
func TwoSumn2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// buradaki çözüm O(n) olarak bulur
func TwoSumn(nums []int, target int) []int {

	imap := make(map[int]int, len(nums))

	for index, element := range nums {
		imap[element] = index
	}

	addUpNum := 0
	for indexCurrentElement, element := range nums {
		addUpNum = target - element
		if addUpNumIndex, ok := imap[addUpNum]; ok && addUpNumIndex != indexCurrentElement {
			return []int{indexCurrentElement, addUpNumIndex}
		}
	}
	return nil
}
