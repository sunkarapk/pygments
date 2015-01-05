# pygments

A pygments wrapper for golang

## Importing

Run this in terminal to install the package

```
go get github.com/pksunkara/pygments
```

Use the following line in source file to import this package

```go
import "github.com/pksunkara/pygments"
```

## Usage

#### Highlighting Code

```go
pygments.Highlight("code", "lexer", "format", "encoding")
```

```go
pygments.Highlight("print \"Hello World!\"", "python", "html", "utf-8")
```

#### Provide Custom Path

```go
pygments.Binary("/path/to/pygments/binary")
```

## Documentation

Visit the docs on [gopkgdoc](http://godoc.org/github.com/pksunkara/pygments.go)

## Testing

```
go test
```

If you like this project, please watch/star this and follow me

## Contributors
Here is a list of [Contributors](http://github.com/pksunkara/pygments.go/contributors)

### TODO

- Improve speed and performance

__I accept pull requests and guarantee a reply back within a day__

## License
MIT/X11

## Bug Reports
Report [here](http://github.com/pksunkara/pygments.go/issues). __Guaranteed reply within a day__.

## Contact
Pavan Kumar Sunkara (pavan.sss1991@gmail.com)

Follow me on [github](https://github.com/users/follow?target=pksunkara), [twitter](http://twitter.com/pksunkara)
