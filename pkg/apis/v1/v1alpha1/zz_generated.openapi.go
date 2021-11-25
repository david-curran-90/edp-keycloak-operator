//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.Composite":               schema_pkg_apis_v1_v1alpha1_Composite(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.Keycloak":                schema_pkg_apis_v1_v1alpha1_Keycloak(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakClient":          schema_pkg_apis_v1_v1alpha1_KeycloakClient(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakClientSpec":      schema_pkg_apis_v1_v1alpha1_KeycloakClientSpec(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakClientStatus":    schema_pkg_apis_v1_v1alpha1_KeycloakClientStatus(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealm":           schema_pkg_apis_v1_v1alpha1_KeycloakRealm(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmRole":       schema_pkg_apis_v1_v1alpha1_KeycloakRealmRole(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmRoleSpec":   schema_pkg_apis_v1_v1alpha1_KeycloakRealmRoleSpec(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmRoleStatus": schema_pkg_apis_v1_v1alpha1_KeycloakRealmRoleStatus(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmSpec":       schema_pkg_apis_v1_v1alpha1_KeycloakRealmSpec(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmStatus":     schema_pkg_apis_v1_v1alpha1_KeycloakRealmStatus(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakSpec":            schema_pkg_apis_v1_v1alpha1_KeycloakSpec(ref),
		"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakStatus":          schema_pkg_apis_v1_v1alpha1_KeycloakStatus(ref),
	}
}

func schema_pkg_apis_v1_v1alpha1_Composite(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_pkg_apis_v1_v1alpha1_Keycloak(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Keycloak is the Schema for the keycloaks API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakSpec", "github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakClient(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakClient is the Schema for the keycloakclients API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakClientSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakClientStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakClientSpec", "github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakClientStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakClientSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakClientSpec defines the desired state of KeycloakClient",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"targetRealm": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"realmRoles": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.RealmRole"),
									},
								},
							},
						},
					},
					"public": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"clientId": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"webUrl": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"protocol": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"attributes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"directAccess": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"advancedProtocolMappers": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"clientRoles": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"audRequired": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"protocolMappers": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.ProtocolMapper"),
									},
								},
							},
						},
					},
				},
				Required: []string{"targetRealm", "secret", "public", "clientId", "webUrl", "directAccess", "advancedProtocolMappers", "audRequired", "protocolMappers"},
			},
		},
		Dependencies: []string{
			"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.ProtocolMapper", "github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.RealmRole"},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakClientStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakClientStatus defines the observed state of KeycloakClient",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"value": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"id": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"value", "id"},
			},
		},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakRealm(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakRealm is the Schema for the keycloakrealms API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmSpec", "github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakRealmRole(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmRoleSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmRoleStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmRoleSpec", "github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.KeycloakRealmRoleStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakRealmRoleSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"realm": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "realm name",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"attributes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Default: "",
													Type:    []string{"string"},
													Format:  "",
												},
											},
										},
									},
								},
							},
						},
					},
					"composite": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"composites": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.Composite"),
									},
								},
							},
						},
					},
				},
				Required: []string{"name", "realm", "description", "attributes", "composite", "composites"},
			},
		},
		Dependencies: []string{
			"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.Composite"},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakRealmRoleStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"id"},
			},
		},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakRealmSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakRealmSpec defines the desired state of KeycloakRealm",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"realmName": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"keycloakOwner": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ssoRealmName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ssoRealmEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"users": {
						SchemaProps: spec.SchemaProps{
							Description: "default (nil, not set) must be true",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.User"),
									},
								},
							},
						},
					},
				},
				Required: []string{"realmName"},
			},
		},
		Dependencies: []string{
			"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.User"},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakRealmStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakRealmStatus defines the observed state of KeycloakRealm",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"available": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakSpec defines the desired state of Keycloak",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"realmName": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"ssoRealmName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"users": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.User"),
									},
								},
							},
						},
					},
				},
				Required: []string{"url", "secret", "realmName"},
			},
		},
		Dependencies: []string{
			"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1.User"},
	}
}

func schema_pkg_apis_v1_v1alpha1_KeycloakStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakStatus defines the observed state of Keycloak",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"connected": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
				},
				Required: []string{"connected"},
			},
		},
	}
}
