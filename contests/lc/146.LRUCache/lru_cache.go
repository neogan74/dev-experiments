package _146_LRUCache

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	size, capacity int
	head, tail     *Node
	cache          map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		cache:    make(map[int]*Node),
	}
}

func (lruc *LRUCache) Get(key int) int {
	if node, exists := lruc.cache[key]; exists {
		lruc.removeNode(node)
		lruc.addToHead(node)
		return node.val
	}
	return -1
}

func (lruc *LRUCache) Put(key int, value int) {
	if node, exists := lruc.cache[key]; exists {
		lruc.removeNode(node)
		node.val = value
		lruc.addToHead(node)
	} else {
		node := &Node{key: key, val: value}
		lruc.cache[key] = node
		lruc.addToHead(node)
		if lruc.size++; lruc.size > lruc.capacity {
			node = lruc.tail.prev
			delete(lruc.cache, node.key)
			lruc.removeNode(node)
			lruc.size--
		}
	}
}

func (lruc *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lruc *LRUCache) addToHead(node *Node) {
	node.next = lruc.head.next
	node.prev = lruc.head
	lruc.head.next = node
	node.next.prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
