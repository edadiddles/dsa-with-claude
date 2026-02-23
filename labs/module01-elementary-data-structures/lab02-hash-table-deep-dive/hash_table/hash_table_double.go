package hash_table

type DoubleHashTable struct {
	tbl []HashTable
	f HashFunc
	m int
}

func (htbl *DoubleHashTable) Init(t HashFuncType, m int) {

}

func (htbl *DoubleHashTable) Insert(k, v int) {
}

func (htbl *DoubleHashTable) Search(k int) {
}

func (htbl *DoubleHashTable) Delete(k int) {
}
