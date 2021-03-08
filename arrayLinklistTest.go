package main

//func removeDuplicates(nums []int) int {
//	pre := nums[0]
//	count := 1
//	for _, data := range nums {
//		if data != pre {
//			pre = data
//			nums[count] = data
//			count++
//		}
//	}
//	return count
//}
//
//func rotate(nums []int, k int) {
//	newNums := make([]int, len(nums))
//	for i, v := range nums {
//		newNums[(i + k) % len(nums)] = v
//	}
//	copy(nums, newNums)
//}

//nums = "----->-->"; k =3
//result = "-->----->";
//
//reverse "----->-->" we can get "<--<-----"
//reverse "<--" we can get "--><-----"
//reverse "<-----" we can get "-->----->"
//this visualization help me figure it out :)
//public void rotate(int[] nums, int k) {
//k %= nums.length;
//reverse(nums, 0, nums.length - 1);
//reverse(nums, 0, k - 1);
//reverse(nums, k, nums.length - 1);
//}
//
//public void reverse(int[] nums, int start, int end) {
//	while (start < end) {
//	int temp = nums[start];
//	nums[start] = nums[end];
//	nums[end] = temp;
//	start++;
//	end--;
//}


//func rotate(nums []int, k int) {
//	k = k % len(nums)
//	reverse(nums, 0, len(nums) - 1)
//	reverse(nums, 0, k - 1)
//	reverse(nums, k, len(nums) - 1)
//}
//
//func reverse(nums []int, start, end int) {
//	//反转数组
//	for start < end {
//		tmp := nums[start]
//		nums[start] = nums[end]
//		nums[end] = tmp
//		start++
//		end--
//	}
//}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	} else if l2 == nil {
//		return l1
//	} else if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists(l1.Next, l2)
//		return l1
//	} else {
//		l2.Next = mergeTwoLists(l1, l2.Next)
//		return l2
//	}
//}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	head := &ListNode{}
//	newList := head
//
//	for l1 != nil || l2 != nil {
//		if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
//			newList.Next = l2
//			newList = l2
//			l2 = l2.Next
//		} else {
//			newList.Next = l1
//			newList = l1
//			l1 = l1.Next
//		}
//	}
//	return head.Next
//}

//迭代 merge两个list
//func mergeList(l1, l2 *ListNode) *ListNode {
//	dummy := &ListNode{}
//	head := dummy
//	for l1 != nil || l2 != nil {
//		if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
//			head.Next = l2
//			head = l2
//			l2 = l2.Next
//		} else {
//			head.Next = l1
//			head = l1
//			l1 = l1.Next
//		}
//	}
//	return dummy.Next
//}

//func isValid(s string) bool {
//	n := len(s)
//	if n % 2 == 1 {
//		return false
//	}
//	pairs := map[byte]byte{
//		')': '(',
//		']': '[',
//		'}': '{',
//	}
//	var stack []byte
//	for i := 0; i < n; i++ {
//		if pairs[s[i]] > 0 {
//			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
//				return false
//			}
//			stack = stack[:len(stack)-1]
//		} else {
//			stack = append(stack, s[i])
//		}
//	}
//	return len(stack) == 0
//}
//
//func isValid(s string) bool {
//	stack := []string{}
//	for _, ch := range s {
//		c := string(ch)
//		if c == "{" || c == "(" || c == "[" {
//			stack = append(stack, c)
//		} else {
//			if len(stack) == 0 {
//				return false
//			}
//			top := stack[len(stack)-1]
//			if top == "(" && c == ")" || top == "[" && c == "]" || top == "{" && c == "}" {
//				stack = stack[:len(stack)-1]
//			} else {
//				return false
//			}
//		}
//	}
//	return len(stack) == 0
//}

//func isValid(str string) bool {
//	var stack []byte
//	n := len(str)
//	for i := 0; i < n; i++ {
//		c := str[i]
//		if c == '{' || c == '[' || c == '(' {
//			stack = append(stack, c)
//		} else {
//			if len(stack) == 0 {
//				return false
//			}
//			top := stack[len(stack) - 1]
//			if top == '(' && c == ')' || top == '{' && c == '}' || top == '[' && c == ']' {
//				stack = stack[:len(stack) - 1]
//			} else {
//				return false
//			}
//		}
//	}
//	return len(stack) == 0
//}


func main() {
}
