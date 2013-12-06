package pygments

import (
	"testing"
)

func TestBinary(t *testing.T) {
	Binary("/usr/bin/pygmentize")
}

func TestWhich(t *testing.T) {
	v := Which()
	e := "/usr/bin/pygmentize"
	if v != e {
		t.Error("Expected", e, ", got", v)
	}
}

func TestHighlighting(t *testing.T) {
	v := Highlight("print \"Hello World!\"", "python", "html", "utf-8")
	e := "<div class=\"highlight\"><pre><span class=\"k\">print</span> <span class=\"s\">&quot;Hello World!&quot;</span>\n</pre></div>\n"
	if v != e {
		t.Error("Expected", e, ", got", v)
	}
}
