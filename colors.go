package gops

// Color contains some colors for terminal representation.
var Color = struct {
	Blue   string
	Green  string
	Yellow string
	Red    string
	Nul    string
}{
	Blue:   "\033[94m",
	Green:  "\033[92m",
	Yellow: "\033[93m",
	Red:    "\033[91m",
	Nul:    "\033[0m",
}
