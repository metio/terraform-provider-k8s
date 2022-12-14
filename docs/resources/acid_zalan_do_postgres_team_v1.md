---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_acid_zalan_do_postgres_team_v1 Resource - terraform-provider-k8s"
subcategory: "acid.zalan.do"
description: |-
  
---

# k8s_acid_zalan_do_postgres_team_v1 (Resource)



## Example Usage

```terraform
resource "k8s_acid_zalan_do_postgres_team_v1" "big" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
    labels = {
      "test" = "abc"
    }
    annotations = {
      "try" = "this"
    }
  }
  spec = {
    additional_members = {
      "team-a" = ["bob", "bill", "barry"]
    }
    additional_superuser_teams = {
      "team-b" = ["alice", "eve", "julia"]
    }
    additional_teams = {
      "team-c" = ["team-1", "team-2", "team-3"]
    }
  }
}

resource "k8s_acid_zalan_do_postgres_team_v1" "small" {
  metadata = {
    name = "test"
  }
  spec = {
    additional_members = {
      "team-a" = ["bob", "bill", "barry"]
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `additional_members` (Map of List of String) Map for teamId and associated additional users
- `additional_superuser_teams` (Map of List of String) Map for teamId and associated additional superuser teams
- `additional_teams` (Map of List of String) Map for teamId and associated additional teams


