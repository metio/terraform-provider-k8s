data "k8s_app_terraform_io_workspace_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    name         = "some-name"
    organization = "some-organization"
    token = {
      secret_key_ref = {
        key = "some-key"
      }
    }
  }
}
