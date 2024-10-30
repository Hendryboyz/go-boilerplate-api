// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/todos": {
            "get": {
                "description": "List All Todo Items",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "List All Todo Items",
                "parameters": [
                    {
                        "type": "string",
                        "example": "henry.chou",
                        "description": "the user id to filter",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "CreateItem Todo Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "CreateItem Todo Item",
                "parameters": [
                    {
                        "description": "CreateItem Todo Item",
                        "name": "todoItem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/todos/{itemId}": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Get Todo Item By Id",
                "parameters": [
                    {
                        "type": "string",
                        "example": "7ae9c676-fc23-47a2-abc1-591ad2859b67",
                        "description": "get item by id",
                        "name": "itemId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "DeleteItem Todo Item",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "DeleteItem Todo Item",
                "parameters": [
                    {
                        "type": "string",
                        "example": "7d105cc8-a709-4a28-ae96-f0270bc5ad20",
                        "description": "the item id to be deleted",
                        "name": "itemId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "UpdateItem Todo Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "UpdateItem Todo Item",
                "parameters": [
                    {
                        "type": "string",
                        "example": "7d105cc8-a709-4a28-ae96-f0270bc5ad20",
                        "description": "the item id to be updated",
                        "name": "itemId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateItem Todo Item",
                        "name": "todoItem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateTodoRequest": {
            "type": "object",
            "required": [
                "description",
                "endDate",
                "startDate"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "format": "string",
                    "example": "say hello to everyone"
                },
                "endDate": {
                    "type": "string",
                    "format": "dateTime"
                },
                "startDate": {
                    "type": "string",
                    "format": "dateTime"
                }
            }
        },
        "dto.UpdateTodoRequest": {
            "type": "object",
            "required": [
                "description",
                "endDate",
                "startDate"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "format": "string",
                    "example": "say hello to everyone"
                },
                "endDate": {
                    "type": "string",
                    "format": "dateTime"
                },
                "startDate": {
                    "type": "string",
                    "format": "dateTime"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
