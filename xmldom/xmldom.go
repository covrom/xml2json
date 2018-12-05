package xmldom

import (
	"encoding/json"
	"encoding/xml"
)

type XMLNode struct {
	XMLName xml.Name
	Content string     `xml:",innerxml"`
	Nodes   []*XMLNode `xml:",any"`
	Attrs   []xml.Attr `xml:"-"`
}

func (n *XMLNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node XMLNode

	return d.DecodeElement((*node)(n), &start)
}

func (n *XMLNode) toMap() map[string]interface{} {
	attrs := make(map[string]interface{})

	attradd := func(key string, val interface{}) {
		if v, ok := attrs[key]; ok {
			if vv, ok := v.([]interface{}); ok {
				vv = append(vv, val)
				attrs[key] = vv
			} else {
				arr := make([]interface{}, 2)
				arr[0] = v
				arr[1] = val
				attrs[key] = arr
			}
		} else {
			attrs[key] = val
		}
	}

	if len(n.Nodes) > 0 {
		for _, node := range n.Nodes {
			attradd(node.XMLName.Local, node.toMap()[node.XMLName.Local])
		}
	}
	if len(n.Attrs) > 0 {
		for _, v := range n.Attrs {
			attradd(v.Name.Local, v.Value)
		}
	}

	res := make(map[string]interface{})
	if len(attrs) == 0 {
		res[n.XMLName.Local] = n.Content
	} else {
		if len(n.Attrs) > 0 && len(n.Nodes) == 0 {
			if len(n.Content) > 0 {
				attradd("_content", n.Content)
			}
			res[n.XMLName.Local] = attrs
		} else {
			res[n.XMLName.Local] = attrs
		}
	}
	return res
}

func (n *XMLNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.toMap())
}
