# The Things Network API

## Languages

| **Language** | **Status**  | **Maintainer** |
| ------------ | ----------- | -------------- |
| C            | structs only (no gRPC) | **core team** |
| C#           | nothing yet; help wanted | you? |
| C++          | nothing yet; help wanted | you? |
| Go           | [sdk available](https://github.com/TheThingsNetwork/go-app-sdk) | **core team** |
| Java         | [sdk available](https://github.com/TheThingsNetwork/java-app-sdk) | [@cambierr](https://github.com/cambierr) |
| Javascript   | [sdk available](https://github.com/TheThingsNetwork/node-app-sdk) | **core team** |
| Objective-C  | nothing yet; help wanted | you? |
| PHP          | files generated; not tested; help wanted | you? |
| Python       | files generated; not tested; help wanted | you? |
| Ruby         | files generated; not tested; help wanted | you? |
| Swift        | files generated; not tested; help wanted | you? |

## Updating the generated files

⚠️ Warning: this will be a complicated process. We'll do our best to improve the documentation.

We're going to need to install quite some stuff, as these protos will be compiled to different languages.

- Clone this repository to `$HOME/src/github.com/TheThingsNetwork/api`
- Set your GOPATH to your HOME dir: `export GOPATH="$HOME"`
- Install [protoc-3.3.0-YOURPLATFORM](https://github.com/google/protobuf/releases/tag/v3.3.0) to `/usr/local`

For C++, C#, Objective-C, PHP, Python, Ruby:

Install [gRPC](https://github.com/grpc/grpc/blob/master/INSTALL.md)

```
git clone https://github.com/grpc/grpc $GOPATH/src/github.com/grpc/grpc
cd $GOPATH/src/github.com/grpc/grpc
make
make install
```

For Go:

- Install Go
- Install the Go generators:

```
go get -u github.com/TheThingsNetwork/api/utils/protoc-gen-gogottn
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

For Node:

- Install Node
- Install the Node generator:

```
npm install grpc-tools
```

For Java:

- Install Java and Gradle
- Install the Java generator:

```
git clone https://github.com/grpc/grpc-java $GOPATH/src/github.com/grpc/grpc-java
cd $GOPATH/src/github.com/grpc/grpc-java/compiler
../gradlew java_pluginExecutable
```

For Swift:

- Install and set up XCode and [Swift](https://swift.org/download/#installation)
- Install the Swift generator:

```
git clone https://github.com/grpc/grpc-swift $GOPATH/src/github.com/grpc/grpc-swift
cd $GOPATH/src/github.com/grpc/grpc-swift/Plugin
make
```

For C:

- Install [protobuf-c](https://github.com/protobuf-c/protobuf-c)

```
git clone https://github.com/protobuf-c/protobuf-c $GOPATH/src/github.com/protobuf-c/protobuf-c
cd $GOPATH/src/github.com/protobuf-c/protobuf-c
./autogen.sh && ./configure && make
```

And finally set the correct PATH:

```
export PATH="$PATH:$GOPATH/bin:$GOPATH/src/github.com/grpc/grpc-swift/Plugin:$GOPATH/src/github.com/grpc/grpc-java/compiler/build/exe/java_plugin:$GOPATH/src/github.com/grpc/grpc/bins/opt:$GOPATH/src/github.com/TheThingsNetwork/protobuf-c/protobuf-c/protoc-c"
```

Now you should be able to update the generated files:

```
make all
```

If you only want to update the files for a specific language, you can use the `protos.go`, `protos.java`, `protos.js` and `protos.swift` targets.
