// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ChePluginLocation":                  schema_pkg_apis_workspaces_v1alpha1_ChePluginLocation(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspace":                       schema_pkg_apis_workspaces_v1alpha1_DevWorkspace(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceSpec":                   schema_pkg_apis_workspaces_v1alpha1_DevWorkspaceSpec(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceStatus":                 schema_pkg_apis_workspaces_v1alpha1_DevWorkspaceStatus(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceTemplate":               schema_pkg_apis_workspaces_v1alpha1_DevWorkspaceTemplate(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceTemplateSpec":           schema_pkg_apis_workspaces_v1alpha1_DevWorkspaceTemplateSpec(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.K8sLikeComponentLocation":           schema_pkg_apis_workspaces_v1alpha1_K8sLikeComponentLocation(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ParentLocation":                     schema_pkg_apis_workspaces_v1alpha1_ParentLocation(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.PolymorphicCommand":                 schema_pkg_apis_workspaces_v1alpha1_PolymorphicCommand(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.PolymorphicComponent":               schema_pkg_apis_workspaces_v1alpha1_PolymorphicComponent(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ProjectSource":                      schema_pkg_apis_workspaces_v1alpha1_ProjectSource(ref),
		"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.VscodeConfigurationCommandLocation": schema_pkg_apis_workspaces_v1alpha1_VscodeConfigurationCommandLocation(ref),
	}
}

