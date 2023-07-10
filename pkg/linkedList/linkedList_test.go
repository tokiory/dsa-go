package linkedList_test

import (
	"algo/pkg/linkedList"
	"testing"
)

func TestNew(t *testing.T) {
	initialSlice := []int{1, 2, 3}
	ll := linkedList.New(initialSlice)
	resultSlice := make([]int, 0, 3)

	ll.Iterate(func(_ int, node *linkedList.Node[int]) {
		resultSlice = append(resultSlice, node.Value)
	})

	for idx := range initialSlice {
		if initialSlice[idx] != resultSlice[idx] {
			t.Fail()
		}
	}
}

func TestLinkedList_Append_First(t *testing.T) {
	ll := linkedList.LinkedList[int]{}
	const NewValue = 1
	ll.Append(linkedList.Node[int]{Value: NewValue})

	if ll.Head.Value != NewValue {
		t.Fail()
	}
}

func TestLinkedList_Append(t *testing.T) {
	ll := linkedList.LinkedList[int]{}
	const NewValue = 1
	const Length = 5

	for idx := 0; idx < Length; idx++ {
		ll.Append(linkedList.Node[int]{Value: NewValue})
	}

	var calculatedLength int
	ll.Iterate(func(_ int, node *linkedList.Node[int]) {
		if node.Value != NewValue {
			t.Fatalf("Wrong value in node")
		}
		calculatedLength++
	})

	if calculatedLength != Length {
		t.Fail()
	}
}

func TestLinkedList_Remove(t *testing.T) {
	ll := linkedList.New([]int{1, 2, 3})
	if last := ll.Remove(); last.Value != 3 {
		t.Fatal("Last element is not 3", last)
	}

	// Check last element
	if last := ll.Last(); last.Value != 2 {
		t.Fatal("Last element is not 2, after deleting", last)
	}

	// Check length
	var length int
	ll.Iterate(func(_ int, n *linkedList.Node[int]) {
		length++
	})

	if length != 2 {
		t.Fatal("Length of linked list is not 2")
	}
}

func TestLinkedList_Find(t *testing.T) {
	ll := linkedList.New([]int{11, 22, 33, 44})

	// Search for existing value
	if found, value := ll.Find(11); !found {
		t.Fatal("Not found an existing element with value 11", value)
	} else {
		t.Log("Found value:", *value)
	}

	// Search for non-existing value
	if found, value := ll.Find(111); found {
		t.Fatal("Found non existing value :(", value)
	}
}

func TestLinkedList_Last(t *testing.T) {
	ll := linkedList.New([]int{1, 2, 3})
	if last := ll.Last(); last.Value != 3 {
		t.Fail()
	}
}

func TestLinkedList_Last_Empty(t *testing.T) {
	ll := linkedList.LinkedList[int]{}
	if last := ll.Last(); last != nil {
		t.Fail()
	}
}
