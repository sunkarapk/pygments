package pygments

import (
	"testing"
)

func TestHighlighting(t *testing.T) {
	v := Highlight("print \"Hello World!\"", "python", "html", "utf-8")
	e := "<div class=\"highlight\"><pre><span class=\"k\">print</span> <span class=\"s\">&quot;Hello World!&quot;</span>\n</pre></div>\n"
	if v != e {
		t.Error("Expected", e, ", got", v)
	}
}
