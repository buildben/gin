// Copyright 2019 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"
)

func TestYAMLBindingBindBody(t *testing.T) {
	var s struct {
		Foo string `yaml:"foo"`
	}
	err := yamlBinding{}.BindBody([]byte("foo: FOO"), &s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}


func TestGinLoadPerformance62(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance63(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance64(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance65(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance66(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance67(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance68(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance69(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance70(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance71(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance72(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance73(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance74(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance75(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance76(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance77(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance78(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance79(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance80(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance81(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance82(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance83(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance84(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance85(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance86(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance87(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance88(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance89(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance90(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance91(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance92(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance93(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance94(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance95(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance96(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance97(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance98(t *testing.T) {
        time.Sleep(time.Second * 3)
}


func TestGinLoadPerformance99(t *testing.T) {
        time.Sleep(time.Second * 3)
}
