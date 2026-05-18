//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	cefTypes "github.com/energye/cef/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICefXmlReader Parent: ICefBaseRefCounted
type ICefXmlReader interface {
	ICefBaseRefCounted
	// MoveToNextNode
	//  Moves the cursor to the next node in the document. This function must be
	//  called at least once to set the current cursor position. Returns true (1)
	//  if the cursor position was set successfully.
	MoveToNextNode() bool // function
	// Close
	//  Close the document. This should be called directly to ensure that cleanup
	//  occurs on the correct thread.
	Close() bool // function
	// HasError
	//  Returns true (1) if an error has been reported by the XML parser.
	HasError() bool // function
	// GetError
	//  Returns the error string.
	GetError() string // function
	// GetType
	//  Returns the node type.
	GetType() cefTypes.TCefXmlNodeType // function
	// GetDepth
	//  Returns the node depth. Depth starts at 0 for the root node.
	GetDepth() int32 // function
	// GetLocalName
	//  Returns the local name. See http://www.w3.org/TR/REC-xml-names/#NT-
	//  LocalPart for additional details.
	GetLocalName() string // function
	// GetPrefix
	//  Returns the namespace prefix. See http://www.w3.org/TR/REC-xml-names/ for
	//  additional details.
	GetPrefix() string // function
	// GetQualifiedName
	//  Returns the qualified name, equal to (Prefix:)LocalName. See
	//  http://www.w3.org/TR/REC-xml-names/#ns-qualnames for additional details.
	GetQualifiedName() string // function
	// GetNamespaceUri
	//  Returns the URI defining the namespace associated with the node. See
	//  http://www.w3.org/TR/REC-xml-names/ for additional details.
	GetNamespaceUri() string // function
	// GetBaseUri
	//  Returns the base URI of the node. See http://www.w3.org/TR/xmlbase/ for
	//  additional details.
	GetBaseUri() string // function
	// GetXmlLang
	//  Returns the xml:lang scope within which the node resides. See
	//  http://www.w3.org/TR/REC-xml/#sec-lang-tag for additional details.
	GetXmlLang() string // function
	// IsEmptyElement
	//  Returns true (1) if the node represents an NULL element. "<a/>" is
	//  considered NULL but "<a></a>" is not.
	IsEmptyElement() bool // function
	// HasValue
	//  Returns true (1) if the node has a text value.
	HasValue() bool // function
	// GetValue
	//  Returns the text value.
	GetValue() string // function
	// HasAttributes
	//  Returns true (1) if the node has attributes.
	HasAttributes() bool // function
	// GetAttributeCount
	//  Returns the number of attributes.
	GetAttributeCount() cefTypes.NativeUInt // function
	// GetAttributeByIndex
	//  Returns the value of the attribute at the specified 0-based index.
	GetAttributeByIndex(index int32) string // function
	// GetAttributeByQName
	//  Returns the value of the attribute with the specified qualified name.
	GetAttributeByQName(qualifiedName string) string // function
	// GetAttributeByLName
	//  Returns the value of the attribute with the specified local name and
	//  namespace URI.
	GetAttributeByLName(localName string, namespaceURI string) string // function
	// GetInnerXml
	//  Returns an XML representation of the current node's children.
	GetInnerXml() string // function
	// GetOuterXml
	//  Returns an XML representation of the current node including its children.
	GetOuterXml() string // function
	// GetLineNumber
	//  Returns the line number for the current node.
	GetLineNumber() int32 // function
	// MoveToAttributeByIndex
	//  Moves the cursor to the attribute at the specified 0-based index. Returns
	//  true (1) if the cursor position was set successfully.
	MoveToAttributeByIndex(index int32) bool // function
	// MoveToAttributeByQName
	//  Moves the cursor to the attribute with the specified qualified name.
	//  Returns true (1) if the cursor position was set successfully.
	MoveToAttributeByQName(qualifiedName string) bool // function
	// MoveToAttributeByLName
	//  Moves the cursor to the attribute with the specified local name and
	//  namespace URI. Returns true (1) if the cursor position was set
	//  successfully.
	MoveToAttributeByLName(localName string, namespaceURI string) bool // function
	// MoveToFirstAttribute
	//  Moves the cursor to the first attribute in the current element. Returns
	//  true (1) if the cursor position was set successfully.
	MoveToFirstAttribute() bool // function
	// MoveToNextAttribute
	//  Moves the cursor to the next attribute in the current element. Returns
	//  true (1) if the cursor position was set successfully.
	MoveToNextAttribute() bool // function
	// MoveToCarryingElement
	//  Moves the cursor back to the carrying element. Returns true (1) if the
	//  cursor position was set successfully.
	MoveToCarryingElement() bool // function
}

