package filelister

type MarkDownSegment struct {
	Children []MarkDownSegment
	Path     string
	Name     string
	IsToc    bool
}
