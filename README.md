Reply
=====

Go package for easily replying to HTTP requests with common response types, setting the appropriate `Content-Type` and `Status` headers as needed.

- HTML 
- Text 
- JSON (using encoding/json)
- XML (using encoding/xml)

# Usage

Reply can be used with pretty much any web framework, as long as you have access to the `http.ResponseWriter`. 

```go
import "github.com/dannyvankooten/reply"
```

```go
// this sets Content-Type (incl. charset) and Status header before writing the response.
func myHandler(w http.ResponseWriter, r *http.Request) {
   reply.HTML(w, http.StatusOK, []byte("Hello world!"))
}

// if you just want to set the Content-Type and Status header, omit the last parameter
func myHandler(w http.ResponseWriter, r *http.Request) {
   reply.HTML(w, http.StatusOK)
}

// JSON
func myHandler(w http.ResponseWriter, r *http.Request) {
   reply.JSON(w, http.StatusOK, map[string]string{"foo": "bar"})
}

// Text
func myHandler(w http.ResponseWriter, r *http.Request) {
   reply.Text(w, http.StatusOK, []byte("Hello world!"))
}

// XML
func myHandler(w http.ResponseWriter, r *http.Request) {
   reply.XML(w, http.StatusOK, map[string]string{"foo": "bar"})
}

// XML with struct
func myHandler(w http.ResponseWriter, r *http.Request) {
   reply.XML(w, http.StatusOK, &myType{ Name: "John Doe" })
}
```

Reply defaults to `UTF-8` as its charset. To override it, set the package global named `Charset`.

```go
reply.Charset = "UTF-16"
```

# License
MIT licensed.
