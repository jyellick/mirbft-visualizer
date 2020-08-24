package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *Network) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col-12"}}}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " <main:Client :MirNodes='c.MirNodes' :EventQueue='c.EventQueue'></main:Client> "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
		vgparent.AppendChild(vgn)
		{
			vgcompKey := vugu.MakeCompKey(0x5F441583D7395FE6, vgiterkey)
			// ask BuildEnv for prior instance of this specific component
			vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Events)
			if vgcomp == nil {
				// create new one if needed
				vgcomp = new(Events)
				vgin.BuildEnv.WireComponent(vgcomp)
			}
			vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
			vgcomp.EventLog = c.EventLog
			vgcomp.Nodes = c.Nodes
			vgcomp.Stepper = c.Stepper
			vgout.Components = append(vgout.Components, vgcomp)
			vgn = &vugu.VGNode{Component: vgcomp}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
		vgparent.AppendChild(vgn)
		for vgiterkeyt, mirNode := range OrderedNodes(c.Nodes) {
			var vgiterkey interface{} = vgiterkeyt
			_ = vgiterkey
			mirNode := mirNode
			_ = mirNode
			{
				vgcompKey := vugu.MakeCompKey(0x5F441583FE5F6034, vgiterkey)
				// ask BuildEnv for prior instance of this specific component
				vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Node)
				if vgcomp == nil {
					// create new one if needed
					vgcomp = new(Node)
					vgin.BuildEnv.WireComponent(vgcomp)
				}
				vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
				vgcomp.MirNode = mirNode
				vgout.Components = append(vgout.Components, vgcomp)
				vgn = &vugu.VGNode{Component: vgcomp}
				vgparent.AppendChild(vgn)
			}
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
