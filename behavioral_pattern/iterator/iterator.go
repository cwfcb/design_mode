package iterator

type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() int
}

type ArrayListIterator struct {
	cursor    int
	arrayList *ArrayList
}

func NewArrayListIterator(list *ArrayList) *ArrayListIterator {
	return &ArrayListIterator{
		cursor:    0,
		arrayList: list,
	}
}

func (a *ArrayListIterator) HasNext() bool {
	return a.cursor != a.arrayList.Length()
}

func (a *ArrayListIterator) Next() {
	a.cursor++
}

func (a *ArrayListIterator) CurrentItem() int {
	return a.arrayList.GetItem(a.cursor)
}
