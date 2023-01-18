package main

import (
	"github.com/jonnyorman/go-all/build"
	"github.com/jonnyorman/go-all/mod"
	"github.com/jonnyorman/go-all/test"
)

func init() {
	build.AddTo(rootCmd)
	mod.AddTo(rootCmd)
	test.AddTo(rootCmd)
}
