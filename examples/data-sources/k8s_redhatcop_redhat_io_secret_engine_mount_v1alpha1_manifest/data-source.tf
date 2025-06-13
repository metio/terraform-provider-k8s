data "k8s_redhatcop_redhat_io_secret_engine_mount_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
