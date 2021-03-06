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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the initial REST API for RIC subscription",
    "title": "RIC subscription",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "0.0.1"
  },
  "host": "hostname",
  "basePath": "/ric/v1",
  "paths": {
    "/config": {
      "get": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "xapp"
        ],
        "summary": "Returns the configuration of all xapps",
        "operationId": "getXappConfigList",
        "responses": {
          "200": {
            "description": "successful query of xApp config",
            "schema": {
              "$ref": "#/definitions/XappConfigList"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/subscriptions": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "query"
        ],
        "summary": "Returns list of subscriptions",
        "operationId": "getAllSubscriptions",
        "responses": {
          "200": {
            "description": "successful query of subscriptions",
            "schema": {
              "$ref": "#/definitions/SubscriptionList"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/subscriptions/policy": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Subscribe and send \"POLICY\" message to RAN to execute a specific POLICY during call processing in RAN after each occurrence of a defined SUBSCRIPTION",
        "operationId": "subscribePolicy",
        "parameters": [
          {
            "description": "Subscription policy parameters",
            "name": "PolicyParams",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PolicyParams"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Subscription successfully created",
            "schema": {
              "$ref": "#/definitions/SubscriptionResponse"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/subscriptions/report": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "report"
        ],
        "summary": "Subscribe a list of X2AP event triggers to receive \"REPORT\" messages sent by RAN or Subscribe to receive the content of gNB NRT table in REPORT message sent by RAN",
        "operationId": "subscribeReport",
        "parameters": [
          {
            "description": "Subscription report parameters",
            "name": "ReportParams",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ReportParams"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Subscription successfully created",
            "schema": {
              "$ref": "#/definitions/SubscriptionResponse"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/subscriptions/{subscriptionId}": {
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "common"
        ],
        "summary": "Unsubscribe X2AP events from Subscription Manager",
        "operationId": "Unsubscribe",
        "parameters": [
          {
            "type": "string",
            "description": "The subscriptionId received in the Subscription Response",
            "name": "subscriptionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Operation done successfully"
          },
          "400": {
            "description": "Invalid subscriptionId supplied"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "ActionParameters": {
      "type": "object",
      "required": [
        "ActionParameterID",
        "ActionParameterValue"
      ],
      "properties": {
        "ActionParameterID": {
          "type": "integer"
        },
        "ActionParameterValue": {
          "type": "boolean"
        }
      }
    },
    "ConfigMetadata": {
      "type": "object",
      "required": [
        "xappName",
        "configType"
      ],
      "properties": {
        "configType": {
          "description": "The type of the content",
          "type": "string",
          "enum": [
            "json",
            "xml",
            "other"
          ]
        },
        "xappName": {
          "description": "Name of the xApp",
          "type": "string"
        }
      }
    },
    "EventTrigger": {
      "type": "object",
      "properties": {
        "ENBId": {
          "type": "string"
        },
        "InterfaceDirection": {
          "type": "integer"
        },
        "PlmnId": {
          "type": "string"
        },
        "ProcedureCode": {
          "type": "integer"
        },
        "TriggerNature": {
          "type": "string",
          "enum": [
            "now",
            "on change"
          ]
        },
        "TypeOfMessage": {
          "type": "integer"
        }
      }
    },
    "EventTriggerList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/EventTrigger"
      }
    },
    "Format1ActionDefinition": {
      "type": "object",
      "required": [
        "StyleID",
        "ActionParameters"
      ],
      "properties": {
        "ActionParameters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ActionParameters"
          }
        },
        "StyleID": {
          "type": "integer"
        }
      }
    },
    "Format2ActionDefinition": {
      "type": "object",
      "required": [
        "RANUeGroupParameters"
      ],
      "properties": {
        "RANUeGroupParameters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RANUeGroupList"
          }
        }
      }
    },
    "ImperativePolicyDefinition": {
      "type": "object",
      "required": [
        "PolicyParameterID",
        "PolicyParameterValue"
      ],
      "properties": {
        "PolicyParameterID": {
          "type": "integer"
        },
        "PolicyParameterValue": {
          "type": "integer"
        }
      }
    },
    "PolicyActionDefinition": {
      "type": "object",
      "properties": {
        "ActionDefinitionFormat2": {
          "$ref": "#/definitions/Format2ActionDefinition"
        }
      }
    },
    "PolicyParams": {
      "type": "object",
      "required": [
        "Meid",
        "RANFunctionID",
        "ClientEndpoint",
        "EventTriggers",
        "PolicyActionDefinitions"
      ],
      "properties": {
        "ClientEndpoint": {
          "type": "string"
        },
        "EventTriggers": {
          "$ref": "#/definitions/EventTriggerList"
        },
        "Meid": {
          "type": "string"
        },
        "PolicyActionDefinitions": {
          "$ref": "#/definitions/PolicyActionDefinition"
        },
        "RANFunctionID": {
          "type": "integer"
        }
      }
    },
    "RANUeGroupList": {
      "type": "object",
      "required": [
        "RANUeGroupID",
        "RANUeGroupDefinition",
        "RANImperativePolicy"
      ],
      "properties": {
        "RANImperativePolicy": {
          "$ref": "#/definitions/ImperativePolicyDefinition"
        },
        "RANUeGroupDefinition": {
          "$ref": "#/definitions/RANUeGroupParams"
        },
        "RANUeGroupID": {
          "type": "integer"
        }
      }
    },
    "RANUeGroupParams": {
      "type": "object",
      "required": [
        "RANParameterID",
        "RANParameterValue"
      ],
      "properties": {
        "RANParameterID": {
          "type": "integer"
        },
        "RANParameterTestCondition": {
          "type": "string",
          "enum": [
            "equal",
            "greaterthan",
            "lessthan",
            "contains",
            "present"
          ]
        },
        "RANParameterValue": {
          "type": "integer"
        }
      }
    },
    "ReportActionDefinition": {
      "type": "object",
      "properties": {
        "ActionDefinitionFormat1": {
          "$ref": "#/definitions/Format1ActionDefinition"
        }
      }
    },
    "ReportParams": {
      "type": "object",
      "required": [
        "RANFunctionID",
        "ClientEndpoint",
        "EventTriggers"
      ],
      "properties": {
        "ClientEndpoint": {
          "type": "string"
        },
        "EventTriggers": {
          "$ref": "#/definitions/EventTriggerList"
        },
        "Meid": {
          "type": "string"
        },
        "RANFunctionID": {
          "type": "integer"
        },
        "ReportActionDefinitions": {
          "$ref": "#/definitions/ReportActionDefinition"
        }
      }
    },
    "SubscriptionData": {
      "type": "object",
      "properties": {
        "Endpoint": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Meid": {
          "type": "string"
        },
        "SubscriptionId": {
          "type": "integer"
        }
      }
    },
    "SubscriptionInstance": {
      "type": "object",
      "required": [
        "RequestorId",
        "InstanceId"
      ],
      "properties": {
        "InstanceId": {
          "type": "integer"
        },
        "RequestorId": {
          "type": "integer"
        }
      }
    },
    "SubscriptionList": {
      "description": "A list of subscriptions",
      "type": "array",
      "items": {
        "$ref": "#/definitions/SubscriptionData"
      }
    },
    "SubscriptionResponse": {
      "type": "object",
      "required": [
        "SubscriptionId",
        "SubscriptionInstances"
      ],
      "properties": {
        "SubscriptionId": {
          "type": "string"
        },
        "SubscriptionInstances": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/SubscriptionInstance"
          }
        }
      }
    },
    "SubscriptionType": {
      "type": "string",
      "enum": [
        "insert",
        "policy",
        "report"
      ]
    },
    "XAppConfig": {
      "type": "object",
      "required": [
        "metadata",
        "config"
      ],
      "properties": {
        "config": {
          "description": "Configuration in JSON format",
          "type": "object"
        },
        "metadata": {
          "$ref": "#/definitions/ConfigMetadata"
        }
      }
    },
    "XappConfigList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/XAppConfig"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the initial REST API for RIC subscription",
    "title": "RIC subscription",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "0.0.1"
  },
  "host": "hostname",
  "basePath": "/ric/v1",
  "paths": {
    "/config": {
      "get": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "xapp"
        ],
        "summary": "Returns the configuration of all xapps",
        "operationId": "getXappConfigList",
        "responses": {
          "200": {
            "description": "successful query of xApp config",
            "schema": {
              "$ref": "#/definitions/XappConfigList"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/subscriptions": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "query"
        ],
        "summary": "Returns list of subscriptions",
        "operationId": "getAllSubscriptions",
        "responses": {
          "200": {
            "description": "successful query of subscriptions",
            "schema": {
              "$ref": "#/definitions/SubscriptionList"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/subscriptions/policy": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Subscribe and send \"POLICY\" message to RAN to execute a specific POLICY during call processing in RAN after each occurrence of a defined SUBSCRIPTION",
        "operationId": "subscribePolicy",
        "parameters": [
          {
            "description": "Subscription policy parameters",
            "name": "PolicyParams",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PolicyParams"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Subscription successfully created",
            "schema": {
              "$ref": "#/definitions/SubscriptionResponse"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/subscriptions/report": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "report"
        ],
        "summary": "Subscribe a list of X2AP event triggers to receive \"REPORT\" messages sent by RAN or Subscribe to receive the content of gNB NRT table in REPORT message sent by RAN",
        "operationId": "subscribeReport",
        "parameters": [
          {
            "description": "Subscription report parameters",
            "name": "ReportParams",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ReportParams"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Subscription successfully created",
            "schema": {
              "$ref": "#/definitions/SubscriptionResponse"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/subscriptions/{subscriptionId}": {
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "common"
        ],
        "summary": "Unsubscribe X2AP events from Subscription Manager",
        "operationId": "Unsubscribe",
        "parameters": [
          {
            "type": "string",
            "description": "The subscriptionId received in the Subscription Response",
            "name": "subscriptionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Operation done successfully"
          },
          "400": {
            "description": "Invalid subscriptionId supplied"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "ActionParameters": {
      "type": "object",
      "required": [
        "ActionParameterID",
        "ActionParameterValue"
      ],
      "properties": {
        "ActionParameterID": {
          "type": "integer"
        },
        "ActionParameterValue": {
          "type": "boolean"
        }
      }
    },
    "ConfigMetadata": {
      "type": "object",
      "required": [
        "xappName",
        "configType"
      ],
      "properties": {
        "configType": {
          "description": "The type of the content",
          "type": "string",
          "enum": [
            "json",
            "xml",
            "other"
          ]
        },
        "xappName": {
          "description": "Name of the xApp",
          "type": "string"
        }
      }
    },
    "EventTrigger": {
      "type": "object",
      "properties": {
        "ENBId": {
          "type": "string"
        },
        "InterfaceDirection": {
          "type": "integer"
        },
        "PlmnId": {
          "type": "string"
        },
        "ProcedureCode": {
          "type": "integer"
        },
        "TriggerNature": {
          "type": "string",
          "enum": [
            "now",
            "on change"
          ]
        },
        "TypeOfMessage": {
          "type": "integer"
        }
      }
    },
    "EventTriggerList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/EventTrigger"
      }
    },
    "Format1ActionDefinition": {
      "type": "object",
      "required": [
        "StyleID",
        "ActionParameters"
      ],
      "properties": {
        "ActionParameters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ActionParameters"
          }
        },
        "StyleID": {
          "type": "integer"
        }
      }
    },
    "Format2ActionDefinition": {
      "type": "object",
      "required": [
        "RANUeGroupParameters"
      ],
      "properties": {
        "RANUeGroupParameters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RANUeGroupList"
          }
        }
      }
    },
    "ImperativePolicyDefinition": {
      "type": "object",
      "required": [
        "PolicyParameterID",
        "PolicyParameterValue"
      ],
      "properties": {
        "PolicyParameterID": {
          "type": "integer"
        },
        "PolicyParameterValue": {
          "type": "integer"
        }
      }
    },
    "PolicyActionDefinition": {
      "type": "object",
      "properties": {
        "ActionDefinitionFormat2": {
          "$ref": "#/definitions/Format2ActionDefinition"
        }
      }
    },
    "PolicyParams": {
      "type": "object",
      "required": [
        "Meid",
        "RANFunctionID",
        "ClientEndpoint",
        "EventTriggers",
        "PolicyActionDefinitions"
      ],
      "properties": {
        "ClientEndpoint": {
          "type": "string"
        },
        "EventTriggers": {
          "$ref": "#/definitions/EventTriggerList"
        },
        "Meid": {
          "type": "string"
        },
        "PolicyActionDefinitions": {
          "$ref": "#/definitions/PolicyActionDefinition"
        },
        "RANFunctionID": {
          "type": "integer"
        }
      }
    },
    "RANUeGroupList": {
      "type": "object",
      "required": [
        "RANUeGroupID",
        "RANUeGroupDefinition",
        "RANImperativePolicy"
      ],
      "properties": {
        "RANImperativePolicy": {
          "$ref": "#/definitions/ImperativePolicyDefinition"
        },
        "RANUeGroupDefinition": {
          "$ref": "#/definitions/RANUeGroupParams"
        },
        "RANUeGroupID": {
          "type": "integer"
        }
      }
    },
    "RANUeGroupParams": {
      "type": "object",
      "required": [
        "RANParameterID",
        "RANParameterValue"
      ],
      "properties": {
        "RANParameterID": {
          "type": "integer"
        },
        "RANParameterTestCondition": {
          "type": "string",
          "enum": [
            "equal",
            "greaterthan",
            "lessthan",
            "contains",
            "present"
          ]
        },
        "RANParameterValue": {
          "type": "integer"
        }
      }
    },
    "ReportActionDefinition": {
      "type": "object",
      "properties": {
        "ActionDefinitionFormat1": {
          "$ref": "#/definitions/Format1ActionDefinition"
        }
      }
    },
    "ReportParams": {
      "type": "object",
      "required": [
        "RANFunctionID",
        "ClientEndpoint",
        "EventTriggers"
      ],
      "properties": {
        "ClientEndpoint": {
          "type": "string"
        },
        "EventTriggers": {
          "$ref": "#/definitions/EventTriggerList"
        },
        "Meid": {
          "type": "string"
        },
        "RANFunctionID": {
          "type": "integer"
        },
        "ReportActionDefinitions": {
          "$ref": "#/definitions/ReportActionDefinition"
        }
      }
    },
    "SubscriptionData": {
      "type": "object",
      "properties": {
        "Endpoint": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Meid": {
          "type": "string"
        },
        "SubscriptionId": {
          "type": "integer"
        }
      }
    },
    "SubscriptionInstance": {
      "type": "object",
      "required": [
        "RequestorId",
        "InstanceId"
      ],
      "properties": {
        "InstanceId": {
          "type": "integer"
        },
        "RequestorId": {
          "type": "integer"
        }
      }
    },
    "SubscriptionList": {
      "description": "A list of subscriptions",
      "type": "array",
      "items": {
        "$ref": "#/definitions/SubscriptionData"
      }
    },
    "SubscriptionResponse": {
      "type": "object",
      "required": [
        "SubscriptionId",
        "SubscriptionInstances"
      ],
      "properties": {
        "SubscriptionId": {
          "type": "string"
        },
        "SubscriptionInstances": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/SubscriptionInstance"
          }
        }
      }
    },
    "SubscriptionType": {
      "type": "string",
      "enum": [
        "insert",
        "policy",
        "report"
      ]
    },
    "XAppConfig": {
      "type": "object",
      "required": [
        "metadata",
        "config"
      ],
      "properties": {
        "config": {
          "description": "Configuration in JSON format",
          "type": "object"
        },
        "metadata": {
          "$ref": "#/definitions/ConfigMetadata"
        }
      }
    },
    "XappConfigList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/XAppConfig"
      }
    }
  }
}`))
}
