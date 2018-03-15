// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1/text_annotation.proto

package vision

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Enum to denote the type of break found. New line, space etc.
type TextAnnotation_DetectedBreak_BreakType int32

const (
	// Unknown break label type.
	TextAnnotation_DetectedBreak_UNKNOWN TextAnnotation_DetectedBreak_BreakType = 0
	// Regular space.
	TextAnnotation_DetectedBreak_SPACE TextAnnotation_DetectedBreak_BreakType = 1
	// Sure space (very wide).
	TextAnnotation_DetectedBreak_SURE_SPACE TextAnnotation_DetectedBreak_BreakType = 2
	// Line-wrapping break.
	TextAnnotation_DetectedBreak_EOL_SURE_SPACE TextAnnotation_DetectedBreak_BreakType = 3
	// End-line hyphen that is not present in text; does not co-occur with
	// `SPACE`, `LEADER_SPACE`, or `LINE_BREAK`.
	TextAnnotation_DetectedBreak_HYPHEN TextAnnotation_DetectedBreak_BreakType = 4
	// Line break that ends a paragraph.
	TextAnnotation_DetectedBreak_LINE_BREAK TextAnnotation_DetectedBreak_BreakType = 5
)

var TextAnnotation_DetectedBreak_BreakType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPACE",
	2: "SURE_SPACE",
	3: "EOL_SURE_SPACE",
	4: "HYPHEN",
	5: "LINE_BREAK",
}
var TextAnnotation_DetectedBreak_BreakType_value = map[string]int32{
	"UNKNOWN":        0,
	"SPACE":          1,
	"SURE_SPACE":     2,
	"EOL_SURE_SPACE": 3,
	"HYPHEN":         4,
	"LINE_BREAK":     5,
}

func (x TextAnnotation_DetectedBreak_BreakType) String() string {
	return proto.EnumName(TextAnnotation_DetectedBreak_BreakType_name, int32(x))
}
func (TextAnnotation_DetectedBreak_BreakType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 1, 0}
}

// Type of a block (text, image etc) as identified by OCR.
type Block_BlockType int32

const (
	// Unknown block type.
	Block_UNKNOWN Block_BlockType = 0
	// Regular text block.
	Block_TEXT Block_BlockType = 1
	// Table block.
	Block_TABLE Block_BlockType = 2
	// Image block.
	Block_PICTURE Block_BlockType = 3
	// Horizontal/vertical line box.
	Block_RULER Block_BlockType = 4
	// Barcode block.
	Block_BARCODE Block_BlockType = 5
)

var Block_BlockType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TEXT",
	2: "TABLE",
	3: "PICTURE",
	4: "RULER",
	5: "BARCODE",
}
var Block_BlockType_value = map[string]int32{
	"UNKNOWN": 0,
	"TEXT":    1,
	"TABLE":   2,
	"PICTURE": 3,
	"RULER":   4,
	"BARCODE": 5,
}

func (x Block_BlockType) String() string {
	return proto.EnumName(Block_BlockType_name, int32(x))
}
func (Block_BlockType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

// TextAnnotation contains a structured representation of OCR extracted text.
// The hierarchy of an OCR extracted text structure is like this:
//     TextAnnotation -> Page -> Block -> Paragraph -> Word -> Symbol
// Each structural component, starting from Page, may further have their own
// properties. Properties describe detected languages, breaks etc.. Please refer
// to the [TextAnnotation.TextProperty][google.cloud.vision.v1.TextAnnotation.TextProperty] message definition below for more
// detail.
type TextAnnotation struct {
	// List of pages detected by OCR.
	Pages []*Page `protobuf:"bytes,1,rep,name=pages" json:"pages,omitempty"`
	// UTF-8 text detected on the pages.
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *TextAnnotation) Reset()                    { *m = TextAnnotation{} }
func (m *TextAnnotation) String() string            { return proto.CompactTextString(m) }
func (*TextAnnotation) ProtoMessage()               {}
func (*TextAnnotation) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *TextAnnotation) GetPages() []*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func (m *TextAnnotation) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// Detected language for a structural component.
type TextAnnotation_DetectedLanguage struct {
	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode" json:"language_code,omitempty"`
	// Confidence of detected language. Range [0, 1].
	Confidence float32 `protobuf:"fixed32,2,opt,name=confidence" json:"confidence,omitempty"`
}

func (m *TextAnnotation_DetectedLanguage) Reset()         { *m = TextAnnotation_DetectedLanguage{} }
func (m *TextAnnotation_DetectedLanguage) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation_DetectedLanguage) ProtoMessage()    {}
func (*TextAnnotation_DetectedLanguage) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 0}
}

