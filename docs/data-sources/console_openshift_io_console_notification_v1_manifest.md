---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_console_openshift_io_console_notification_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "console.openshift.io"
description: |-
  ConsoleNotification is the extension for configuring openshift web console notifications.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).
---

# k8s_console_openshift_io_console_notification_v1_manifest (Data Source)

ConsoleNotification is the extension for configuring openshift web console notifications.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).

## Example Usage

```terraform
data "k8s_console_openshift_io_console_notification_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) ConsoleNotificationSpec is the desired console notification configuration. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `text` (String) text is the visible text of the notification.

Optional:

- `background_color` (String) backgroundColor is the color of the background for the notification as CSS data type color.
- `color` (String) color is the color of the text for the notification as CSS data type color.
- `link` (Attributes) link is an object that holds notification link details. (see [below for nested schema](#nestedatt--spec--link))
- `location` (String) location is the location of the notification in the console. Valid values are: 'BannerTop', 'BannerBottom', 'BannerTopBottom'.

<a id="nestedatt--spec--link"></a>
### Nested Schema for `spec.link`

Required:

- `href` (String) href is the absolute secure URL for the link (must use https)
- `text` (String) text is the display text for the link