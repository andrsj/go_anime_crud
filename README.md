# Anime Go CRUD

```note
/cmd
    /server
        main.go
/internal
    /domain
    |
    | The domain layer is usually designed to be independent of the infrastructure, 
    | such as databases or web frameworks, and can be easily tested in isolation. 
    |
    |---/model
    |   |
    |   | Just model description
    |   |
    |   |---anime.go
    |   --------------
    |   /repositories    // repository uses driver
    |   |
    |   | The objects that is responsible for managing a collection of data or entities
    |   |
    |   |---interface.go
    |   |---errors.go
    |   |
    |   /memory          // It's map driver XD
    |       /create.go
    |       /get.go
    |       /update.go
    |       /delete.go
    |
    |       /tools.go
    |       /storage.go
    |       -----------

    |
    // Haven't realized
    /infrastructure
    |
    | The infrastructure package in a software project typically includes 
    | the implementation details of the system, such as databases, network protocols, 
    | and other technical concerns that aren't part of the core business logic of the application.
    |
    |---/database        // of DB drivers
    |   |---/mysql
    |   |   |---anime.go
    |   |---database.go
    |   -----------------
    |


    |
    /delivery
    |---/rest
    |   |---/api
    |       |---api.go     // description and construction of anime api
    |       |---create.go
    |       |---get.go
    |       |---update.go
    |       |---delete.go
    |
    /services // or usecases
    |
    | The usecases layer typically orchestrates the interaction between 
    | the domain and infrastructure layers to achieve a specific goal.
    |
    |---/anime
    |   |---anime.go  // description and construction of anime service
    |   |---create.go
    |   |---delete.go
    |   |---update.go
    |   |---get.go
    |   ------------
```
