include "include.thrift"

namespace go example

struct ExampleRequest {
    1: i32 FieldInt32
    2: optional bool FieldBool
    3: list<string> FieldListString
    4: map<string,i32> FieldMapStringInt32
    5: InnerAStruct FieldInnerAStruct
}

struct InnerAStruct {
    1: i32 FieldInt32
    2: string FieldString
    3: list<InnerBStruct> FieldListInnerBStruct
    4: map<string,InnerBStruct> FieldMapInnerBStruct
}

struct InnerBStruct {
    1: i32 FieldInt32
    2: string FieldString
    3: map<string,include.IncludeStruct> FieldMapStringIncludeStruct
}

struct ExampleResponse {
    1: i32 FieldInt32
    2: bool FieldBool
    3: list<string> FieldListString
    4: map<string,i32> FieldMapStringInt32
    5: InnerAStruct FieldInnerAStruct
}

service ExampleService {
    ExampleResponse Example(1: ExampleRequest req)
}