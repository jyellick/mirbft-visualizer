package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *Node) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

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
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-12"}}}
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
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-12 p-2"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
					vgparent.AppendChild(vgn)
					{
						vgcompKey := vugu.MakeCompKey(0x5F3C9671E76F5600, fmt.Sprintf("sequences-%d", c.MirNode.Node.Config.ID))
						// ask BuildEnv for prior instance of this specific component
						vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Sequences)
						if vgcomp == nil {
							// create new one if needed
							vgcomp = new(Sequences)
							vgin.BuildEnv.WireComponent(vgcomp)
						}
						vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
						vgcomp.MirNode = c.MirNode
						vgout.Components = append(vgout.Components, vgcomp)
						vgn = &vugu.VGNode{Component: vgcomp}
						vgparent.AppendChild(vgn)
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row border"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-4 border-right"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
					vgparent.AppendChild(vgn)
					{
						vgcompKey := vugu.MakeCompKey(0x5F3C967164D5559C, fmt.Sprintf("actions-%d", c.MirNode.Node.Config.ID))
						// ask BuildEnv for prior instance of this specific component
						vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Actions)
						if vgcomp == nil {
							// create new one if needed
							vgcomp = new(Actions)
							vgin.BuildEnv.WireComponent(vgcomp)
						}
						vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
						vgcomp.MirNode = c.MirNode
						vgout.Components = append(vgout.Components, vgcomp)
						vgn = &vugu.VGNode{Component: vgcomp}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-4 border-right"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
					vgparent.AppendChild(vgn)
					{
						vgcompKey := vugu.MakeCompKey(0x5F3C96711D75430A, fmt.Sprintf("epochchanges-%d", c.MirNode.Node.Config.ID))
						// ask BuildEnv for prior instance of this specific component
						vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*EpochChanges)
						if vgcomp == nil {
							// create new one if needed
							vgcomp = new(EpochChanges)
							vgin.BuildEnv.WireComponent(vgcomp)
						}
						vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
						vgcomp.MirNode = c.MirNode
						vgout.Components = append(vgout.Components, vgcomp)
						vgn = &vugu.VGNode{Component: vgcomp}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-4 border-left"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
					vgparent.AppendChild(vgn)
					{
						vgcompKey := vugu.MakeCompKey(0x5F3C9671A198B8CE, fmt.Sprintf("checkpoints-%d", c.MirNode.Node.Config.ID))
						// ask BuildEnv for prior instance of this specific component
						vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Checkpoints)
						if vgcomp == nil {
							// create new one if needed
							vgcomp = new(Checkpoints)
							vgin.BuildEnv.WireComponent(vgcomp)
						}
						vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
						vgcomp.MirNode = c.MirNode
						vgout.Components = append(vgout.Components, vgcomp)
						vgn = &vugu.VGNode{Component: vgcomp}
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
