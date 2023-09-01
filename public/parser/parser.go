package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
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
	parentStack   []*DocumentItem
}

func (d *Document) Push(di *DocumentItem) {
	d.parentStack = append(d.parentStack, di)
}

func (d *Document) Pop() (*DocumentItem, bool) {
	lenStack := len(d.parentStack)
	if lenStack == 0 {
		return nil, false
	}
	di := d.parentStack[lenStack-1]
	d.parentStack = d.parentStack[0 : lenStack-1]
	return di, true
}

func (d *Document) Current() (*DocumentItem, bool) {
	lenStack := len(d.parentStack)
	if lenStack == 0 {
		return nil, false
	}
	di := d.parentStack[lenStack-1]
	return di, true
}

func (d *Document) getNodeRecur(node ast.Node, parent *DocumentItem, fn func(parent *DocumentItem, item *DocumentItem, depht int), depth int) error {
	if node == nil {
		return errors.New("empty node")
	}

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
		di := &DocumentItem{Type: typeName, HeaderName: v.HeadingID, Level: v.Level}
		if fn != nil {
			fn(parent, di, depth+1)
		}
		current, ok := d.Current()
		if !ok {
			return errors.New("bad header level")
		}
		if current.Level >= di.Level {
			level := current.Level
			for level >= di.Level {
				// поднимаемся на уровень выше
				var ok bool
				_, ok = d.Pop()
				if !ok {
					return errors.New("bad header level")
				}
				current, ok = d.Current()
				if !ok {
					return errors.New("bad header level")
				}
				level = current.Level
			}
			d.Push(di)
			current.DocumentItems = append(current.DocumentItems, di)
		} else {
			d.Push(di)
			current.DocumentItems = append(current.DocumentItems, di)
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
		current, ok := d.Current()
		if !ok {
			return errors.New("bad header level")
		}
		current.DocumentItems = append(current.DocumentItems, di)
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
		d.Push(di)
		parent = di
	default:
		fmt.Printf("TypeName %v\r\n", typeName)
	}
	for _, child := range node.GetChildren() {
		d.getNodeRecur(child, parent, fn, depth+1)
	}
	return nil
}

func ParseMD(md []byte, fn func(parent *DocumentItem, item *DocumentItem, depht int)) (*Document, error) {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)
	d := &Document{}
	err := d.getNodeRecur(doc, nil, fn, 0)

	return d, err
}
