data "k8s_autoscaling_karmada_io_federated_hpa_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    max_replicas = 5
    scale_target_ref = {
      api_version = "v1"
      kind        = "Deployment"
      name        = "some-deployment"
    }
  }
}
