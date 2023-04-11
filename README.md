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
* Set of proto files for each mentioned above service. They can be used to learn about API and can help in
  generating client library in practically any programming language.
* Dedicated Golang client libraries for all the services. They were generated using proto files by dedicated protoc
  compilers from Goten.
* Docs in HTML format located [here](./docs/apis), with API reference for each service.
* [Examples](./examples/cmd) - they use our Golang libraries.

In simple words, EdgeLQ SDK provides code for client-based applications talking to EdgeLQ services.

## Repository structure

[goten sdk]: https://github.com/cloudwan/goten-sdk
Each service contains its own directory in this repo. Internal structure is practically same everywhere, as Goten
enforces some common standards. All the components reference [Goten SDK] library.


Code organization within single service has following pattern:

### \<service-name\>/proto/\<version\>
Provides description of APIs, resources and other objects in protobuf format. Files there can be used to learn about service
or be used in order to generate client libraries in other programming languages.

Separation of APIs and resources into subpackages allows building more light-weight applications.

### \<service-name\>/resources/\<version\>/\<resource-name\>
Golang module for single resource. Contains definition of it and all helper objects - dedicated Filter, FieldMask,
set of FieldPaths etc.

### \<service-name\>/client/\<version\>/\<api-name\>
Contains basic (almost like raw, but type safe) client for communication with RPC API and definitions for most
request/response objects.

### \<service-name\>/client/\<version\>/\<service-name\>
Contains sum of all clients for given service. May not be recommended if we want to minimize size of application runtime.

### \<service-name\>/access/\<version\>/\<resource-name\>
Contains high-level client-based components like Watcher - dedicated for each resource.

## How to use
EdgeLQ utilizes [Goten SDK], which contains basic instructions. It does not require much
extra - in order to develop Golang application with this SDK, you need to install Go and include SDK in your list of dependencies.
You can use [examples](./examples) provided in this repository as a help.

In order to generate SDK in other language, you need to follow steps regarding protoc installation, as described in Goten
SDK. The only extra thing to do, is to include directory with this repository (on your machine) as additional include proto
path.

## Notes about examples
Examples provided in this repository, in order to be executed, require access to some currently running service. You
need user or service account ready on hand. This user/service account needs to have a granted permission in a project where
you want to execute the examples. It is assumed that, at the very least, you have all of those, otherwise you would not be
able to use cuttle or UI dashboard. SDK gives an access on the code-level, so one step further. Here you just need some additional,
extra information on how to execute examples properly.

### Endpoint
Endpoint should contain domain with service name and port number, for example "devices.stg01b.edgelq.com:443".

### Authentication & Authorization
Every action requires authentication and authorization. For this reason, example has to be executed as user or service
account. If you want to execute example as user, you need to get access token. For service account, you need to get a
file with service account key. Every example requires one or another to work.

#### Getting Access Token
In order to get fresh access token for user, login to any dashboard like Squid. Open developer tools, find HTTP headers for
any request dashboards sends to edgeLQ service (devices, iam, applications...). Find "authorization" header. It should have
a format "Bearer $VALUE". This VALUE will be your access token. You can copy it and use it with example.

Of course, this token has limited value, so it will need to be checked again in some near future.

Same token could be obtained from cuttle credentials (in ~/.config/goten.cuttle.json as default path). You need to find
your user entry and access token.

#### Service account key file
This is JSON file with email, key ID, private key value. File format can be found [here](./examples/files/service_account_credentials_template.json).

If you dont have private key for any service account, but you have user, then you can create your own, its simple. This
is how you can make with cuttle:

First, List your projects:
```shell
cuttle iam list-my projects
```

Chose some project ID, taken from "name" field. Of course, strip "projects/" prefix. Now, I assume that your user
can create service accounts. In that case, make some (with key):

```shell
cuttle iam create service-account my-test-svcacc-for-examples [--region $REGION_ID] --project $PROJECT_ID
```

Of course, "my-test-svcacc-for-examples" is $SERVICE_ACCOUNT_ID. You can replace with any value you like. You can specify
REGION if you like, but you can skip. If you do, it will create in the nearest available region, which is fine. You need
to look at response object. You will see fully qualified name with REGION_ID assigned. You will also have now EMAIL value,
which is exactly what you need to have in service account key credentials file!

Now, you can create service account key. With cuttle help, it will automatically create service account key file for you,
so you dont need template really:

```shell
cuttle iam create service-account-key --project $PROJECT_ID --region $REGION_ID --service-account $SERVICE_ACCOUNT_ID \
  --algorithm RSA_4096 --credentials-output-file my-service-account-key-file.json
```

As you may have noticed, this command will create file for you.
If you want to create private key on your own machine, its also possible (see cuttle iam create service-account-key --help).

Service account created in this way will not be functional without role bindings. If you want to allow this
account to do something useful, you need to create role binding. For example, this will allow service account to do
anything it likes in devices service:
```shell
cuttle iam create role-binding --project $PROJECT_ID --role 'roles/devices-admin' --member 'serviceAccount:$EMAIL'
```

Value EMAIL must be same as in service account key file.
