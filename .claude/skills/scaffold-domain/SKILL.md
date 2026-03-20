---
name: scaffold-domain
description: >
  Scaffolds a new domain end-to-end in this Go boilerplate project: HTTP CRUD controller
  with swag annotations, protobuf message and service definitions, buf-generated gRPC
  handler, versioned HTTP routes, and full Wire DI wiring. Use this skill whenever the
  user says "scaffold a new domain", "add a new domain", "create a new resource",
  "add [noun] API", or anything that implies adding a new vertical slice (controller +
  proto + gRPC + routes + DI) to this project. Also trigger when they mention a new
  model and want HTTP and gRPC endpoints for it.
---

Scaffold a new domain end-to-end (HTTP controller, proto, gRPC handler, routes, Wire DI).

Follow these steps in order, pausing to ask the user where indicated.

## Step 0 — Pre-flight check

Before asking for any inputs, ask this yes/no question:

> "Have you already created the domain model struct in `internal/app/model/`?"

- If **no** — stop and tell the user to create the model file first, then come back. Do not continue.
- If **yes** — proceed to gather inputs.

## Step 1 — Gather inputs


Ask the user for these three things before doing anything else:
1. **Domain name** (e.g. `bill`, `customer`) → `{{domain}}` (lowercase), `{{Domain}}` (PascalCase)
2. **API version** (e.g. `1`) → `{{version}}`
3. **Domain model** — list files in `internal/app/model/` and ask which one is the model for this domain → `{{Model}}`

Do not proceed until all three are confirmed.

---

## Step 2 — Domain package

Create the directory `internal/app/{{domain}}/` with three files modelled after `internal/app/todo/`.

**`repository.go`** — read the template from `assets/repository.go.tmpl` and substitute `{{domain}}`, `{{Domain}}`, and `{{Model}}`.

**`{{domain}}_service.go`** — read the template from `assets/service.go.tmpl` and substitute `{{domain}}`, `{{Domain}}`, and `{{Model}}`.

**`{{domain}}_service_test.go`** — empty test file:
```go
package {{domain}}
```

After creating the package, also do the following so the service is wired as a provider:

- In **`internal/app/api/http/handlers/{{domain}}_controller.go`** (Step 3), accept `{{domain}}.{{Domain}}Service` in the constructor — not a pointer, it's a value type.
- In **`internal/app/api/grpc/{{domain}}.go`** (Step 5), accept `{{domain}}.{{Domain}}Service` in `NewGRPC{{Domain}}Handler` — same value type.

---

## Step 3 — HTTP Controller

Create `internal/app/api/http/handlers/{{domain}}_controller.go`.

Follow the same structure as `internal/app/api/http/handlers/todo_controller.go`:
- `{{Domain}}Controller` struct with `service` (domain service type) and `defaultUser string`
- `New{{Domain}}Controller` constructor
- Five handlers: `CreateItem`, `FindAll`, `GetItem`, `UpdateItem`, `DeleteItem`
- Every service call passes `ctx` (`*gin.Context`) as the first argument
- Swag annotations on every handler:
  - `@Router` → `/{{domain}}` and `/{{domain}}/{itemId}`
  - `@Success` → `{object} model.{{Model}}` or `{array} model.{{Model}}`
  - `@Failure` → `400` and `500` where applicable
  - `@Param` body → fully-qualified DTO (e.g. `dto.Create{{Domain}}Request`)

---

## Step 4 — Proto definitions

Create `proto/{{domain}}/v{{version}}/` with two files.

**`{{domain}}.proto`** — messages mirroring `model.{{Model}}`:
- `{{Domain}}` message matching the Go struct fields (`google.protobuf.Timestamp` for time, `string` for UUID)
- `Create{{Domain}}Request` — omit `id`, `user_id`, `created_at`, `updated_at`
- `Update{{Domain}}Request` — all fields optional; include `item_id`
- `List{{Domain}}sRequest` — `user_id` field
- `List{{Domain}}sResponse` — `repeated {{Domain}} items`
- `Get{{Domain}}Request`, `Delete{{Domain}}Request` — `item_id` field
- `Delete{{Domain}}Response` — empty
- `{{Domain}}Response` — single `{{Domain}} item`

**`{{domain}}_service.proto`** — service:
- `package {{domain}}.v{{version}};`
- `option go_package = "v{{version}}/{{domain}};{{domain}}";`
- Imports: `google/api/annotations.proto` and `proto/{{domain}}/v{{version}}/{{domain}}.proto`
- `{{Domain}}Service` with five RPCs (Create, List, Get, Update, Delete)
- `option (google.api.http)` on each RPC → `/v{{version}}/{{domain}}` and `/v{{version}}/{{domain}}/{item_id}`
- Doc comments on the service and every RPC

---

## Step 5 — Generate protobuf Go code

```bash
buf generate --path proto/{{domain}}
```

Verify files appear under `pkg/proto/{{domain}}/v{{version}}/` before continuing.

---

## Step 6 — gRPC handler

Create `internal/app/api/grpc/{{domain}}.go`:
- `GRPC{{Domain}}Handler` embedding `Unimplemented{{Domain}}ServiceServer` from the generated package
- `{{domain}}Service` field of the domain service type
- `NewGRPC{{Domain}}Handler` constructor
- Implement all five RPC methods with `context.Context` as the first param

Update `internal/app/api/grpc/server.go`:
- Add `*GRPC{{Domain}}Handler` parameter to `NewGrpcServer`
- Register: `{{domain}}grpc.Register{{Domain}}ServiceServer(s, {{domain}}Handler)`

---

## Step 7 — HTTP routes

Check whether `internal/app/api/http/routes/v{{version}}/` exists.

**Directory does not exist** — create two files:
- `index.go` — read from `assets/routes_index.go.tmpl`, substitute `{{domain}}`, `{{Domain}}`, `{{version}}`
- `{{domain}}.go` — read from `assets/routes.go.tmpl`, substitute `{{domain}}`, `{{Domain}}`, `{{version}}`

**Directory already exists** — add only:
- `{{domain}}.go` from `assets/routes.go.tmpl`
- Add `*handlers.{{Domain}}Controller` parameter to the existing `RegisterRouterApi` and call `set{{Domain}}Routes`

---

## Step 8 — Wire up HTTP router

Update `internal/app/api/http/index.go` (`NewApiRouter`):
- Add `*handlers.{{Domain}}Controller` as a parameter
- Pass it to `v{{version}}.RegisterRouterApi`

---

## Step 9 — Wire up inject files

**`internal/app/inject.go`** — add the service provider to `ApiProviderSet`:
```go
var ApiProviderSet = wire.NewSet(
    repositoryBindings,
    todo.NewTodoService,
    {{domain}}.New{{Domain}}Service, // add this
    api.ProviderSet,
)
```

**`internal/app/api/inject.go`** — add the HTTP controller and gRPC handler providers:
```go
var grpcHandlerProvider = wire.NewSet(
    grpc.NewGRPCTodoHandler,
    grpc.NewGRPC{{Domain}}Handler, // add this
)

var httpHandlerProviderSet = wire.NewSet(
    httpHandlers.NewTodoController,
    httpHandlers.New{{Domain}}Controller, // add this
)
```
