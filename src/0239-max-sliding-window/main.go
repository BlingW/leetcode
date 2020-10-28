package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	result := make([]int, 0)

	fmt.Println(nums)
	left, right := 0, 0
	dq := NewDeque(l)
	for i := 0; i < l; i ++ {
		fmt.Println(left, right)
		if i > 0 && nums[i] >= nums[dq.Tail().val] {
			// 维护一个单调递减的队列
			// 对于每个元素 如果大于等于队列的最后一个元素，就把队尾的元素去除，直到队列为空或者队尾元素大于当前元素，再把当前元素追加到队尾
			// 队列第一个元素就是最大值
			for dq.Tail() != nil && nums[i] >= nums[dq.Tail().val] {
				dq.DeleteLast()
			}
		}
		dq.InsertLast(i)
		if right-left == k-1 {
			result = append(result, nums[dq.Head().val])
			left++
			right++
			if dq.Head().val < left {
				dq.DeleteFront()
			}
		} else {
			right++
		}
		fmt.Println("result:", result)
		fmt.Println("dq:", dq.ToList())
		fmt.Println("--------")
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
