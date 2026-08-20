package parser

type Limits struct{ MaxBytes, MaxDepth, MaxRefs int }

func DefaultLimits() Limits { return Limits{MaxBytes: 2 << 20, MaxDepth: 32, MaxRefs: 100} }
