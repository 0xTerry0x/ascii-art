<h1>ASCII-ART</h1>


<h2>A small Go program that converts text into ASCII-art using different banner styles.</h2>

<h3>Usage:</h3>

```
go run . "<text>" [style]
```

<h3>Examples:</h3>

```
go run . "Hello"
go run . "Hello\nWorld"
go run . "Hello" --shadow
go run . "Hello" --thinkertoy
```

<h3>🎨 Styles


Use one of the following:</h3>


<h2>

--standard (default)


--shadow


--thinkertoy

</h2>


<h3>

If no style is provided, standard is used.


📏 Supported Characters

</h3>

<h2>

The program supports printable ASCII characters (32–126).


Anything outside that range (non-printable or non-ASCII) is:


skipped during rendering, and their indexes are listed at the end.

</h2>



<h3>🔄 Newlines</h3>


<h2>To print multiple lines, include the literal sequence:</h2>


```
\n
```


<h3>Example:</h3>

```
go run . "Hello\nWorld"
```


<h3>

⚠️ Errors


You may see error messages for:

</h3>

<h2>

missing text argument

invalid style

empty text input

</h2>
