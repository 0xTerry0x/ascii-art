<h1 align="center">🎨 ASCII-ART</h1>

<h3 align="center">A small Go program that converts text into ASCII-art using different banner styles.</h3>

---

<h2> Usage</h2>

### Quick Start

Run directly from the project root:

```bash
go run ./cmd/ascii-art "<text>" --[style]
```

Or build the binary first:

```bash
go build ./cmd/ascii-art
./ascii-art "<text>" --[style]
```



---

<h2> Examples</h2>

```bash
go run ./cmd/ascii-art "Hello"


 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/




go run ./cmd/ascii-art "Hello" --shadow

_|    _|          _| _|
_|    _|   _|_|   _| _|   _|_|
_|_|_|_| _|_|_|_| _| _| _|    _|
_|    _| _|       _| _| _|    _|
_|    _|   _|_|_| _| _|   _|_|


go run ./cmd/ascii-art "Hello" --thinkertoy

o  o     o o
|  |     | |
O--O o-o | | o-o
|  | |-' | | | |
o  o o-o o o o-o

```


---

<h2> Styles</h2>

Available banner styles:

<ul>
  <li><b>standard</b> (default)</li>
  <li><b>shadow</b></li>
  <li><b>thinkertoy</b></li>
</ul>

If no style is provided, <b>standard</b> is used.

---

<h2> Supported Characters</h2>

The program supports all printable ASCII characters
(ASCII codes <b>32–126</b>).

Non-printable or out-of-range characters are:

- skipped during rendering
- and their indexes are printed at the end

---

<h2> Newlines</h2>

To print multiple lines, include the literal sequence:

```
\n
```


Example:

```
go run ./cmd/ascii-art "Hello\nWorld"

 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/


                           _       _
                          | |     | |
__      __   ___    _ __  | |   __| |
\ \ /\ / /  / _ \  | '__| | |  / _` |
 \ V  V /  | (_) | | |    | | | (_| |
  \_/\_/    \___/  |_|    |_|  \__,_|


```



---

<h2>⚠️ Error Handling</h2>

Errors may occur for:

- empty text argument
- invalid style
- missing input

---

<h2>Project Structure</h2>

```
ascii-art/
├── cmd/ascii-art/          # Main application entry point
│   └── main.go
├── internal/
│   ├── ascii/              # Core ASCII art logic
│   │   ├── ascii.go        # Art generation
│   │   ├── args.go         # Argument parsing
│   │   ├── filter.go       # Character validation
│   │   ├── printer.go      # Output formatting
│   │   ├── read-file.go    # Banner file loading
│   │   ├── art_test.go     # Unit tests for art generation
│   │   └── banner_test.go  # Banner loading tests
│   └── assets/             # Banner font files
│       ├── standard.txt
│       ├── shadow.txt
│       └── thinkertoy.txt
├── main_test.go            # Integration tests
├── edge_cases_test.go      # Edge case tests
├── Makefile                # Build automation
├── go.mod                  # Go module definition
├── README.md               # Project documentation
├── CHANGELOG.md            # Version history
├── CONTRIBUTING.md         # Contribution guidelines
└── LICENSE                 # MIT License
```

---

<h2>Testing</h2>

Run all tests:

```bash
go test ./...
```

Run specific test suites:

```bash
# Unit tests for internal package
go test ./internal/ascii -v

# Integration tests
go test -v -run TestMain

# Edge case tests
go test -v -run TestEdge
```

