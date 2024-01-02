// Code generated by hamlet - DO NOT EDIT.

package main

import (
	"bytes"
	"context"
	"github.com/stackus/hamlet"
	"io"
)

// For most attribute names you can include the name in the list
// of attributes just as you expect it to appear in the HTML. Names
// that contain alphanumeric characters, dashes (-), and
// underscores (_) are all acceptable as-is.

func SimpleNames() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<a href=\"https://github.com/stackus/hamlet\" data-foo=\"bar\" odd_name=\"baz\" _=\"I&#39;m a _hyperscript attribute!\">Hamlet</a>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}

// For more complex names, such as data attributes, you can use
// enclose the name in in double quotes or backticks.
// - Names that start with an at sign (@).
// - Names that contain a colon (:).
// - Names that contain a question mark (?).
// The names will be rendered into the HTML without the quotes.

func ComplexNames() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<a href=\"https://github.com/stackus/hamlet\" :class=\"show ? &#39;&#39; : &#39;hidden&#39;\" @click=\"show = !show\">Hamlet</a>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}