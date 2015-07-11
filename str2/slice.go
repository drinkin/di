package str2

// Slice is and alias for a slice of strings
type Slice []string

// Contains returns true if `s` contains `v`
func (s Slice) Contains(v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}
