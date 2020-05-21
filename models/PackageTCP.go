package models

type From struct {
	Hash string
	Addr string
}

type To struct {
	Addr string 
}

type Head struct {
	Title string
	Mode string
}

type Body struct {
	Data string
	Branch string
}

type PackageTCP struct {
	From From
	To To
	Head Head
	Body Body
}
