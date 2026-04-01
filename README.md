# ⚠️ DEPRECATED

> [!WARNING]
> **This repository is deprecated.**
> We are now moving to a new architecture using **Redis Streams** for event handling and processing.
> Please refer to the new implementation for the latest updates.

# Architecture Documentation

## Overview

This is a DID (Decentralized Identifier) Backend API built with Go using Clean Architecture principles. The API provides endpoints for managing:

- **Account Abstraction**: Wallet creation and operation execution
- **DID Root**: DID resolution, key/claim management, and data operations
- **DID Alias**: Alias resolution and domain/subdomain registration
- **DID Verifiable Credentials (DID VC)**: Credential metadata, validation, and token management
- **Helper**: Utility functions for encoding parameters and conversions

The API is built using the Fiber framework and follows a layered architecture pattern with clear separation of concerns.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Folder Structure & Responsibilities](#folder-structure--responsibilities)
3. [Architecture Diagram](#architecture-diagram)
4. [Data Flow](#data-flow)
5. [Dependency Graph](#dependency-graph)
6. [API Modules](#api-modules)

---

## Getting Started

### Prerequisites

- Go 1.25.5 or higher
- Git with SSH key configured (for submodule access) or use HTTPS alternative

### Cloning the Repository

This project uses Git submodules for managing external SDK dependencies. All submodules are located in `internal/infrastructure/lib/` and use the `go-be-v1` branch.

#### Option 1: Clone with SSH (Administrator Access)

If you have SSH access configured (the SSH key is private and known only by administrators):

```bash
# Clone the main repository
git clone git@ssh.dev.azure.com:v3/dattabot/Blockchain%202024/go-be.git

# Navigate to the project directory
cd go-be

# Initialize all submodules
git submodule update --init --recursive

# For each submodule, checkout the go-be-v1 branch
# This ensures all submodules are on the correct branch
git submodule foreach 'git checkout go-be-v1 || git fetch origin go-be-v1:go-be-v1 && git checkout go-be-v1'

# Update all submodules to the latest commit on go-be-v1 branch
git submodule update --remote
```

#### Option 2: Clone with HTTPS (Alternative Access Method)

If you don't have SSH access, you can use HTTPS by modifying the submodule URLs. **Note:** You may need to authenticate with Azure DevOps using a Personal Access Token (PAT) if the repositories are private.

```bash
# Clone the main repository
git clone https://dev.azure.com/dattabot/Blockchain%202024/_git/go-be

# Navigate to the project directory
cd go-be

# Update submodule URLs to use HTTPS
# Replace USERNAME with your Azure DevOps username, or use PAT in the URL format:
# https://USERNAME@dev.azure.com/dattabot/Blockchain%202024/_git/REPO_NAME
git config submodule."internal/infrastructure/lib/core-general-sdk".url https://dev.azure.com/dattabot/Blockchain%202024/_git/core-general-sdk
git config submodule."internal/infrastructure/lib/core-account-abstraction-sdk".url https://dev.azure.com/dattabot/Blockchain%202024/_git/core-account-abstraction-sdk
git config submodule."internal/infrastructure/lib/did-root-sdk".url https://dev.azure.com/dattabot/Blockchain%202024/_git/did-root-sdk
git config submodule."internal/infrastructure/lib/did-alias-sdk".url https://dev.azure.com/dattabot/Blockchain%202024/_git/did-alias-sdk
git config submodule."internal/infrastructure/lib/did-verifiable-credentials-sdk".url https://dev.azure.com/dattabot/Blockchain%202024/_git/did-verifiable-credentials-sdk

# Initialize all submodules
git submodule update --init --recursive

# For each submodule, checkout the go-be-v1 branch
# This ensures all submodules are on the correct branch
git submodule foreach 'git checkout go-be-v1 || git fetch origin go-be-v1:go-be-v1 && git checkout go-be-v1'

# Update all submodules to the latest commit on go-be-v1 branch
git submodule update --remote
```

### Submodules Overview

The following submodules are used in this project (all on `go-be-v1` branch):

| Submodule                        | Path                                                         | Purpose                        |
| -------------------------------- | ------------------------------------------------------------ | ------------------------------ |
| `core-general-sdk`               | `internal/infrastructure/lib/core-general-sdk`               | General blockchain utilities   |
| `core-account-abstraction-sdk`   | `internal/infrastructure/lib/core-account-abstraction-sdk`   | Account abstraction SDK        |
| `did-root-sdk`                   | `internal/infrastructure/lib/did-root-sdk`                   | DID Root SDK                   |
| `did-alias-sdk`                  | `internal/infrastructure/lib/did-alias-sdk`                  | DID Alias SDK                  |
| `did-verifiable-credentials-sdk` | `internal/infrastructure/lib/did-verifiable-credentials-sdk` | DID Verifiable Credentials SDK |

### Working with Submodules

After initial setup, if you need to update submodules to the latest version of the `go-be-v1` branch:

```bash
# Update all submodules to the latest commit on go-be-v1 branch
git submodule update --remote --merge

# Or update a specific submodule
git submodule update --remote internal/infrastructure/lib/core-general-sdk
```

To checkout a specific branch for all submodules:

```bash
# Switch all submodules to go-be-v1 branch
git submodule foreach 'git checkout go-be-v1'
```

### Building the Project

After cloning and initializing submodules:

```bash
# Install dependencies
go mod download

# Build the application
go build -o bin/server cmd/server/main.go

# Run the server
./bin/server
```

Or use `go run`:

```bash
go run cmd/server/main.go
```

### Configuration

Create a `.env` file in the root directory with the required configuration variables. See `internal/config/config.go` for required environment variables.

---

## Folder Structure & Responsibilities

### 📁 `cmd/` - Application Entry Point

**Purpose**: Contains the main application entry point and bootstrapping logic.

**Responsibilities**:

- Initializes application configuration
- Sets up blockchain connections
- Initializes SDKs (Account Abstraction SDK, DID Root SDK, DID Alias SDK, DID VC SDK)
- Configures and starts the HTTP server
- Handles global error handling

**Key Files**:

- `cmd/server/main.go`: Main entry point that orchestrates application startup

**Dependencies**:

- Imports: `internal/config`, `internal/delivery/http/router`, `internal/infrastructure/sdk/*`, `pkg/logger`, `internal/validator`

---

### 📁 `internal/config/` - Configuration Management

**Purpose**: Centralized configuration loading, validation, and access.

**Responsibilities**:

- Loads environment variables from `.env` files
- Validates required configuration values
- Provides singleton access to configuration throughout the application
- Manages blockchain network configurations (HNS - Hybrid Name Service addresses)
- Manages application settings (port, log level, version, timeouts, RPC endpoints)
- Configures HNS for all modules: Account Abstraction, DID Root, DID Alias, DID VC

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
  - `wallet_factory.go`: Wallet creation input structures
  - `wallet.go`: Wallet-related entities
  - `entrypoint.go`: Operation execution input structures
- `didalias/`:
  - `aliasfactory.go`: DID Alias factory operation entities
- `didroot/`:
  - `root_storage.go`: DID Root storage operation entities
- `didvc/`:
  - `nftbase.go`: Verifiable Credentials NFT base entities
  - `vcstorage.go`: Verifiable Credentials storage entities
- `helper/`:
  - `entity.go`: Helper utility domain entities
  - `didalias.go`: DID Alias encoding helper entities
  - `didroot.go`: DID Root encoding helper entities
  - `didvc.go`: DID VC encoding helper entities

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

- `accountabstraction.go`: `AccountAbstractionRepository` - Wallet creation, operation execution, and wallet validation
- `didroot.go`: `DIDRootRepository` - DID resolution, key/claim management, data operations, ownership verification
- `didalias.go`: `DIDAliasRepository` - Alias resolution, owner/DID retrieval, namehash, registration period, TLD registration
- `didvc.go`: `DIDVCRepository` - Credential metadata, validation, token management, approval checks
- `helper.go`: `HelperRepository` - String conversion, parameter encoding for all DID modules (Root, Alias, VC)

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
  - `handle_ops.go`: Operation execution use case
  - `is_valid_wallet.go`: Wallet validation use case
- `didroot/`:
  - `usecase.go`: DID Root use case structure
  - `resolve_did.go`: DID resolution use case
  - `verify_ownership.go`: Ownership verification use case
  - `get_key.go`, `get_key_by_index.go`, `get_key_by_did.go`: Key retrieval use cases
  - `get_claim.go`, `get_claim_by_did.go`, `verify_claim.go`: Claim management use cases
  - `get_data.go`, `get_did_key_data_count.go`, `didhex32_index.go`, `index_did.go`: Data operations use cases
  - `get_original_key.go`: Original key retrieval use case
- `didalias/`:
  - `usecase.go`: DID Alias use case structure
  - `resolve_alias.go`, `resolve_alias_from_string.go`: Alias resolution use cases
  - `get_alias_status.go`, `get_alias_status_from_string.go`: Alias status retrieval use cases
  - `get_owner.go`, `get_owner_from_string.go`: Owner retrieval use cases
  - `get_did.go`, `get_did_from_string.go`: DID retrieval use cases
  - `get_namehash.go`: Namehash calculation use case
  - `get_registration_period.go`: Registration period retrieval use case
  - `register_tld.go`: Top-level domain registration use case
- `didvc/`:
  - `usecase.go`: DID VC use case structure
  - `get_metadata.go`: Credential metadata retrieval use case
  - `is_credential_valid.go`: Credential validation use case
  - `get_credentials_with_metadata.go`: Credentials with metadata retrieval use case
  - `get_unclaimed_tokenid.go`, `total_token_tobe_claimed_did.go`, `get_tobe_claimed_token_did.go`: Token claim management use cases
  - `is_approved_for_all.go`: Approval check use case
  - `get_identity_token_count.go`, `get_certificate_token_count.go`: Token count retrieval use cases
  - `get_identity_token_ids.go`, `get_certificate_token_ids.go`: Token ID retrieval use cases
  - `get_all_identity_token_ids.go`, `get_all_certificate_token_ids.go`: All token IDs retrieval use cases
- `helper/`:
  - `usecase.go`: Helper use case structure
  - `didroot.go`: DID Root parameter encoding use cases
  - `didalias.go`: DID Alias parameter encoding use cases
  - `didvc.go`: DID VC parameter encoding use cases

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
  - `entrypoint.go`: Entry point operations for handleOps
  - **Implements**: `repository.AccountAbstractionRepository`
- **`didroot/`**:
  - `sdk.go`: DID Root SDK initialization and singleton management
  - `root_storage.go`: Root storage operations (implements repository interface)
  - **Implements**: `repository.DIDRootRepository`
- **`didalias/`**:
  - `sdk.go`: DID Alias SDK initialization and singleton management
  - `aliasfactory.go`: Alias factory operations (implements repository interface)
  - **Implements**: `repository.DIDAliasRepository`
- **`didvc/`**:
  - `sdk.go`: DID VC SDK initialization and singleton management
  - `nftbase.go`: NFT base operations for credentials (implements repository interface)
  - `vcstorage.go`: VC storage operations (implements repository interface)
  - **Implements**: `repository.DIDVCRepository`

#### `internal/infrastructure/lib/` - External Libraries

- Contains local copies/vendored external libraries:
  - `core-account-abstraction-sdk/`: Account abstraction SDK
  - `core-general-sdk/`: General blockchain utilities
  - `did-root-sdk/`: DID Root SDK
  - `did-alias-sdk/`: DID Alias SDK
  - `did-verifiable-credentials-sdk/`: DID Verifiable Credentials SDK

#### `internal/infrastructure/helper/` - Infrastructure Helpers

- `helper.go`: Infrastructure-level helper implementations
- `didroot.go`: DID Root helper operations
- `didalias.go`: DID Alias helper operations
- `didvc.go`: DID VC helper operations
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
    - `wallet.go`: Wallet validation handler
    - `entrypoint.go`: Entry point operations handler
  - `didroot/`: DID Root HTTP handlers
    - `handler.go`: Handler structure
    - `root_storage.go`: Root storage operations handler
  - `didalias/`: DID Alias HTTP handlers
    - `handler.go`: Handler structure
    - `aliasfactory.go`: Alias factory operations handler
  - `didvc/`: DID VC HTTP handlers
    - `handler.go`: Handler structure
    - `nftbase.go`: NFT base operations handler
    - `vcstorage.go`: VC storage operations handler
  - `helper/`: Helper endpoint handlers
    - `handler.go`: Helper handler structure
    - `didroot.go`: DID Root encoding handlers
    - `didalias.go`: DID Alias encoding handlers
    - `didvc.go`: DID VC encoding handlers
    - `types.go`: Helper response types

  **Responsibilities**:

  - Parse HTTP requests
  - Validate input data
  - Call use cases
  - Format HTTP responses
  - Handle HTTP errors

- **`router/`** - Route Definitions:

  - `router.go`: Main router setup and middleware configuration
    - Health check endpoint (`/health`)
    - Swagger documentation endpoint (`/api/v1/swagger/*`)
  - `accountabstraction/`: Account abstraction routes
    - `router.go`: Route group setup
    - `wallet_factory.go`: Wallet factory routes
    - `wallet.go`: Wallet routes
    - `entrypoint.go`: Entry point routes
  - `didroot/`: DID Root routes
    - `router.go`: Route group setup
    - `root_storage.go`: Root storage routes
  - `didalias/`: DID Alias routes
    - `router.go`: Route group setup
    - `aliasfactory.go`: Alias factory routes
  - `didvc/`: DID VC routes
    - `router.go`: Route group setup
    - `nftbase.go`: NFT base routes
    - `vcstorage.go`: VC storage routes
  - `helper/`: Helper routes
    - `router.go`: Helper route setup (encoding endpoints for all DID modules)

  **Responsibilities**:

  - Define URL patterns and HTTP methods
  - Wire handlers to routes
  - Configure middleware (CORS, logging, recovery)
  - Dependency injection (creates use cases, handlers)

#### `internal/delivery/event/` - Event-Driven Delivery

- `accountabstraction/wallet_factory.go`: Event handlers for blockchain events (wallet deployment events)
- `didalias/aliasfactory.go`: Event handlers for alias factory events
- `handleops/`: Operation execution event handlers
  - `handleops.go`: Main handleOps orchestrator
  - `didalias.go`: DID Alias handleOps events
  - `didroot.go`: DID Root handleOps events
  - `didvc.go`: DID VC handleOps events

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

### Example: Creating a Wallet (Account Abstraction)

```
1. HTTP Request
   Client → POST /api/v1/account-abstraction/wallet-factory/create
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

### Example: Resolving a DID (DID Root)

```
1. HTTP Request
   Client → GET /api/v1/did-root/resolve?did=did:example:123

2. Router Layer (internal/delivery/http/router/didroot/)
   • Gets DID Root SDK instance
   • Creates use case with repository
   • Creates handler with use case
   • Routes to handler

3. Handler Layer (internal/delivery/http/handler/didroot/)
   • Parses query parameters
   • Calls use case: uc.ResolveDID(ctx, input)

4. Use Case Layer (internal/usecase/didroot/)
   • Calls repository: repo.ResolveDID(ctx, input)
   • Returns DID document

5. Repository Implementation (internal/infrastructure/sdk/didroot/)
   • Calls RootStorage.ResolveDID()
   • Returns DID document with keys, claims, data

6. Response
   • DID document formatted and returned to client
```

### Example: Resolving an Alias (DID Alias)

```
1. HTTP Request
   Client → GET /api/v1/did-alias/resolve?alias=example.domain

2. Router Layer (internal/delivery/http/router/didalias/)
   • Gets DID Alias SDK instance
   • Creates use case with repository
   • Creates handler with use case
   • Routes to handler

3. Handler Layer (internal/delivery/http/handler/didalias/)
   • Parses query parameters
   • Calls use case: uc.ResolveAlias(ctx, input)

4. Use Case Layer (internal/usecase/didalias/)
   • Calls repository: repo.Resolve(ctx, input)
   • Returns DID hash

5. Repository Implementation (internal/infrastructure/sdk/didalias/)
   • Calls AliasFactory.Resolve()
   • Returns associated DID hash

6. Response
   • DID hash formatted and returned to client
```

### Example: Getting Credential Metadata (DID VC)

```
1. HTTP Request
   Client → GET /api/v1/did-vc/get-metadata?tokenId=123

2. Router Layer (internal/delivery/http/router/didvc/)
   • Gets DID VC SDK instance
   • Creates use case with repository
   • Creates handler with use case
   • Routes to handler

3. Handler Layer (internal/delivery/http/handler/didvc/)
   • Parses query parameters
   • Calls use case: uc.GetMetadata(ctx, input)

4. Use Case Layer (internal/usecase/didvc/)
   • Calls repository: repo.GetMetadata(ctx, input)
   • Returns credential metadata

5. Repository Implementation (internal/infrastructure/sdk/didvc/)
   • Calls NFTBase.GetMetadata()
   • Returns credential metadata

6. Response
   • Metadata formatted and returned to client
```

### Example: Helper Operation (String to Hex32)

```
1. HTTP Request
   Client → POST /api/v1/helper/string-2-hex32
   Body: { "input": "..." }

2. Router Layer (internal/delivery/http/router/helper/)
   • Creates helper repository implementation
   • Creates use case with repository
   • Creates handler with use case
   • Routes to handler

3. Handler Layer (internal/delivery/http/handler/helper/)
   • Parses request
   • Calls use case: uc.StringToHex32(input)

4. Use Case Layer (internal/usecase/helper/)
   • Calls repository: repo.StringToHex32(input)

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
  ├─→ internal/infrastructure/sdk/accountabstraction (initialization)
  ├─→ internal/infrastructure/sdk/didroot (initialization)
  ├─→ internal/infrastructure/sdk/didalias (initialization)
  ├─→ internal/infrastructure/sdk/didvc (initialization)
  ├─→ internal/validator
  └─→ pkg/logger

internal/infrastructure/sdk/accountabstraction/
  ├─→ internal/repository (implements AccountAbstractionRepository)
  ├─→ internal/domain/accountabstraction/*
  └─→ external SDKs (from lib/core-account-abstraction-sdk/)

internal/infrastructure/sdk/didroot/
  ├─→ internal/repository (implements DIDRootRepository)
  ├─→ internal/domain/didroot/*
  └─→ external SDKs (from lib/did-root-sdk/)

internal/infrastructure/sdk/didalias/
  ├─→ internal/repository (implements DIDAliasRepository)
  ├─→ internal/domain/didalias/*
  └─→ external SDKs (from lib/did-alias-sdk/)

internal/infrastructure/sdk/didvc/
  ├─→ internal/repository (implements DIDVCRepository)
  ├─→ internal/domain/didvc/*
  └─→ external SDKs (from lib/did-verifiable-credentials-sdk/)

internal/infrastructure/helper/
  ├─→ internal/repository (implements HelperRepository)
  ├─→ internal/domain/helper/*
  ├─→ internal/config
  └─→ utils/

internal/delivery/event/
  ├─→ internal/infrastructure/sdk/* (uses SDKs)
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

## API Modules

The API is organized into the following modules, each accessible under `/api/v1/`:

### System Endpoints

- `GET /health` - Health check endpoint
- `GET /api/v1/swagger/*` - Swagger API documentation

### Account Abstraction (`/api/v1/account-abstraction/`)

**Wallet Factory:**

- `POST /wallet-factory/create` - Create a new wallet

**Wallet:**

- `GET /wallet/is-valid` - Check if a wallet is valid

**Entry Point:**

- `POST /entrypoint/handle-ops` - Execute operations through entry point

### DID Root (`/api/v1/did-root/`)

**Root Storage Operations:**

- `GET /resolve-did` - Resolve a DID to get its document
- `GET /verify-did-ownership` - Verify DID ownership
- `GET /get-key` - Get a key by hash
- `GET /get-keys-by-did` - Get all keys associated with a DID
- `GET /get-claim` - Get a claim by hash
- `GET /get-claims-by-did` - Get all claims associated with a DID
- `GET /verify-claim` - Verify a claim
- `GET /get-data` - Get data associated with a DID and key
- `GET /get-did-key-data-count` - Get count of data entries for a DID key
- `GET /get-did-key-data-by-index` - Get data by index for a DID key
- `GET /get-original-key` - Get original key by hash
- `GET /did-index-map` - Map DID hex32 to index
- `GET /did-index-map-reverse` - Map index to DID hex32

### DID Alias (`/api/v1/did-alias/`)

**Alias Factory Operations:**

- `GET /resolve` - Resolve alias namehash to DID
- `GET /resolve-from-string` - Resolve alias string to DID
- `GET /status` - Get alias status by namehash
- `GET /status-from-string` - Get alias status by string
- `GET /owner` - Get alias owner by namehash
- `GET /owner-from-string` - Get alias owner by string
- `GET /did` - Get DID associated with alias by namehash
- `GET /did-from-string` - Get DID associated with alias by string
- `GET /namehash` - Calculate namehash for a domain string
- `GET /registration-period` - Get registration period for an alias
- `POST /register-tld` - Register a top-level domain

### DID Verifiable Credentials (`/api/v1/did-vc/`)

**NFT Base Operations:**

- `GET /get-metadata` - Get credential metadata by token ID
- `GET /is-credential-valid` - Check if a credential is valid
- `GET /get-credentials-with-metadata` - Get credentials with metadata for a DID
- `GET /get-unclaimed-token-id` - Get unclaimed token ID
- `GET /total-tokens-to-be-claimed-by-did` - Get total tokens to be claimed by DID
- `GET /get-to-be-claimed-tokens-by-did` - Get list of tokens to be claimed by DID
- `GET /is-approved-for-all` - Check if operator is approved for all tokens

**VC Storage Operations:**

- `GET /get-identity-token-count` - Get count of identity tokens for a DID
- `GET /get-certificate-token-count` - Get count of certificate tokens for a DID
- `GET /get-identity-token-ids` - Get identity token IDs for a DID (paginated)
- `GET /get-certificate-token-ids` - Get certificate token IDs for a DID (paginated)
- `GET /get-all-identity-token-ids` - Get all identity token IDs for a DID
- `GET /get-all-certificate-token-ids` - Get all certificate token IDs for a DID
- `GET /get-did-root-storage` - Get DID root storage address

### Helper (`/api/v1/helper/`)

**General Utilities:**

- `POST /string-2-hex32` - Convert string to hex32

**DID Root Encoding:**

- `POST /encode-create-did-param` - Encode create DID parameters
- `POST /encode-update-did-param` - Encode update DID parameters
- `POST /encode-deactivate-did-param` - Encode deactivate DID parameters
- `POST /encode-reactivate-did-param` - Encode reactivate DID parameters
- `POST /encode-transfer-did-owner-param` - Encode transfer DID owner parameters
- `POST /encode-store-data-param` - Encode store data parameters
- `POST /encode-delete-data-param` - Encode delete data parameters
- `POST /encode-add-key-param` - Encode add key parameters
- `POST /encode-remove-key-param` - Encode remove key parameters
- `POST /encode-add-claim-param` - Encode add claim parameters
- `POST /encode-remove-claim-param` - Encode remove claim parameters

**DID Alias Encoding:**

- `POST /encode-set-did-root-storage-param` - Encode set DID root storage parameters
- `POST /encode-register-domain-param` - Encode register domain parameters
- `POST /encode-register-subdomain-param` - Encode register subdomain parameters
- `POST /encode-set-did-param` - Encode set DID parameters
- `POST /encode-extend-registration-param` - Encode extend registration parameters
- `POST /encode-revoke-alias-param` - Encode revoke alias parameters
- `POST /encode-unrevoke-alias-param` - Encode unrevoke alias parameters
- `POST /encode-transfer-alias-ownership-param` - Encode transfer alias ownership parameters

**DID VC Encoding:**

- `POST /encode-issue-credential-param` - Encode issue credential parameters
- `POST /encode-burn-credential-param` - Encode burn credential parameters
- `POST /encode-update-metadata-param` - Encode update metadata parameters
- `POST /encode-revoke-credential-param` - Encode revoke credential parameters
- `POST /encode-claim-credential-param` - Encode claim credential parameters

For detailed API documentation, visit `/api/v1/swagger/` when the server is running.

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

Each feature has code organized across layers:

**Account Abstraction:**

- `domain/accountabstraction/` - Domain entities
- `usecase/accountabstraction/` - Business logic
- `infrastructure/sdk/accountabstraction/` - SDK implementations
- `delivery/http/handler/accountabstraction/` - HTTP handlers
- `delivery/http/router/accountabstraction/` - HTTP routes
- `delivery/event/accountabstraction/` - Event handlers

**DID Root:**

- `domain/didroot/` - Domain entities
- `usecase/didroot/` - Business logic
- `infrastructure/sdk/didroot/` - SDK implementations
- `delivery/http/handler/didroot/` - HTTP handlers
- `delivery/http/router/didroot/` - HTTP routes
- `delivery/event/handleops/didroot.go` - HandleOps event handlers

**DID Alias:**

- `domain/didalias/` - Domain entities
- `usecase/didalias/` - Business logic
- `infrastructure/sdk/didalias/` - SDK implementations
- `delivery/http/handler/didalias/` - HTTP handlers
- `delivery/http/router/didalias/` - HTTP routes
- `delivery/event/didalias/` - Event handlers
- `delivery/event/handleops/didalias.go` - HandleOps event handlers

**DID Verifiable Credentials (DID VC):**

- `domain/didvc/` - Domain entities
- `usecase/didvc/` - Business logic
- `infrastructure/sdk/didvc/` - SDK implementations
- `delivery/http/handler/didvc/` - HTTP handlers
- `delivery/http/router/didvc/` - HTTP routes
- `delivery/event/handleops/didvc.go` - HandleOps event handlers

**Helper (Cross-cutting utilities):**

- `domain/helper/` - Domain entities
- `usecase/helper/` - Business logic
- `infrastructure/helper/` - Infrastructure implementations
- `delivery/http/handler/helper/` - HTTP handlers
- `delivery/http/router/helper/` - HTTP routes

This allows features to be developed and maintained independently while following consistent architectural patterns.