func (m *TextAnnotation_DetectedLanguage) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *TextAnnotation_DetectedLanguage) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Detected start or end of a structural component.
type TextAnnotation_DetectedBreak struct {
	// Detected break type.
	Type TextAnnotation_DetectedBreak_BreakType `protobuf:"varint,1,opt,name=type,enum=google.cloud.vision.v1.TextAnnotation_DetectedBreak_BreakType" json:"type,omitempty"`
	// True if break prepends the element.
	IsPrefix bool `protobuf:"varint,2,opt,name=is_prefix,json=isPrefix" json:"is_prefix,omitempty"`
}

func (m *TextAnnotation_DetectedBreak) Reset()                    { *m = TextAnnotation_DetectedBreak{} }
func (m *TextAnnotation_DetectedBreak) String() string            { return proto.CompactTextString(m) }
func (*TextAnnotation_DetectedBreak) ProtoMessage()               {}
func (*TextAnnotation_DetectedBreak) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 1} }

func (m *TextAnnotation_DetectedBreak) GetType() TextAnnotation_DetectedBreak_BreakType {
	if m != nil {
		return m.Type
	}
	return TextAnnotation_DetectedBreak_UNKNOWN
}

func (m *TextAnnotation_DetectedBreak) GetIsPrefix() bool {
	if m != nil {
		return m.IsPrefix
	}
	return false
}

// Additional information detected on the structural component.
type TextAnnotation_TextProperty struct {
	// A list of detected languages together with confidence.
	DetectedLanguages []*TextAnnotation_DetectedLanguage `protobuf:"bytes,1,rep,name=detected_languages,json=detectedLanguages" json:"detected_languages,omitempty"`
	// Detected start or end of a text segment.
	DetectedBreak *TextAnnotation_DetectedBreak `protobuf:"bytes,2,opt,name=detected_break,json=detectedBreak" json:"detected_break,omitempty"`
}

func (m *TextAnnotation_TextProperty) Reset()                    { *m = TextAnnotation_TextProperty{} }
func (m *TextAnnotation_TextProperty) String() string            { return proto.CompactTextString(m) }
func (*TextAnnotation_TextProperty) ProtoMessage()               {}
func (*TextAnnotation_TextProperty) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 2} }

func (m *TextAnnotation_TextProperty) GetDetectedLanguages() []*TextAnnotation_DetectedLanguage {
	if m != nil {
		return m.DetectedLanguages
	}
	return nil
}

func (m *TextAnnotation_TextProperty) GetDetectedBreak() *TextAnnotation_DetectedBreak {
	if m != nil {
		return m.DetectedBreak
	}
	return nil
}

