<h1>ASCII-ART</h1>


<h2>A small Go program that converts text into ASCII-art using different banner styles.</h2>

<h3>Usage:

```
go run . "<text>" [style]
```

Examples:

```
go run . "Hello"
go run . "Hello\nWorld"
go run . "Hello" --shadow
go run . "Hello" --thinkertoy
```

🎨 Styles

Use one of the following:


--standard (default)

--shadow

--thinkertoy


If no style is provided, standard is used.


📏 Supported Characters

The program supports printable ASCII characters (32–126).

Anything outside that range (non-printable or non-ASCII) is:

skipped during rendering, and their indexes are listed at the end.



🔄 Newlines

To print multiple lines, include the literal sequence:

\n


Example:

```
go run . "Hello\nWorld"
```


⚠️ Errors

You may see error messages for:

missing text argument

invalid style

empty text input
