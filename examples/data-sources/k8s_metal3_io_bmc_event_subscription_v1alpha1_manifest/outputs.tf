output "manifests" {
  value = {
    "example" = data.k8s_metal3_io_bmc_event_subscription_v1alpha1_manifest.example.yaml
  }
}
