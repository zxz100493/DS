package hash

import (
	"math"
	"sync"

	"github.com/cespare/xxhash/v2"
)

// refer https://cc.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6b0q1te6sq
const (
	expandFactor = 0.75
)

type HashMap struct {
	array        []*keyPairs
	capacity     int
	len          int
	capacityMask int
	lock         sync.Mutex
}

type keyPairs struct {
	key   string
	value interface{}
	next  *keyPairs
}

// init hashmap
func NewHashMap(capacity int) *HashMap {
	defaultCapacity := 1 << 4
	if capacity <= defaultCapacity {
		capacity = defaultCapacity
	} else {
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity)))))
	}

	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)
	m.capacity = capacity
	m.capacityMask = capacity - 1
	return m
}

func (m *HashMap) Len() int {
	return m.len
}

// get hash value of key
var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

func (m *HashMap) hashIndex(key string, mask int) int {
	hash := hashAlgorithm([]byte(key))
	index := hash & uint64(mask)
	return int(index)
}

func (m *HashMap) Put(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capacityMask)

	element := m.array[index]

	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		var lastPairs *keyPairs
		for element != nil {
			// if there has been an element already, update it
			if element.key == key {
				element.value = value
				return
			}
			lastPairs = element
			element = element.next
		}
		// if can't find element put new element to then end of hashmap
		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}
	newLen := m.len + 1

	// if over the expand factor expand it
	if float64(newLen)/float64(m.capacity) >= expandFactor {
		newM := new(HashMap)
		newM.array = make([]*keyPairs, 2*m.capacity, 2*m.capacity)
		newM.capacity = 2 * m.capacity
		newM.capacityMask = 2*m.capacity - 1

		// copy old hashmap to new
		for _, pairs := range m.array {
			for pairs != nil {
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
	}

	m.len = newLen
}

func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capacityMask)

	element := m.array[index]

	for element != nil {
		if element.key == key {
			return element.value, true
		}
		element = element.next
	}
	return
}

func (m *HashMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capacityMask)

	element := m.array[index]

	if element == nil {
		return
	}
	// 否则查看链表第一个元素的键是否匹配：element.key == key，如果匹配，
	// 那么对链表头部进行替换，链表的第二个元素补位成为链表头部：
	if element.key == key {
		m.array[index] = element.next
		m.len = m.len - 1
		return
	}

	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next
			m.len = m.len - 1
			return
		}
		element = nextElement
		nextElement = nextElement.next
	}
}
