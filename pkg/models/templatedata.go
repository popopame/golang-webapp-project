package models 

// TemplateData Hold data sent to template
type TemplateData struct{
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string

}