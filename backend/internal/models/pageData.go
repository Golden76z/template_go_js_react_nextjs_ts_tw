package models

// PageData: data structure passed to templates, it holds the common data for all pages
type PageData struct {
	Title     string
	Header    string
	Content   interface{}
	IsError   bool
	ErrorCode int
}

// CustomError defines a custom error type with status code and message
type CustomError struct {
	StatusCode int
	Message    string
}

// Error Interface, implements the error interface for CustomError
func (e *CustomError) Error() string {
	return e.Message
}
