package hash_table


type HashTable []KVPairs

type KVPairs struct {
	key int
	val int
}

type ChainingHashTable struct {
	tbl []HashTable
	f HashFunc
	m int
}

func (htbl *ChainingHashTable) Init(t HashFuncType, m int) {
	htbl.f = SelectHashFunc(t)
	htbl.m = m
	htbl.tbl = make([]HashTable, m)
}

func (htbl *ChainingHashTable) Insert(k, v int) {
	idx := htbl.f(k, htbl.m)
	v_chain := htbl.tbl[idx]
	if v_chain == nil {
		v_chain = make(HashTable, 0)
	}
	htbl.tbl[idx] = append(v_chain, KVPairs{ key: k, val: v })
}

func (htbl *ChainingHashTable) Search(k int) int {
	idx := htbl.f(k, htbl.m)
	v_chain := htbl.tbl[idx]
	var v int
	for _, kv_pair := range v_chain {
		if kv_pair.key == k {
			v = kv_pair.val
			break
		}
	}

	return v
}

func (htbl *ChainingHashTable) Delete(k int) {
	idx := htbl.f(k, htbl.m)
	v_chain := htbl.tbl[idx]
	for i, kv_pair := range v_chain {
		if kv_pair.key == k {
			htbl.tbl[idx] = append(v_chain[:i], v_chain[i+1:]...)
			break
		}
	}
}
