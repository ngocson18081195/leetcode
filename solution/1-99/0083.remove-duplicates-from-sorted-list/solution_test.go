package leetcode

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data1 := []int{1, 1, 2}
		got := deleteDuplicates(MakeListNote(data1))
		want := MakeListNote([]int{1, 2})

		if !Equal(got, want) {
			fmt.Print("GOT:")
			PrintList(got)
			fmt.Print("WANT:")
			PrintList(want)
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := deleteDuplicates(&ListNode{})
		want := &ListNode{}
		if !Equal(got, want) {
			fmt.Print("GOT:")
			PrintList(got)
			fmt.Print("WANT:")
			PrintList(want)
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}

func Equal(x, y *ListNode) bool {
	for x == nil || y == nil {
		if x == nil && y != nil {
			fmt.Println(x, y)
			return false
		}
		if x != nil && y == nil {
			fmt.Println(x, y)
			return false
		}

		if x.Val != y.Val {
			fmt.Println(x, y)
			return false
		}
		x = x.Next
		y = y.Next
	}
	return true
}

func PrintList(x *ListNode) {
	for x != nil {
		fmt.Print(x.Val, " ")
		x = x.Next
	}
	fmt.Println()
}

func MakeListNote(x []int) *ListNode {
	list := &ListNode{}
	head := list

	list.Val = x[0]
	for i := 1; i < len(x); i++ {
		list.Next = &ListNode{}
		list = list.Next
		list.Val = x[i]
	}
	return head
}
