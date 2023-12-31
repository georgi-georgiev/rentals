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
        "/rentals": {
            "get": {
                "description": "by query",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all rentals",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "100",
                        "name": "price_min",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "500",
                        "name": "price_max",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "10",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "0",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "1,2,3",
                        "name": "ids",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "40.7128,-74.0060",
                        "name": "near",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name DESC",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/app.RentalResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.HTTPErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.HTTPErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/rentals/{rental_id}": {
            "get": {
                "description": "by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get rental",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "rental_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.RentalResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.HTTPErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.HTTPErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.HTTPErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.HTTPError": {
            "type": "object",
            "properties": {
                "argument": {
                    "type": "string",
                    "example": "number"
                },
                "detail": {
                    "type": "string",
                    "example": "The parameters: limit, date were not provided"
                },
                "expression": {
                    "type": "string",
                    "example": "greater"
                },
                "field": {
                    "type": "string",
                    "example": "email"
                },
                "placement": {
                    "type": "string",
                    "example": "field"
                },
                "reason": {
                    "type": "string",
                    "example": "invalidParameter"
                },
                "reason_code": {
                    "type": "integer",
                    "example": 150
                },
                "title": {
                    "type": "string",
                    "example": "required parameters are missing"
                }
            }
        },
        "app.HTTPErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/app.HTTPError"
                    }
                },
                "status": {
                    "type": "integer",
                    "example": 400
                },
                "trace_id": {
                    "type": "string",
                    "example": "EJplcsCHuLu"
                }
            }
        },
        "app.LocationResponse": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                },
                "state": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "app.PriceResponse": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                }
            }
        },
        "app.RentalResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "length": {
                    "type": "number"
                },
                "location": {
                    "$ref": "#/definitions/app.LocationResponse"
                },
                "make": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "$ref": "#/definitions/app.PriceResponse"
                },
                "primary_image_url": {
                    "type": "string"
                },
                "sleeps": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/app.UserResponse"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "app.UserResponse": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
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
