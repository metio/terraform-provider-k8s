data "k8s_vpcresources_k8s_aws_security_group_policy_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    pod_selector = {
      match_labels = {
        role = "my-role"
      }
    }
    security_groups = {
      group_ids = ["my_pod_security_group_id"]
    }
  }
}
