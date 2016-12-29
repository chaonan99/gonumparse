// This is a small example to compare two string concatenate method in Go

package main

import (
    "bytes"
    "testing"
)

func BenchmarkStradd(b *testing.B) {
    a := ""
    for n := 0; n < b.N; n++ {
        a += "xxx"
    }
}

func BenchmarkBytes(b *testing.B) {
    var r bytes.Buffer
    for n := 0; n < b.N; n++ {
        r.WriteString("xxx")
    }
    _ = r.String()
}
