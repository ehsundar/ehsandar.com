---
title: "Go code generation approaches"
date: 2023-07-14T13:37:09+03:30
draft: false

ShowToc: true
TocOpen: true
---

# Code Generators

It's a quite common pattern in Go projects to use code generators.
Code generators or in short term, codegens are used to provide type safety and eliminate runtime unnecessary processing.
Without codegens, we need to use [reflect](https://pkg.go.dev/reflect) package at runtime.

## Our Problem

Imagine we have a `json` serialized byte stream received from a user or read from a file.

```json
{
    "first_name": "Amir",
    "last_name": "Ehsandar"
}
```

How can we **Unmarshal** this arbitrary data structure into a `struct`?
In this article, we discuss two different approaches to **unmarshal** this serialized binary into
corresponding `struct`.
We also compare three different stages for generating codes, on developer's machine, on CI or every time before build
process begins.

# Unmarshalling

We can unmarshal our json into a struct using reflection.
The reflection API provides the developer with a few function to iterate over a `struct` exported fields and inspect
their
data types.
We also can use codegens to generate codes for unmarshalling our data into a struct.
There are a few open source modules available on GitHub which we'll mention later on.

## Unmarshalling JSON using the reflection API

This is how builtin json package works.
Here's an example usage of `reflect` to iterate through an arbitrary struct:

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type anotherStructure struct {
		Field string
	}

	type sampleStructure struct {
		Name   string
		Age    int
		Height float32

		Another anotherStructure
	}

	sType := reflect.TypeOf(sampleStructure{})

	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)

		fmt.Printf("%s -> %s\n", field.Name, field.Type.Name())
	}
}
```

Output:

```text
Name -> string
Age -> int
Height -> float32
Another -> anotherStructure
```

## Marshalling struct to JSON using generated code

To eliminate this overhead at runtime, there are a number of third-party module to preprocess to be marshaled/unmarshalled structures.
We use [easyjson](https://github.com/mailru/easyjson) as an example here.
Code generate for nearly the same struct mentioned above looks like this:

```go
func easyjson89aae3efEncodeGithubComEhsundarBlahEasyjsonsample(out *jwriter.Writer, in sampleStructure) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	{
		const prefix string = ",\"height\":"
		out.RawString(prefix)
		out.Float32(float32(in.Height))
	}
	out.RawByte('}')
}
```

This generated code saves a lot of CPU time at runtime,
instead takes a while to generate this code and needs a few extra steps to develop go code.

You can see easyjson's benchmarks [here](https://github.com/mailru/easyjson#unmarshaling)
compared to a few other popular tools.

| lib      | json size | MB/s | allocs/op | B/op  |
|:---------|:----------|-----:|----------:|------:|
| standard | regular   | 22   | 218       | 10229 |
| standard | small     | 9.7  | 14        | 720   |
|          |           |      |           |       |
| easyjson | regular   | 125  | 128       | 9794  |
| easyjson | small     | 67   | 3         | 128   |
|          |           |      |           |       |
| ffjson   | regular   | 66   | 141       | 9985  |
| ffjson   | small     | 17.6 | 10        | 488   |
|          |           |      |           |       |
| codec    | regular   | 55   | 434       | 19299 |
| codec    | small     | 29   | 7         | 336   |
|          |           |      |           |       |
| ujson    | regular   | 103  | N/A       | N/A   |

# Code Generation Stage

## Generation on the Developer's Machine

TODO

## Generation on CI and Commit on Another Repository

## Generation before Build Process

# Conclusion

| Stage | Suitable for                                                                                                                                      |
|-------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| Local | - small projects <br>- generated codes only used by the same module                                                                               |
| CI    | - another repository is needed for the generated code to be committed at <br>- add extra complexity and generated code are available with a delay |
| Build | - clone repository can not be built without generating codes <br>- there is no change history for generated codes in git                          |
