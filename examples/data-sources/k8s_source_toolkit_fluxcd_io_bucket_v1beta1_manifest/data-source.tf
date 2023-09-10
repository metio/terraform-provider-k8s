data "k8s_source_toolkit_fluxcd_io_bucket_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
