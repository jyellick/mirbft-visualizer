package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *Root) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "html", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "head", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "title", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "MirBFT Visualizer"}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "link", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "rel", Val: "stylesheet"}, vugu.VGAttribute{Namespace: "", Key: "href", Val: "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"}}}
			vgout.AppendCSS(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "script", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "src", Val: "https://code.jquery.com/jquery-3.2.1.slim.min.js"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.AppendJS(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "script", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "src", Val: "https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.AppendJS(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "script", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "src", Val: "https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.AppendJS(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "meta", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "name", Val: "viewport"}, vugu.VGAttribute{Namespace: "", Key: "content", Val: "width=device-width, initial-scale=1, shrink-to-fit=no"}}}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "body", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "container-fluid"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row m-2"}}}
				vgparent.AppendChild(vgn)
				{
					vghtml := fmt.Sprint("\n              \x3Cdiv class=\"column\"\x3E\n                  \x3Ch1\x3EMirBFT Visualization\x3C/h1\x3E\n              \x3C/div\x3E\n          ")
					vgn.InnerHTML = &vghtml
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
				vgparent.AppendChild(vgn)
				if !c.Bootstrapped {
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row m-2"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
						vgparent.AppendChild(vgn)
						{
							vgcompKey := vugu.MakeCompKey(0x5E5F0F2315ED0B51, vgiterkey)
							// ask BuildEnv for prior instance of this specific component
							vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Bootstrap)
							if vgcomp == nil {
								// create new one if needed
								vgcomp = new(Bootstrap)
							}
							vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
							vgcomp.Bootstrap = c.Bootstrap
							vgcomp.NodeCountDefault = 4
							vgout.Components = append(vgout.Components, vgcomp)
							vgn = &vugu.VGNode{Component: vgcomp}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
						vgparent.AppendChild(vgn)
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
				vgparent.AppendChild(vgn)
				if c.Bootstrapped {
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row m-2"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n              "}
						vgparent.AppendChild(vgn)
						{
							vgcompKey := vugu.MakeCompKey(0x5E5F0F239CE3F682, vgiterkey)
							// ask BuildEnv for prior instance of this specific component
							vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*Network)
							if vgcomp == nil {
								// create new one if needed
								vgcomp = new(Network)
							}
							vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
							vgcomp.Parameters = c.BootstrapParameters
							vgout.Components = append(vgout.Components, vgcomp)
							vgn = &vugu.VGNode{Component: vgcomp}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n          "}
						vgparent.AppendChild(vgn)
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      "}
				vgparent.AppendChild(vgn)
			}
		}
	}
	return vgout
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ vjson.RawMessage
var _ js.Value
