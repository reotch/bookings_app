package models

// TemplateData stores data that might be sent to templates to be rendered
type TemplateData struct {
	StringMap map[string]string      // e.g. names, email addresses, etc.
	IntMap    map[string]int         // e.g. number of days
	FloatMap  map[string]float32     // e.g. dollar amounts, results of calcs
	Data      map[string]interface{} // covers any other data types; uses `interface{}` type
	CSRFToken string                 // Cross-Site Request Forgery token for web sec
	// Flashed messages to users
	Flash   string
	Warning string
	Error   string
}