// Detected page from OCR.
type Page struct {
	// Additional information detected on the page.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// Page width in pixels.
	Width int32 `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	// Page height in pixels.
	Height int32 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	// List of blocks of text, images etc on this page.
	Blocks []*Block `protobuf:"bytes,4,rep,name=blocks" json:"blocks,omitempty"`
	// Confidence of the OCR results on the page. Range [0, 1].
	Confidence float32 `protobuf:"fixed32,5,opt,name=confidence" json:"confidence,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Page) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Page) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Page) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Page) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *Page) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Logical element on the page.
type Block struct {
	// Additional information detected for the block.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the block.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//
	// * when the text is horizontal it might look like:
	//
	//         0----1
	//         |    |
	//         3----2
	//
	// * when it's rotated 180 degrees around the top-left corner it becomes:
	//
	//         2----3
	//         |    |
	//         1----0
	//
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of paragraphs in this block (if this blocks is of type text).
	Paragraphs []*Paragraph `protobuf:"bytes,3,rep,name=paragraphs" json:"paragraphs,omitempty"`
	// Detected block type (text, image etc) for this block.
	BlockType Block_BlockType `protobuf:"varint,4,opt,name=block_type,json=blockType,enum=google.cloud.vision.v1.Block_BlockType" json:"block_type,omitempty"`
	// Confidence of the OCR results on the block. Range [0, 1].
	Confidence float32 `protobuf:"fixed32,5,opt,name=confidence" json:"confidence,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Block) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Block) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Block) GetParagraphs() []*Paragraph {
	if m != nil {
		return m.Paragraphs
	}
	return nil
}

func (m *Block) GetBlockType() Block_BlockType {
	if m != nil {
		return m.BlockType
	}
	return Block_UNKNOWN
}

func (m *Block) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Structural unit of text representing a number of words in certain order.
type Paragraph struct {
	// Additional information detected for the paragraph.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the paragraph.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of words in this paragraph.
	Words []*Word `protobuf:"bytes,3,rep,name=words" json:"words,omitempty"`
	// Confidence of the OCR results for the paragraph. Range [0, 1].
	Confidence float32 `protobuf:"fixed32,4,opt,name=confidence" json:"confidence,omitempty"`
}

func (m *Paragraph) Reset()                    { *m = Paragraph{} }
func (m *Paragraph) String() string            { return proto.CompactTextString(m) }
func (*Paragraph) ProtoMessage()               {}
func (*Paragraph) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *Paragraph) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Paragraph) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Paragraph) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

func (m *Paragraph) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// A word representation.
type Word struct {
	// Additional information detected for the word.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the word.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// List of symbols in the word.
	// The order of the symbols follows the natural reading order.
	Symbols []*Symbol `protobuf:"bytes,3,rep,name=symbols" json:"symbols,omitempty"`
	// Confidence of the OCR results for the word. Range [0, 1].
	Confidence float32 `protobuf:"fixed32,4,opt,name=confidence" json:"confidence,omitempty"`
}

func (m *Word) Reset()                    { *m = Word{} }
func (m *Word) String() string            { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()               {}
func (*Word) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Word) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Word) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Word) GetSymbols() []*Symbol {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *Word) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// A single symbol representation.
type Symbol struct {
	// Additional information detected for the symbol.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The bounding box for the symbol.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
	// The actual UTF-8 representation of the symbol.
	Text string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	// Confidence of the OCR results for the symbol. Range [0, 1].
	Confidence float32 `protobuf:"fixed32,4,opt,name=confidence" json:"confidence,omitempty"`
}

func (m *Symbol) Reset()                    { *m = Symbol{} }
func (m *Symbol) String() string            { return proto.CompactTextString(m) }
func (*Symbol) ProtoMessage()               {}
func (*Symbol) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *Symbol) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Symbol) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Symbol) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Symbol) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func init() {
	proto.RegisterType((*TextAnnotation)(nil), "google.cloud.vision.v1.TextAnnotation")
	proto.RegisterType((*TextAnnotation_DetectedLanguage)(nil), "google.cloud.vision.v1.TextAnnotation.DetectedLanguage")
	proto.RegisterType((*TextAnnotation_DetectedBreak)(nil), "google.cloud.vision.v1.TextAnnotation.DetectedBreak")
	proto.RegisterType((*TextAnnotation_TextProperty)(nil), "google.cloud.vision.v1.TextAnnotation.TextProperty")
	proto.RegisterType((*Page)(nil), "google.cloud.vision.v1.Page")
	proto.RegisterType((*Block)(nil), "google.cloud.vision.v1.Block")
	proto.RegisterType((*Paragraph)(nil), "google.cloud.vision.v1.Paragraph")
	proto.RegisterType((*Word)(nil), "google.cloud.vision.v1.Word")
	proto.RegisterType((*Symbol)(nil), "google.cloud.vision.v1.Symbol")
	proto.RegisterEnum("google.cloud.vision.v1.TextAnnotation_DetectedBreak_BreakType", TextAnnotation_DetectedBreak_BreakType_name, TextAnnotation_DetectedBreak_BreakType_value)
	proto.RegisterEnum("google.cloud.vision.v1.Block_BlockType", Block_BlockType_name, Block_BlockType_value)
}

func init() { proto.RegisterFile("google/cloud/vision/v1/text_annotation.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcf, 0x6f, 0xd3, 0x4a,
	0x10, 0x7e, 0x6e, 0xec, 0x34, 0x9e, 0xb4, 0x91, 0xdf, 0xbe, 0xa7, 0x2a, 0x0a, 0xa5, 0x2a, 0x01,
	0x44, 0x0f, 0xc8, 0x51, 0x53, 0x10, 0x48, 0x20, 0xa4, 0x38, 0x35, 0xb4, 0x6a, 0x94, 0x58, 0xdb,
	0x44, 0xe5, 0xc7, 0xc1, 0xf2, 0x8f, 0xad, 0x63, 0x35, 0xf5, 0x5a, 0xb6, 0xdb, 0x26, 0xff, 0x0d,
	0xff, 0x13, 0x12, 0x27, 0xae, 0x9c, 0xb9, 0x02, 0x27, 0xe4, 0xb5, 0x9d, 0x26, 0x01, 0x53, 0x40,
	0x1c, 0x7a, 0xb1, 0x76, 0x26, 0xdf, 0x7c, 0x33, 0xdf, 0x8c, 0x27, 0x6b, 0xb8, 0xef, 0x50, 0xea,
	0x8c, 0x48, 0xc3, 0x1a, 0xd1, 0x33, 0xbb, 0x71, 0xee, 0x86, 0x2e, 0xf5, 0x1a, 0xe7, 0xdb, 0x8d,
	0x88, 0x8c, 0x23, 0xdd, 0xf0, 0x3c, 0x1a, 0x19, 0x91, 0x4b, 0x3d, 0xd9, 0x0f, 0x68, 0x44, 0xd1,
	0x5a, 0x82, 0x96, 0x19, 0x5a, 0x4e, 0xd0, 0xf2, 0xf9, 0x76, 0x6d, 0x3d, 0x65, 0x31, 0x7c, 0xb7,
	0x71, 0x19, 0x14, 0x26, 0x51, 0xb5, 0xbb, 0x39, 0x39, 0x1c, 0x42, 0x4f, 0x49, 0x14, 0x4c, 0x12,
	0x58, 0xfd, 0x13, 0x0f, 0x95, 0x3e, 0x19, 0x47, 0xad, 0x29, 0x01, 0x6a, 0x82, 0xe0, 0x1b, 0x0e,
	0x09, 0xab, 0xdc, 0x66, 0x61, 0xab, 0xdc, 0x5c, 0x97, 0x7f, 0x9c, 0x5f, 0xd6, 0x0c, 0x87, 0xe0,
	0x04, 0x8a, 0x10, 0xf0, 0x71, 0xf1, 0xd5, 0xa5, 0x4d, 0x6e, 0x4b, 0xc4, 0xec, 0x5c, 0x3b, 0x02,
	0x69, 0x97, 0x44, 0xc4, 0x8a, 0x88, 0xdd, 0x31, 0x3c, 0xe7, 0xcc, 0x70, 0x08, 0xba, 0x0d, 0xab,
	0xa3, 0xf4, 0xac, 0x5b, 0xd4, 0x26, 0x55, 0x8e, 0x05, 0xac, 0x64, 0xce, 0x36, 0xb5, 0x09, 0xda,
	0x00, 0xb0, 0xa8, 0x77, 0xec, 0xda, 0xc4, 0xb3, 0x08, 0xa3, 0x5c, 0xc2, 0x33, 0x9e, 0xda, 0x47,
	0x0e, 0x56, 0x33, 0x66, 0x25, 0x20, 0xc6, 0x09, 0xc2, 0xc0, 0x47, 0x13, 0x3f, 0x61, 0xab, 0x34,
	0x9f, 0xe5, 0x55, 0x3c, 0x2f, 0x54, 0x9e, 0xe3, 0x90, 0xd9, 0xb3, 0x3f, 0xf1, 0x09, 0x66, 0x5c,
	0xe8, 0x06, 0x88, 0x6e, 0xa8, 0xfb, 0x01, 0x39, 0x76, 0xc7, 0xac, 0x88, 0x12, 0x2e, 0xb9, 0xa1,
	0xc6, 0xec, 0xba, 0x05, 0xe2, 0x14, 0x8f, 0xca, 0xb0, 0x3c, 0xe8, 0x1e, 0x74, 0x7b, 0x47, 0x5d,
	0xe9, 0x1f, 0x24, 0x82, 0x70, 0xa8, 0xb5, 0xda, 0xaa, 0xc4, 0xa1, 0x0a, 0xc0, 0xe1, 0x00, 0xab,
	0x7a, 0x62, 0x2f, 0x21, 0x04, 0x15, 0xb5, 0xd7, 0xd1, 0x67, 0x7c, 0x05, 0x04, 0x50, 0xdc, 0x7b,
	0xa5, 0xed, 0xa9, 0x5d, 0x89, 0x8f, 0xf1, 0x9d, 0xfd, 0xae, 0xaa, 0x2b, 0x58, 0x6d, 0x1d, 0x48,
	0x42, 0xed, 0x1d, 0x07, 0x2b, 0x71, 0xc9, 0x5a, 0x40, 0x7d, 0x12, 0x44, 0x13, 0x74, 0x0c, 0xc8,
	0x4e, 0x6b, 0xd6, 0xb3, 0x8e, 0x65, 0x63, 0x7a, 0xf4, 0x9b, 0xa2, 0xb3, 0x91, 0xe0, 0x7f, 0xed,
	0x05, 0x4f, 0x88, 0xde, 0x40, 0x65, 0x9a, 0xc7, 0x8c, 0x65, 0x32, 0xfd, 0xe5, 0xe6, 0x83, 0x3f,
	0x69, 0x2c, 0x5e, 0xb5, 0x67, 0xcd, 0xfa, 0x07, 0x0e, 0xf8, 0xf8, 0xd5, 0x41, 0x3d, 0x28, 0xf9,
	0xa9, 0x32, 0x36, 0xb8, 0x72, 0x73, 0xe7, 0x17, 0xf9, 0x67, 0x9b, 0x82, 0xa7, 0x24, 0xe8, 0x7f,
	0x10, 0x2e, 0x5c, 0x3b, 0x1a, 0xb2, 0x6a, 0x05, 0x9c, 0x18, 0x68, 0x0d, 0x8a, 0x43, 0xe2, 0x3a,
	0xc3, 0xa8, 0x5a, 0x60, 0xee, 0xd4, 0x42, 0x0f, 0xa1, 0x68, 0x8e, 0xa8, 0x75, 0x12, 0x56, 0x79,
	0xd6, 0xc0, 0x9b, 0x79, 0xc9, 0x95, 0x18, 0x85, 0x53, 0xf0, 0xc2, 0xcb, 0x29, 0x2c, 0xbe, 0x9c,
	0xf5, 0xb7, 0x05, 0x10, 0x58, 0xc4, 0xdf, 0xd7, 0xf7, 0x02, 0x56, 0x4c, 0x7a, 0xe6, 0xd9, 0xae,
	0xe7, 0xe8, 0x26, 0x1d, 0xa7, 0x43, 0xb9, 0x93, 0x5b, 0x77, 0x8a, 0xd5, 0xe8, 0x68, 0x82, 0xcb,
	0x59, 0xa4, 0x42, 0xc7, 0xa8, 0x05, 0xe0, 0x1b, 0x81, 0xe1, 0x04, 0x86, 0x3f, 0x0c, 0xab, 0x05,
	0x26, 0xff, 0x56, 0xfe, 0x9a, 0xa7, 0x48, 0x3c, 0x13, 0x84, 0x9e, 0x03, 0xb0, 0x86, 0xe8, 0x6c,
	0xef, 0x78, 0xb6, 0x77, 0xf7, 0x7e, 0xda, 0xc1, 0xe4, 0xc9, 0x16, 0x4c, 0x34, 0xb3, 0xe3, 0x95,
	0xed, 0xc4, 0x20, 0x4e, 0xe3, 0xe6, 0x17, 0xad, 0x04, 0x7c, 0x5f, 0x7d, 0xd9, 0x97, 0xb8, 0x78,
	0xe5, 0xfa, 0x2d, 0xa5, 0x13, 0xaf, 0x58, 0x19, 0x96, 0xb5, 0xfd, 0x76, 0x7f, 0x80, 0xe3, 0xdd,
	0x12, 0x41, 0xc0, 0x83, 0x8e, 0x8a, 0x25, 0x3e, 0xf6, 0x2b, 0x2d, 0xdc, 0xee, 0xed, 0xaa, 0x92,
	0x50, 0xff, 0xc2, 0x81, 0x38, 0x55, 0x75, 0x8d, 0xc7, 0xd4, 0x04, 0xe1, 0x82, 0x06, 0x76, 0x36,
	0xa1, 0xdc, 0x3f, 0xe2, 0x23, 0x1a, 0xd8, 0x38, 0x81, 0x2e, 0xf4, 0x93, 0xff, 0xae, 0x9f, 0x5f,
	0x39, 0xe0, 0x63, 0xfc, 0x35, 0x96, 0xfd, 0x18, 0x96, 0xc3, 0xc9, 0xa9, 0x49, 0x47, 0x99, 0xf0,
	0x8d, 0x3c, 0x8e, 0x43, 0x06, 0xc3, 0x19, 0xfc, 0x4a, 0xf1, 0xef, 0x39, 0x28, 0x26, 0x31, 0xd7,
	0x58, 0x7e, 0x76, 0x95, 0x16, 0x2e, 0xaf, 0xd2, 0xab, 0x84, 0x29, 0x11, 0xd4, 0x2c, 0x7a, 0x9a,
	0x93, 0x4b, 0xf9, 0x6f, 0x5e, 0x81, 0x16, 0x5f, 0xfc, 0x1a, 0xf7, 0xfa, 0x69, 0x0a, 0x77, 0x68,
	0x7c, 0x97, 0xc8, 0x34, 0x70, 0x1a, 0x0e, 0xf1, 0xd8, 0x67, 0x41, 0x23, 0xf9, 0xc9, 0xf0, 0xdd,
	0x70, 0xf1, 0x03, 0xe2, 0x49, 0x72, 0xfa, 0xcc, 0x71, 0x66, 0x91, 0x61, 0x77, 0xbe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0x4a, 0xe7, 0xb0, 0xcf, 0x08, 0x00, 0x00,
}
