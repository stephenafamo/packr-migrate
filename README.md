# Packr Migrate

## Usage

### Read Packr box with NewWithSourceInstance

```go
import (
  "github.com/gobuffalo/packr/v2"
  "github.com/golang-migrate/migrate/v4"
  "github.com/stephenafamo/packr-migrate"
)

func main() {
  // wrap assets into Resource
  box := packr.New("myBox", "./migrations")

  d, err := packr_migrate.WithInstance(box)
  m, err := migrate.NewWithSourceInstance("packr", d, "database://foobar")
  m.Up() // run your migrations and handle the errors above of course
}
```

