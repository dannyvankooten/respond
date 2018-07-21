Respond
=====

Go package for easily responding to HTTP requests with common response types, setting the appropriate `Content-Type` and `Status` headers as needed.

- HTML 
- Text 
- JSON (using encoding/json)
- XML (using encoding/xml)

# Usage

Response can be used with pretty much any web framework. As long as you have access to a `http.ResponseWriter`, you are good to go. 

```go
import "github.com/dannyvankooten/response"
```

```go
// this sets Content-Type (incl. charset) and Status header before writing the response.
func myHandler(w http.ResponseWriter, r *http.Request) {
   response.HTML(w, http.StatusOK, []byte("Hello world!"))
}

// if you just want to set the Content-Type and Status header, omit the last parameter
func myHandler(w http.ResponseWriter, r *http.Request) {
   response.HTML(w, http.StatusOK)
}

// JSON
func myHandler(w http.ResponseWriter, r *http.Request) {
   response.JSON(w, http.StatusOK, map[string]string{"foo": "bar"})
}

// Text
func myHandler(w http.ResponseWriter, r *http.Request) {
   response.Text(w, http.StatusOK, []byte("Hello world!"))
}

// XML
func myHandler(w http.ResponseWriter, r *http.Request) {
   response.XML(w, http.StatusOK, map[string]string{"foo": "bar"})
}

// XML with struct
func myHandler(w http.ResponseWriter, r *http.Request) {
   response.XML(w, http.StatusOK, &myType{ Name: "John Doe" })
}
```

Response defaults to `UTF-8` as its charset. To override it, set the package global named `Charset`.

```go
response.Charset = "UTF-16"
```

# License
MIT licensed.
