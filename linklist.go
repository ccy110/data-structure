package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//func hasCycle(head *ListNode) bool {
//	if head == nil ||  head.Next == nil {
//		return false
//	}
//	slow, fast := head, head.Next
//	for slow != fast {
//
//		if fast == nil ||  fast.Next == nil {
//			return false
//		}
//
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return true
//}

//func hasCycle(head *ListNode) bool {
//	if head == nil {
//		return false
//	}
//	slow, fast := head, head
//	for fast.Next != nil && fast.Next.Next != nil {
//		fast = fast.Next.Next
//		slow = slow.Next
//		if slow == fast {
//			return true
//		}
//	}
//	return false
//}

//func hasCycle(head *ListNode) bool {
//	if head == nil {
//		return false
//	}
//	slow, fast := head, head
//
//	for fast.Next != nil && fast.Next.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//		if slow == fast {
//			return true
//		}
//	}
//	return false
//}

//recursion 递归
//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	newHead := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return newHead
//}

//尾递归
//func reverseList(head *ListNode) *ListNode {
//	return reverse(head, head)
//}
//
//func reverse(root, node *ListNode) *ListNode {
//	if root == nil || node.Next == nil {
//		return root
//	}
//	//node.Next = node.Next.Next
//	//node.Next.Next = root
//	//root = node.Next 不能如上这样写 是错误的
//	node.Next, node.Next.Next, root = node.Next.Next, root, node.Next
//
//	return reverse(root, node)
//}

////iteration 迭代
//func reverseList(head *ListNode) *ListNode{
//	var pre *ListNode
//	cur := head
//	for cur != nil {
//		next := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = next
//	}
//	return pre
//}

//func hasCycle(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return false
//	}
//	slow, fast := head, head.Next
//	for slow != fast {
//		if fast == nil || fast.Next == nil { //肯定是fast先到达链表尾部，所以fast或者fast.Next 为链表尾部
//			return false
//		}
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return true
//}

//func swapPairs(head *ListNode) *ListNode {
//	//此处虚拟头结点, 不虚拟头结点呢？？？？
//	dummy := &ListNode{
//		Next: head,
//	}
//	prev := dummy
//	cur := head
//
//	for cur != nil && cur.Next != nil {
//		next := cur.Next
//
//		cur.Next = next.Next
//		next.Next = cur
//		prev.Next = next
//
//		prev = cur
//		cur = cur.Next
//	}
//
//	return dummy.Next
//}

////两两交换
//func swapPairs(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	pre := head
//	cur := head
//	for next == nil {
//		next = pre
//
//		cur = next.Next
//		next = cur.Next
//	}
//
//	return head
//}
//
//func detectCycle(head *ListNode) *ListNode {
//	seen := map[*ListNode]struct{}{}
//	for head != nil {
//		if _, ok := seen[head]; ok {
//			return head
//		}
//		seen[head] = struct{}{}
//		head = head.Next
//	}
//	return nil
//}

//func reverseKGroup(head *ListNode, k int) *ListNode {
//	hair := &ListNode{Next: head}
//	pre := hair
//	for head != nil {
//		tail := pre
//		for i := 0; i < k; i++ {
//			tail = tail.Next
//			if tail == nil {
//				return hair.Next
//			}
//		}
//		nex := tail.Next
//		head, tail = myReverse(head, tail)
//		pre.Next = head
//		tail.Next = nex
//		pre = tail
//		head = tail.Next
//	}
//	return hair.Next
//}
//
//func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
//	prev := tail.Next
//	p := head
//	for prev != tail {
//		nex := p.Next
//		p.Next = prev
//		prev = p
//		p = nex
//	}
//	return tail, head
//}
//
//func reverseGroup(head *ListNode, k int) *ListNode {
//	hair := &ListNode{Next: head}
//	pre := hair
//	for head != nil {
//		tail := pre
//		for i := 0; i < k; i++ {
//			tail = tail.Next
//			if tail == nil {
//				return hair.Next
//			}
//		}
//		nex := tail.Next
//		head, tail = myReverse(head, tail)
//		pre.Next = head
//		tail.Next = nex
//		pre = tail
//		head = tail.Next
//	}
//	return hair.Next
//}
//
//func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
//	prev := tail.Next
//	p := head
//	for prev != tail {
//		nex := p.Next
//		p.Next = prev
//		prev = p
//		p = nex
//	}
//	return tail, head
//}

//func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
//	prev := tail.Next
//	p := head
//	for prev != tail {
//		nex := p.Next
//		p.Next = prev
//		prev = p
//		p = nex
//	}
//	return tail, head
//}

//func reverseGroup(head *ListNode, k int) *ListNode {
//	hair := &ListNode{Next: head}
//	pre := hair
//	for head != nil {
//		tail := pre
//		for i := 0; i < k; i++ {
//			tail = tail.Next
//			if tail == nil {
//				return hair.Next
//			}
//		}
//		next := tail.Next
//		//将k个的头和尾 返回
//		head, tail = myReverse(head, tail)
//		pre.Next = head
//		tail.Next = next
//		pre = tail
//		head = tail.Next
//	}
//	return hair.Next
//}
//
////反转链表 双指针反转
//func myReverse(head, tail *ListNode) (*ListNode, *ListNode){
//	prev := tail.Next
//	p := head
//	for prev != tail {
//		nex := p.Next
//		p.Next = prev
//		prev = p
//		p = nex
//	}
//	return tail, head
//}

//在遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。
//在更改引用之前，还需要存储后一个节点。最后返回新的头引用 双指针引用
//func reverseList(head *ListNode) *ListNode{
//	var pre *ListNode
//	cur := head
//	for cur != nil {
//		next := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = next
//	}
//	return pre
//}

//func reverseList(head *ListNode) *ListNode {
//	var pre *ListNode
//	cur := head
//	for cur != nil {
//		next := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = next
//	}
//	return pre
//}

//func detectCycle(head *ListNode) *ListNode {
//	slow, fast := head, head
//	for {
//		if fast == nil || fast.Next == nil {
//			return nil
//		}
//		fast = fast.Next.Next
//		slow = slow.Next
//		if fast == slow {
//			break
//		}
//	}
//
//	fast = head
//	for fast != slow {
//		fast = fast.Next
//		slow = slow.Next
//	}
//	return fast
//}

func main() {

}
