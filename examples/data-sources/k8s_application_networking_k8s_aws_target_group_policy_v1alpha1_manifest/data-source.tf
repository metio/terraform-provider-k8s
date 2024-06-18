data "k8s_application_networking_k8s_aws_target_group_policy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    target_ref = {
      group = "v1"
      kind  = "Pod"
      name  = "some-pod"
    }
  }
}
