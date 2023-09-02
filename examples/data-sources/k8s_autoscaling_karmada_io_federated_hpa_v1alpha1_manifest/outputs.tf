output "manifests" {
  value = {
    "example" = data.k8s_autoscaling_karmada_io_federated_hpa_v1alpha1_manifest.example.yaml
  }
}
