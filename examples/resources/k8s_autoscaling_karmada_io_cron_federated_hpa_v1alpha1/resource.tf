resource "k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
