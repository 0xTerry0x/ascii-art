<h1 align="center">🎨 ASCII-ART</h1>

<h3 align="center">A small Go program that converts text into ASCII-art using different banner styles.</h3>

---

<h2> Usage</h2>

```
go run . "<text>" [style]
```



---

<h2> Examples</h2>

```
go run . "Hello"


 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/




go run . "Hello" --shadow

_|    _|          _| _|
_|    _|   _|_|   _| _|   _|_|
_|_|_|_| _|_|_|_| _| _| _|    _|
_|    _| _|       _| _| _|    _|
_|    _|   _|_|_| _| _|   _|_|


go run . "Hello" --thinkertoy

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

<h2>🔤 Supported Characters</h2>

The program supports all printable ASCII characters
(ASCII codes <b>32–126</b>).

Non-printable or out-of-range characters are:

- skipped during rendering
- and their indexes are printed at the end

---

<h2>🔄 Newlines</h2>

To print multiple lines, include the literal sequence:

```
\n
```


Example:

```
go run . "Hello\nWorld"

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


<h4 align="center"></h4>
