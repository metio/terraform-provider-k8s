---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "imaging-ingestion.alvearie.org/v1alpha1"
description: |-
  Provides a proxied DIMSE Application Entity (AE) in the cluster for C-STORE operations to a storage space
---

# k8s_imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1 (Resource)

Provides a proxied DIMSE Application Entity (AE) in the cluster for C-STORE operations to a storage space

## Example Usage

```terraform
resource "k8s_imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1" "example" {
  metadata = {
    name = "ingestion"
  }
  spec = {
    application_entity_title : "DICOM-INGEST"
    bucket_config_name : "imaging-ingestion"
    bucket_secret_name : "imaging-ingestion"
    dicom_event_driven_ingestion_name : "core"
    dimse_service : {}
    image_pull_policy : "Always"
    nats_secure : true
    nats_subject_root : "DIMSE"
    nats_token_secret : "ingestion-nats-secure-bound-token"
    nats_url : "nats-secure.imaging-ingestion.svc.cluster.local:4222"
    provider_name : "provider"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) DimseIngestionServiceSpec defines the desired state of DimseIngestionService (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `dicom_event_driven_ingestion_name` (String) DICOM Event Driven Ingestion Name

Optional:

- `application_entity_title` (String) Application Entity Title
- `bucket_config_name` (String) Bucket Config Name
- `bucket_secret_name` (String) Bucket Secret Name
- `dimse_service` (Attributes) DIMSE Service Spec (see [below for nested schema](#nestedatt--spec--dimse_service))
- `image_pull_policy` (String) Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
- `image_pull_secrets` (Attributes List) Image Pull Secrets (see [below for nested schema](#nestedatt--spec--image_pull_secrets))
- `nats_secure` (Boolean) Make NATS Connection Secure
- `nats_subject_root` (String) NATS Subject Root
- `nats_token_secret` (String) NATS Token Secret Name
- `nats_url` (String) NATS URL
- `provider_name` (String) Provider Name

<a id="nestedatt--spec--dimse_service"></a>
### Nested Schema for `spec.dimse_service`

Optional:

- `image` (String) Image


<a id="nestedatt--spec--image_pull_secrets"></a>
### Nested Schema for `spec.image_pull_secrets`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?

