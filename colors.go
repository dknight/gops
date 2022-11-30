package main

// Color contains some colors for terminal representation.
var Color = struct {
	Blue   string
	Yellow string
	Green  string
	Red    string
	Nul    string
}{
	Blue:   "\033[94m",
	Yellow: "\033[93m",
	Green:  "\033[92m",
	Red:    "\033[91m",
	Nul:    "\033[0m",
}
