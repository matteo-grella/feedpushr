package model

// FilterProps constain properties of a filter
type FilterProps map[string]interface{}

// Filter is the filter interface
type Filter interface {
	DoFilter(article *Article) error
	GetDef() FilterDef
}

// FilterDefCollection is an array of filter definition
type FilterDefCollection []*FilterDef

// FilterDef contains filter definition
type FilterDef struct {
	ID int `json:"id"`
	Spec
	Tags    []string    `json:"tags,omitempty"`
	Props   FilterProps `json:"props:omitempty"`
	Enabled bool        `json:"enabled"`
}
