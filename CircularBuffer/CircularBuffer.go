package CircularBuffer

const arraySize = 10

type CircularBuffer[T any] struct {
	data    [arraySize]T
	pointer int
}

func (b *CircularBuffer[T]) InsertValue(i T) {
	if b.pointer == len(b.data) {
		b.pointer = 0
	}
	b.data[b.pointer] = i
	b.pointer += 1
}

func (b *CircularBuffer[T]) GetValues() [arraySize]T {
	return b.data
}

func (b *CircularBuffer[T]) GetValuesFromPosition(i int) ([arraySize]T, bool) {
	var out [arraySize]T

	if i >= len(out) {
		return out, false
	}

	for u := 0; u < len(out); u++ {
		if i >= len(b.data) {
			i = 0
		}
		out[u] = b.data[i]
		i += 1
	}
	return out, true
}
