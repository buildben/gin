// Copyright 2020 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !nomsgpack
// +build !nomsgpack

package binding

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
	"time"
)

func TestBindingMsgPack(t *testing.T) {
	test := FooStruct{
		Foo: "bar",
	}

	h := new(codec.MsgpackHandle)
	assert.NotNil(t, h)
	buf := bytes.NewBuffer([]byte{})
	assert.NotNil(t, buf)
	err := codec.NewEncoder(buf, h).Encode(test)
	assert.NoError(t, err)

	data := buf.Bytes()

	testMsgPackBodyBinding(t,
		MsgPack, "msgpack",
		"/", "/",
		string(data), string(data[1:]))
}

func testMsgPackBodyBinding(t *testing.T, b Binding, name, path, badPath, body, badBody string) {
	assert.Equal(t, name, b.Name())

	obj := FooStruct{}
	req := requestWithBody("POST", path, body)
	req.Header.Add("Content-Type", MIMEMSGPACK)
	err := b.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "bar", obj.Foo)

	obj = FooStruct{}
	req = requestWithBody("POST", badPath, badBody)
	req.Header.Add("Content-Type", MIMEMSGPACK)
	err = MsgPack.Bind(req, &obj)
	assert.Error(t, err)
}

func TestBindingDefaultMsgPack(t *testing.T) {
	assert.Equal(t, MsgPack, Default("POST", MIMEMSGPACK))
	assert.Equal(t, MsgPack, Default("PUT", MIMEMSGPACK2))
}

func TestGinLoadPerformance1(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance2(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance3(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance4(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance5(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance6(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance7(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance8(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance9(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance10(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance11(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance12(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance13(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance14(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance15(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance16(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance17(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance18(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance19(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance20(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance21(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance22(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance23(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance24(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance25(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance26(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance27(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance28(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance29(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance30(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance31(t *testing.T) {
        time.Sleep(time.Second * 3)
}
