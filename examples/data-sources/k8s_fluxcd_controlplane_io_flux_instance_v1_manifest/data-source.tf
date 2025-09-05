data "k8s_fluxcd_controlplane_io_flux_instance_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
