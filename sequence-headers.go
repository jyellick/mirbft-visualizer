package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

var _ vugu.ComponentType = (*SequenceHeaders)(nil)

func (comp *SequenceHeaders) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*SequenceHeadersData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "m-0 p-0"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "table", DataAtom: vugu.VGAtom(365829), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "table table-bordered my-auto"}, vugu.VGAttribute{Namespace: "", Key: "style", Val: "border-style: hidden"}}}
		parent.AppendChild(n)
		{
			parent := n
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "thead", DataAtom: vugu.VGAtom(208901), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", DataAtom: vugu.VGAtom(52226), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				{
					parent := n
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
					n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "th", DataAtom: vugu.VGAtom(87554), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "rowspan", Val: "2"}}}
					parent.AppendChild(n)
					{
						parent := n
						n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Buckets", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
						parent.AppendChild(n)
					}
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
					n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "th", DataAtom: vugu.VGAtom(87554), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
					n.Props = vugu.Props{
						"colspan": data.Status.HighWatermark - data.Status.LowWatermark,
					}
					{
						parent := n
						n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: " Sequences ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
						parent.AppendChild(n)
					}
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
				}
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", DataAtom: vugu.VGAtom(52226), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				{
					parent := n
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
					for i := data.Status.LowWatermark; i < data.Status.HighWatermark; i++ {
						n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "th", DataAtom: vugu.VGAtom(87554), Namespace: "", Attr: []vugu.VGAttribute(nil)}
						parent.AppendChild(n)
						n.InnerHTML = fmt.Sprint(i)
					}
					n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
				}
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tbody", DataAtom: vugu.VGAtom(9989), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
				for i, bucketStatus := range data.Status.Buckets {
					n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tr", DataAtom: vugu.VGAtom(52226), Namespace: "", Attr: []vugu.VGAttribute(nil)}
					parent.AppendChild(n)
					{
						parent := n
						n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
						parent.AppendChild(n)
						n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", DataAtom: vugu.VGAtom(37378), Namespace: "", Attr: []vugu.VGAttribute(nil)}
						parent.AppendChild(n)
						n.InnerHTML = fmt.Sprint(fmt.Sprintf("Bucket %d", i))
						n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
						parent.AppendChild(n)
						for _, seqState := range bucketStatus.Sequences {
							n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "td", DataAtom: vugu.VGAtom(37378), Namespace: "", Attr: []vugu.VGAttribute(nil)}
							parent.AppendChild(n)
							n.InnerHTML = fmt.Sprint(SeqStateToChar(seqState))
							{
								parent := n
								n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: " ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
								parent.AppendChild(n)
							}
						}
						n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n      ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
						parent.AppendChild(n)
					}
				}
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
			n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "tbody", DataAtom: vugu.VGAtom(9989), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
			{
				parent := n
				n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n  ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
				parent.AppendChild(n)
			}
		}
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
	}
	return
}

type SequenceHeaders struct {}

func init() { vugu.RegisterComponentType("sequence-headers", &SequenceHeaders{}) }