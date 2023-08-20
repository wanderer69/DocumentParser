package parser

import (
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
)

func getNodeType(node ast.Node) string {
	s := fmt.Sprintf("%T", node)
	s = strings.TrimSuffix(s, "()")
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx+1:]
	}
	return s
}

type DocumentItem struct {
	Type       string
	HeaderName string
	Level      int
	Text       string
	Url        string
	Image      string

	DocumentItems []*DocumentItem
}

type Document struct {
	DocumentItems []*DocumentItem
	//	parent        *DocumentItem
}

func (d *Document) getNodeRecur(node ast.Node, parent *DocumentItem, fn func(parent *DocumentItem, item *DocumentItem, depht int), depth int) {
	if node == nil {
		return
	}

	//content := shortenString(getContent(node), 40)
	typeName := getNodeType(node)
	switch v := node.(type) {
	case *ast.Link:
		//fmt.Printf("Link %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Url: string(v.Destination)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.Image:
		//fmt.Printf("Image %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Image: string(v.Destination)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.List:
		//fmt.Printf("List %#v\r\n", v)
		di := &DocumentItem{Type: typeName}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
		/*
			if v.Start > 1 {
				content += fmt.Sprintf("start=%d ", v.Start)
			}
			if v.Tight {
				content += "tight "
			}
			if v.IsFootnotesList {
				content += "footnotes "
			}
			flags := getListFlags(v.ListFlags)
			if len(flags) > 0 {
				content += "flags=" + flags + " "
			}
		*/
	case *ast.ListItem:
		//fmt.Printf("ListItem %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
		/*
			if v.Tight {
				content += "tight "
			}
			if v.IsFootnotesList {
				content += "footnotes "
			}
			flags := getListFlags(v.ListFlags)
			if len(flags) > 0 {
				content += "flags=" + flags + " "
			}
		*/
	case *ast.CodeBlock:
		//fmt.Printf("CodeBlock %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.Heading:
		//fmt.Printf("Heading %#v\r\n", v)
		di := &DocumentItem{Type: typeName, HeaderName: v.HeadingID}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.Text:
		//fmt.Printf("Text %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.HorizontalRule:
		//fmt.Printf("HorizontalRule %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.Paragraph:
		//fmt.Printf("Paragraph %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.Table:
		//fmt.Printf("Table %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.TableBody:
		//fmt.Printf("TableBody %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.TableRow:
		//fmt.Printf("TableRow %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.TableCell:
		//fmt.Printf("TableCell %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		if parent != nil {
			parent.DocumentItems = append(parent.DocumentItems, di)
		}
		parent = di
	case *ast.Document:
		//fmt.Printf("Document %#v\r\n", v)
		di := &DocumentItem{Type: typeName, Text: string(v.Literal)}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		d.DocumentItems = append(d.DocumentItems, di)
		parent = di
	default:
		fmt.Printf("TypeName %v\r\n", typeName)
	}
	for _, child := range node.GetChildren() {
		d.getNodeRecur(child, parent, fn, depth+1)
	}
}

func ParseMD(md []byte, fn func(parent *DocumentItem, item *DocumentItem, depht int)) (*Document, error) {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)
	//printRecur(doc, "\t", 0)
	d := &Document{}
	d.getNodeRecur(doc, nil, fn, 0)

	return d, nil
}
