package main

import "testing"
import "github.com/stretchr/testify/assert"

func Test_generateFilterName(t *testing.T) {
	tests := []struct {
		wasmName string
		expect   string
	}{
		{
			wasmName: "foo.wasm",
			expect:   "foo_wasm_filter",
		},
		{
			wasmName: "bar.wasm",
			expect:   "bar_wasm_filter",
		},
		{
			wasmName: "",
			expect:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.wasmName, func(t *testing.T) {
			assert.Equal(t, tt.expect, generateFilterName("", tt.wasmName))
		})
	}
}
