data "k8s_redhatcop_redhat_io_kubernetes_secret_engine_role_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
