data "k8s_app_k8s_io_application_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
