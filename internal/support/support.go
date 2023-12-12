// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package support provides support functions for protoveneer.
package support

import (
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"
)

// TransformSlice applies f to each element of from and returns
// a new slice with the results.
func TransformSlice[From, To any](from []From, f func(From) To) []To {
	if from == nil {
		return nil
	}
	to := make([]To, len(from))
	for i, e := range from {
		to[i] = f(e)
	}
	return to
}

// TransformMapValues applies f to each value of from, returning a new map.
// It does not change the keys.
func TransformMapValues[K comparable, VFrom, VTo any](from map[K]VFrom, f func(VFrom) VTo) map[K]VTo {
	if from == nil {
		return nil
	}
	to := map[K]VTo{}
	for k, v := range from {
		to[k] = f(v)
	}
	return to
}

// ZeroToNil returns nil if x is the zero value for T,
// or &x otherwise.
func ZeroToNil[T comparable](x T) *T {
	var z T
	if x == z {
		return nil
	}
	return &x
}

// NilToZero returns the zero value for T if x is nil,
// or *x otherwise.
func NilToZero[T any](x *T) T {
	if x == nil {
		var z T
		return z
	}
	return *x
}

// MapToStructPB converts a map into a structpb.Struct.
func MapToStructPB(m map[string]any) *structpb.Struct {
	if m == nil {
		return nil
	}
	s, err := structpb.NewStruct(m)
	if err != nil {
		panic(fmt.Errorf("support.MapToProto: %w", err))
	}
	return s
}

// MapFromStructPB converts a structpb.Struct to a map.
func MapFromStructPB(p *structpb.Struct) map[string]any {
	if p == nil {
		return nil
	}
	return p.AsMap()
}