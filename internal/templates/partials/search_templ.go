// Code generated by templ@v0.2.282 DO NOT EDIT.

package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Search(term string) templ.Component {
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
		_, err = templBuffer.WriteString("<form")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" method=\"GET\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" action=\"/todos\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"inline [&amp;:has(+ul:empty)]:hidden\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<label")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"flex items-center\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<span")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"text-lg font-bold\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Text
		var_2 := `Search`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" id=\"search\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" name=\"search\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" value=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(term))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" type=\"text\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" placeholder=\"Begin typing to search...\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-get=\"/todos\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-target=\"#todos\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-trigger=\"keyup changed, search\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-replace=\"innerHTML\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"ml-2 grow\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
