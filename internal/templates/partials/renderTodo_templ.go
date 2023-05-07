// Code generated by templ@v0.2.282 DO NOT EDIT.

package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// GoExpression
import (
	"github.com/stackus/todos/internal/domain"
)

func RenderTodo(todo *domain.Todo) templ.Component {
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
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"block py-2 border-b-4 border-dotted border-red-900 draggable\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<form")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" method=\"POST\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" action=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString("/todos/" + todo.ID.String() + "/delete"))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"inline\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<button")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" type=\"submit\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-target=\"closest div\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-swap=\"outerHTML\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-delete=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString("/todos/" + todo.ID.String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"focus:outline focus:outline-red-500 focus:outline-4 mr-2\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Text
		var_2 := `❌`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</form>")
		if err != nil {
			return err
		}
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
		_, err = templBuffer.WriteString(" action=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString("/todos/" + todo.ID.String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"inline\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<button")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" type=\"submit\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-target=\"closest div\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-swap=\"outerHTML\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-get=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString("/todos/" + todo.ID.String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"focus:outline focus:outline-red-500 focus:outline-4 mr-2\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Text
		var_3 := `📝`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</form>")
		if err != nil {
			return err
		}
		// Element (standard)
		// Element CSS
		var var_4 = []any{"inline", templ.KV("line-through", todo.Completed)}
		err = templ.RenderCSSItems(ctx, templBuffer, var_4...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<form")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" method=\"POST\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" action=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString("/todos/" + todo.ID.String() + "/edit"))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-target=\"closest div\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-swap=\"outerHTML\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_4).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" type=\"hidden\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" name=\"completed\"")
		if err != nil {
			return err
		}
		if !todo.Completed {
			// Element Attributes
			_, err = templBuffer.WriteString(" value=\"true\"")
			if err != nil {
				return err
			}
		} else {
			// Element Attributes
			_, err = templBuffer.WriteString(" value=\"false\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" type=\"hidden\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" name=\"description\"")
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
		_, err = templBuffer.WriteString(templ.EscapeString(todo.Description))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<noscript>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" type=\"submit\"")
		if err != nil {
			return err
		}
		if todo.Completed {
			// Element Attributes
			_, err = templBuffer.WriteString(" value=\"Set as Not Completed\"")
			if err != nil {
				return err
			}
		} else {
			// Element Attributes
			_, err = templBuffer.WriteString(" value=\"Set as Completed\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" class=\"mr-2\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</noscript>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<span")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" hx-patch=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString("/todos/" + todo.ID.String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// StringExpression
		var var_5 string = todo.Description
		_, err = templBuffer.WriteString(templ.EscapeString(var_5))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</form>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" type=\"hidden\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" name=\"id\"")
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
		_, err = templBuffer.WriteString(templ.EscapeString(todo.ID.String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}