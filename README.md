# struct-fmt

## package structfmt

A [GoDoc](https://godoc.org/github.com/mezcel/struct-fmt), json to struct, conversion package. 

### About

* This is a Go package used to port Scripture Rosary data from a ```.json``` file into Go structs.
* The imported ```.json``` file  must have the same ER Schema used in:
    > My Rosary Apps \
    > [python-curses]( http://github.com/mezcel/python-curses ), [printf-time]( http://github.com/mezcel/printf-time ), [jq-tput-terminal](https://github.com/mezcel/jq-tput-terminal), or [electron-container](https://github.com/mezcel/electron-container).

### Install

* Install Go: [download](https://golang.org/dl/)
* Install the ```struct-fmt``` Go package
    ```sh
    ## Install struct-fmt
    go get github.com/mezcel/struct-fmt
    ```

### Code Usage

* Include in a Go script.
    ```go
    // Import this package into a Go program
    import "github.com/mezcel/struct-fmt"

    // or
    import structfmt "github.com/mezcel/struct-fmt"
    ```

### Demo Examples

* [README.md](example/README.md)
    * CLI Application [Terminal](https://asciinema.org/a/343751)
    * Web Server with Webpage [YouTube](https://www.youtube.com/watch?v=diSGGO_kDZ0)
    * Native desktop window UI
* [Screenshots](example/screenshots/README.md)