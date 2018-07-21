/*
MIT License

Copyright (c) 2018 Danny van Kooten

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package respond

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (
	// ContentType HTTP header name for defining the content type
	ContentType = "Content-Type"

	// ContentHTML HTTP header value for HTML data
	ContentHTML = "text/html"

	// ContentJSON HTTP header value for JSON data
	ContentJSON = "application/json"

	// ContentText header value for Text data.
	ContentText = "text/plain"

	// ContentXML header value for XML data.
	ContentXML = "text/xml"
)

var Charset = "UTF-8"

// HTML executes the template and writes to the responsewriter
func HTML(w http.ResponseWriter, statusCode int, data ...[]byte) error {
	w.Header().Set(ContentType, ContentHTML+"; charset="+Charset)
	w.WriteHeader(statusCode)

	for i := range data {
		w.Write(data[i])
	}

	return nil
}

// JSON renders the data as a JSON HTTP response to the ResponseWriter
func JSON(w http.ResponseWriter, statusCode int, data ...interface{}) error {
	w.Header().Set(ContentType, ContentJSON+"; charset="+Charset)
	w.WriteHeader(statusCode)

	for i := range data {
		err := json.NewEncoder(w).Encode(data[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// XML writes the data as a XML HTTP response to the ResponseWriter
func XML(w http.ResponseWriter, statusCode int, data ...interface{}) error {
	w.Header().Set(ContentType, ContentXML+"; charset="+Charset)
	w.WriteHeader(statusCode)

	// do nothing if nil data
	for i := range data {
		err := xml.NewEncoder(w).Encode(data[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// Text writes the data as a JSON HTTP response to the ResponseWriter
func Text(w http.ResponseWriter, statusCode int, data ...[]byte) error {
	w.Header().Set(ContentType, ContentText+"; charset="+Charset)
	w.WriteHeader(statusCode)

	for i := range data {
		w.Write(data[i])
	}

	return nil
}
