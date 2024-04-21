data "k8s_policy_kubeedge_io_service_account_access_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
