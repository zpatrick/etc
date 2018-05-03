# Errors

- dont' use `type T error`, since that will always pass type assertions.
Instead, use `type T struct` and have `func (t *T) Error() string`

- todo - 
