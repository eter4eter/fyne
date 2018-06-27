package canvas

import "image/color"

import "github.com/fyne-io/fyne"
import "github.com/fyne-io/fyne/theme"

// Text describes a text primitive in a Fyne canvas
type Text struct {
	Size      fyne.Size      // The current size of the Text - the font will not scale to this Size
	Position  fyne.Position  // The current position of the Text
	Alignment fyne.TextAlign // The alignment of the text content within the current size
	Options   Options        // Options to pass to the renderer

	Color    color.RGBA // The main text draw colour
	Text     string     // The string content of this Text
	FontSize int        // Size of the font - if the Canvas scale is 1.0 this will be equivalent to point size
	Bold     bool       // Should the text be bold
	Italic   bool       // Should the text be italic
}

// CurrentSize gets the current size of this text object
func (t *Text) CurrentSize() fyne.Size {
	return t.Size
}

// Resize sets a new size for the text object
func (t *Text) Resize(size fyne.Size) {
	t.Size = size
}

// CurrentPosition gets the current position of this text object, relative to it's parent / canvas
func (t *Text) CurrentPosition() fyne.Position {
	return t.Position
}

// Move the text object to a new position, relative to it's parent / canvas
func (t *Text) Move(pos fyne.Position) {
	t.Position = pos
}

// MinSize returns the minimum size of this text objet based on it's font size and content.
// This is normally determined by the render implementation.
func (t *Text) MinSize() fyne.Size {
	return fyne.GetDriver().RenderedTextSize(t.Text, t.FontSize)
}

// NewText returns a new Text implementation
func NewText(text string, color color.RGBA) *Text {
	return &Text{
		Color:    color,
		Text:     text,
		FontSize: theme.TextSize(),
	}
}
