/*
    Copyright 2021 Rabia Research Team and Developers

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/
package serialization

import (
	"fmt"
	"os"
	"testing"
)

/*
	KeyNum          1    2	    3      4     ...
	Message Size(B)	8    16	    24     32    ...
*/

var KeyNum = 7 // KeyNum  = number of int64 data types (8 bytes each)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Test_Generate_GoBinMsg(t *testing.T) {
	f, err := os.Create("./gobin_msg.go")
	check(err)
	_, _ = f.WriteString("package serialization\n")
	_, _ = f.WriteString("\n")
	_, _ = f.WriteString("// ~/go/src/gobin-codegen/bin/bi ~/go/src/rc3/serialization/gobin_msg.go > ~/go/src/rc3/serialization/gobin_gen.go \n")
	_, _ = f.WriteString("\n")
	_, _ = f.WriteString("type GoBinMsg struct {\n")
	for i := 0; i < KeyNum; i++ {
		_, _ = f.WriteString(fmt.Sprintf("\tKey%d int64\n", i))
	}
	_, _ = f.WriteString("}")
	_ = f.Close()
}

func Test_Generate_ProtoMsg(t *testing.T) {
	f, err := os.Create("./proto_msg.proto")
	check(err)
	_, _ = f.WriteString("syntax = \"proto3\";\n\n" +
		"/*\n" +
		"  protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/google/protobuf --go_out=. ./proto_msg.proto\n" +
		"*/\n\n" +
		"package serialization;\n" +
		"option go_package = \"../serialization\";\n\n")

	_, _ = f.WriteString("message ProtoMsg {\n")
	for i := 0; i < KeyNum; i++ {
		_, _ = f.WriteString(fmt.Sprintf("\tint64 Key%d = %d; \n", i, i))
	}
	_, _ = f.WriteString("}")

	_ = f.Close()
}
