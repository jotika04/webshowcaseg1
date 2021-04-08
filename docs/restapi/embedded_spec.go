// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Showcase",
    "version": "1.0.0"
  },
  "host": "showcase.com",
  "basePath": "/v1",
  "paths": {
    "/login": {
      "post": {
        "description": "Returns token for authorized User",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Login",
        "parameters": [
          {
            "description": "Login Payload",
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful login",
            "schema": {
              "$ref": "#/definitions/LoginSuccess"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "description": "To register a new user",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Register",
        "parameters": [
          {
            "description": "Registeration Payload",
            "name": "signup",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful registeration",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "LoginInfo": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "RegisterUser": {
      "type": "object",
      "required": [
        "userFirstName",
        "userLastName",
        "email",
        "username",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "userFirstName": {
          "type": "string"
        },
        "userLastName": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "SuccessResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Showcase",
    "version": "1.0.0"
  },
  "host": "showcase.com",
  "basePath": "/v1",
  "paths": {
    "/login": {
      "post": {
        "description": "Returns token for authorized User",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Login",
        "parameters": [
          {
            "description": "Login Payload",
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful login",
            "schema": {
              "$ref": "#/definitions/LoginSuccess"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "description": "To register a new user",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Register",
        "parameters": [
          {
            "description": "Registeration Payload",
            "name": "signup",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful registeration",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "LoginInfo": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "RegisterUser": {
      "type": "object",
      "required": [
        "userFirstName",
        "userLastName",
        "email",
        "username",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "userFirstName": {
          "type": "string"
        },
        "userLastName": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "SuccessResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    }
  }
}`))
}
