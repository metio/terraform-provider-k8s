data "k8s_monitoring_coreos_com_probe_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
