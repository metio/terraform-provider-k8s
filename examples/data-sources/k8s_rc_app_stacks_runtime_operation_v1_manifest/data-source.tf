data "k8s_rc_app_stacks_runtime_operation_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
