data "k8s_ucp_dev_resource_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
