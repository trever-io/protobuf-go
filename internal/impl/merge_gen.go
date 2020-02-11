// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Code generated by generate-types. DO NOT EDIT.

package impl

import ()

func mergeBool(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.Bool() = *src.Bool()
}

func mergeBoolNoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	v := *src.Bool()
	if v != false {
		*dst.Bool() = v
	}
}

func mergeBoolPtr(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	p := *src.BoolPtr()
	if p != nil {
		v := *p
		*dst.BoolPtr() = &v
	}
}

func mergeBoolSlice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	ds := dst.BoolSlice()
	ss := src.BoolSlice()
	*ds = append(*ds, *ss...)
}

func mergeInt32(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.Int32() = *src.Int32()
}

func mergeInt32NoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	v := *src.Int32()
	if v != 0 {
		*dst.Int32() = v
	}
}

func mergeInt32Ptr(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	p := *src.Int32Ptr()
	if p != nil {
		v := *p
		*dst.Int32Ptr() = &v
	}
}

func mergeInt32Slice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	ds := dst.Int32Slice()
	ss := src.Int32Slice()
	*ds = append(*ds, *ss...)
}

func mergeUint32(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.Uint32() = *src.Uint32()
}

func mergeUint32NoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	v := *src.Uint32()
	if v != 0 {
		*dst.Uint32() = v
	}
}

func mergeUint32Ptr(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	p := *src.Uint32Ptr()
	if p != nil {
		v := *p
		*dst.Uint32Ptr() = &v
	}
}

func mergeUint32Slice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	ds := dst.Uint32Slice()
	ss := src.Uint32Slice()
	*ds = append(*ds, *ss...)
}

func mergeInt64(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.Int64() = *src.Int64()
}

func mergeInt64NoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	v := *src.Int64()
	if v != 0 {
		*dst.Int64() = v
	}
}

func mergeInt64Ptr(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	p := *src.Int64Ptr()
	if p != nil {
		v := *p
		*dst.Int64Ptr() = &v
	}
}

func mergeInt64Slice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	ds := dst.Int64Slice()
	ss := src.Int64Slice()
	*ds = append(*ds, *ss...)
}

func mergeUint64(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.Uint64() = *src.Uint64()
}

func mergeUint64NoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	v := *src.Uint64()
	if v != 0 {
		*dst.Uint64() = v
	}
}

func mergeUint64Ptr(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	p := *src.Uint64Ptr()
	if p != nil {
		v := *p
		*dst.Uint64Ptr() = &v
	}
}

func mergeUint64Slice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	ds := dst.Uint64Slice()
	ss := src.Uint64Slice()
	*ds = append(*ds, *ss...)
}

func mergeFloat32(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.Float32() = *src.Float32()
}

func mergeFloat32NoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	v := *src.Float32()
	if v != 0 {
		*dst.Float32() = v
	}
}

func mergeFloat32Ptr(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	p := *src.Float32Ptr()
	if p != nil {
		v := *p
		*dst.Float32Ptr() = &v
	}
}

func mergeFloat32Slice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	ds := dst.Float32Slice()
	ss := src.Float32Slice()
	*ds = append(*ds, *ss...)
}

func mergeFloat64(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.Float64() = *src.Float64()
}

func mergeFloat64NoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	v := *src.Float64()
	if v != 0 {
		*dst.Float64() = v
	}
}

func mergeFloat64Ptr(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	p := *src.Float64Ptr()
	if p != nil {
		v := *p
		*dst.Float64Ptr() = &v
	}
}

func mergeFloat64Slice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	ds := dst.Float64Slice()
	ss := src.Float64Slice()
	*ds = append(*ds, *ss...)
}

func mergeString(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.String() = *src.String()
}

func mergeStringNoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	v := *src.String()
	if v != "" {
		*dst.String() = v
	}
}

func mergeStringPtr(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	p := *src.StringPtr()
	if p != nil {
		v := *p
		*dst.StringPtr() = &v
	}
}

func mergeStringSlice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	ds := dst.StringSlice()
	ss := src.StringSlice()
	*ds = append(*ds, *ss...)
}