func schema_pkg_apis_workspaces_v1alpha1_ChePluginLocation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"locationType": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of plugin location",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"registryEntry": {
						SchemaProps: spec.SchemaProps{
							Description: "Location of an entry inside a plugin registry",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.RegistryEntryPluginLocation"),
						},
					},
					"uri": {
						SchemaProps: spec.SchemaProps{
							Description: "Location defined as an URI",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "locationType",
							"fields-to-discriminateBy": map[string]interface{}{
								"registryEntry": "RegistryEntry",
								"uri":           "Uri",
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.RegistryEntryPluginLocation"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_DevWorkspace(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DevWorkspace is the Schema for the devworkspaces API",
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
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceSpec", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_DevWorkspaceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DevWorkspaceSpec defines the desired state of DevWorkspace",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"started": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"routingClass": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceTemplateSpec"),
						},
					},
				},
				Required: []string{"started"},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceTemplateSpec"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_DevWorkspaceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DevWorkspaceStatus defines the observed state of DevWorkspace",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"workspaceId": {
						SchemaProps: spec.SchemaProps{
							Description: "Id of the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"mainIdeUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "URL at which the Worksace Editor can be joined",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"additionalInfo": {
						SchemaProps: spec.SchemaProps{
							Description: "AdditionalInfo",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
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
				Required: []string{"workspaceId"},
			},
		},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_DevWorkspaceTemplate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DevWorkspaceTemplate is the Schema for the devworkspacetemplates API",
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
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceTemplateSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.DevWorkspaceTemplateSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_DevWorkspaceTemplateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Structure of the workspace. This is also the specification of a workspace template.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"parent": {
						SchemaProps: spec.SchemaProps{
							Description: "Parent workspace template",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.Parent"),
						},
					},
					"commands": {
						SchemaProps: spec.SchemaProps{
							Description: "Predefined, ready-to-use, workspace-related commands",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.Command"),
									},
								},
							},
						},
					},
					"projects": {
						SchemaProps: spec.SchemaProps{
							Description: "Projects worked on in the workspace, containing names and sources locations",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.Project"),
									},
								},
							},
						},
					},
					"components": {
						SchemaProps: spec.SchemaProps{
							Description: "List of the workspace components, such as editor and plugins, user-provided containers, or other types of components",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.Component"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.Command", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.Component", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.Parent", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.Project"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_K8sLikeComponentLocation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Mandatory name that allows referencing the component in commands, or inside a parent",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"locationType": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of Kubernetes-like location",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "Location in a plugin registry",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"inlined": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the plugin definition",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name"},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "locationType",
							"fields-to-discriminateBy": map[string]interface{}{
								"inlined": "Inlined",
								"name":    "Name",
								"url":     "Url",
							},
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_ParentLocation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Location from where the parent workspace structure is retrieved",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"locationType": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of parent location",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"uri": {
						SchemaProps: spec.SchemaProps{
							Description: "Uri of a Devfile yaml file",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"registryEntry": {
						SchemaProps: spec.SchemaProps{
							Description: "Entry in a registry (base URL + ID) that contains a Devfile yaml file",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.RegistryEntryParentLocation"),
						},
					},
					"kubernetes": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.KubernetesCustomResourceParentLocation"),
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "locationType",
							"fields-to-discriminateBy": map[string]interface{}{
								"kubernetes":    "Kubernetes",
								"registryEntry": "RegistryEntry",
								"uri":           "Uri",
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.KubernetesCustomResourceParentLocation", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.RegistryEntryParentLocation"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_PolymorphicCommand(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of workspace command",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"exec": {
						SchemaProps: spec.SchemaProps{
							Description: "Exec command",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ExecCommand"),
						},
					},
					"vscodeTask": {
						SchemaProps: spec.SchemaProps{
							Description: "VscodeTask command",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.VscodeConfigurationCommand"),
						},
					},
					"vscodeLaunch": {
						SchemaProps: spec.SchemaProps{
							Description: "VscodeLaunch command",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.VscodeConfigurationCommand"),
						},
					},
					"composite": {
						SchemaProps: spec.SchemaProps{
							Description: "Composite command",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CompositeCommand"),
						},
					},
					"custom": {
						SchemaProps: spec.SchemaProps{
							Description: "Custom command",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CustomCommand"),
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "type",
							"fields-to-discriminateBy": map[string]interface{}{
								"composite":    "Composite",
								"custom":       "Custom",
								"exec":         "Exec",
								"vscodeLaunch": "VscodeLaunch",
								"vscodeTask":   "VscodeTask",
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CompositeCommand", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CustomCommand", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ExecCommand", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.VscodeConfigurationCommand"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_PolymorphicComponent(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of project source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"container": {
						SchemaProps: spec.SchemaProps{
							Description: "Container component",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ContainerComponent"),
						},
					},
					"cheEditor": {
						SchemaProps: spec.SchemaProps{
							Description: "CheEditor component",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CheEditorComponent"),
						},
					},
					"chePlugin": {
						SchemaProps: spec.SchemaProps{
							Description: "ChePlugin component",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ChePluginComponent"),
						},
					},
					"kubernetes": {
						SchemaProps: spec.SchemaProps{
							Description: "Kubernetes component",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.KubernetesComponent"),
						},
					},
					"openshift": {
						SchemaProps: spec.SchemaProps{
							Description: "Openshift component",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.OpenshiftComponent"),
						},
					},
					"custom": {
						SchemaProps: spec.SchemaProps{
							Description: "Custom component",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CustomComponent"),
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "type",
							"fields-to-discriminateBy": map[string]interface{}{
								"cheEditor":  "CheEditor",
								"chePlugin":  "ChePlugin",
								"container":  "Container",
								"custom":     "Custom",
								"kubernetes": "Kubernetes",
								"openshift":  "Openshift",
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CheEditorComponent", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ChePluginComponent", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ContainerComponent", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CustomComponent", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.KubernetesComponent", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.OpenshiftComponent"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_ProjectSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"sourceType": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of project source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"git": {
						SchemaProps: spec.SchemaProps{
							Description: "Project's Git source",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.GitProjectSource"),
						},
					},
					"github": {
						SchemaProps: spec.SchemaProps{
							Description: "Project's GitHub source",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.GithubProjectSource"),
						},
					},
					"zip": {
						SchemaProps: spec.SchemaProps{
							Description: "Project's Zip source",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ZipProjectSource"),
						},
					},
					"custom": {
						SchemaProps: spec.SchemaProps{
							Description: "Project's Custom source",
							Ref:         ref("github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CustomProjectSource"),
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "sourceType",
							"fields-to-discriminateBy": map[string]interface{}{
								"custom": "Custom",
								"git":    "Git",
								"github": "Github",
								"zip":    "Zip",
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.CustomProjectSource", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.GitProjectSource", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.GithubProjectSource", "github.com/che-incubator/devworkspace-api/pkg/apis/workspaces/v1alpha1.ZipProjectSource"},
	}
}

func schema_pkg_apis_workspaces_v1alpha1_VscodeConfigurationCommandLocation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"locationType": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of Vscode configuration command location",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "Location as an absolute of relative URL",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"inlined": {
						SchemaProps: spec.SchemaProps{
							Description: "Embedded content of the vscode configuration file",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "locationType",
							"fields-to-discriminateBy": map[string]interface{}{
								"inlined": "Inlined",
								"url":     "Url",
							},
						},
					},
				},
			},
		},
	}
}
