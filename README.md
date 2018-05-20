# md2html

a simple Markdown to HTML converter

## How to use

```bash
$ go get github.com/russross/blackfriday
$ go build md2html.go
$ ./md2html xxx.md > xxx.html
```

## Example

```hello.md
# hello

```hello.c
main()
{
	printf("Hello!");
}
```
```

```hello.html
<h1>hello</h1>

<pre><code class="language-hello.c">main()
{
	printf(&quot;Hello!&quot;);
}
</code></pre>
```
