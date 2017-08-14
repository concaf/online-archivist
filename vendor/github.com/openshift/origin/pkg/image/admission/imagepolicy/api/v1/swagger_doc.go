package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_GroupResource = map[string]string{
	"":         "GroupResource represents a resource in a specific group.",
	"resource": "Resource is the name of an admission resource to process, e.g. 'statefulsets'.",
	"group":    "Group is the name of the group the resource is in, e.g. 'apps'.",
}

func (GroupResource) SwaggerDoc() map[string]string {
	return map_GroupResource
}

var map_ImageCondition = map[string]string{
	"":     "ImageCondition defines the conditions for matching a particular image source. The conditions below are all required (logical AND). If Reject is specified, the condition is false if all conditions match, and true otherwise.",
	"name": "Name is the name of this policy rule for reference. It must be unique across all rules.",
	"ignoreNamespaceOverride": "IgnoreNamespaceOverride prevents this condition from being overridden when the `alpha.image.policy.openshift.io/ignore-rules` is set on a namespace and contains this rule name.",
	"onResources":             "OnResources determines which resources this applies to. Defaults to 'pods' for ImageExecutionPolicyRules.",
	"invertMatch":             "InvertMatch means the value of the condition is logically inverted (true -> false, false -> true).",
	"matchIntegratedRegistry": "MatchIntegratedRegistry will only match image sources that originate from the configured integrated registry.",
	"matchRegistries":         "MatchRegistries will match image references that point to the provided registries. The image registry must match at least one of these strings.",
	"skipOnResolutionFailure": "SkipOnResolutionFailure allows the subsequent conditions to be bypassed if the integrated registry does not have access to image metadata (no image exists matching the image digest).",
	"matchDockerImageLabels":  "MatchDockerImageLabels checks against the resolved image for the presence of a Docker label. All conditions must match.",
	"matchImageLabels":        "MatchImageLabels checks against the resolved image for a label. All conditions must match.",
	"matchImageAnnotations":   "MatchImageAnnotations checks against the resolved image for an annotation. All conditions must match.",
}

func (ImageCondition) SwaggerDoc() map[string]string {
	return map_ImageCondition
}

var map_ImageExecutionPolicyRule = map[string]string{
	"":       "ImageExecutionPolicyRule determines whether a provided image may be used on the platform.",
	"reject": "Reject means this rule, if it matches the condition, will cause an immediate failure. No other rules will be considered.",
}

func (ImageExecutionPolicyRule) SwaggerDoc() map[string]string {
	return map_ImageExecutionPolicyRule
}

var map_ImagePolicyConfig = map[string]string{
	"":                "ImagePolicyConfig is the configuration for control of images running on the platform.",
	"resolveImages":   "ResolveImages indicates the default image resolution behavior.  If a rewriting policy is chosen, then the image pull specs will be updated.",
	"resolutionRules": "ResolutionRules allows more specific image resolution rules to be applied per resource. If empty, it defaults to allowing local image stream lookups - \"mysql\" will map to the image stream tag \"mysql:latest\" in the current namespace if the stream supports it.",
	"executionRules":  "ExecutionRules determine whether the use of an image is allowed in an object with a pod spec. By default, these rules only apply to pods, but may be extended to other resource types. If all execution rules are negations, the default behavior is allow all. If any execution rule is an allow, the default behavior is to reject all.",
}

func (ImagePolicyConfig) SwaggerDoc() map[string]string {
	return map_ImagePolicyConfig
}

var map_ImageResolutionPolicyRule = map[string]string{
	"":               "ImageResolutionPolicyRule describes resolution rules based on resource.",
	"targetResource": "TargetResource is the identified group and resource. If Resource is *, this rule will apply to all resources in that group.",
	"localNames":     "LocalNames will allow single segment names to be interpreted as namespace local image stream tags, but only if the target image stream tag has the \"resolveLocalNames\" field set.",
}

func (ImageResolutionPolicyRule) SwaggerDoc() map[string]string {
	return map_ImageResolutionPolicyRule
}

var map_ValueCondition = map[string]string{
	"":      "ValueCondition reflects whether the following key in a map is set or has a given value.",
	"key":   "Key is the name of a key in a map to retrieve.",
	"set":   "Set indicates the provided key exists in the map. This field is exclusive with Value.",
	"value": "Value indicates the provided key has the given value. This field is exclusive with Set.",
}

func (ValueCondition) SwaggerDoc() map[string]string {
	return map_ValueCondition
}
