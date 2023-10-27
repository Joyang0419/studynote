package hashTable

import (
	"math/rand"
	"time"
)

/*
Insert:
時間複雜度: O(1)
查詢 map 是否包含元素的時間複雜度為 O(1)
在 slice 尾部添加元素的時間複雜度為 O(1)
在 map 中添加元素的時間複雜度為 O(1)
空間複雜度: O(1) (如果不考慮隨著元素增加而增加的儲存空間)

Remove:
時間複雜度: O(1)
查詢 map 是否包含元素的時間複雜度為 O(1)
從 slice 中移除元素(通過交換最後一個元素)的時間複雜度為 O(1)
更新 map 中的元素的時間複雜度為 O(1)
空間複雜度: O(1)

GetRandom:
時間複雜度: O(1)
從 slice 中隨機選擇一個元素的時間複雜度為 O(1)
空間複雜度: O(1)

整體空間複雜度:
map 會根據存儲的元素數量增加，其空間複雜度為 O(N)，其中 N 為插入的元素數。
slice 也會隨著存儲的元素數量增加，其空間複雜度為 O(N)。
總的空間複雜度為 O(N)。

結論:
時間複雜度: O(1) (對於每一個操作)
空間複雜度: O(N) (N 為插入的元素數)

desc:
使用一個 map 來存儲元素和它在 slice 中的索引。
使用一個 slice 來存儲元素。
Insert 操作時，只需在 slice 尾部添加元素，並在 map 中更新該元素的索引。
Delete 操作時，可以將要刪除的元素與 slice 的最後一個元素交換，然後刪除最後一個元素，並在 map 中更新該元素的索引。
GetRandom 操作時，只需隨機選取 slice 中的一個索引即可。
*/

type RandomizedSet struct {
	data     []int
	position map[int]int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		position: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.position[val]; exists {
		return false
	}
	this.data = append(this.data, val)
	this.position[val] = len(this.data) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.position[val]
	if !exists {
		return false
	}
	last := this.data[len(this.data)-1]

	this.data[idx] = last
	this.position[last] = idx

	delete(this.position, val)
	this.data = this.data[:len(this.data)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
