# Hexagonal Todo API

golang todo api using Hexagonal architecture

# Project Structure

* cmd
  * rest-server: entry point for restful api server
* api: [oapi codegen](https://github.com/deepmap/oapi-codegen/) stuff goes here
* internal
  * core: core business process
    * domain: entity
    * port: interface to glue service adapter with core
    * service: use case goes here
  * adapters: service adapter
    * config: save the application config
    * rest-api: restapi driver actor
    * storage: driven actor for all storage capability
      * pgsql: all stuff needed for pgsql goes here like: schema, repo, and migration

# Acknowledgement

* https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij

# Todo

* query builder: https://github.com/Masterminds/squirrel
* auth