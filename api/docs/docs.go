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
        "/api/series": {
            "get": {
                "description": "Returns a list of series stored in the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "series"
                ],
                "summary": "Retrieves all series",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Series"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Allows adding a series to the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "series"
                ],
                "summary": "Creates a new series",
                "parameters": [
                    {
                        "description": "Series data",
                        "name": "series",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Series"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.Series"
                        }
                    }
                }
            }
        },
        "/api/series/{id}": {
            "get": {
                "description": "Returns a series stored in the database based on the provided ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "series"
                ],
                "summary": "Retrieves a series by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Series"
                        }
                    }
                }
            },
            "put": {
                "description": "Allows modifying a series in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "series"
                ],
                "summary": "Updates a series",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated series data",
                        "name": "series",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Series"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Series"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a series based on its ID.",
                "tags": [
                    "series"
                ],
                "summary": "Deletes a series",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Series deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/series/{id}/downvote": {
            "patch": {
                "description": "Decreases the ranking field of a series by 1.",
                "tags": [
                    "series"
                ],
                "summary": "Downvotes a series",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/series/{id}/episode": {
            "patch": {
                "description": "Increments the last_episode_watched field of a series.",
                "tags": [
                    "series"
                ],
                "summary": "Increments the last episode watched",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/series/{id}/status": {
            "patch": {
                "description": "Updates the status field of a series in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "series"
                ],
                "summary": "Update the status of a series",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New status",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ApiResponse"
                        }
                    }
                }
            }
        },
        "/api/series/{id}/upvote": {
            "patch": {
                "description": "Increases the ranking field of a series by 1.",
                "tags": [
                    "series"
                ],
                "summary": "Upvote a series",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ApiResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Message related to the response",
                    "type": "string"
                },
                "series": {
                    "description": "Optional series data, included if available",
                    "allOf": [
                        {
                            "$ref": "#/definitions/main.Series"
                        }
                    ]
                },
                "success": {
                    "description": "Indicates whether the request was successful",
                    "type": "boolean"
                }
            }
        },
        "main.Series": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "lastEpisodeWatched": {
                    "type": "integer"
                },
                "ranking": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "totalEpisodes": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Series API",
	Description:      "API for series management.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
