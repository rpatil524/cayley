package proto

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/quad/pquads"
)

//go:generate protoc --proto_path=$GOPATH/src:. --gogo_out=. serializations.proto

// GetNativeValue returns the value stored in Node.
func (m *NodeData) GetNativeValue() quad.Value {
	if m == nil {
		return quad.Raw("")
	} else if m.Value == nil {
		if m.Name == "" {
			return quad.Raw("")
		}
		return quad.Raw(m.Name)
	}
	return m.Value.ToNative()
}

func (m *NodeData) Upgrade() {
	if m.Value == nil {
		m.Value = pquads.MakeValue(quad.Raw(m.Name))
		m.Name = ""
	}
}
