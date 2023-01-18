module github.com/jonnyorman/go-all/test

go 1.19

replace github.com/jonnyorman/go-all/command => ../command

require (
	github.com/jonnyorman/go-all/command v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.6.1
)

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.7.0 // indirect
)
