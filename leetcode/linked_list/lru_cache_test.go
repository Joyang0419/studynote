package linked_list

import (
	"container/list"
)

/*
time complexity:
查找 (Get 操作): 由于使用哈希表来存储键值对的引用，
查找的时间复杂度是 O(1)。
但是，在查找后还需要将找到的元素移动到双向链表的前面，
这个操作也是 O(1) 的时间复杂度，因为双向链表的元素移动是常数时间的操作。

插入/更新 (Put 操作):
如果是更新已存在的键的值，
那么哈希表的查找时间复杂度是 O(1)，
移动元素到双向链表前面的时间复杂度也是 O(1)。
如果是插入新的键值对，
那么哈希表的插入时间复杂度是 O(1)，
双向链表的插入时间复杂度也是 O(1)。

但是，如果缓存已满，
还需要从双向链表和哈希表中删除最少使用的元素，这个操作的时间复杂度也是 O(1)。

space complexity:
哈希表的空间复杂度: 是 O(n)，其中 n 是缓存的容量。
双向链表的空间复杂度: 也是 O(n)，其中 n 是缓存的容量。
综合来看，LRU 缓存的总空间复杂度是 O(n)。

description:
LRU (Least Recently Used)
缓存的实现通常涉及两个主要数据结构：一个哈希表用于快速查找，
和一个双向链表用于快速添加、删除和移动元素。下面是对时间复杂度和空间复杂度的分析

capacity來表示缓存的容量，
一個cache哈希表來保存每個鍵值對，並且用一個list雙向鏈表來保持鍵值對的順序。
Get和Put方法分別用來獲取和設置鍵值對，並且保證在每次操作時，
最近訪問或插入的鍵值對總是放在雙向鏈表的前面。
當缓存超過容量時，我們從雙向鏈表的後面刪除最少使用的鍵值對。

快速摘要:
雙向鍊錶 && 一本map去紀錄(map[int]*list.Element)
*/

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if element, ok := c.cache[key]; ok {
		c.list.MoveToFront(element)
		return element.Value.(Pair).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if element, ok := c.cache[key]; ok {
		element.Value = Pair{key: key, value: value}
		c.list.MoveToFront(element)
	} else {
		// 當長度超越cache設定的容量時
		if c.list.Len() >= c.capacity {
			// 刪除map中最後一個 linked list的key
			delete(c.cache, c.list.Back().Value.(Pair).key)
			// 刪除最後一個element
			c.list.Remove(c.list.Back())
		}
		c.list.PushFront(Pair{key: key, value: value})
		c.cache[key] = c.list.Front()
	}
}
