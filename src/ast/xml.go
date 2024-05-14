package ast

import (
	"encoding/xml"
)


func (n *BaseNode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start_ := xml.StartElement{
		Name: xml.Name{Local: "BaseNode"},
		Attr: []xml.Attr{

		},
	}
	e.EncodeToken(start_)
	for _, child := range n.Children() {
		xml.Marshal(child)
	}
	return nil
}

func (dl *DeclarationList) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	start_ := xml.StartElement{Name: xml.Name{Local: "Decls"}}
	e.EncodeToken(start_)
	// e.Encode(dl.Decls)
	e.EncodeToken(start_.End())
	return nil
}

func (*StatementList) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*Binding) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*FunctionDefinition) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*CallExpr) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*OpExpr) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*IfExpression) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*NameAccess) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*LiteralInt) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*LiteralFloat) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*LiteralString) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}

func (*LiteralBool) MarshalXML(e xml.Encoder, start xml.StartElement) error {
	return nil
}
