data "k8s_redhatcop_redhat_io_group_alias_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
