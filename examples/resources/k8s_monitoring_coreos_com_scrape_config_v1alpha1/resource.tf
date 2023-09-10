resource "k8s_monitoring_coreos_com_scrape_config_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
