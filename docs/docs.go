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
        "/api/v1/account/get-all": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all accounts. Requires admin role.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get all accounts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Total number of items per page",
                        "name": "total",
                        "in": "query"
                    },
                    {
                        "description": "Get all accounts request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.AccountFilter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get all accounts successful",
                        "schema": {
                            "$ref": "#/definitions/common.success_response"
                        }
                    },
                    "400": {
                        "description": "Get all accounts failure",
                        "schema": {
                            "$ref": "#/definitions/common.error_response"
                        }
                    }
                }
            }
        },
        "/api/v1/account/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create new account with user's info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "register new account",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.Account_Register"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created new account successfully",
                        "schema": {
                            "$ref": "#/definitions/common.success_response"
                        }
                    },
                    "400": {
                        "description": "Create failure",
                        "schema": {
                            "$ref": "#/definitions/common.error_response"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login-google-account": {
            "post": {
                "description": "Get user information from Google using OAuth2 token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get user information from Google",
                "parameters": [
                    {
                        "description": "OAuth2 token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/oauth.TokenResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User information retrieved successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.success_response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": true
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.error_response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.error_response"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "Log in to the account with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "login to account",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.RequestLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful login",
                        "schema": {
                            "$ref": "#/definitions/common.success_response"
                        }
                    },
                    "400": {
                        "description": "login failure",
                        "schema": {
                            "$ref": "#/definitions/common.error_response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "account.AccountFilter": {
            "type": "object",
            "properties": {
                "account_role": {
                    "type": "string",
                    "enum": [
                        "HR",
                        "ADMIN"
                    ]
                },
                "created_time_from": {
                    "type": "string"
                },
                "created_time_to": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
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
        },
        "account.Account_Register": {
            "type": "object",
            "properties": {
                "account_role": {
                    "enum": [
                        "HR",
                        "ADMIN"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/account.EnumAccountRole"
                        }
                    ]
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "account.EnumAccountRole": {
            "type": "string",
            "enum": [
                "HR",
                "ADMIN"
            ],
            "x-enum-varnames": [
                "HR",
                "ADMIN"
            ]
        },
        "account.RequestLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "common.error_response": {
            "type": "object",
            "properties": {
                "log": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status-code": {
                    "type": "integer"
                }
            }
        },
        "common.success_response": {
            "type": "object",
            "properties": {
                "data": {},
                "filter": {},
                "message": {
                    "type": "string"
                },
                "paging": {},
                "status": {
                    "type": "integer"
                }
            }
        },
        "oauth.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expiry": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Human Resources Management System",
	Description:      "Human Resource Management (HRM)\nServer is a comprehensive software or hardware system designed to manage all aspects of personnel-related activities within an organization.\nIt encompasses employee information management, timekeeping, payroll, training and development, performance management, document management, employee interaction, security, compliance.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
