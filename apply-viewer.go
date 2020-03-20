package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *ApplyViewer) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "w-100 card"}}}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n   "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "card-header w-100 p-0"}, vugu.VGAttribute{Namespace: "", Key: "role", Val: "group"}}}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "text-left btn btn-outline-primary w-100"}, vugu.VGAttribute{Namespace: "", Key: "type", Val: "button"}, vugu.VGAttribute{Namespace: "", Key: "data-toggle", Val: "collapse"}, vugu.VGAttribute{Namespace: "", Key: "aria-expanded", Val: "false"}}}
			vgparent.AppendChild(vgn)
			vgn.Attr = append(vgn.Attr, vugu.VGAttribute{Key: "aria-controls", Val: fmt.Sprint(fmt.Sprintf("%p", c.Apply))})
			vgn.Attr = append(vgn.Attr, vugu.VGAttribute{Key: "data-target", Val: fmt.Sprint(fmt.Sprintf("#%p", c.Apply))})
			{
				vghtml := fmt.Sprint(c.ApplySummary)
				vgn.InnerHTML = &vghtml
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "mp-0 collapse"}}}
		vgparent.AppendChild(vgn)
		vgn.Attr = append(vgn.Attr, vugu.VGAttribute{Key: "id", Val: fmt.Sprint(fmt.Sprintf("%p", c.Apply))})
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n     "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "card-body"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
				vgparent.AppendChild(vgn)
				if len(c.Actions.Preprocess) > 0 {
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "strong", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vghtml := fmt.Sprint("Preprocess")
							vgn.InnerHTML = &vghtml
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
						vgparent.AppendChild(vgn)
						for i, preprocess := range c.Actions.Preprocess {
							var vgiterkey interface{} = i
							_ = vgiterkey
							i := i
							_ = i
							preprocess := preprocess
							_ = preprocess
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row"}}}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
								vgparent.AppendChild(vgn)
								{
									vgcompKey := vugu.MakeCompKey(0x5E74BD9B8C2C4834, vgiterkey)
									// ask BuildEnv for prior instance of this specific component
									vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*PreprocessViewer)
									if vgcomp == nil {
										// create new one if needed
										vgcomp = new(PreprocessViewer)
									}
									vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
									vgcomp.Request = preprocess
									vgcomp.Result = c.Apply.Preprocessed[i]
									vgout.Components = append(vgout.Components, vgcomp)
									vgn = &vugu.VGNode{Component: vgcomp}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
								vgparent.AppendChild(vgn)
							}
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n     "}
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
