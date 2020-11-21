package gocopy

type mytype []int

func (s mytype) Copy() mytype {
	return append(s[:0:0], s...)
}

func (s mytype) CopyWithMake() mytype {
	var c = make(mytype, len(s))
	copy(c, s)
	return c
}
