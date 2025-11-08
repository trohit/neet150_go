// https://neetcode.io/problems/duplicate-integer?list=neetcode150
func hasDuplicate(nums []int) bool {
    ss := make(map[int]bool)
    for _, v := range nums {
        if ss[v] {
            return true
        }
        ss[v] = true
    }
    return false
}
