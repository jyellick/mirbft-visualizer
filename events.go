package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *Events) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row border rounded mb-2"}, vugu.VGAttribute{Namespace: "", Key: "style", Val: "border-width:3px"}}}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col"}}}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "h4", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				{
					vghtml := fmt.Sprint(fmt.Sprintf("Events (time is %d)", c.EventLog.FakeTime))
					vgn.InnerHTML = &vghtml
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row"}, vugu.VGAttribute{Namespace: "", Key: "style", Val: "margin-right: 10px"}}}
					vgparent.AppendChild(vgn)
					{
						vghtml := fmt.Sprint("\n          \x3Cdiv class=\"col-1\"\x3E\x3Cstrong\x3ETarget\x3C/strong\x3E\x3C/div\x3E\n          \x3Cdiv class=\"col-1\"\x3E\x3Cstrong\x3EType\x3C/strong\x3E\x3C/div\x3E\n          \x3Cdiv class=\"col-8\"\x3E\x3Cstrong\x3EDetails\x3C/strong\x3E\x3C/div\x3E\n          \x3Cdiv class=\"col-2\"\x3E\x3Cstrong\x3EOccursAt\x3C/strong\x3E\x3C/div\x3E\n        ")
						vgn.InnerHTML = &vghtml
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row mp-0"}, vugu.VGAttribute{Namespace: "", Key: "style", Val: "display:block;position:relative;height:200px;overflow-y:scroll;"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col mp-0"}}}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
							vgparent.AppendChild(vgn)
							for iterator, event := c.Events(); event != nil; event = iterator.Next() {
								var vgiterkey interface{} = iterator
								_ = vgiterkey
								iterator := iterator
								_ = iterator
								event := event
								_ = event
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row border rounded"}}}
								vgparent.AppendChild(vgn)
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
									vgparent.AppendChild(vgn)
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-1 my-auto"}}}
									vgparent.AppendChild(vgn)
									{
										vghtml := fmt.Sprint(event.Event.Target)
										vgn.InnerHTML = &vghtml
									}
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
									vgparent.AppendChild(vgn)
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-1 my-auto"}}}
									vgparent.AppendChild(vgn)
									{
										vghtml := fmt.Sprint(EventType(event.Event))
										vgn.InnerHTML = &vghtml
									}
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
									vgparent.AppendChild(vgn)
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-8 my-auto"}}}
									vgparent.AppendChild(vgn)
									{
										vgparent := vgn
										_ = vgparent
										{
											vgcompKey := vugu.MakeCompKey(0x5F3C9671547F0065, vgiterkey)
											// ask BuildEnv for prior instance of this specific component
											vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*EventDetails)
											if vgcomp == nil {
												// create new one if needed
												vgcomp = new(EventDetails)
												vgin.BuildEnv.WireComponent(vgcomp)
											}
											vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
											vgcomp.Event = event
											vgcomp.PlaybackNode = c.EventNode(event)
											vgout.Components = append(vgout.Components, vgcomp)
											vgn = &vugu.VGNode{Component: vgcomp}
											vgparent.AppendChild(vgn)
										}
									}
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
									vgparent.AppendChild(vgn)
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-2 my-auto"}}}
									vgparent.AppendChild(vgn)
									{
										vghtml := fmt.Sprint(event.Event.Time)
										vgn.InnerHTML = &vghtml
									}
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
									vgparent.AppendChild(vgn)
								}
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row p-2"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "form", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "container-fluid"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-row"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group col-md-4"}}}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "label", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							{
								vghtml := fmt.Sprint("Target")
								vgn.InnerHTML = &vghtml
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "select", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-control"}}}
							vgparent.AppendChild(vgn)
							vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
								EventType:	"change",
								Func:		func(event *vugu.DOMEvent) { c.SetNodeFilter(event) },
								// TODO: implement capture, etc. mostly need to decide syntax
							})
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
								vgparent.AppendChild(vgn)
								for i := 0; i < len(c.Nodes); i++ {
									var vgiterkey interface{} = i
									_ = vgiterkey
									i := i
									_ = i
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute(nil)}
									vgparent.AppendChild(vgn)
									vgn.Attr = append(vgn.Attr, vugu.VGAttribute{Key: "value", Val: fmt.Sprint(i)})
									{
										vghtml := fmt.Sprint(fmt.Sprintf("Node %d", i))
										vgn.InnerHTML = &vghtml
									}
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "selected", Val: ""}, vugu.VGAttribute{Namespace: "", Key: "value", Val: "all"}}}
								vgparent.AppendChild(vgn)
								{
									vghtml := fmt.Sprint("All")
									vgn.InnerHTML = &vghtml
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
								vgparent.AppendChild(vgn)
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn-group w-100"}, vugu.VGAttribute{Namespace: "", Key: "role", Val: "group"}}}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "type", Val: "button"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn btn-primary"}}}
								vgparent.AppendChild(vgn)
								vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
									EventType:	"click",
									Func:		func(event *vugu.DOMEvent) { c.StepNext(event) },
									// TODO: implement capture, etc. mostly need to decide syntax
								})
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Step Next"}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
								vgparent.AppendChild(vgn)
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group col-md-4"}}}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn-group w-100 align-bottom"}, vugu.VGAttribute{Namespace: "", Key: "role", Val: "group"}}}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "type", Val: "button"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn btn-primary"}}}
								vgparent.AppendChild(vgn)
								vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
									EventType:	"click",
									Func:		func(event *vugu.DOMEvent) { c.StepInstant(event) },
									// TODO: implement capture, etc. mostly need to decide syntax
								})
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Step Instant"}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
								vgparent.AppendChild(vgn)
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group col-md-4"}}}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "label", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							{
								vghtml := fmt.Sprint("Preview Window")
								vgn.InnerHTML = &vghtml
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "select", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-control"}}}
							vgparent.AppendChild(vgn)
							vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
								EventType:	"change",
								Func:		func(event *vugu.DOMEvent) { c.SetStepWindow(event) },
								// TODO: implement capture, etc. mostly need to decide syntax
							})
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "100ms"}}}
								vgparent.AppendChild(vgn)
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "100ms"}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "selected", Val: ""}, vugu.VGAttribute{Namespace: "", Key: "value", Val: "500ms"}}}
								vgparent.AppendChild(vgn)
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "500ms"}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "option", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "value", Val: "2000ms"}}}
								vgparent.AppendChild(vgn)
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "2000ms"}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
								vgparent.AppendChild(vgn)
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn-group w-100"}, vugu.VGAttribute{Namespace: "", Key: "role", Val: "group"}}}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "step-window"}, vugu.VGAttribute{Namespace: "", Key: "type", Val: "button"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "btn btn-primary"}}}
								vgparent.AppendChild(vgn)
								vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
									EventType:	"click",
									Func:		func(event *vugu.DOMEvent) { c.StepStepWindow(event) },
									// TODO: implement capture, etc. mostly need to decide syntax
								})
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: " Step Window"}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
								vgparent.AppendChild(vgn)
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
					vgparent.AppendChild(vgn)
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
