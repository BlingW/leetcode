package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, 0)
	dq := NewDeque(len(nums))
	left, right := 0, 0
	for i := range nums {
		if !dq.IsEmpty() {
			for nums[i] > nums[dq.Tail().val] {
				dq.DeleteLast()
				if dq.IsEmpty() {
					break
				}
			}
		}
		dq.InsertLast(i)
		if i < k - 1 {
			right++
		} else {
			result = append(result, nums[dq.head.val])
			left++
			right++
			if deque.Head != nil && dq.head.val < left {
				dq.DeleteFront()
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

type duLNode struct {
	val int
	pre *duLNode
	nex *duLNode
}

type deque struct {
	head   *duLNode
	tail   *duLNode
	curLen int
	maxLen int
}

func NewDeque(l int) *deque {
	return &deque{
		curLen: 0,
		maxLen: l,
	}
}

func (dq deque) IsEmpty() bool {
	if dq.head == nil && dq.tail == nil {
		return true
	}
	return false
}

func (dq deque) IsFull() bool {
	if dq.curLen == dq.maxLen {
		return true
	}
	return false
}

func (dq deque) Head() *duLNode {
	return dq.head
}

func (dq deque) Tail() *duLNode {
	return dq.tail
}

func (dq *deque) InsertFront(i int) bool {
	if dq.IsFull() {
		return false
	}
	dln := &duLNode{
		val: i,
	}
	if dq.IsEmpty() {
		dq.head, dq.tail = dln, dln
		dq.curLen++
		return true
	}
	dln.nex = dq.head
	dq.head.pre = dln
	dq.head = dln
	dq.curLen++
	return true
}

func (dq *deque) InsertLast(i int) bool {
	if dq.IsFull() {
		return false
	}
	dln := &duLNode{
		val: i,
	}
	if dq.IsEmpty() {
		dq.head, dq.tail = dln, dln
		dq.curLen++
		return true
	}
	dln.pre = dq.tail
	dq.tail.nex = dln
	dq.tail = dln
	dq.curLen++
	return true
}

func (dq *deque) DeleteFront() bool {
	if dq.IsEmpty() {
		return false
	}

	if dq.head.nex == nil {
		dq.head, dq.tail = nil, nil
		dq.curLen--
		return true
	}

	lastHead := dq.head
	dq.head = lastHead.nex
	lastHead.nex = nil
	dq.head.pre = nil
	dq.curLen--

	return true
}

func (dq *deque) DeleteLast() bool {
	if dq.IsEmpty() {
		return false
	}

	if dq.tail.pre == nil {
		dq.head, dq.tail = nil, nil
		dq.curLen--
		return true
	}

	lastTail := dq.tail
	dq.tail = lastTail.pre
	lastTail.pre = nil
	dq.tail.nex = nil
	dq.curLen--

	return true
}

func (dq deque) ToList() []int {
	if dq.IsEmpty() {
		return nil
	}

	res := make([]int, dq.curLen)
	cur := dq.Head()
	i := 0
	for cur != nil {
		res[i] = cur.val
		cur = cur.nex
		i++
	}

	return res
}
