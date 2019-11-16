package runtime

// for now noos/thumb supports only single CPU

type cpumtx struct{}

func (mt *cpumtx) lock()   {}
func (mt *cpumtx) unlock() {}
