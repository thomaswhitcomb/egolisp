package main

type envir map[string]*cell

type envirs struct {
	prev *envirs
	env  map[string]*cell
}

func newEnvirs(e *envirs) *envirs {
	ne := &envirs{e, make(envir)}
	ne.prev = e
	return ne
}
func (es envirs) find(s string) *cell {
	if cellp, ok := es.env[s]; ok {
		return cellp
	} else {
		if es.prev != nil {
			return es.prev.find(s)
		} else {
			return nil
		}
	}
}
func (es envirs) put(s string, v *cell) {
	es.env[s] = v
}
