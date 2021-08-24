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
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The GA Exporter API makes use of the [Core Reporting API version \n3.0](https://developers.google.com/analytics/devguides/reporting/core/v3/reference) \nto access information about visitors to a web property. It pulls data from \nan existing and preconfigured \n[reporting view](https://support.google.com/analytics/answer/2649553?hl=en) \nin a Google Analytics account.\n\n# Errors\nThe API uses standard HTTP status codes to indicate the success or failure\nof the API call. The body of the response will be JSON in the following\nformat:\n` + "`" + `` + "`" + `` + "`" + `\n{\n  \"code\": 405,\n  \"message\": \"method POST is not allowed, but [GET] are\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n",
    "title": "Google Analytics Exporter API",
    "contact": {
      "name": "Ga Exporter API mantainers",
      "url": "https://github.com/quintsys/ga-exporter",
      "email": "contact@quintsys.com"
    },
    "version": "0.0.1"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "Visitors"
        ],
        "operationId": "VisitorList",
        "responses": {
          "200": {
            "description": "Returns data about visitors\n",
            "schema": {
              "type": "array",
              "maxItems": 5000,
              "items": {
                "$ref": "#/definitions/visitor"
              }
            }
          },
          "405": {
            "description": "Method Not Allowed",
            "schema": {
              "$ref": "#/definitions/error_method_not_allowed"
            }
          },
          "406": {
            "description": "Not Acceptable",
            "schema": {
              "$ref": "#/definitions/error_not_acceptable"
            }
          },
          "429": {
            "description": "Too many request",
            "schema": {
              "$ref": "#/definitions/error_too_many_requests"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error_internal_server_error"
            }
          },
          "503": {
            "description": "Service Unavailable",
            "schema": {
              "$ref": "#/definitions/error_service_unavailable"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64",
          "maximum": 599,
          "minimum": 400
        },
        "message": {
          "type": "string",
          "maxLength": 255,
          "minLength": 1
        }
      }
    },
    "error_internal_server_error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 500
        },
        "message": {
          "type": "string",
          "example": "Internal Server Error"
        }
      }
    },
    "error_method_not_allowed": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 405
        },
        "message": {
          "type": "string",
          "example": "method POST is not allowed, but [GET] are"
        }
      }
    },
    "error_not_acceptable": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 406
        },
        "message": {
          "type": "string",
          "example": "unsupported media type requested, only [application/json] are available"
        }
      }
    },
    "error_service_unavailable": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 503
        },
        "message": {
          "type": "string",
          "example": "service is not available"
        }
      }
    },
    "error_too_many_requests": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 429
        },
        "message": {
          "type": "string",
          "example": "too many requests"
        }
      }
    },
    "visitor": {
      "type": "object",
      "required": [
        "clientId"
      ],
      "properties": {
        "adContent": {
          "type": "string",
          "maxLength": 255
        },
        "adGroup": {
          "type": "string",
          "maxLength": 255
        },
        "adMatchedQuery": {
          "type": "string",
          "maxLength": 255
        },
        "campaign": {
          "type": "string",
          "maxLength": 100
        },
        "clientId": {
          "type": "string",
          "maxLength": 36,
          "minLength": 1,
          "x-order": 0
        },
        "keyword": {
          "type": "string",
          "maxLength": 255
        },
        "medium": {
          "type": "string",
          "maxLength": 50
        },
        "source": {
          "type": "string",
          "maxLength": 50
        }
      }
    }
  },
  "tags": [
    {
      "description": "Visitors are extracted from a configured Google Analytics view.\n - Each unique visitor is represented by a unique ` + "`" + `ClientId` + "`" + `.\n - The list of dimensions to be exported is predefined and can't be \n modified using configuration parameters.\n",
      "name": "Visitors"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The GA Exporter API makes use of the [Core Reporting API version \n3.0](https://developers.google.com/analytics/devguides/reporting/core/v3/reference) \nto access information about visitors to a web property. It pulls data from \nan existing and preconfigured \n[reporting view](https://support.google.com/analytics/answer/2649553?hl=en) \nin a Google Analytics account.\n\n# Errors\nThe API uses standard HTTP status codes to indicate the success or failure\nof the API call. The body of the response will be JSON in the following\nformat:\n` + "`" + `` + "`" + `` + "`" + `\n{\n  \"code\": 405,\n  \"message\": \"method POST is not allowed, but [GET] are\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n",
    "title": "Google Analytics Exporter API",
    "contact": {
      "name": "Ga Exporter API mantainers",
      "url": "https://github.com/quintsys/ga-exporter",
      "email": "contact@quintsys.com"
    },
    "version": "0.0.1"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "Visitors"
        ],
        "operationId": "VisitorList",
        "responses": {
          "200": {
            "description": "Returns data about visitors\n",
            "schema": {
              "type": "array",
              "maxItems": 5000,
              "items": {
                "$ref": "#/definitions/visitor"
              }
            }
          },
          "405": {
            "description": "Method Not Allowed",
            "schema": {
              "$ref": "#/definitions/error_method_not_allowed"
            }
          },
          "406": {
            "description": "Not Acceptable",
            "schema": {
              "$ref": "#/definitions/error_not_acceptable"
            }
          },
          "429": {
            "description": "Too many request",
            "schema": {
              "$ref": "#/definitions/error_too_many_requests"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error_internal_server_error"
            }
          },
          "503": {
            "description": "Service Unavailable",
            "schema": {
              "$ref": "#/definitions/error_service_unavailable"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64",
          "maximum": 599,
          "minimum": 400
        },
        "message": {
          "type": "string",
          "maxLength": 255,
          "minLength": 1
        }
      }
    },
    "error_internal_server_error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 500
        },
        "message": {
          "type": "string",
          "example": "Internal Server Error"
        }
      }
    },
    "error_method_not_allowed": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 405
        },
        "message": {
          "type": "string",
          "example": "method POST is not allowed, but [GET] are"
        }
      }
    },
    "error_not_acceptable": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 406
        },
        "message": {
          "type": "string",
          "example": "unsupported media type requested, only [application/json] are available"
        }
      }
    },
    "error_service_unavailable": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 503
        },
        "message": {
          "type": "string",
          "example": "service is not available"
        }
      }
    },
    "error_too_many_requests": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 429
        },
        "message": {
          "type": "string",
          "example": "too many requests"
        }
      }
    },
    "visitor": {
      "type": "object",
      "required": [
        "clientId"
      ],
      "properties": {
        "adContent": {
          "type": "string",
          "maxLength": 255,
          "minLength": 0
        },
        "adGroup": {
          "type": "string",
          "maxLength": 255,
          "minLength": 0
        },
        "adMatchedQuery": {
          "type": "string",
          "maxLength": 255,
          "minLength": 0
        },
        "campaign": {
          "type": "string",
          "maxLength": 100,
          "minLength": 0
        },
        "clientId": {
          "type": "string",
          "maxLength": 36,
          "minLength": 1,
          "x-order": 0
        },
        "keyword": {
          "type": "string",
          "maxLength": 255,
          "minLength": 0
        },
        "medium": {
          "type": "string",
          "maxLength": 50,
          "minLength": 0
        },
        "source": {
          "type": "string",
          "maxLength": 50,
          "minLength": 0
        }
      }
    }
  },
  "tags": [
    {
      "description": "Visitors are extracted from a configured Google Analytics view.\n - Each unique visitor is represented by a unique ` + "`" + `ClientId` + "`" + `.\n - The list of dimensions to be exported is predefined and can't be \n modified using configuration parameters.\n",
      "name": "Visitors"
    }
  ]
}`))
}
