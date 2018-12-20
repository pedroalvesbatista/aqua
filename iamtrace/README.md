# IAMTrace

With IAMTrace, will be a lot easier to administrate IAM roles for a bunch of roles regarding variated environment's, routes, pipeline and everything related to authentication in service accounts.

## Basic Structure

To follow most than one scenario at same time, and seeking to start small to lean, IAMTrace grabs :

* A folder to keep all the API's key's and all necessary auth - `/aqua/.iamtrace/keys`, `/aqua/.iamtrace/certs`, `/aqua/.iamtrace/roles` and any other specific artifact to keep on guard;
* A Lambtrone specific directory (`/aqua/.iamtrace/lambtrone`) to keep anything related to it. When calling any specific "trone" or artifact to serve a role, lambtrone component must pass `iamtraceRole := IAMTrace.trone{id}`
