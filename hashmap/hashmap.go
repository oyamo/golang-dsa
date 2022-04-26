package hashmap

type Node[K any, V any] struct {
	Key  K
	Val  V
	Hash uint32
	Next *Node
}

type HashMap[K any, V any] struct {
	Capacity uint32
	Size     uint32
	Buckets  []*Node
}

func NewHashMap[K any, V any](capacity uint32) *HashMap[K, V] {
	return &HashMap[K, V]{
		Capacity: capacity,
		Size:     0,
		Buckets:  make([]*Node, capacity),
	}
}
