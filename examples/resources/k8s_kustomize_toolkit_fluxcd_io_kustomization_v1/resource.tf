resource "k8s_kustomize_toolkit_fluxcd_io_kustomization_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
