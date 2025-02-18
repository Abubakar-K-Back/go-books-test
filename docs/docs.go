// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/books": {
            "get": {
                "summary": "Get all books",
                "description": "Get a paginated list of books",
                "tags": ["Books"],
                "parameters": [
                    {"name": "limit", "in": "query", "type": "integer", "required": false, "description": "Number of books to retrieve"},
                    {"name": "offset", "in": "query", "type": "integer", "required": false, "description": "Offset for pagination"}
                ],
                "responses": {
                    "200": {"description": "Success", "schema": {"type": "array", "items": {"$ref": "#/definitions/models.Book"}}}
                }
            },
            "post": {
                "summary": "Create a book",
                "description": "Add a new book to the database",
                "tags": ["Books"],
                "parameters": [
                    {"name": "book", "in": "body", "required": true, "schema": {"$ref": "#/definitions/models.Book"}}
                ],
                "responses": {
                    "201": {"description": "Created", "schema": {"$ref": "#/definitions/models.Book"}},
                    "400": {"description": "Invalid request data", "schema": {"type": "object"}}
                }
            }
        },
        "/books/{id}": {
            "get": {
                "summary": "Get a book by ID",
                "description": "Retrieve a specific book",
                "tags": ["Books"],
                "parameters": [
                    {"name": "id", "in": "path", "type": "integer", "required": true, "description": "Book ID"}
                ],
                "responses": {
                    "200": {"description": "Success", "schema": {"$ref": "#/definitions/models.Book"}},
                    "404": {"description": "Book not found", "schema": {"type": "object"}}
                }
            },
            "put": {
                "summary": "Update a book",
                "description": "Modify an existing book",
                "tags": ["Books"],
                "parameters": [
                    {"name": "id", "in": "path", "type": "integer", "required": true, "description": "Book ID"},
                    {"name": "book", "in": "body", "required": true, "schema": {"$ref": "#/definitions/models.Book"}}
                ],
                "responses": {
                    "200": {"description": "Updated", "schema": {"$ref": "#/definitions/models.Book"}},
                    "400": {"description": "Invalid request data", "schema": {"type": "object"}},
                    "404": {"description": "Book not found", "schema": {"type": "object"}}
                }
            },
            "delete": {
                "summary": "Delete a book",
                "description": "Remove a book from the database",
                "tags": ["Books"],
                "parameters": [
                    {"name": "id", "in": "path", "type": "integer", "required": true, "description": "Book ID"}
                ],
                "responses": {
                    "200": {"description": "Book deleted", "schema": {"type": "object"}},
                    "404": {"description": "Book not found", "schema": {"type": "object"}}
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Book Management API",
	Description:      "A REST API for managing books using Gin, PostgreSQL, Kafka, and Redis.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

