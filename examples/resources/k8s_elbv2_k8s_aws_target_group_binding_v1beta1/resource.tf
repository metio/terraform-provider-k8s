resource "k8s_elbv2_k8s_aws_target_group_binding_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
