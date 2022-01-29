# EdgeLQ SDK

## Overview
EdgeLQ is a development platform and environment of reusable, generic microservices, which enables rapid application
development. Each EdgeLQ service utilizes Goten framework, but its API is gRPC/REST compatible.

Microservices available in EdgeLQ:
* Identity and Access Management (IAM) - provides access to users/service account, role management and permission checks.
  It is used by other microservices for authorization.
* Audit - Generic service that provides view into any activity by any user or service account. Any interesting API call
  to any EdgeLQ based service is reported to Audit - including calls to Audit itself.
* Monitoring - Provides generic time-series endpoint, allowing various application to store any timestamp based data
  conforming to registered (by the same applications) descriptors.
* Devices - Provides device management (physical or virtual machines). It also enables connecting into devices.
* Applications - It is generic service managing (any) applications installed on devices.
* Proxies - Provides generic proxy feature for easy tunneling.
* Secrets - Manages secrets, that can be used by other services (like Applications).
* Meta - provides information about state of deployment (like list of regions and endpoints).

EdgeLQ SDK provides:
* Set of proto files for each mentioned above service. They can be used to learn about API, and also can help in
  generating client library in practically any programming language.
* Dedicated Golang client libraries for all the services. They were generated based on proto files by dedicated protoc
  compilers from Goten.
* Docs in HTML format - this is [the place](./docs/apis) where you can obtain full API reference for each service.
* [Examples](./examples/cmd) - they use our Golang libraries.

In simple words, EdgeLQ SDK provides code for client-based applications talking to EdgeLQ services.

## Repository structure
Each service contains its own directory in this repo. Internal structure is practically same everywhere, as Goten
enforces some common standards. All the components reference [Goten SDK library](github.com/cloudwan/goten-sdk).

As an additional note, EdgeLQ SDK contains only client-side elements.
Code organization within single service has following pattern:

### \<service-name\>/proto/\<version\>
Provides description of APIs, resources and other objects in protobuf format. Those can be used to learn about service
and also they can be used to generate clients in other programming languages.

Separation of APIs and resources allows building more light-weight applications, because they may be picking only those
packages they are interested in.

### \<service-name\>/resources/\<version\>/\<resource-name\>
Golang module for single resource. Contains definition of resource and all helper objects - dedicated Filter, FieldMask,
set of FieldPaths etc.

### \<service-name\>/client/\<version\>/\<api-name\>
Contains basic (almost like raw, but type safe) client for communication with RPC API and definitions for most
request/response objects.

### \<service-name\>/client/\<version\>/\<service-name\>
Contains sum of all clients for given service. May not be recommended if we want to minimize size of application runtime.

### \<service-name\>/access/\<version\>/\<resource-name\>
Contains high-level client-based components like Watcher - dedicated for each resource.

## How to use
EdgeLQ utilizes [Goten SDK](github.com/cloudwan/goten-sdk), which contains basic instructions. It does not require much
extra - in order to develop Golang application, you need to install Go and include EdgeLQ SDK in your list of dependencies
(file go.mod!). You can use [examples](./examples) provided in this repository as a help.

In order to generate library in own language, you need to follow steps regarding protoc installation, as described in Goten
SDK. The only extra thing to do, is to include directory with this repository (on your machine) as additional include proto
path.
