data "k8s_shipwright_io_build_run_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
