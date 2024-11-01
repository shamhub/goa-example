package compute

import (
	"context"
	"log"

	calc "github.com/shamhub/goa-example/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Divide implements divide.
func (s *calcsrvc) Divide(ctx context.Context, p *calc.DividePayload) (res *calc.DivideResult, err error) {
	quotient := p.Dividend / p.Divisor
	remainder := p.Dividend % p.Divisor
	res = &calc.DivideResult{Quotient: quotient, Reminder: remainder}
	s.logger.Print("calc.divide")
	return
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	s.logger.Print("calc.add")
	return
}
