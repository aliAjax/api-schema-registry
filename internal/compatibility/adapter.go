package compatibility

type Rule interface{ Check(Result) error }
