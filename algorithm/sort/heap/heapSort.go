package heap

type Heap struct {
	Size  int
	Array []int
}

func NewHeap(array []int) *Heap {
	h := new(Heap)
	h.Array = array
	return h
}

func (h *Heap) Push(x int) {
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}

	i := h.Size

	for i > 0 {
		parent := (i - 1) / 2
		if x <= h.Array[parent] {
			break
		}
		h.Array[i] = h.Array[parent]
		i = parent
	}
	h.Array[i] = x
	h.Size++
}

func (h *Heap) Pop() int {
	if h.Size == 0 {
		return -1
	}
	ret := h.Array[0]

	h.Size--
	x := h.Array[h.Size]
	h.Array[h.Size] = ret

	i := 0
	for {
		a := 2*i + 1
		b := 2*i + 2

		if a >= h.Size {
			break
		}

		if b < h.Size && h.Array[b] > h.Array[a] {
			a = b
		}

		if x >= h.Array[a] {
			break
		}

		h.Array[i] = h.Array[a]

		i = a
	}

	h.Array[i] = x
	return ret
}
