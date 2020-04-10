package p21_merge_two_sorted_lists

import "fmt"

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var c1 *ListNode = l1
	var c2 *ListNode = l2
	var mergedList *ListNode
	var p *ListNode
	for c1 != nil || c2 != nil {
		if c1 == nil && c2 != nil {
			move(&p, &mergedList, c2.Val)
			c2 = c2.Next
			continue
		}
		if c2 == nil && c1 != nil {
			move(&p, &mergedList, c1.Val)
			c1 = c1.Next
			continue
		}
		if c1.Val < c2.Val {
			move(&p, &mergedList, c1.Val)
			c1 = c1.Next
			continue
		} else {
			move(&p, &mergedList, c2.Val)
			c2 = c2.Next
			continue
		}
	}
	return mergedList
}

func move(p **ListNode, mergedList **ListNode, val int) {
	if *p == nil {
		*p = &ListNode{val, nil}
		*mergedList = *p
	} else {
		(*p).Next = &ListNode{val, nil}
		*p = (*p).Next
	}
}

func PrintList(l1 *ListNode) {
	tmp := l1
	for tmp != nil {
		fmt.Printf("%p -> %v |==> ", tmp, tmp)
		tmp = tmp.Next
	}
	fmt.Println("nil")
}
