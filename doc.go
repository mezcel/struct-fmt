# struct-fmt

## About

* This is a Go package used to port Scripture Rosary data from a ```.json``` file into Go structs.
* The imported ```.json``` file  must have the same ER Schema used in: [python-curses]( http://github.com/mezcel/python-curses ), [printf-time]( http://github.com/mezcel/printf-time ), [jq-tput-terminal](https://github.com/mezcel/jq-tput-terminal), or [electron-container](https://github.com/mezcel/electron-container).
    * Demo Json: [demos/json](demos/json)
* [GoDoc](https://godoc.org/github.com/mezcel/struct-fmt) 
    * <b>struct-fmt</b>: <u>github.com/mezcel/struct-fmt</u>

### Use Case / Demonstration

* [Demo README](example/README.md)
    * [Terminal](https://asciinema.org/a/343751)
    * Web Server with Webpage
    * Native desktop window UI
* [Screenshots](example/screenshots/README.md)

---

## Install

* Install Go: [download](https://golang.org/dl/)
* Install the ```struct-fmt``` Go package
    ```sh
    ## Install struct-fmt
    go get github.com/mezcel/struct-fmt
    ```

---

## Code Usage

* Include in a Go script.
    ```go
    // Import this package into a Go program
    import "github.com/mezcel/struct-fmt"

    // or
    import structfmt "github.com/mezcel/struct-fmt"
    ```