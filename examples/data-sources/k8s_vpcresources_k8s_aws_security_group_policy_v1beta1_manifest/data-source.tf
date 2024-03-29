data "k8s_vpcresources_k8s_aws_security_group_policy_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
