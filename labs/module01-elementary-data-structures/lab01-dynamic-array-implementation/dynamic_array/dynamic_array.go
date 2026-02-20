package dynamic_array

type GrowthOp int

const (
	ADD GrowthOp = iota
	MULTIPLY
)

type GrowthStrategy struct {
	op     GrowthOp
	factor float32
}

type DynamicArray struct {
	buf          []int
	length       int
	growth_strat GrowthStrategy
}

func (arr *DynamicArray) Init(c int, f float32, op GrowthOp) {
	arr.buf = make([]int, c)
	arr.length = 0
	arr.growth_strat = GrowthStrategy{ op, f }
}

func (arr *DynamicArray) Push(value int) {
	arr.check_growth()
	
	arr.buf[arr.length] = value
	arr.length += 1
}

func (arr *DynamicArray) Pop() int {
	arr.length -= 1
	v := arr.buf[arr.length]
	return v
}

func (arr *DynamicArray) Get(index int) int {
	return arr.buf[index]
}

func (arr *DynamicArray) Set(index, value int) {
	arr.buf[index] = value
}

func (arr *DynamicArray) Insert(index, value int) {
	arr.check_growth()

	for i := arr.length; i > index; i-- {
		arr.buf[i] = arr.buf[i-1]
	}
	arr.buf[index] = value
	arr.length += 1
}

func (arr *DynamicArray) Delete(index int) {
	arr.length -= 1
	for i := index; i < arr.length; i++ {
		arr.buf[i] = arr.buf[i+1]
	}
}

func (arr *DynamicArray) Size() int {
	return arr.length
}

func (arr *DynamicArray) Capacity() int {
	return len(arr.buf)
}

func (arr *DynamicArray) check_growth() {
	if arr.Size() != arr.Capacity() {
		return
	}

	tmp := arr.buf
	cap := float32(arr.Capacity()*2)
	switch arr.growth_strat.op {
	case ADD:
		cap = float32(arr.Capacity()) + arr.growth_strat.factor
	case MULTIPLY:
		cap = float32(arr.Capacity()) * arr.growth_strat.factor
	}

	arr.buf = make([]int, int(cap))
	copy(arr.buf[0:len(tmp)], tmp[:])
}
