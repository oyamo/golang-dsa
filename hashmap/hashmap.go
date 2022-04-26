package hashmap

import (
	"hash/fnv"
)

type Node[V any] struct {
	Key  string
	Val  V
	Hash uint32
	Next *Node[V]
}

type HashMap[V any] struct {
	Capacity   uint32
	Size       uint32
	Buckets    []*Node[V]
	LoadFactor float32
}

func hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func NewHashMap[V any](capacity uint32, lf float32) *HashMap[V] {
	return &HashMap[V]{
		Capacity:   capacity,
		Size:       0,
		Buckets:    make([]*Node[V], capacity),
		LoadFactor: lf,
	}
}

func (hm *HashMap[V]) Resize() {
	hm.Size++
	if float32(hm.Size)/
		float32(hm.Capacity) >= hm.LoadFactor {
		hm.Capacity += 16
		newBucket := make([]*Node[V], hm.Capacity)
		_ = copy(newBucket, hm.Buckets)
		hm.Buckets = newBucket
	}
}

func (hm *HashMap[V]) Put(key string, val V) {
	h := hash(key)

	newNode := &Node[V]{
		Key:  key,
		Val:  val,
		Hash: h,
	}

	index := h & (hm.Capacity - 1)

	if hm.Buckets[index] == nil {
		hm.Buckets[index] = newNode
		hm.Resize()
	} else if hm.Buckets[index] != nil && hm.Buckets[index].Key == key {
		// Replace, so size doesn't change
		hm.Buckets[index] = newNode
	} else if hm.Buckets[index] != nil && hm.Buckets[index].Key != key {
		newNode.Next = hm.Buckets[index]
		hm.Buckets[index] = newNode
		hm.Resize()
	}
}

func (hm *HashMap[V]) Get(k string) (bool, *V) {
	h := hash(k) & (hm.Capacity - 1)
	if h > hm.Capacity || hm.Buckets[h] == nil {
		return false, nil
	}

	head := hm.Buckets[h]
	for head != nil {
		if head.Key == k {
			return true, &head.Val
		}
		head = head.Next
	}

	return false, nil
}

func (hm *HashMap[V]) Keys() []string {
	keys := make([]string, 0)
	for _, v := range hm.Buckets {
		head := v
		for head != nil {
			keys = append(keys, head.Key)
			head = head.Next
		}
	}

	return keys
}
