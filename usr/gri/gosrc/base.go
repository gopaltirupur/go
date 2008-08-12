// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Base for the decls.go tests.

package base

export type Foo int

export type Bar *float;

export type Node struct {
  left, right *Node;
  val bool;
  f Foo;
  const, type, var, package int;
}

export func (p *Node) case(x int) {};

export type I interface {
  func();
}
