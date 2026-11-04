data "k8s_application_networking_k8s_aws_vpc_association_policy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }

  spec = {
    target_ref = {
      group = "some-value"
      kind  = "some-value"
      name  = "some-value"
    }
  }
}
