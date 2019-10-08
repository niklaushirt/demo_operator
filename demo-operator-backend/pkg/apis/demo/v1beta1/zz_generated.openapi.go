// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"demo-operator-backend/pkg/apis/demo/v1beta1.MyDemoBackend":       schema_pkg_apis_demo_v1beta1_MyDemoBackend(ref),
		"demo-operator-backend/pkg/apis/demo/v1beta1.MyDemoBackendStatus": schema_pkg_apis_demo_v1beta1_MyDemoBackendStatus(ref),
	}
}

func schema_pkg_apis_demo_v1beta1_MyDemoBackend(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MyDemoBackend is the Schema for the mydemobackends API",
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
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("demo-operator-backend/pkg/apis/demo/v1beta1.MyDemoBackendSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("demo-operator-backend/pkg/apis/demo/v1beta1.MyDemoBackendStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"demo-operator-backend/pkg/apis/demo/v1beta1.MyDemoBackendSpec", "demo-operator-backend/pkg/apis/demo/v1beta1.MyDemoBackendStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_demo_v1beta1_MyDemoBackendStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MyDemoBackendStatus defines the observed state of MyDemoBackend",
				Properties: map[string]spec.Schema{
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"nodes"},
			},
		},
		Dependencies: []string{},
	}
}
