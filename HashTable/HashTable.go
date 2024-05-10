package HashTable

import (
	"cmp"
	"hash/fnv"
	"unsafe"
)

type TableItem[K cmp.Ordered, V any] struct {
	key  K
	data V
	next *TableItem[K, V]
}

type HashTable[K cmp.Ordered, V any] struct {
	data [256]*TableItem[K, V]
}

func (table *HashTable[K, V]) Add(key K, i V) {
	position := generateHash(key)
	if table.data[position] == nil {
		table.data[position] = &TableItem[K, V]{key: key, data: i}
		return
	}
	current := table.data[position]
	for current.next != nil {
		current = current.next
	}
	current.next = &TableItem[K, V]{key: key, data: i}
}

func (table *HashTable[K, V]) Get(key K) (V, bool) {
	position := generateHash(key)
	current := table.data[position]
	for current != nil {
		if cmp.Compare(current.key, key) == 0 {
			return current.data, true
		}
		current = current.next
	}
	return 0, false
}

func (table *HashTable[K, V]) Set(key K, value V) bool {
	position := generateHash(key)
	current := table.data[position]
	for current != nil {
		if cmp.Compare(current.key, key) == 0 {
			current.data = value
			return true
		}
		current = current.next
	}
	return false
}

func (table *HashTable[K, V]) Remove(key K) bool {
	position := generateHash(key)
	if table.data[position] == nil {
		return false
	}
	if cmp.Compare(table.data[position].key, key) == 0 {
		table.data[position] = table.data[position].next
		return true
	}
	current := table.data[position]
	for current.next != nil {
		if cmp.Compare(current.next.key, key) == 0 {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func generateHash[K any](s K) uint8 {
	hash := fnv.New32a()
	b := unsafe.Slice(&s, unsafe.Sizeof(s))
	hash.Write(b)
	return uint8(hash.Sum32() % 256)
}
