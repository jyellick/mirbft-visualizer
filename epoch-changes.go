package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *EpochChanges) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

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
				vghtml := fmt.Sprint("\x3Ctr\x3E\n      \x3Cth\x3EEpoch\x3C/th\x3E\n      \x3Cth\x3EEpoch Changes\x3C/th\x3E\n      \x3Cth\x3EEchos\x3C/th\x3E\n      \x3Cth\x3EReadies\x3C/th\x3E\n      \x3Cth\x3ESuspicions\x3C/th\x3E\n    \x3C/tr\x3E")
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
				for _, et := range c.MirNode.Status.EpochChanger.EpochTargets {
					var vgiterkey interface{} = fmt.Sprintf("epoch-%d", et.Number)
					_ = vgiterkey
					et := et
					_ = et
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "sm"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vghtml := fmt.Sprint(et.Number)
							vgn.InnerHTML = &vghtml
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vghtml := fmt.Sprint(fmt.Sprintf("%v", prettyEpochChanges(et.EpochChanges)))
							vgn.InnerHTML = &vghtml
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vghtml := fmt.Sprint(fmt.Sprintf("%v", et.Echos))
							vgn.InnerHTML = &vghtml
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vghtml := fmt.Sprint(fmt.Sprintf("%v", et.Readies))
							vgn.InnerHTML = &vghtml
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vghtml := fmt.Sprint(fmt.Sprintf("%v", et.Suspicions))
							vgn.InnerHTML = &vghtml
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
						vgparent.AppendChild(vgn)
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
