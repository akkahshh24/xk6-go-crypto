package statement

import "go.k6.io/k6/js/modules"

func init() {
	modules.Register("k6/x/xk6-go-crypto", new(Statement))
}

type Statement struct {
	StatementResult string
}

func (s *Statement) GetStatement() string {
	return "Hello from xk6-go-crypto extension!"
}
