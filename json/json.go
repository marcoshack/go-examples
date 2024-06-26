package json

type AStruct struct {
	Attr1   string         `json:"attribute1"`
	Another *AnotherStruct `json:"anotherStruct"`
}

type AnotherStruct struct {
	AnotherAttr1 string `json:"anotherAttribute1"`
}
