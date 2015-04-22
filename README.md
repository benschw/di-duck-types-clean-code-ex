Example repo to walk through layering in dependency injection and duck typing for cleaner, more testable code

### default

- Shows straight forward way to discover an address
- To test, you have to build and run against an instance of Consul


### di

- Refactored to have a `Discover` function into which you can inject your dns resolver
- Makes testing easier, but you still can't test `Discover` without also running Consul and resolving a real SRV record

### duck

- Adds `AddressGetter` interface so that tests can use a stub implementation
- We don't want to test dns-clb-go, only our implementation. This facilitates that
- Didn't have to include unnecessary interface in core library
- Allows implementation to only depend on the parts of the library we are using (cleaner code reuse)


### server

- Example client/server implementation where client discovers server's address
- `AddressGetter` interface allows different load balancer implementations to be injected 
- We can component test our web service locally without running Consul
