// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Proprietary"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns service info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "Returns service info",
                "operationId": "healthckeck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Healthcheck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            }
        },
        "/sections": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Finds all sections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "Finds all sections",
                "operationId": "get-sections",
                "parameters": [
                    {
                        "description": "sections request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.SectionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Section"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Description": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Операция выполнена успешно"
                },
                "reason": {
                    "type": "string"
                },
                "stacktrace": {
                    "type": "string"
                }
            }
        },
        "api.HTTPError": {
            "type": "object",
            "properties": {
                "archive": {
                    "type": "boolean",
                    "example": false
                },
                "count_inner_themes": {
                    "type": "integer",
                    "example": 0
                },
                "count_outer_themes": {
                    "type": "integer",
                    "example": 0
                },
                "date_archive": {
                    "type": "string",
                    "example": "2016-02-20"
                },
                "description": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Description"
                    }
                },
                "inner_themes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.InnerTheme"
                    }
                },
                "name_razdel": {
                    "type": "string",
                    "example": "qwerty"
                },
                "otdel_razdel": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/api.Otdel"
                    }
                },
                "outer_themes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.OuterTheme"
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "api.Healthcheck": {
            "type": "object",
            "properties": {
                "DB_addr": {
                    "type": "string",
                    "example": "192.168.8.250"
                },
                "DB_name": {
                    "type": "string",
                    "example": "dbqueue_korenovsk_actual"
                },
                "DB_time": {
                    "type": "string",
                    "example": "2020-11-23 15:47:00.900"
                },
                "name": {
                    "type": "string",
                    "example": "Сервис получения информации по разделу"
                },
                "root_path": {
                    "type": "string",
                    "example": "section-info"
                },
                "version": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "api.InnerTheme": {
            "type": "object",
            "properties": {
                "id_theme": {
                    "type": "integer",
                    "example": 123
                },
                "name_theme": {
                    "type": "string"
                }
            }
        },
        "api.Otdel": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "string"
                },
                "windows": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "api.OuterTheme": {
            "type": "object",
            "properties": {
                "id_theme": {
                    "type": "integer",
                    "example": 123
                },
                "name_theme": {
                    "type": "string"
                },
                "tax": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "api.Section": {
            "type": "object",
            "properties": {
                "archive": {
                    "type": "boolean",
                    "example": false
                },
                "count_inner_themes": {
                    "type": "integer",
                    "example": 0
                },
                "count_outer_themes": {
                    "type": "integer",
                    "example": 0
                },
                "date_archive": {
                    "type": "string",
                    "example": "2016-02-20"
                },
                "description": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Description"
                    }
                },
                "inner_themes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.InnerTheme"
                    }
                },
                "name_razdel": {
                    "type": "string",
                    "example": "qwerty"
                },
                "otdel_razdel": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/api.Otdel"
                    }
                },
                "outer_themes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.OuterTheme"
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "api.SectionRequest": {
            "type": "object",
            "properties": {
                "id_operator": {
                    "type": "integer",
                    "example": 2
                },
                "id_otdel": {
                    "type": "integer",
                    "example": 123
                },
                "id_razdel": {
                    "type": "integer",
                    "example": 123123
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Content program rest gateway API",
	Description: "This service allows webui to access content program functionality",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
