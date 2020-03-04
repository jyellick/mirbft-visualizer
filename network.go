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
		{
			vgcompKey := vugu.MakeCompKey(0x5E5FF1EC359D6794, vgiterkey)
			// ask BuildEnv for prior instance of this specific component
			vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Events)
			if vgcomp == nil {
				// create new one if needed
				vgcomp = new(Events)
			}
			vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
			vgcomp.EventQueue = c.EventQueue
			vgout.Components = append(vgout.Components, vgcomp)
			vgn = &vugu.VGNode{Component: vgcomp}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  "}
		vgparent.AppendChild(vgn)
		for vgiterkeyt, mirNode := range c.MirNodes {
			var vgiterkey interface{} = vgiterkeyt
			_ = vgiterkey
			mirNode := mirNode
			_ = mirNode
			{
				vgcompKey := vugu.MakeCompKey(0x5E5FF1EC17BA1481, vgiterkey)
				// ask BuildEnv for prior instance of this specific component
				vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Node)
				if vgcomp == nil {
					// create new one if needed
					vgcomp = new(Node)
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
