package ast

import (
	"encoding/xml"
)

func (n BaseNode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start_ := xml.StartElement{Name: xml.Name{Local: "BaseNode"}}
	e.EncodeToken(start_)
	for _, child := range n.Children() {
		e.Encode(child)
	}
	return nil
}

func (defn FunctionDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start_ := xml.StartElement{Name: xml.Name{Local: "FunctionDefinition"}}
	e.EncodeToken(start_)

	name := xml.StartElement{Name: xml.Name{Local: "Name"}}
	e.EncodeToken(name)
	e.Encode(defn.Name)
	e.EncodeToken(name.End())

	params := xml.StartElement{Name: xml.Name{Local: "Parameters"}}
	e.EncodeToken(params)
	for _, param := range defn.Parameters {
		e.Encode(param)
	}
	e.EncodeToken(params.End())

	e.Encode(defn.Body.StatementList)
	// body := xml.StartElement{Name: xml.Name{Local: "Body"}}
	// e.EncodeToken(body)
	// for _, stmt := range defn.Body.Children() {
	// 	e.Encode(stmt)
	// }
	// e.EncodeToken(body.End())

	e.EncodeToken(start_.End())
	return nil
}

func (sl StatementList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	stmts := xml.StartElement{Name: xml.Name{Local: "StatementList"}}
	e.EncodeToken(stmts)
	e.EncodeToken(stmts.End())
	return nil
}

func (b Binding) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (call CallExpr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (expr OpExpr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (ifExpr IfExpression) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (na NameAccess) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (li LiteralInt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (lf LiteralFloat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (ls LiteralString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (lb LiteralBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}
