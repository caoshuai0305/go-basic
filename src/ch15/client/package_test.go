package client

import (
	"basic/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
}
