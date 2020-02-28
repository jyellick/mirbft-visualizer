package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *Actions) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "table", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "table table-sm"}}}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "thead", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
			{
				vghtml := fmt.Sprint("\x3Ctr\x3E\x3Cth\x3EAction\x3C/th\x3E\x3Cth\x3EOutstanding\x3C/th\x3E\x3C/tr\x3E")
				vgn.InnerHTML = &vghtml
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tbody", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Broadcasts")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint(len(c.Actions.Broadcast))
						vgn.InnerHTML = &vghtml
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Unicasts")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint(len(c.Actions.Unicast))
						vgn.InnerHTML = &vghtml
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Preprocess")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint(len(c.Actions.Preprocess))
						vgn.InnerHTML = &vghtml
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Process")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint(len(c.Actions.Process))
						vgn.InnerHTML = &vghtml
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("QEntries")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint(len(c.Actions.QEntries))
						vgn.InnerHTML = &vghtml
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("PEntries")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint(len(c.Actions.PEntries))
						vgn.InnerHTML = &vghtml
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Commits")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint(len(c.Actions.Commits))
						vgn.InnerHTML = &vghtml
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}, vugu.VGAttribute{Namespace: "", Key: "style", Val: "fontWeight:\"bold\""}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Total")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint(ActionsLength(c.Actions))
						vgn.InnerHTML = &vghtml
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n"}
		vgparent.AppendChild(vgn)
	}
	return vgout
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ vjson.RawMessage
var _ js.Value
