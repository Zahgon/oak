package debugstream

const (
	unmatched          = iota
	potentialDuplicate = iota
	matched            = iota
)

const (
	prefixLen = 4
)

// jaroDecreased is a lightly modified version of Jaro
// While it takes inspiration from JaroWinkler this seemed more fun.
// Since this is intended for commands the lengths of the strings will be short.
// Presuppose that users will miss the end of a command  or misappend extra data.
// Modified approach for domain that diverges from Jaro-Winkler's prefix strategy.
func jaroDecreased(candidate, registered string) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

//  probably shouldnt let it get to this step due to upstream constraints but adding for completeness

// denoted as match if within

// check for a potential match of every character in the candidate

// for our purposes lets be less mean to extra duplicates

// dont boost if its super low otherwise ful vs help will have a high boost.
