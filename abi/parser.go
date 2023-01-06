package abi

import (
	"bytes"
	"encoding/xml"
)

var METHOD = `




<internal name="transfer">
        <input>
            transfer#5fcc3d14 query_id:uint64 new_owner:MsgAddress response_destination:MsgAddress custom_payload:(Maybe ^Cell) forward_amount:(VarUInteger 16) forward_payload:(Either Cell ^Cell) = InternalMsgBody;
        </input>
        <ouput>
            ownership_assigned#05138d91 query_id:uint64 prev_owner:MsgAddress forward_payload:(Either Cell ^Cell) = InternalMsgBody;
        </ouput>
        <ouput>
            excesses#d53276db query_id:uint64 = InternalMsgBody;
        </ouput>
    </internal>
`

type Interface struct {
	Name      string      `xml:"name,attr"`
	Group     string      `xml:"group,attr"`
	Methods   []GetMethod `xml:"get_method"`
	Internals []Internal  `xml:"internal"`
	Types     string      `xml:"types"`
}

type Internal struct {
	Name    string   `xml:"name,attr"`
	Input   string   `xml:"input"`
	Outputs []string `xml:"output"`
}

type GetMethod struct {
	Tag   xml.Name
	Input struct {
		StackValues []StackRecord `xml:",any"`
	} `xml:"input"`
	Name        string            `xml:"name,attr"`
	Callback    bool              `xml:"callback,attr"`
	FixedLength bool              `xml:"fixed_length,attr"`
	Output      []GetmethodOutput `xml:"output"`
}

type GetmethodOutput struct {
	Version     string        `xml:"version"`
	FixedLength bool          `xml:"fixed_length"`
	Stack       []StackRecord `xml:",any"`
}

type StackRecord struct {
	XMLName xml.Name
	Name    string `xml:"name,attr"`
	Type    string `xml:",chardata"`
}

func ParseInterface(s []byte) ([]Interface, error) {
	var i struct {
		List []Interface `xml:"interface"`
	}

	if !bytes.HasPrefix(bytes.TrimSpace(s), []byte("<interfaces>")) {
		s = append(append([]byte("<interfaces>"), s...), []byte("</interfaces>")...)
	}
	err := xml.Unmarshal(s, &i)
	return i.List, err
}

func ParseMethod(s []byte) (GetMethod, error) {
	var m GetMethod
	err := xml.Unmarshal(s, &m)
	return m, err
}