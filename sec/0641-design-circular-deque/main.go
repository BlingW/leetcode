package main

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


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) *deque {
	return &deque{
		curLen: 0,
		maxLen: k,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (dq *deque) InsertFront(value int) bool {
	if dq.IsFull() {
		return false
	}
	dln := &duLNode{
		val: value,
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


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (dq *deque) InsertLast(value int) bool {
	if dq.IsFull() {
		return false
	}
	dln := &duLNode{
		val: value,
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


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
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


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
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


/** Get the front item from the deque. */
func (dq *deque) GetFront() int {
	if dq.head == nil {
		return -1
	}
	return dq.head.val
}


/** Get the last item from the deque. */
func (dq *deque) GetRear() int {
	if dq.tail == nil {
		return -1
	}
	return dq.tail.val
}


/** Checks whether the circular deque is empty or not. */
func (dq *deque) IsEmpty() bool {
	if dq.head == nil && dq.tail == nil {
		return true
	}
	return false
}


/** Checks whether the circular deque is full or not. */
func (dq *deque) IsFull() bool {
	if dq.curLen == dq.maxLen {
		return true
	}
	return false
}
