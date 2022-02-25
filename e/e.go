package e

import (
	"strconv"
	"strings"
)

type Element interface {
	Parse() string
}

type Elements []Element

func (e Elements) Parse() string {
	sb := strings.Builder{}
	for _, v := range e {
		sb.WriteString(v.Parse())
	}
	return sb.String()
}

type Int struct {
	Value int
}

func (i Int) Parse() string {
	return strconv.Itoa(i.Value)
}

type Float struct {
	Value float64
}

func (f Float) Parse() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

type Text struct {
	Value string
}

func (t Text) Parse() string {
	return t.Value
}

type Header struct {
	Level int
	Value Element
}

func (h Header) Parse() string {
	return "<h" + strconv.Itoa(h.Level) + ">" + h.Value.Parse() + "</h" + strconv.Itoa(h.Level) + ">"
}

type Paragraph struct {
	Value Element
}

func (p Paragraph) Parse() string {
	return "<p>" + p.Value.Parse() + "</p>"
}

type Italic struct {
	Value Element
}

func (i Italic) Parse() string {
	return "<i>" + i.Value.Parse() + "</i>"
}

type Bold struct {
	Value Element
}

func (b Bold) Parse() string {
	return "<b>" + b.Value.Parse() + "</b>"
}

type TableHeader struct {
	Values []Element
}

func (th TableHeader) Parse() string {
	sb := strings.Builder{}
	sb.WriteString("<thead>")
	sb.WriteString("<tr>")
	for _, v := range th.Values {
		sb.WriteString("<th>" + v.Parse() + "</th>")
	}
	sb.WriteString("</tr>")
	sb.WriteString("</thead>")
	return sb.String()
}

type TableData struct {
	Values []Element
}

func (td TableData) Parse() string {
	sb := strings.Builder{}
	sb.WriteString("<td>")
	for _, v := range td.Values {
		sb.WriteString(v.Parse())
	}
	sb.WriteString("</td>")
	return sb.String()
}

type TableRow struct {
	Values []TableData
}

func (tr TableRow) Parse() string {
	sb := strings.Builder{}
	sb.WriteString("<tr>")
	for _, v := range tr.Values {
		sb.WriteString(v.Parse())
	}
	sb.WriteString("</tr>")
	return sb.String()
}

type Table struct {
	Headers TableHeader
	Rows    []TableRow
}

func (t Table) Parse() string {
	sb := strings.Builder{}
	sb.WriteString("<table>")
	sb.WriteString(t.Headers.Parse())
	for _, v := range t.Rows {
		sb.WriteString(v.Parse())
	}
	sb.WriteString("</table>")
	return sb.String()
}

type Br struct{}

func (b Br) Parse() string {
	return "<br>"
}

type Hr struct{}

func (h Hr) Parse() string {
	return "<hr>"
}

type Body struct {
	Values []Element
}

func (b Body) Parse() string {
	sb := strings.Builder{}
	sb.WriteString("<body>")
	for _, v := range b.Values {
		sb.WriteString(v.Parse())
	}
	sb.WriteString("</body>")
	return sb.String()
}

type Head struct {
	Values []Element
}

func (h Head) Parse() string {
	sb := strings.Builder{}
	sb.WriteString("<head>")
	for _, v := range h.Values {
		sb.WriteString(v.Parse())
	}
	sb.WriteString("</head>")
	return sb.String()
}

type Script struct {
	Src string
}

func (s Script) Parse() string {
	return "<script src=\"" + s.Src + "\"></script>"
}

type Link struct {
	Href string
}

func (l Link) Parse() string {
	return "<link href=\"" + l.Href + "\" rel=\"stylesheet\">"
}

type Document struct {
	Head Head
	Body Body
}

func (d Document) Parse() string {
	sb := strings.Builder{}
	sb.WriteString("<!DOCTYPE html>")
	sb.WriteString("<html>")
	sb.WriteString(d.Head.Parse())
	sb.WriteString(d.Body.Parse())
	sb.WriteString("</html>")
	return sb.String()
}
