# go-mobile-collection

This is a fork of [scisci/go-mobile-collection](https://github.com/scisci/go-mobile-collection) that has been refactored to suit our needs. It is in its infancy and may be released as a separate package once it reaches maturity.

The reason for this is that go mobile doesn't currently support slices. So a current work around is to wrap a go slice with another type and expose methods that operate on that slice. So a wrapper may have methods like push, pop, insert, count, etc for modifying or reading from the array.

## Usage
What this repo does is build a command line utility that can then be automatically called using go generate semantics.

- Find the file you that contains the struct definition you want to have a collection wrapper for.
- At the top of the file add the following: `//go:generate go-mobile-collection $GOFILE`
- Before the struct, add a comment to flag it:
```
// @collection-wrapper
type Example struct {
 ExampleField string
}
```
- When you build your project a new file should now be generated called ($GOFILE)_collection.go that contains the automatically generated definitions.

