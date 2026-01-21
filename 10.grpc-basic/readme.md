# gRPC with Go – Important Commands & Their Purpose

This README is a **practical cheat sheet** for gRPC + Protobuf in Go.
It explains **what command to run, when to run it, and why it exists**.

---

## 1️⃣ Go Module Initialization

```bash
go mod init grpc-basic
```

### Purpose
- Initializes a Go module
- Required for imports, versioning, and dependency management

### When to use
- Once, when starting a new Go project

---

## 2️⃣ Check Go Environment

```bash
go version
go env GOPATH
```

### Purpose
- Verify Go installation
- Ensure `GOPATH/bin` exists (needed for protoc plugins)

---

## 3️⃣ Install Protobuf Compiler (`protoc`)

### macOS
```bash
brew install protobuf
```

### Ubuntu / Debian
```bash
sudo apt install protobuf-compiler
```

### Verify
```bash
protoc --version
```

### Purpose
- `protoc` converts `.proto` → language-specific code

---

## 4️⃣ Install Go Protobuf Plugin

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### Purpose
- Generates Go structs for **protobuf messages**

### Generates
- `*.pb.go`

---

## 5️⃣ Install Go gRPC Plugin

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Purpose
- Generates Go code for **gRPC services**

### Generates
- `*_grpc.pb.go`

---

## 6️⃣ Update PATH (VERY IMPORTANT)

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Purpose
- Allows `protoc` to find Go plugins

### Symptom if missing
```text
protoc-gen-go: program not found
```

---

## 7️⃣ Recommended Project Structure

```text
grpc-basic/
├── idl/                # .proto files only
│   ├── user/v1/user.proto
│   └── post/v1/post.proto
│
├── proto/              # generated Go code
│   ├── user/v1/*.pb.go
│   └── post/v1/*.pb.go
│
├── server/
├── client/
└── go.mod
```

---

## 8️⃣ Generate Go Code from Proto (MOST IMPORTANT)

```bash
protoc \
  --proto_path=idl \
  --go_out=. \
  --go-grpc_out=. \
  idl/user/v1/user.proto \
  idl/post/v1/post.proto
```

rm -rf proto grpc-project
mkdir proto

protoc \
  --proto_path=idl \
  --go_out=proto \
  --go_opt=paths=source_relative \
  --go-grpc_out=proto \
  --go-grpc_opt=paths=source_relative \
  idl/user/v1/user.proto \
  idl/post/v1/post.proto


  protoc \
  --proto_path=idl \
  --go_out=proto \
  --go_opt=paths=source_relative \
  --go-grpc_out=proto \
  --go-grpc_opt=paths=source_relative \
  $(find idl -name "*.proto")


### What Each Flag Means

| Flag | Purpose |
|-----|--------|
| `--proto_path` | Where `.proto` imports are resolved |
| `--go_out` | Where message code is written |
| `--go-grpc_out` | Where service code is written |

---

## 9️⃣ Understanding Output vs go_package

### `.proto`
```proto
option go_package = "grpc-basic/proto/user/v1;userpb";
```

### Meaning
- Controls **Go import path**
- Controls **package name inside Go files**

### Important Rule
```
go_package → Go import path
--go_out    → filesystem location
```

---

## 🔟 Run gRPC Server

```bash
go run server/main.go
```

### Purpose
- Starts gRPC server
- Opens TCP port (e.g. `:50051`)

---

## 1️⃣1️⃣ Run gRPC Client

```bash
go run client/main.go
```

### Purpose
- Calls gRPC server
- Tests RPC functionality

---

## 1️⃣2️⃣ Re-generate Code (After Proto Change)

```bash
protoc --proto_path=idl --go_out=. --go-grpc_out=. idl/**/*.proto
```

### When to run
- Any time `.proto` changes

---

## 1️⃣3️⃣ Clean Generated Code (Optional)

```bash
rm -rf proto/
```

### Purpose
- Force fresh generation
- Useful when refactoring `go_package`

---

## 1️⃣4️⃣ Useful Debug Commands

```bash
go list ./...
go mod tidy
```

### Purpose
- Validate module imports
- Clean unused dependencies

---

## 🧠 Mental Model Summary

```
.proto        → API contract
protoc        → code generator
go_package    → Go import path
--go_out      → file location
gRPC server   → implements interface
gRPC client   → calls interface
```

---

## ⭐ Recommended VS Code Extensions

- **Buf** (best protobuf experience)
- Protobuf Support (basic syntax highlighting)

---

## 🏁 Final Advice

- Never edit `.pb.go` manually
- Always version your proto (`v1`, `v2`)
- Keep IDL and generated code separate

---


## auto .pb.go generator

✅ Why Buf is better

Auto-discovers all .proto files

No path duplication ever

Enforces best practices

One command generation

Industry standard (Google, Uber, Netflix)

✅ Step 1 — Install Buf
brew install bufbuild/buf/buf


(or download binary)

✅ Step 2 — buf.yaml (repo root/idl)
version: v1
name: grpc-project

✅ Step 3 — buf.gen.yaml at root
version: v1
plugins:
  - plugin: go
    out: proto
    opt:
      - paths=source_relative

  - plugin: go-grpc
    out: proto
    opt:
      - paths=source_relative

✅ Step 4 — Run generation
buf generate idl



##Bazel

1️⃣ What Bazel gives you (why people use it)

With Bazel:

✅ Automatically discovers .proto files

✅ Generates .pb.go and _grpc.pb.go

✅ Handles imports correctly (no duplicate dirs, ever)

✅ Caches builds (VERY fast after first run)

✅ Works the same in CI, locally, anywhere

Think of Bazel as:

Make + Go modules + protoc + Buf + cache + CI discipline

2️⃣ How Bazel does proto → Go

Bazel does NOT run protoc directly.

Instead, it uses:

rules_proto

rules_go

rules_go_grpc

These rules:

Understand proto dependencies

Generate Go code as build artifacts

Avoid filesystem pollution (no random folders)


4️⃣ Step-by-step Bazel setup
🔹 Step 1: Install Bazel
brew install bazel

OR for newer version
`
brew unlink bazel
brew install bazelisk
brew link --force bazelisk
`

🔹 Step 2: create MODULE.bazel file (repo root)

🔹 Step 3: idl/user/v1/BUILD.bazel

🔹 Step 4: idl/post/v1/BUILD.bazel
🔹 Step 4.5: idl/BUILD.bazel

🔹 Step 5: Build everything
bazel clean --expunge
bazel build //idl:all_grpc


5️⃣ Where are the .pb.go files?

They live in Bazel’s output tree:
bazel-bin/




