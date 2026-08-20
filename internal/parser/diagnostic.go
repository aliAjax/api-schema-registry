package parser

type Diagnostic struct {
	Level, Code, Message string
	Line, Column         int
}

func Error(code, msg string, line, col int) Diagnostic {
	return Diagnostic{Level: "error", Code: code, Message: msg, Line: line, Column: col}
}
func Warning(code, msg string, line, col int) Diagnostic {
	return Diagnostic{Level: "warning", Code: code, Message: msg, Line: line, Column: col}
}
func HasErrors(ds []Diagnostic) bool {
	for _, d := range ds {
		if d.Level == "error" {
			return true
		}
	}
	return false
}
