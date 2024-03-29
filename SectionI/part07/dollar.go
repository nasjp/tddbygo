package money

// Dollar is dollar
type Dollar struct {
	*Money
}

// NewDollar is dollar constructor
func NewDollar(a int) *Dollar {
	return &Dollar{
		&Money{
			amount: a,
		},
	}
}

// Times multiplier dollar
func (d *Dollar) Times(m int) *Dollar {
	return &Dollar{
		&Money{
			amount: d.amount * m,
		},
	}
}
