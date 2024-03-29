data "k8s_workload_codeflare_dev_app_wrapper_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
