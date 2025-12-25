# Architecture Documentation

## Table of Contents
1. [Folder Structure & Responsibilities](#folder-structure--responsibilities)
2. [Architecture Diagram](#architecture-diagram)
3. [Data Flow](#data-flow)
4. [Dependency Graph](#dependency-graph)

---

## Folder Structure & Responsibilities

### 📁 `cmd/` - Application Entry Point

**Purpose**: Contains the main application entry point and bootstrapping logic.

**Responsibilities**:
- Initializes application configuration
- Sets up blockchain connections
- Initializes SDKs (Account Abstraction SDK, DID Root SDK)
- Configures and starts the HTTP server
- Handles global error handling

**Key Files**:
- `cmd/server/main.go`: Main entry point that orchestrates application startup

**Dependencies**:
- Imports: `internal/config`, `internal/delivery/http/router`, `internal/infrastructure/sdk/*`, `pkg/logger`

---

### 📁 `internal/config/` - Configuration Management

**Purpose**: Centralized configuration loading, validation, and access.

**Responsibilities**:
- Loads environment variables from `.env` files
- Validates required configuration values
- Provides singleton access to configuration throughout the application
- Manages blockchain network configurations (HNS - Hybrid Name Service addresses)
- Manages application settings (port, log level, version, timeouts)

**Key Files**:
- `config.go`: Main configuration structure and loading logic
- `blockchain.go`: Blockchain-specific configuration
- `sync.go`: Thread-safe configuration access

**Dependencies**:
- Imports: `internal/infrastructure/sdk/*` (for HNS config types)
- No dependencies from other internal packages

---

### 📁 `internal/domain/` - Domain Layer (Business Entities)

**Purpose**: Core business domain models, entities, and pure business logic structures.

**Responsibilities**:
- Defines domain entities without external dependencies
- Contains pure business logic structures
- Represents core business concepts (Wallet, Transaction Hash, DID operations)
- No dependencies on infrastructure, frameworks, or external libraries (except for domain-specific types)

**Key Subfolders**:
- `accountabstraction/`:
  - `entity.go`: Transaction hash entities
  - `wallet_factory.go`: Wallet creation input structures
  - `wallet.go`: Wallet-related entities
  - `entrypoint.go`: Operation execution input structures
- `helper/`:
  - `entity.go`: Helper utility domain entities

**Key Principles**:
- Pure Go structures
- No business logic dependencies on external packages
- Entities represent business concepts

---

### 📁 `internal/repository/` - Repository Interface Layer

**Purpose**: Defines data access abstraction through interfaces (contracts).

**Responsibilities**:
- Declares repository interfaces that define data access contracts
- Enables dependency inversion principle
- Allows business logic (usecases) to depend on abstractions, not concrete implementations
- Supports testing through interface mocking

**Key Files**:
- `interface.go`: Defines all repository interfaces:
  - `AccountAbstractionRepository`: Wallet creation and operation execution
  - `HelperRepository`: String conversion and encoding utilities

**Key Principles**:
- Interfaces define **what** operations are available
- Concrete implementations live in `internal/infrastructure/`
- Enables swapping implementations without changing business logic

---

### 📁 `internal/usecase/` - Application Use Cases (Business Logic Orchestration)

**Purpose**: Orchestrates application-specific business operations.

**Responsibilities**:
- Implements application-specific use cases
- Coordinates between repositories to execute business logic
- Handles logging and error management
- Validates business rules
- Transforms domain entities as needed

**Key Subfolders**:
- `accountabstraction/`:
  - `usecase.go`: Use case structure and initialization
  - `create_account.go`: Wallet creation use case
  - `execute_operation.go`: Operation execution use case
- `helper/`:
  - `usecase.go`: Helper use case structure
  - `string_2_byte_32.go`: String to byte32 conversion use case
  - `encode_did_param_input.go`: DID parameter encoding use case

**Dependencies**:
- Imports: `internal/domain/*`, `internal/repository`, `pkg/logger`
- Depends on repository interfaces (not implementations)

---

### 📁 `internal/infrastructure/` - Infrastructure Layer (External Integrations)

**Purpose**: Implements technical details and external service integrations.

**Responsibilities**:
- Wraps external SDKs and libraries
- Implements repository interfaces
- Handles blockchain interactions
- Manages external service connections
- Provides infrastructure-level utilities

**Key Subfolders**:

#### `internal/infrastructure/sdk/` - SDK Wrappers
- **`accountabstraction/`**:
  - `sdk.go`: SDK initialization and singleton management
  - `wallet_factory.go`: Wallet factory operations (implements repository interface)
  - `wallet.go`: Wallet operations (implements repository interface)
  - **Implements**: `repository.AccountAbstractionRepository`
- **`didroot/`**:
  - `sdk.go`: DID Root SDK initialization and operations

#### `internal/infrastructure/lib/` - External Libraries
- Contains local copies/vendored external libraries:
  - `core-account-abstraction-sdk/`: Account abstraction SDK
  - `core-general-sdk/`: General blockchain utilities
  - `did-root-sdk/`: DID Root SDK

#### `internal/infrastructure/helper/` - Infrastructure Helpers
- `helper.go`: Infrastructure-level helper implementations
- **Implements**: `repository.HelperRepository`

**Dependencies**:
- Imports: `internal/domain/*`, `internal/repository`, `internal/config`
- Uses external SDKs from `lib/` folder

---

### 📁 `internal/delivery/` - Delivery Layer (API Layer)

**Purpose**: Exposes application functionality via different protocols (HTTP, Events).

**Responsibilities**:
- Receives external requests
- Validates input
- Routes requests to appropriate handlers
- Formats responses
- Handles HTTP-specific concerns (status codes, headers)
- Processes events from external sources

**Key Subfolders**:

#### `internal/delivery/http/` - HTTP API
- **`handler/`** - Request Handlers (Controller Layer):
  - `accountabstraction/`: Account abstraction HTTP handlers
    - `handler.go`: Handler structure
    - `wallet_factory.go`: Wallet creation handler
    - `wallet.go`: Wallet operation handler
    - `types.go`: Response types
  - `helper/`: Helper endpoint handlers
    - `handler.go`: Helper handler implementations
    - `types.go`: Helper response types
  
  **Responsibilities**:
  - Parse HTTP requests
  - Validate input data
  - Call use cases
  - Format HTTP responses
  - Handle HTTP errors

- **`router/`** - Route Definitions:
  - `router.go`: Main router setup and middleware configuration
  - `accountabstraction/`: Account abstraction routes
    - `router.go`: Route group setup
    - `wallet_factory.go`: Wallet factory routes
    - `wallet.go`: Wallet routes
  - `helper/`: Helper routes
    - `router.go`: Helper route setup
  - `didroot/`: DID Root routes
  
  **Responsibilities**:
  - Define URL patterns and HTTP methods
  - Wire handlers to routes
  - Configure middleware (CORS, logging, recovery)
  - Dependency injection (creates use cases, handlers)

#### `internal/delivery/event/` - Event-Driven Delivery
- `accountabstraction/wallet_factory.go`: Event handlers for blockchain events (wallet deployment events)

**Dependencies**:
- Imports: `internal/usecase/*`, `internal/domain/*`, `internal/config`, `pkg/response`
- Uses Fiber framework for HTTP handling

---

### 📁 `pkg/` - Public Packages (Reusable Components)

**Purpose**: Contains reusable packages that can be imported by other projects.

**Key Subfolders**:
- **`logger/`**:
  - `logger.go`: Logging utilities (InfoLogger, ErrorLogger)
  - Provides standardized logging interface
  
- **`response/`**:
  - `response.go`: HTTP response formatting utilities
  - Provides `Success()` and `Error()` helper functions
  - Standardizes API response format with metadata (timestamp, version)

**Key Principles**:
- Public API that can be used by external projects
- Should not depend on `internal/` packages
- Provides utility functions for common operations

---

### 📁 `utils/` - Application Constants

**Purpose**: Shared constants and enums across the application.

**Key Files**:
- `const.go`: Defines operation type constants:
  - `TypeGeneralExecute`, `TypeCreateDID`, `TypeUpdateDID`, etc.
  - Used for DID and account abstraction operations

**Usage**:
- Referenced by domain logic and infrastructure implementations
- Provides centralized constant definitions

---

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────────────┐
│                          CLIENT / EXTERNAL                               │
└────────────────────────────┬────────────────────────────────────────────┘
                             │ HTTP Requests
                             ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                      DELIVERY LAYER (HTTP API)                           │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │  internal/delivery/http/router/                                 │   │
│  │  • Route definitions                                            │   │
│  │  • Middleware (CORS, Logging, Recovery)                         │   │
│  │  • Dependency Injection                                         │   │
│  └────────────────────┬────────────────────────────────────────────┘   │
│                       │ Wires to                                        │
│  ┌────────────────────▼────────────────────────────────────────────┐   │
│  │  internal/delivery/http/handler/                                │   │
│  │  • Parse HTTP requests                                          │   │
│  │  • Validate input                                               │   │
│  │  • Call use cases                                               │   │
│  │  • Format responses                                             │   │
│  └────────────────────┬────────────────────────────────────────────┘   │
└───────────────────────┼─────────────────────────────────────────────────┘
                        │ Uses
                        ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                     APPLICATION LAYER (USE CASES)                        │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │  internal/usecase/                                              │   │
│  │  • Orchestrates business operations                             │   │
│  │  • Coordinates repositories                                     │   │
│  │  • Handles logging & errors                                     │   │
│  │  • Business rule validation                                     │   │
│  └────────────────────┬────────────────────────────────────────────┘   │
└───────────────────────┼─────────────────────────────────────────────────┘
                        │ Depends on (interface)
                        ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                       DOMAIN LAYER                                       │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │  internal/domain/                                               │   │
│  │  • Business entities                                            │   │
│  │  • Pure data structures                                         │   │
│  │  • No external dependencies                                     │   │
│  └─────────────────────────────────────────────────────────────────┘   │
│                                                                         │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │  internal/repository/                                           │   │
│  │  • Interface definitions (contracts)                            │   │
│  │  • Dependency inversion                                         │   │
│  └────────────────────┬────────────────────────────────────────────┘   │
└───────────────────────┼─────────────────────────────────────────────────┘
                        │ Implements
                        ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                    INFRASTRUCTURE LAYER                                  │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │  internal/infrastructure/sdk/                                   │   │
│  │  • SDK wrappers                                                 │   │
│  │  • Implements repository interfaces                             │   │
│  │  • Blockchain interactions                                      │   │
│  └────────────────────┬────────────────────────────────────────────┘   │
│                       │ Uses                                            │
│  ┌────────────────────▼────────────────────────────────────────────┐   │
│  │  internal/infrastructure/lib/                                   │   │
│  │  • External SDKs (vendored)                                     │   │
│  │  • Account Abstraction SDK                                      │   │
│  │  • DID Root SDK                                                 │   │
│  │  • Blockchain utilities                                         │   │
│  └────────────────────┬────────────────────────────────────────────┘   │
└───────────────────────┼─────────────────────────────────────────────────┘
                        │ Connects to
                        ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                    EXTERNAL SERVICES                                     │
│  • Blockchain Network                                                   │
│  • Smart Contracts                                                      │
│  • External APIs                                                        │
└─────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────┐
│                    CROSS-CUTTING CONCERNS                                │
│  ┌────────────────────┐  ┌────────────────────┐  ┌──────────────────┐ │
│  │  pkg/logger/       │  │  pkg/response/     │  │  internal/config/│ │
│  │  • Logging         │  │  • HTTP responses  │  │  • Configuration │ │
│  └────────────────────┘  └────────────────────┘  └──────────────────┘ │
│                                                                         │
│  ┌──────────────────────────────────────────────────────────────────┐ │
│  │  utils/                                                           │ │
│  │  • Constants                                                      │ │
│  └──────────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## Data Flow

### Example: Creating a Wallet

```
1. HTTP Request
   Client → POST /api/v1/account-abstraction/create
   Body: { "privKey": "...", "input": {...} }

2. Router Layer (internal/delivery/http/router/accountabstraction/)
   • Matches route pattern
   • Applies middleware (CORS, logging, recovery)
   • Gets SDK instance (which implements repository interface)
   • Creates use case with repository
   • Creates handler with use case
   • Routes request to handler

3. Handler Layer (internal/delivery/http/handler/accountabstraction/)
   • Parses HTTP request body into domain entity (CreateWalletInput)
   • Validates request format
   • Calls use case: uc.CreateWallet(ctx, input)

4. Use Case Layer (internal/usecase/accountabstraction/)
   • Receives domain entity (CreateWalletInput)
   • Calls repository: repo.CreateWallet(ctx, input)
   • Handles errors and logging
   • Returns domain entity (TxHash)

5. Repository Implementation (internal/infrastructure/sdk/accountabstraction/)
   • SDK implements repository interface
   • Creates wallet using blockchain library
   • Calls WalletFactory.DeployWallet()
   • Returns transaction hash(es)

6. External Blockchain
   • Smart contract execution
   • Transaction submission
   • Event emission

7. Response Flow (back up the layers)
   • Handler checks transaction status
   • Decodes blockchain events
   • Formats response using pkg/response
   • Returns JSON response to client
```

### Example: Helper Operation (String to Byte32)

```
1. HTTP Request
   Client → POST /api/v1/helper/string-2-byte32
   Body: { "input": "..." }

2. Router Layer (internal/delivery/http/router/helper/)
   • Creates helper repository implementation
   • Creates use case with repository
   • Creates handler with use case
   • Routes to handler

3. Handler Layer (internal/delivery/http/handler/helper/)
   • Parses request
   • Calls use case: uc.StringToByte32(input)

4. Use Case Layer (internal/usecase/helper/)
   • Calls repository: repo.StringToByte32(input)

5. Repository Implementation (internal/infrastructure/helper/)
   • Performs conversion using blockchain utilities
   • Returns result

6. Response
   • Formatted and returned to client
```

---

## Dependency Graph

```
cmd/server/main.go
  ├─→ internal/config
  │     └─→ internal/infrastructure/sdk/* (for config types)
  ├─→ internal/delivery/http/router
  │     ├─→ internal/usecase/*
  │     │     ├─→ internal/repository (interface)
  │     │     ├─→ internal/domain/*
  │     │     └─→ pkg/logger
  │     ├─→ internal/delivery/http/handler/*
  │     │     ├─→ internal/usecase/*
  │     │     ├─→ internal/domain/*
  │     │     ├─→ pkg/response
  │     │     └─→ internal/config
  │     └─→ internal/config
  └─→ internal/infrastructure/sdk/*
        ├─→ internal/domain/*
        └─→ external SDKs (from lib/)

internal/infrastructure/helper/
  ├─→ internal/repository (implements interface)
  ├─→ internal/domain/*
  ├─→ internal/config
  └─→ utils/

internal/infrastructure/sdk/accountabstraction/
  ├─→ internal/repository (implements AccountAbstractionRepository)
  ├─→ internal/domain/*
  └─→ external SDKs
```

### Dependency Rules

1. **Outward Dependencies Only**: Each layer can only depend on layers closer to the center (domain)
2. **Domain is Pure**: `internal/domain/` has no dependencies on other internal packages
3. **Repository Pattern**: Use cases depend on repository interfaces, implementations are in infrastructure
4. **No Circular Dependencies**: Architecture enforces unidirectional flow
5. **Public Packages**: `pkg/` should not depend on `internal/` packages

---

## Key Architectural Principles

### 1. Clean Architecture / Hexagonal Architecture
- Clear separation of concerns
- Business logic independent of frameworks
- Testability through dependency inversion

### 2. Dependency Inversion
- High-level modules (use cases) don't depend on low-level modules (infrastructure)
- Both depend on abstractions (repository interfaces)

### 3. Single Responsibility
- Each package has a clear, single purpose
- Folders organize related functionality

### 4. Dependency Injection
- Dependencies are injected through constructors
- Makes code testable and flexible

### 5. Interface Segregation
- Repository interfaces define clear contracts
- Implementations can be swapped without changing business logic

---

## Module Organization

### Feature-Based Organization
Each feature (e.g., `accountabstraction`, `helper`) has code organized across layers:
- `domain/accountabstraction/` - Domain entities
- `usecase/accountabstraction/` - Business logic
- `infrastructure/sdk/accountabstraction/` - SDK implementations
- `delivery/http/handler/accountabstraction/` - HTTP handlers
- `delivery/http/router/accountabstraction/` - HTTP routes

This allows features to be developed and maintained independently while following consistent architectural patterns.

