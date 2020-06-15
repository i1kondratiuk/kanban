package value

type BodyText string

// Comment represents a comment
type Comment struct {
	BodyText BodyText `json:"bodyText"`
}
