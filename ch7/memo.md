# Chapter7: Go web services

## 7.1: Introduction to web services
Web services types

- SOAP-base: used in enterprise system
  - Simple Object Access Protocol (originally)
  - good: well-supported, many custom utilities, robust, secure
  - bad: large and complicated, hard to troubleshoot
  - **Utility-driven**
  - RPC-form
  - Response message is packed into *envelop*
- REST-base: used in publicly available services
  - REpresentational State Transfer, evolved along with the idea of OOP
  - good: fast and flexible, can use simple format (like JSON)
  - **Data-driven**
  - HTTP method behaves like *verbs*
  - HTTP methods used in REST correspond to CRUD operations in database
  - REST is related to just API design, not response message.
  - When complicated services or some operation/action is required
	- Example: activate user account
	- x: `ACTIVATE /user/456 HTTP/1.1`
	- o: convert action to resource, `POST /user/456/activation HTTP/1.1`
	- o: convert action to resouce property, `PATCH /user/456 HTTP/1.1`
- XML-RPC

## 7.4: Parse and generate XML by golang
XML is used in not describing but sending/receiving data. Officially recommended and defined in W3C.

struct-tag defines correspondance between XML and struct, is conposed of key-value pairs

``` xml
<?xml version="1.0" encoding="UTF-8"?>
<post id="1">
  <content>Hello, World!</content>
  <author id="2">Sau Sheng</author>
</post>      
```

Above XML is packed into struct

``` go
type Post struct {
	XMLName xml.Name `xml:"post"`      // name of XML element itself
	Id      string   `xml:"id,attr"`   // attribute of XML element
	Content string   `xml:"content"`   // low-level element (without mode-flag)
	Author  Author   `xml:"author"`    // low-level element (with attribute id)
	Xml     string   `xml:",innerxml"` // raw XML from XML elements
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"` // character data of XML element
}
```

## 7.5: Parse and generate JSON by golang
JSON (JavaScript Serialized Object Notation) is JS-based data-format. Stands along the idea that *readable for both human and machine*. Firstly defined by Douglas Crockford, now specified in RFC 7159 and ECMA-404. JSON is frequently used in REST-base Web services.

There are many front-end applications using JSON. Such app is natural to use JSON to interact with backend applications.

Go can parse a JSON file with both Decoder and Unmarshal. We can use them properly with input.
Decoder is appropriate for the case of streaming data from `io.Reader` (like Body of `http.Request`).
Unmarshal is appropriate for the case of character data or in-memory data.

## 7.6: Build Go web service

`$ go run server.go data.go`

Interact REST API via curl command (example: retrieve one post)

`$ curl -i -X GET http://127.0.0.1:8080/post/1`

## 7.7: Summary
- Recently, Go is mainly used for building web services
- 2 types of web services exist, SOAP-base and JSON-base
- SOAP is a protocol for handling structured data defined in XML. SOAP-based services are not as popular as REST-based because WSDL message is possibly complicated.
- REST-based web services make resource public and allow action for them
- Generating and parsing XML and JSON is quite similar. We can create struct to generate (unmarshal) them, and convert (marshal) them into struct.
