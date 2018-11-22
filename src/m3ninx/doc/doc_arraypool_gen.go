// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package doc

import (
	"github.com/m3db/m3x/pool"
)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// DocumentArrayPool provides a pool for document slices.
type DocumentArrayPool interface {
	// Init initializes the array pool, it needs to be called
	// before Get/Put use.
	Init()

	// Get returns the a slice from the pool.
	Get() []Document

	// Put returns the provided slice to the pool.
	Put(elems []Document)
}

type DocumentFinalizeFn func([]Document) []Document

type DocumentArrayPoolOpts struct {
	Options     pool.ObjectPoolOptions
	Capacity    int
	MaxCapacity int
	FinalizeFn  DocumentFinalizeFn
}

type DocumentArrPool struct {
	opts DocumentArrayPoolOpts
	pool pool.ObjectPool
}

func NewDocumentArrayPool(opts DocumentArrayPoolOpts) DocumentArrayPool {
	if opts.FinalizeFn == nil {
		opts.FinalizeFn = defaultDocumentFinalizerFn
	}
	p := pool.NewObjectPool(opts.Options)
	return &DocumentArrPool{opts, p}
}

func (p *DocumentArrPool) Init() {
	p.pool.Init(func() interface{} {
		return make([]Document, 0, p.opts.Capacity)
	})
}

func (p *DocumentArrPool) Get() []Document {
	return p.pool.Get().([]Document)
}

func (p *DocumentArrPool) Put(arr []Document) {
	arr = p.opts.FinalizeFn(arr)
	if max := p.opts.MaxCapacity; max > 0 && cap(arr) > max {
		return
	}
	p.pool.Put(arr)
}

func defaultDocumentFinalizerFn(elems []Document) []Document {
	var empty Document
	for i := range elems {
		elems[i] = empty
	}
	elems = elems[:0]
	return elems
}

type DocumentArr []Document

func (elems DocumentArr) grow(n int) []Document {
	if cap(elems) < n {
		elems = make([]Document, n)
	}
	elems = elems[:n]
	// following compiler optimized memcpy impl
	// https://github.com/golang/go/wiki/CompilerOptimizations#optimized-memclr
	var empty Document
	for i := range elems {
		elems[i] = empty
	}
	return elems
}
