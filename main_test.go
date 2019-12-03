package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyStrToInt(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt("5")
	assert.Nil(err)
	assert.Equal(5, result)
}

func TestZero(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt("0")
	assert.Nil(err)
	assert.Equal(0, result)
}

func TestNegative(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt("-3")
	assert.Nil(err)
	assert.Equal(-3, result)
}

func TestEmpty(t *testing.T) {
	assert := assert.New(t)
	_, err := myStrToInt("")
	assert.NotNil(err)
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	_, err := myStrToInt("epam")
	assert.NotNil(err)
}

func TestALotOfZero(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt("0000")
	assert.Nil(err)
	assert.Equal(0, result)
}

func TestWithZeroPrefix(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt("07")
	assert.Nil(err)
	assert.Equal(7, result)
}






func Test2MyStrToInt(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt2("5")
	assert.Nil(err)
	assert.Equal(5, result)
}

func Test2Zero(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt2("0")
	assert.Nil(err)
	assert.Equal(0, result)
}

func Test2Negative(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt2("-3")
	assert.Nil(err)
	assert.Equal(-3, result)
}

func Test2Empty(t *testing.T) {
	assert := assert.New(t)
	_, err := myStrToInt2("")
	assert.NotNil(err)
}

func Test2String(t *testing.T) {
	assert := assert.New(t)
	_, err := myStrToInt2("epam")
	assert.NotNil(err)
}

func Test2ALotOfZero(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt2("0000")
	assert.Nil(err)
	assert.Equal(0, result)
}

func Test2WithZeroPrefix(t *testing.T) {
	assert := assert.New(t)
	result, err := myStrToInt2("07")
	assert.Nil(err)
	assert.Equal(7, result)
}

func BenchmarkMyStrToInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		myStrToInt("13")
	}
}

func BenchmarkMyStrToInt2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		myStrToInt2("13")
	}
}