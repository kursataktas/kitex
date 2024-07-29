/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package thrift

import (
	"context"
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"

	"github.com/cloudwego/gopkg/protocol/thrift"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/stretchr/testify/require"

	"github.com/cloudwego/kitex/internal/generic/proto"
	"github.com/cloudwego/kitex/pkg/generic/descriptor"
	"github.com/cloudwego/kitex/pkg/remote"
)

var (
	stringInput = "hello world"
	binaryInput = []byte(stringInput)
)

func Test_nextReader(t *testing.T) {
	type args struct {
		tt  descriptor.Type
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    reader
		wantErr bool
	}{
		// TODO: Add test cases.
		{"void", args{tt: descriptor.VOID, t: &descriptor.TypeDescriptor{Type: descriptor.VOID}, opt: &readerOption{}}, readVoid, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nextReader(tt.args.tt, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("nextReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if fmt.Sprintf("%p", got) != fmt.Sprintf("%p", tt.want) {
				t.Errorf("nextReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readVoid(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"void", args{t: &descriptor.TypeDescriptor{Type: descriptor.VOID, Struct: &descriptor.StructDescriptor{}}}, descriptor.Void{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeVoid(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeVoid() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readVoid(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readVoid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readVoid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readDouble(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readDouble", args{t: &descriptor.TypeDescriptor{Type: descriptor.DOUBLE}}, 1.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeFloat64(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeFloat64() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readDouble(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDouble() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readBool(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readBool", args{t: &descriptor.TypeDescriptor{Type: descriptor.BOOL}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeBool(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeBool() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readBool(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readByte(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readByte", args{t: &descriptor.TypeDescriptor{Type: descriptor.BYTE}, opt: &readerOption{}}, int8(1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeInt8(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeInt8() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readByte(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readInt16(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readInt16", args{t: &descriptor.TypeDescriptor{Type: descriptor.I16}, opt: &readerOption{}}, int16(1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeInt16(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeInt16() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readInt16(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readInt32(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readInt32", args{t: &descriptor.TypeDescriptor{Type: descriptor.I32}}, int32(1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeInt32(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeInt32() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readInt32(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readInt64(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readInt64", args{t: &descriptor.TypeDescriptor{Type: descriptor.I64}}, int64(1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeInt64(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeInt64() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readInt64(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readString(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readString", args{t: &descriptor.TypeDescriptor{Type: descriptor.STRING}}, stringInput, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeString(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeString() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readString(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readBinary64String(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readBase64Binary", args{t: &descriptor.TypeDescriptor{Name: "binary", Type: descriptor.STRING}}, base64.StdEncoding.EncodeToString(binaryInput), false}, // read base64 string from binary field
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeBase64Binary(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeBase64Binary() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readBase64Binary(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readBinary(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readBinary", args{t: &descriptor.TypeDescriptor{Name: "binary", Type: descriptor.STRING}}, binaryInput, false}, // read base64 string from binary field
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeBinary(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeBinary() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readBinary(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readList(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"readList", args{t: &descriptor.TypeDescriptor{Type: descriptor.LIST, Elem: &descriptor.TypeDescriptor{Type: descriptor.STRING}}}, []interface{}{stringInput, stringInput, stringInput}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeList(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeList() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readList(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readMap(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		writer  func(ctx context.Context, val interface{}, out *thrift.BinaryWriter, t *descriptor.TypeDescriptor, opt *writerOption) error
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"readMap",
			args{t: &descriptor.TypeDescriptor{Type: descriptor.MAP, Key: &descriptor.TypeDescriptor{Type: descriptor.STRING}, Elem: &descriptor.TypeDescriptor{Type: descriptor.STRING}, Struct: &descriptor.StructDescriptor{}}, opt: &readerOption{}},
			writeInterfaceMap,
			map[interface{}]interface{}{"hello": "world"},
			false,
		},
		{
			"readJsonMap",
			args{t: &descriptor.TypeDescriptor{Type: descriptor.MAP, Key: &descriptor.TypeDescriptor{Type: descriptor.STRING}, Elem: &descriptor.TypeDescriptor{Type: descriptor.STRING}, Struct: &descriptor.StructDescriptor{}}, opt: &readerOption{forJSON: true}},
			writeStringMap,
			map[string]interface{}{"hello": "world"},
			false,
		},
		{
			"readJsonMapWithInt16Key",
			args{t: &descriptor.TypeDescriptor{Type: descriptor.MAP, Key: &descriptor.TypeDescriptor{Type: descriptor.I16}, Elem: &descriptor.TypeDescriptor{Type: descriptor.BOOL}, Struct: &descriptor.StructDescriptor{}}, opt: &readerOption{forJSON: true}},
			writeStringMap,
			map[string]interface{}{"16": false},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := tt.writer(context.Background(), tt.want, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeInterfaceMap() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readMap(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readStruct(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	tests := []struct {
		name    string
		args    args
		input   interface{}
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"readStruct with setFieldsForEmptyStruct",
			args{t: &descriptor.TypeDescriptor{
				Type: descriptor.STRUCT,
				Struct: &descriptor.StructDescriptor{
					FieldsByID: map[int32]*descriptor.FieldDescriptor{
						1:  {Name: "string", Type: &descriptor.TypeDescriptor{Type: descriptor.STRING, Name: "string"}},
						2:  {Name: "byte", Type: &descriptor.TypeDescriptor{Type: descriptor.BYTE, Name: "byte"}},
						3:  {Name: "i8", Type: &descriptor.TypeDescriptor{Type: descriptor.I08, Name: "i8"}},
						4:  {Name: "i16", Type: &descriptor.TypeDescriptor{Type: descriptor.I16, Name: "i16"}},
						5:  {Name: "i32", Type: &descriptor.TypeDescriptor{Type: descriptor.I32, Name: "i32"}},
						6:  {Name: "i64", Type: &descriptor.TypeDescriptor{Type: descriptor.I64, Name: "i64"}},
						7:  {Name: "double", Type: &descriptor.TypeDescriptor{Type: descriptor.DOUBLE, Name: "double"}},
						8:  {Name: "list", Type: &descriptor.TypeDescriptor{Type: descriptor.LIST, Name: "list", Elem: &descriptor.TypeDescriptor{Type: descriptor.STRING, Name: "binary"}}},
						9:  {Name: "set", Type: &descriptor.TypeDescriptor{Type: descriptor.SET, Name: "set", Elem: &descriptor.TypeDescriptor{Type: descriptor.BOOL}}},
						10: {Name: "map", Type: &descriptor.TypeDescriptor{Type: descriptor.MAP, Name: "map", Key: &descriptor.TypeDescriptor{Type: descriptor.STRING, Name: "string"}, Elem: &descriptor.TypeDescriptor{Type: descriptor.DOUBLE}}},
						11: {Name: "struct", Type: &descriptor.TypeDescriptor{Type: descriptor.STRUCT}},
						12: {Name: "list-list", Type: &descriptor.TypeDescriptor{Type: descriptor.LIST, Name: "list", Elem: &descriptor.TypeDescriptor{Type: descriptor.LIST, Name: "list", Elem: &descriptor.TypeDescriptor{Type: descriptor.STRING, Name: "binary"}}}},
						13: {Name: "map-list", Type: &descriptor.TypeDescriptor{Type: descriptor.MAP, Name: "map", Key: &descriptor.TypeDescriptor{Type: descriptor.STRING, Name: "string"}, Elem: &descriptor.TypeDescriptor{Type: descriptor.LIST, Name: "list", Elem: &descriptor.TypeDescriptor{Type: descriptor.STRING, Name: "binary"}}}},
					},
				},
			}, opt: &readerOption{setFieldsForEmptyStruct: 1}},
			map[string]interface{}{
				"string":    "",
				"byte":      byte(0),
				"i8":        int8(0),
				"i16":       int16(0),
				"i32":       int32(0),
				"i64":       int64(0),
				"double":    float64(0),
				"list":      [][]byte(nil),
				"set":       []bool(nil),
				"map":       map[string]float64(nil),
				"struct":    map[string]interface{}(nil),
				"list-list": [][][]byte(nil),
				"map-list":  map[string][][]byte(nil),
			},
			map[string]interface{}{
				"string":    "",
				"byte":      byte(0),
				"i8":        int8(0),
				"i16":       int16(0),
				"i32":       int32(0),
				"i64":       int64(0),
				"double":    float64(0),
				"list":      [][]byte(nil),
				"set":       []bool(nil),
				"map":       map[string]float64(nil),
				"struct":    map[string]interface{}(nil),
				"list-list": [][][]byte(nil),
				"map-list":  map[string][][]byte(nil),
			},
			false,
		},
		{
			"readStruct with setFieldsForEmptyStruct optional",
			args{t: &descriptor.TypeDescriptor{
				Type: descriptor.STRUCT,
				Struct: &descriptor.StructDescriptor{
					FieldsByID: map[int32]*descriptor.FieldDescriptor{
						1:  {Name: "string", Type: &descriptor.TypeDescriptor{Type: descriptor.STRING, Name: "string"}},
						2:  {Name: "byte", Type: &descriptor.TypeDescriptor{Type: descriptor.BYTE, Name: "byte"}, Optional: true},
						3:  {Name: "i8", Type: &descriptor.TypeDescriptor{Type: descriptor.I08, Name: "i8"}},
						4:  {Name: "i16", Type: &descriptor.TypeDescriptor{Type: descriptor.I16, Name: "i16"}},
						5:  {Name: "i32", Type: &descriptor.TypeDescriptor{Type: descriptor.I32, Name: "i32"}},
						6:  {Name: "i64", Type: &descriptor.TypeDescriptor{Type: descriptor.I64, Name: "i64"}},
						7:  {Name: "double", Type: &descriptor.TypeDescriptor{Type: descriptor.DOUBLE, Name: "double"}},
						8:  {Name: "list", Type: &descriptor.TypeDescriptor{Type: descriptor.LIST, Name: "list", Elem: &descriptor.TypeDescriptor{Type: descriptor.I08}}},
						9:  {Name: "set", Type: &descriptor.TypeDescriptor{Type: descriptor.SET, Name: "set", Elem: &descriptor.TypeDescriptor{Type: descriptor.I16}}},
						10: {Name: "map", Type: &descriptor.TypeDescriptor{Type: descriptor.MAP, Name: "map", Key: &descriptor.TypeDescriptor{Type: descriptor.I32}, Elem: &descriptor.TypeDescriptor{Type: descriptor.I64}}},
						11: {Name: "struct", Type: &descriptor.TypeDescriptor{Type: descriptor.STRUCT}, Optional: true},
					},
				},
			}, opt: &readerOption{setFieldsForEmptyStruct: 1}},
			map[string]interface{}{
				"string": "",
				"i8":     int8(0),
				"i16":    int16(0),
				"i32":    int32(0),
				"i64":    int64(0),
				"double": float64(0),
				"list":   []int8(nil),
				"set":    []int16(nil),
				"map":    map[int32]int64(nil),
			},
			map[string]interface{}{
				"string": "",
				"i8":     int8(0),
				"i16":    int16(0),
				"i32":    int32(0),
				"i64":    int64(0),
				"double": float64(0),
				"list":   []int8(nil),
				"set":    []int16(nil),
				"map":    map[int32]int64(nil),
			},
			false,
		},
		{
			"readStruct",
			args{t: &descriptor.TypeDescriptor{
				Type: descriptor.STRUCT,
				Struct: &descriptor.StructDescriptor{
					FieldsByID: map[int32]*descriptor.FieldDescriptor{
						1: {Name: "hello", Type: &descriptor.TypeDescriptor{Type: descriptor.STRING}},
					},
					FieldsByName: map[string]*descriptor.FieldDescriptor{
						"hello": {Name: "hello", ID: 1, Type: &descriptor.TypeDescriptor{Type: descriptor.STRING}},
					},
				},
			}},
			map[string]interface{}{"hello": "world"},
			map[string]interface{}{"hello": "world"},
			false,
		},
		{
			"readStructError",
			args{t: &descriptor.TypeDescriptor{
				Type: descriptor.STRUCT,
				Struct: &descriptor.StructDescriptor{
					FieldsByID: map[int32]*descriptor.FieldDescriptor{
						1: {Name: "strList", Type: &descriptor.TypeDescriptor{Type: descriptor.LIST, Elem: &descriptor.TypeDescriptor{Type: descriptor.I64}}},
					},
					FieldsByName: map[string]*descriptor.FieldDescriptor{
						"strList": {Name: "strList", ID: 1, Type: &descriptor.TypeDescriptor{Type: descriptor.LIST, Elem: &descriptor.TypeDescriptor{Type: descriptor.STRING}}},
					},
				},
			}},
			map[string]interface{}{},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			err := writeStruct(context.Background(), tt.input, w, tt.args.t, &writerOption{})
			if err != nil {
				t.Errorf("writeStruct() error = %v", err)
			}
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readStruct(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_readHTTPResponse(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	resp := descriptor.NewHTTPResponse()
	resp.Body = map[string]interface{}{"hello": "world"}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"readHTTPResponse",
			args{t: &descriptor.TypeDescriptor{
				Type: descriptor.STRUCT,
				Struct: &descriptor.StructDescriptor{
					FieldsByID: map[int32]*descriptor.FieldDescriptor{
						1: {
							Name:        "hello",
							Type:        &descriptor.TypeDescriptor{Type: descriptor.STRING},
							HTTPMapping: descriptor.DefaultNewMapping("hello"),
						},
					},
				},
			}},
			resp,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			w.WriteFieldBegin(thrift.TType(descriptor.STRING), 1)
			w.WriteString("world")
			w.WriteFieldStop()
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readHTTPResponse(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readHTTPResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readHTTPResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readHTTPResponseWithPbBody(t *testing.T) {
	type args struct {
		t   *descriptor.TypeDescriptor
		opt *readerOption
	}
	desc, err := getRespPbDesc()
	if err != nil {
		t.Error(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"readHTTPResponse",
			args{t: &descriptor.TypeDescriptor{
				Type: descriptor.STRUCT,
				Struct: &descriptor.StructDescriptor{
					FieldsByID: map[int32]*descriptor.FieldDescriptor{
						1: {
							ID:          1,
							Name:        "msg",
							Type:        &descriptor.TypeDescriptor{Type: descriptor.STRING},
							HTTPMapping: descriptor.DefaultNewMapping("msg"),
						},
					},
				},
			}, opt: &readerOption{pbDsc: desc}},
			map[int]interface{}{
				1: "hello world",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := thrift.NewBinaryWriter()
			w.WriteFieldBegin(thrift.TType(descriptor.STRING), 1)
			w.WriteString("hello world")
			w.WriteFieldStop()
			in := thrift.NewBinaryReader(remote.NewReaderBuffer(w.Bytes()))
			got, err := readHTTPResponse(context.Background(), in, tt.args.t, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("readHTTPResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			respGot := got.(*descriptor.HTTPResponse)
			if respGot.ContentType != descriptor.MIMEApplicationProtobuf {
				t.Errorf("expected content type: %v, got: %v", descriptor.MIMEApplicationProtobuf, respGot.ContentType)
			}
			body := respGot.GeneralBody.(proto.Message)
			for fieldID, expectedVal := range tt.want {
				val, err := body.TryGetFieldByNumber(fieldID)
				if err != nil {
					t.Errorf("get by fieldID [%v] err: %v", fieldID, err)
				}
				if val != expectedVal {
					t.Errorf("expected field value: %v, got: %v", expectedVal, val)
				}
			}
		})
	}
}

func getRespPbDesc() (proto.MessageDescriptor, error) {
	path := "main.proto"
	content := `
	package kitex.test.server;

	message BizResp {
		optional string msg = 1;
	}
	`

	var pbParser protoparse.Parser
	pbParser.Accessor = protoparse.FileContentsFromMap(map[string]string{path: content})
	fds, err := pbParser.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	return fds[0].GetMessageTypes()[0], nil
}
