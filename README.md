# CRUD SERVER
Golang test webserver for validation. The test server can do basic CRUD operation.

## How to Run

## Supported operations

## Documentation
Documentation of the project can be generated using the godoc.

* Append $PWD to GOPATH before running the godoc server.

* Run the godoc webserver as 
    ```
        godoc -play -http=":6060"
    ```

* To view documentation of exposed packages, view the webpage in the browser
    ```
    http://localhost:6060/pkg/
    ```

* To view all project including the internal packages,
    ```
        http://localhost:6060/pkg/?m=all
    ```