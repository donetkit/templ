// Code generated by templ@(devel) DO NOT EDIT.

package testtextwhitespace

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func WhitespaceIsAddedWithinTemplStatements() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<p>")
		if err != nil {
			return err
		}
		// Text
		var_2 := `This is some text.`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		// Whitespace (normalised)
		_, err = templBuffer.WriteString(` `)
		if err != nil {
			return err
		}
		// If
		if true {
			// Text
			var_3 := `So is this.`
			_, err = templBuffer.WriteString(var_3)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</p>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// GoExpression
const WhitespaceIsAddedWithinTemplStatementsExpected = `<p>This is some text. So is this.</p>`

func InlineElementsAreNotPadded() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<p>")
		if err != nil {
			return err
		}
		// Text
		var_5 := `Inline text `
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<b>")
		if err != nil {
			return err
		}
		// Text
		var_6 := `is spaced properly`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</b>")
		if err != nil {
			return err
		}
		// Whitespace (normalised)
		_, err = templBuffer.WriteString(` `)
		if err != nil {
			return err
		}
		// Text
		var_7 := `without adding extra spaces.`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// GoExpression
const InlineElementsAreNotPaddedExpected = `<p>Inline text <b>is spaced properly</b> without adding extra spaces.</p>`

func WhiteSpaceInHTMLIsNormalised() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_8 := templ.GetChildren(ctx)
		if var_8 == nil {
			var_8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<p>")
		if err != nil {
			return err
		}
		// Text
		var_9 := `newlines and other whitespace are stripped`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		// Whitespace (normalised)
		_, err = templBuffer.WriteString(` `)
		if err != nil {
			return err
		}
		// Text
		var_10 := `but it is normalised`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		// Whitespace (normalised)
		_, err = templBuffer.WriteString(` `)
		if err != nil {
			return err
		}
		// Text
		var_11 := `like HTML.`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// GoExpression
const WhiteSpaceInHTMLIsNormalisedExpected = `<p>newlines and other whitespace are stripped but it is normalised like HTML.</p>`

func WhiteSpaceAroundValues() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_12 := templ.GetChildren(ctx)
		if var_12 == nil {
			var_12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<p>")
		if err != nil {
			return err
		}
		// Text
		var_13 := `templ allows `
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		// StringExpression
		var var_14 string = "strings"
		_, err = templBuffer.WriteString(templ.EscapeString(var_14))
		if err != nil {
			return err
		}
		// Whitespace (normalised)
		_, err = templBuffer.WriteString(` `)
		if err != nil {
			return err
		}
		// Text
		var_15 := `to be included in sentences.`
		_, err = templBuffer.WriteString(var_15)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// GoExpression
const WhiteSpaceAroundValuesExpected = `<p>templ allows strings to be included in sentences.</p>`

