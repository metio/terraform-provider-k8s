data "k8s_infra_contrib_fluxcd_io_terraform_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