// ICefXmlReaderRef Parent: ICefXmlReader ICefBaseRefCountedRef
type ICefXmlReaderRef interface {
	ICefXmlReader
	ICefBaseRefCountedRef
	AsIntfXmlReader() uintptr
}

type TCefXmlReaderRef struct {
	TCefBaseRefCountedRef
}

func (m *TCefXmlReaderRef) MoveToNextNode() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) Close() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) HasError() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) GetError() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetType() cefTypes.TCefXmlNodeType {
	if !m.IsValid() {
		return 0
	}
	r := cefXmlReaderRefAPI().SysCallN(5, m.Instance())
	return cefTypes.TCefXmlNodeType(r)
}

func (m *TCefXmlReaderRef) GetDepth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefXmlReaderRefAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCefXmlReaderRef) GetLocalName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetPrefix() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetQualifiedName() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetNamespaceUri() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetBaseUri() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetXmlLang() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) IsEmptyElement() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) HasValue() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(14, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) GetValue() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) HasAttributes() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(16, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) GetAttributeCount() cefTypes.NativeUInt {
	if !m.IsValid() {
		return 0
	}
	r := cefXmlReaderRefAPI().SysCallN(17, m.Instance())
	return cefTypes.NativeUInt(r)
}

func (m *TCefXmlReaderRef) GetAttributeByIndex(index int32) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(18, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetAttributeByQName(qualifiedName string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(19, m.Instance(), api.PasStr(qualifiedName), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetAttributeByLName(localName string, namespaceURI string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(20, m.Instance(), api.PasStr(localName), api.PasStr(namespaceURI), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetInnerXml() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetOuterXml() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	cefXmlReaderRefAPI().SysCallN(22, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCefXmlReaderRef) GetLineNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r := cefXmlReaderRefAPI().SysCallN(23, m.Instance())
	return int32(r)
}

func (m *TCefXmlReaderRef) MoveToAttributeByIndex(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(24, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) MoveToAttributeByQName(qualifiedName string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(25, m.Instance(), api.PasStr(qualifiedName))
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) MoveToAttributeByLName(localName string, namespaceURI string) bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(26, m.Instance(), api.PasStr(localName), api.PasStr(namespaceURI))
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) MoveToFirstAttribute() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(27, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) MoveToNextAttribute() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(28, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) MoveToCarryingElement() bool {
	if !m.IsValid() {
		return false
	}
	r := cefXmlReaderRefAPI().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TCefXmlReaderRef) AsIntfXmlReader() uintptr {
	return m.GetIntfPointer(0)
}

// XmlReaderRef  is static instance
var XmlReaderRef _XmlReaderRefClass

// _XmlReaderRefClass is class type defined by TCefXmlReaderRef
type _XmlReaderRefClass uintptr

func (_XmlReaderRefClass) UnWrap(data uintptr) (result ICefXmlReader) {
	var resultPtr uintptr
	cefXmlReaderRefAPI().SysCallN(30, uintptr(data), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefXmlReaderRef(resultPtr)
	return
}

func (_XmlReaderRefClass) New(stream ICefStreamReader, encodingType cefTypes.TCefXmlEncodingType, uRI string) (result ICefXmlReader) {
	var resultPtr uintptr
	cefXmlReaderRefAPI().SysCallN(31, base.GetObjectUintptr(stream), uintptr(encodingType), api.PasStr(uRI), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCefXmlReaderRef(resultPtr)
	return
}

// NewXmlReaderRef class constructor
func NewXmlReaderRef(data uintptr) ICefXmlReaderRef {
	var xmlReaderPtr uintptr // ICefXmlReader
	r := cefXmlReaderRefAPI().SysCallN(0, uintptr(data), uintptr(base.UnsafePointer(&xmlReaderPtr)))
	ret := AsCefXmlReaderRef(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, xmlReaderPtr)
	}
	return ret
}

var (
	cefXmlReaderRefOnce   base.Once
	cefXmlReaderRefImport *imports.Imports = nil
)

func cefXmlReaderRefAPI() *imports.Imports {
	cefXmlReaderRefOnce.Do(func() {
		cefXmlReaderRefImport = api.NewDefaultImports()
		cefXmlReaderRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCefXmlReaderRef_Create", 0), // constructor NewXmlReaderRef
			/* 1 */ imports.NewTable("TCefXmlReaderRef_MoveToNextNode", 0), // function MoveToNextNode
			/* 2 */ imports.NewTable("TCefXmlReaderRef_Close", 0), // function Close
			/* 3 */ imports.NewTable("TCefXmlReaderRef_HasError", 0), // function HasError
			/* 4 */ imports.NewTable("TCefXmlReaderRef_GetError", 0), // function GetError
			/* 5 */ imports.NewTable("TCefXmlReaderRef_GetType", 0), // function GetType
			/* 6 */ imports.NewTable("TCefXmlReaderRef_GetDepth", 0), // function GetDepth
			/* 7 */ imports.NewTable("TCefXmlReaderRef_GetLocalName", 0), // function GetLocalName
			/* 8 */ imports.NewTable("TCefXmlReaderRef_GetPrefix", 0), // function GetPrefix
			/* 9 */ imports.NewTable("TCefXmlReaderRef_GetQualifiedName", 0), // function GetQualifiedName
			/* 10 */ imports.NewTable("TCefXmlReaderRef_GetNamespaceUri", 0), // function GetNamespaceUri
			/* 11 */ imports.NewTable("TCefXmlReaderRef_GetBaseUri", 0), // function GetBaseUri
			/* 12 */ imports.NewTable("TCefXmlReaderRef_GetXmlLang", 0), // function GetXmlLang
			/* 13 */ imports.NewTable("TCefXmlReaderRef_IsEmptyElement", 0), // function IsEmptyElement
			/* 14 */ imports.NewTable("TCefXmlReaderRef_HasValue", 0), // function HasValue
			/* 15 */ imports.NewTable("TCefXmlReaderRef_GetValue", 0), // function GetValue
			/* 16 */ imports.NewTable("TCefXmlReaderRef_HasAttributes", 0), // function HasAttributes
			/* 17 */ imports.NewTable("TCefXmlReaderRef_GetAttributeCount", 0), // function GetAttributeCount
			/* 18 */ imports.NewTable("TCefXmlReaderRef_GetAttributeByIndex", 0), // function GetAttributeByIndex
			/* 19 */ imports.NewTable("TCefXmlReaderRef_GetAttributeByQName", 0), // function GetAttributeByQName
			/* 20 */ imports.NewTable("TCefXmlReaderRef_GetAttributeByLName", 0), // function GetAttributeByLName
			/* 21 */ imports.NewTable("TCefXmlReaderRef_GetInnerXml", 0), // function GetInnerXml
			/* 22 */ imports.NewTable("TCefXmlReaderRef_GetOuterXml", 0), // function GetOuterXml
			/* 23 */ imports.NewTable("TCefXmlReaderRef_GetLineNumber", 0), // function GetLineNumber
			/* 24 */ imports.NewTable("TCefXmlReaderRef_MoveToAttributeByIndex", 0), // function MoveToAttributeByIndex
			/* 25 */ imports.NewTable("TCefXmlReaderRef_MoveToAttributeByQName", 0), // function MoveToAttributeByQName
			/* 26 */ imports.NewTable("TCefXmlReaderRef_MoveToAttributeByLName", 0), // function MoveToAttributeByLName
			/* 27 */ imports.NewTable("TCefXmlReaderRef_MoveToFirstAttribute", 0), // function MoveToFirstAttribute
			/* 28 */ imports.NewTable("TCefXmlReaderRef_MoveToNextAttribute", 0), // function MoveToNextAttribute
			/* 29 */ imports.NewTable("TCefXmlReaderRef_MoveToCarryingElement", 0), // function MoveToCarryingElement
			/* 30 */ imports.NewTable("TCefXmlReaderRef_UnWrap", 0), // static function UnWrap
			/* 31 */ imports.NewTable("TCefXmlReaderRef_New", 0), // static function New
		}
	})
	return cefXmlReaderRefImport
}
