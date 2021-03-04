package main

//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]

//双指针 i j
	//func moveZeros(nums []int) []int {
	//	j := 0
	//	for i, e := range nums {
	//		if e != 0 {
	//
	//			nums[i], nums[j] = nums[j], nums[i]
	//			//nums[j] = nums[i]
	//			//if i != j {
	//			//	nums[i] = 0
	//			//}
	//
	//			j++
	//		}
	//		//nums[i] = 0
	//	}
	//	return nums
//}

//[1,8,6,2,5,4,8,3,7]

//func maxArea(height []int) int {
//	maxArea := 0
//	for i, j := 0, len(height) - 1; i < j; {
//		minHeight := 0
//		if height[i] < height[j] {
//			minHeight = height[i]
//			i++
//		} else {
//			minHeight = height[j]
//			j--
//		}
//		tmpArea := minHeight * (j - i + 1)
//		maxArea = max(tmpArea, maxArea)
//	}
//	return maxArea
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//前提 是排序
//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	var numsList [][]int
//	for k := 0; k < len(nums) - 2; k++ {
//		if nums[k] > 0	{
//			break
//		}
//		if k > 0 && nums[k] == nums[k - 1] {
//			continue
//		}
//		for i, j := k + 1, len(nums) - 1; i < j ; {
//			s := nums[k] + nums[i] + nums[j]
//			if s > 0 {
//				j--
//			} else if s < 0 {
//				i++
//			} else {
//				e := []int{nums[k], nums[i], nums[j]}
//				numsList = append(numsList, e)
//				//防止重复 需过滤重复的num 注意是有序的！！！！！！！！！！！
//				i++
//				for ; i < j && nums[i] == nums[i - 1]; i++ {}
//				j--
//				for ; i < j && nums[j] == nums[j + 1]; j-- {}
//
//			}
//		}
//	}
//	return numsList
//}

//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	res := [][]int{}
//
//	for i := 0; i < len(nums)-2; i++ {
//		n1 := nums[i]
//		if n1 > 0 {
//			break
//		}
//
//		if i > 0 && n1 == nums[i-1] {
//			continue
//		}
//
//		l, r := i+1, len(nums)-1
//		for l < r {
//			n2, n3 := nums[l], nums[r]
//			if n1+n2+n3 == 0 {
//				res = append(res, []int{n1, n2, n3})
//				for l < r && nums[l] == n2 {
//					l++
//				}
//				for l < r && nums[r] == n3 {
//					r--
//				}
//			} else if n1+n2+n3 < 0 {
//				l++
//			} else {
//				r--
//			}
//		}
//	}
//	return res
//}

//func threeSum(nums []int) [][]int {
//	var results [][]int
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-2; i++ {
//
//		if i > 0 && nums[i] == nums[i-1] {
//			continue//To prevent the repeat
//		}
//
//		target, left, right := -nums[i], i+1, len(nums)-1
//
//		for left < right {
//			sum := nums[left] + nums[right]
//
//			if sum == target {
//				results = append(results, []int{nums[i], nums[left], nums[right]})
//				left++
//				right--
//				for left < right && nums[left] == nums[left-1] {
//					left++
//				}
//				for left < right && nums[right] == nums[right+1] {
//					right--
//				}
//			} else if sum > target {
//				right--
//			} else if sum < target {
//				left++
//			}
//		}
//	}
//	return results
//}

//爬阶梯
//func climbStairs(num int) int {
//	if num < 4 {
//		return num
//	}
//	return climbStairs(num - 1) + climbStairs(num - 2)
//}

//将fibonacci改为循环的模式
//func climbStairs(n int) int {
//	if n <= 2 {
//		return n
//	}
//	pre, cur, next := 1, 2, 3
//	for i := 3; i <= n; i++ {
//		next = pre + cur
//		pre = cur
//		cur = next
//	}
//	return next
//}


func main() {

	//fmt.Println(climbStairs(6))
	//nums := []int{-1,0,3,0,1,10,1}
	//fmt.Println(threeSum(nums))

	//height := []int{1,8,6,2,5,4,8,3,7}
	//fmt.Println(maxArea(height))

	//nums := []int{0,1,0,3,12}
	//fmt.Println(moveZeros(nums))
}

