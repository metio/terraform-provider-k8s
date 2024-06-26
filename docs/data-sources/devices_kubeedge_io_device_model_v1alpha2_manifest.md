---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_devices_kubeedge_io_device_model_v1alpha2_manifest Data Source - terraform-provider-k8s"
subcategory: "devices.kubeedge.io"
description: |-
  DeviceModel is the Schema for the device model API
---

# k8s_devices_kubeedge_io_device_model_v1alpha2_manifest (Data Source)

DeviceModel is the Schema for the device model API

## Example Usage

```terraform
data "k8s_devices_kubeedge_io_device_model_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) DeviceModelSpec defines the model / template for a device.It is a blueprint which describes the device capabilities and access mechanism via property visitors. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `properties` (Attributes List) Required: List of device properties. (see [below for nested schema](#nestedatt--spec--properties))
- `protocol` (String) Required for DMI: Protocol name used by the device.

<a id="nestedatt--spec--properties"></a>
### Nested Schema for `spec.properties`

Optional:

- `description` (String) The device property description.
- `name` (String) Required: The device property name.
- `type` (Attributes) Required: PropertyType represents the type and data validation of the property. (see [below for nested schema](#nestedatt--spec--properties--type))

<a id="nestedatt--spec--properties--type"></a>
### Nested Schema for `spec.properties.type`

Optional:

- `boolean` (Attributes) (see [below for nested schema](#nestedatt--spec--properties--type--boolean))
- `bytes` (Attributes) (see [below for nested schema](#nestedatt--spec--properties--type--bytes))
- `double` (Attributes) (see [below for nested schema](#nestedatt--spec--properties--type--double))
- `float` (Attributes) (see [below for nested schema](#nestedatt--spec--properties--type--float))
- `int` (Attributes) (see [below for nested schema](#nestedatt--spec--properties--type--int))
- `string` (Attributes) (see [below for nested schema](#nestedatt--spec--properties--type--string))

<a id="nestedatt--spec--properties--type--boolean"></a>
### Nested Schema for `spec.properties.type.boolean`

Optional:

- `access_mode` (String) Required: Access mode of property, ReadWrite or ReadOnly.
- `default_value` (Boolean)


<a id="nestedatt--spec--properties--type--bytes"></a>
### Nested Schema for `spec.properties.type.bytes`

Optional:

- `access_mode` (String) Required: Access mode of property, ReadWrite or ReadOnly.


<a id="nestedatt--spec--properties--type--double"></a>
### Nested Schema for `spec.properties.type.double`

Optional:

- `access_mode` (String) Required: Access mode of property, ReadWrite or ReadOnly.
- `default_value` (Number)
- `maximum` (Number)
- `minimum` (Number)
- `unit` (String) The unit of the property


<a id="nestedatt--spec--properties--type--float"></a>
### Nested Schema for `spec.properties.type.float`

Optional:

- `access_mode` (String) Required: Access mode of property, ReadWrite or ReadOnly.
- `default_value` (Number)
- `maximum` (Number)
- `minimum` (Number)
- `unit` (String) The unit of the property


<a id="nestedatt--spec--properties--type--int"></a>
### Nested Schema for `spec.properties.type.int`

Optional:

- `access_mode` (String) Required: Access mode of property, ReadWrite or ReadOnly.
- `default_value` (Number)
- `maximum` (Number)
- `minimum` (Number)
- `unit` (String) The unit of the property


<a id="nestedatt--spec--properties--type--string"></a>
### Nested Schema for `spec.properties.type.string`

Optional:

- `access_mode` (String) Required: Access mode of property, ReadWrite or ReadOnly.
- `default_value` (String)
