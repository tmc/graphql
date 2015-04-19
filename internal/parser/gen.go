// requires pigeon and goexports
// go get golang.org/x/tools/cmd/goimports github.com/PuerkitoBio/pigeon
//go:generate pigeon -o parser.go ../../graphql.peg
//go:generate goimports -w parser.go

package parser
