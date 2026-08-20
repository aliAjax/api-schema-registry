package validation

type FormatChecker interface{ Check(string, string) bool }
