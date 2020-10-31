package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(3)
	obj.Insert(2)
	obj.Insert(1)
	obj.Insert(2)
	obj.Insert(3)
	obj.Insert(1)

	obj.Remove(2)
	obj.Remove(2)
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}

type RandomizedCollection struct {
	m    map[int]map[int]int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		m:    make(map[int]map[int]int),
		nums: make([]int, 0),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	res := true
	this.nums = append(this.nums, val)
	if _, exist := this.m[val]; exist {
		res = false
		this.m[val][len(this.nums)-1] = -1
	} else {
		this.m[val] = map[int]int{len(this.nums) - 1: -1}
	}
	fmt.Println(this)
	return res
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	res := false
	if _, exist := this.m[val]; exist {
		for k := range this.m[val] {
			res = true
			this.nums[k], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[k]
			this.nums = this.nums[:len(this.nums)-1]
			delete(this.m[val], k)
			if len(this.m[val]) == 0 {
				delete(this.m, val)
			}
			break
		}
	}
	fmt.Println(this)
	return res
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums)-1)]
}
