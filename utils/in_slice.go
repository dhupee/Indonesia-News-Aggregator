package utils

// isInSlice checks if a given word is present in a slice of strings.
//
// Parameters:
// - word: a string representing the word to search for.
// - slice: a slice of strings to search in.
//
// Returns:
// - a boolean value indicating whether the word is present in the slice.
func IsInSlice(word string, slice []string) bool {
    for _, value := range slice {
        if value == word {
            return true
        }
    }
    return false
}
