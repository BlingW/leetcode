package main

import "fmt"

type MyCircularDeque struct {
	length int
	data   []int
	head   int
	rear   int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		length:k+1,
		data:make([]int,k+1), //空一个位置区分满和空
		head:0,
		rear:0,
	}
}

func maxSlidingWindow1(nums []int, k int) []int {
	l := len(nums)
	result := make([]int, 0)

	fmt.Println(nums)
	left, right := 0, 0
	dq := Constructor(l)
	for i := 0; i < l; i ++ {
		fmt.Println(left, right)
		if dq.IsEmpty() {
			goto rl
		}
		// 维护一个单调递减的队列
		// 对于每个元素 如果大于等于队列的最后一个元素，就把队尾的元素去除，直到队列为空或者队尾元素大于当前元素，再把当前元素追加到队尾
		// 队列第一个元素就是最大值
		if nums[i] >= nums[dq.GetRear()] {
			for !dq.IsEmpty() && nums[i] >= nums[dq.GetRear()] {
				dq.DeleteLast()
			}
		}
	rl:
		dq.InsertLast(i)
		if right-left < k - 1 {
			right++
		} else {
			result = append(result, nums[dq.GetFront()])
			left++
			right++
			if dq.GetFront() < left {
				dq.DeleteFront()
			}
		}
		fmt.Println("result:", result)
		fmt.Println("dq:", dq.data)
		fmt.Println("--------")
	}

	return result
}

func main() {
	nums := []int{1}
	fmt.Println(maxSlidingWindow1(nums, 1))
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		if this.rear == this.length-1 {
			this.rear = 0
		} else {
			this.rear++
		}
		this.data[this.head] = value
		return true
	}

	if this.head == 0 {
		this.head = this.length - 1
	} else {
		this.head--
	}
	this.data[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.data[this.rear] = value
		if this.rear == this.length-1 {
			this.rear = 0
		} else {
			this.rear++
		}
		return true
	}

	this.data[this.rear] = value
	if this.rear == this.length-1 {
		this.rear = 0
	} else {
		this.rear++
	}
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.length-1 {
		this.head = 0
	} else {
		this.head++
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.rear == 0 {
		this.rear = this.length - 1
	} else {
		this.rear--
	}
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]

}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.rear == 0 {
		return this.data[this.length-1]
	}
	return this.data[this.rear-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.rear
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.rear+1)%this.length == this.head
}
