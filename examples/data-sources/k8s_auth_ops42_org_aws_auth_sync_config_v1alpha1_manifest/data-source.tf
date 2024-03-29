data "k8s_auth_ops42_org_aws_auth_sync_config_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
