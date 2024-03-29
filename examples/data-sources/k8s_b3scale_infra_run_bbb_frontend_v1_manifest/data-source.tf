data "k8s_b3scale_infra_run_bbb_frontend_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
