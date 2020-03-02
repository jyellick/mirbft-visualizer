package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *NodeControl) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n"}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "form", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "label", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				{
					vghtml := fmt.Sprint("Processing ")
					vgn.InnerHTML = &vghtml
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "select", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-control"}, vugu.VGAttribute{Namespace: "", Key: "selected", Val: "manual"}}}
				vgparent.AppendChild(vgn)
				vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
					EventType:	"click",
					Func:		func(event *vugu.DOMEvent) { c.SwitchProcessing(event) },
					// TODO: implement capture, etc. mostly need to decide syntax
				})
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " <option selected value=\"0\">Automatic (immediate)</option> "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "0"}}}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Automatic (immediate)")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "500ms"}}}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Automatic (500ms delay)")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "selected", Val: ""}, vugu.VGAttribute{Namespace: "", Key: "value", Val: "1000ms"}}}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Automatic (1000ms delay)")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "2000ms"}}}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Automatic (2000ms delay)")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "manual"}}}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("Manual")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn-group w-100"}, vugu.VGAttribute{Namespace: "", Key: "role", Val: "group"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n     "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "type", Val: "submit"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn btn-primary"}}}
					vgparent.AppendChild(vgn)
					vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
						EventType:	"click",
						Func:		func(event *vugu.DOMEvent) { c.Process(event) },
						// TODO: implement capture, etc. mostly need to decide syntax
					})
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Process"}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "label", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				{
					vghtml := fmt.Sprint("Ticking ")
					vgn.InnerHTML = &vghtml
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "select", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-control"}}}
				vgparent.AppendChild(vgn)
				vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
					EventType:	"click",
					Func:		func(event *vugu.DOMEvent) { c.SwitchTicking(event) },
					// TODO: implement capture, etc. mostly need to decide syntax
				})
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "1s"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Automatic (1s delay)"}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "selected", Val: ""}, vugu.VGAttribute{Namespace: "", Key: "value", Val: "5s"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Automatic (5s delay)"}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "15s"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Automatic (15s delay)"}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "manual"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Manual"}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn-group w-100"}, vugu.VGAttribute{Namespace: "", Key: "role", Val: "group"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n     "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "type", Val: "submit"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn btn-primary"}}}
					vgparent.AppendChild(vgn)
					vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
						EventType:	"click",
						Func:		func(event *vugu.DOMEvent) { c.Tick(event) },
						// TODO: implement capture, etc. mostly need to decide syntax
					})
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Tick"}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n"}
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
