package packr_migrate

import (
	"fmt"
	"io"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/golang-migrate/migrate/v4/source"
)

func init() {
	source.Register("packr", &Packr{})
}

func WithInstance(box *packr.Box) (source.Driver, error) {
	bn := &Packr{
		box: box,
		migrations:  source.NewMigrations(),
	}

	for _, fi := range box.List() {
		m, err := source.DefaultParse(fi)
		if err != nil {
			continue // ignore files that we can't parse
		}

		if !bn.migrations.Append(m) {
			return nil, fmt.Errorf("unable to parse file %v", fi)
		}
	}

	return bn, nil
}

type Packr struct {
	box        *packr.Box
	migrations  *source.Migrations
}

func (p *Packr) Open(url string) (source.Driver, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (p *Packr) Close() error {
	return nil
}

func (p *Packr) First() (version uint, err error) {
	if v, ok := p.migrations.First(); !ok {
		return 0, &os.PathError{Op: "first", Path: p.box.Path, Err: os.ErrNotExist}
	} else {
		return v, nil
	}
}

func (p *Packr) Prev(version uint) (prevVersion uint, err error) {
	if v, ok := p.migrations.Prev(version); !ok {
		return 0, &os.PathError{Op: fmt.Sprintf("prev for version %v", version), Path: p.box.Path, Err: os.ErrNotExist}
	} else {
		return v, nil
	}
}

func (p *Packr) Next(version uint) (nextVersion uint, err error) {
	if v, ok := p.migrations.Next(version); !ok {
		return 0, &os.PathError{Op: fmt.Sprintf("next for version %v", version), Path: p.box.Path, Err: os.ErrNotExist}
	} else {
		return v, nil
	}
}

func (p *Packr) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := p.migrations.Up(version); ok {
		r, err := p.box.Open(m.Raw)
		if err != nil {
			return nil, "", err
		}
		return r, m.Identifier, nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read version %v", version), Path: p.box.Path, Err: os.ErrNotExist}
}

func (p *Packr) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := p.migrations.Down(version); ok {
		r, err := p.box.Open(m.Raw)
		if err != nil {
			return nil, "", err
		}
		return r, m.Identifier, nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read version %v", version), Path: p.box.Path, Err: os.ErrNotExist}
}
