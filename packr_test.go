package packr_migrate

import (
	"testing"

	"github.com/gobuffalo/packr/v2"
	st "github.com/golang-migrate/migrate/v4/source/testing"
)

func Test(t *testing.T) {
	// wrap assets into Resource first
	Box := packr.New("testbox", "./testdata/")

	d, err := WithInstance(Box)
	if err != nil {
		t.Fatal(err)
	}
	st.Test(t, d)
}

func TestOpen(t *testing.T) {
	b := &Packr{}
	_, err := b.Open("")
	if err == nil {
		t.Fatal("expected err, because it's not implemented yet")
	}
}