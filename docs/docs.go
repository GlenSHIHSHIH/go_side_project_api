// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/admin/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backstage"
                ],
                "summary": "Backstage login",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backstagedto.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backstagedto.LoginResponseDTO"
                        }
                    }
                }
            }
        },
        "/backstage/jwt/refreshtoken": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backstage"
                ],
                "summary": "Backstage RefreshToken",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backstagedto.JwtRefTokenDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backstagedto.LoginResponseDTO"
                        }
                    }
                }
            }
        },
        "/backstage/user/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backstage"
                ],
                "summary": "Backstage UserLogin",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backstagedto.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResponseDTO"
                        }
                    }
                }
            }
        },
        "/carousel/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Forestage"
                ],
                "summary": "Carousel",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forestagedto.CarouselDTO"
                        }
                    }
                }
            }
        },
        "/category/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Forestage"
                ],
                "summary": "Category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forestagedto.CategoryDTO"
                        }
                    }
                }
            }
        },
        "/forestage/config": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Forestage"
                ],
                "summary": "Forestage config",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forestagedto.BaseForestageConfigDTO"
                        }
                    }
                }
            }
        },
        "/production/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Forestage"
                ],
                "summary": "Production list",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "int default",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            15,
                            20,
                            30,
                            40,
                            50
                        ],
                        "type": "integer",
                        "description": "int enums",
                        "name": "pageLimit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "string enums",
                        "name": "sort",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "PName",
                            "PId",
                            "PCategory",
                            "PCreTime"
                        ],
                        "type": "string",
                        "description": "string enums",
                        "name": "sortColumn",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "",
                        "description": "string default",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "",
                        "description": "string default",
                        "name": "searchCategory",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forestagedto.ProductionDTO"
                        }
                    }
                }
            }
        },
        "/production/rank/{count}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Forestage"
                ],
                "summary": "Production rank",
                "operationId": "10",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "count",
                        "name": "count",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forestagedto.ProductionRankDTO"
                        }
                    }
                }
            }
        },
        "/production/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Forestage"
                ],
                "summary": "Production detail",
                "operationId": "1",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/forestagedto.ProductionDetailDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "backstagedto.JwtInfoDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "minimum": 6
                },
                "name": {
                    "type": "string",
                    "minLength": 4
                }
            }
        },
        "backstagedto.JwtRefTokenDTO": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "backstagedto.JwtTokenDTO": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "backstagedto.LoginDTO": {
            "type": "object",
            "properties": {
                "loginName": {
                    "type": "string",
                    "minLength": 4
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "backstagedto.LoginResponseDTO": {
            "type": "object",
            "properties": {
                "authorityJwt": {
                    "$ref": "#/definitions/backstagedto.JwtTokenDTO"
                },
                "userInfo": {
                    "$ref": "#/definitions/backstagedto.JwtInfoDTO"
                }
            }
        },
        "backstagedto.UserDTO": {
            "type": "object",
            "properties": {
                "UserType": {
                    "description": "是否為系統用戶",
                    "type": "boolean"
                },
                "email": {
                    "description": "Email",
                    "type": "string"
                },
                "loginName": {
                    "description": "登入帳號",
                    "type": "string",
                    "minLength": 4
                },
                "name": {
                    "description": "使用者名稱",
                    "type": "string",
                    "minLength": 4
                },
                "password": {
                    "description": "密碼",
                    "type": "string",
                    "minLength": 6
                },
                "remark": {
                    "description": "備註",
                    "type": "string"
                },
                "status": {
                    "description": "帳號狀態(false停用 true正常)",
                    "type": "boolean"
                }
            }
        },
        "dto.BaseResponseDTO": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "dto.PageDTO": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "pageLimit": {
                    "type": "integer"
                },
                "search": {
                    "type": "string"
                },
                "searchCategory": {
                    "type": "string"
                },
                "sort": {
                    "type": "string"
                },
                "sortColumn": {
                    "type": "string"
                }
            }
        },
        "forestagedto.BaseConfigDTO": {
            "type": "object",
            "properties": {
                "imgUrl": {
                    "type": "string"
                }
            }
        },
        "forestagedto.BaseForestageConfigDTO": {
            "type": "object",
            "properties": {
                "baseConfig": {
                    "$ref": "#/definitions/forestagedto.BaseConfigDTO"
                }
            }
        },
        "forestagedto.CarouselDTO": {
            "type": "object",
            "properties": {
                "carousels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/forestagedto.CarouselData"
                    }
                }
            }
        },
        "forestagedto.CarouselData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "forestagedto.CategoryDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "forestagedto.ProductionDTO": {
            "type": "object",
            "properties": {
                "pageData": {
                    "$ref": "#/definitions/dto.PageDTO"
                },
                "productionList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/forestagedto.ProductionData"
                    }
                }
            }
        },
        "forestagedto.ProductionData": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "images": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "options": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "priceMin": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "forestagedto.ProductionDetailDTO": {
            "type": "object",
            "properties": {
                "production": {
                    "$ref": "#/definitions/forestagedto.ProductionDetailData"
                }
            }
        },
        "forestagedto.ProductionDetailData": {
            "type": "object",
            "properties": {
                "attribute": {
                    "type": "string"
                },
                "categories": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "historicalSold": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "images": {
                    "type": "string"
                },
                "likedCount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "options": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "priceMin": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "forestagedto.ProductionRankDTO": {
            "type": "object",
            "properties": {
                "productionList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/forestagedto.ProductionRankData"
                    }
                }
            }
        },
        "forestagedto.ProductionRankData": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "categories": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "historicalSold": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "images": {
                    "type": "string"
                },
                "likedCount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "options": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "priceMin": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
