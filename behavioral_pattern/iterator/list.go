package iterator

type List interface {
	Iterator() Iterator
	Length() int
	GetItem(index int) int
}

type ArrayList struct {
	nums []int
}

func NewArrayList(vals ...int) *ArrayList {
	nums := make([]int, len(vals))
	for i, num := range vals {
		nums[i] = num
	}
	return &ArrayList{nums: nums}
}

func (a *ArrayList) Iterator() Iterator {
	return NewArrayListIterator(a)
}

func (a *ArrayList) Length() int {
	return len(a.nums)
}

func (a *ArrayList) GetItem(index int) int {
	return a.nums[index]
}
