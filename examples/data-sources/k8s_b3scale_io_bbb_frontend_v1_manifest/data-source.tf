data "k8s_b3scale_io_bbb_frontend_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
