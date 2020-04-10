package p21_merge_two_sorted_lists

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	n104 := ListNode{7, nil}
	n103 := ListNode{5, &n104}
	n102 := ListNode{3, &n103}
	n101 := ListNode{1, &n102}
	l1 := &n101
	n203 := ListNode{6, nil}
	n202 := ListNode{4, &n203}
	n201 := ListNode{2, &n202}
	l2 := &n201
	PrintList(l1)
	PrintList(l2)
	res := mergeTwoLists(l1, l2)
	PrintList(res)
}
