package tags

// BSONTag returns the `bson` struct tag of the given struct's field
// or any errors that occur.
func BSONTag(data interface{}, fieldName string) (string, error) {
	return Tag(data, "bson", fieldName)
}

// MustHaveBSONTag returns the `bson` struct tag of the specified field.
// Panics if the requested tag of field does not exist.
func MustHaveBSONTag(data interface{}, fieldName string) string {
	return MustHaveTag(data, "bson", fieldName)
}
