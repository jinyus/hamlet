// Code generated by hamlet - DO NOT EDIT.

package __kitchensink

import (
	"bytes"
	"context"
	"github.com/stackus/hamlet"
	"io"
)

func SiteLayout() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<html lang=\"en\"><head><title>Hamlet</title></head><body><h1>Hamlet</h1><p>A HAML-like template engine for Go.</p>"); __err != nil {
			return
		}
		if __err = __children.Render(ctx, __buf); __err != nil {
			return
		}
		if _, __err = __buf.WriteString("</body></html>"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __buf.WriteTo(__w)
		}
		return
	})
}

func HomePage() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		__var1 := hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
			__buf, __isBuf := __w.(*bytes.Buffer)
			if !__isBuf {
				__buf = hamlet.GetBuffer()
				defer hamlet.ReleaseBuffer(__buf)
			}
			if _, __err = __buf.WriteString("<p>This is the home page.</p>"); __err != nil {
				return
			}
			if !__isBuf {
				_, __err = io.Copy(__w, __buf)
			}
			return
		})
		if __err = SiteLayout().Render(hamlet.PushChildren(ctx, __var1), __w); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __buf.WriteTo(__w)
		}
		return
	})
}