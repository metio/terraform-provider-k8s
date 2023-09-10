data "k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
