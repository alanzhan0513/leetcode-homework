func countRangeSum(nums []int, lower int, upper int) int {
    if len(nums) == 0 {
        return 0
    }
    var count int
    var sumArray = []int{nums[0]}
    for i, num := range nums {
        if num >= lower && num <= upper{
            count++
        }
        if i == 0 {
            continue
        }
        sum := sumArray[i-1]+num
        if sum >= lower && sum <= upper{
            count++
        }
        sumArray = append(sumArray,sum)
    }
    
    for i := 1; i < len(sumArray)-1; i++ {
        for j := i+1; j < len(sumArray); j ++{
            compareVal := sumArray[j] - sumArray[i-1]
            if compareVal >= lower && compareVal <= upper {
                count++
            }
        }
    }
    
    return count
}