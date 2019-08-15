package grx

//StringSlice StringSlice
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//byteSlice byteSlice
type byteSlice []byte

func (p byteSlice) Len() int           { return len(p) }
func (p byteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p byteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
