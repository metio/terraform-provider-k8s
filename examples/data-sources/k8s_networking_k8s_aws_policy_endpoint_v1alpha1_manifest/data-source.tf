data "k8s_networking_k8s_aws_policy_endpoint_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
