data "k8s_fluxcd_controlplane_io_flux_report_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
