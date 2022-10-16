package validators

// IValidatable interface
type IValidatable interface {
	IsValid(insc string) bool
}
