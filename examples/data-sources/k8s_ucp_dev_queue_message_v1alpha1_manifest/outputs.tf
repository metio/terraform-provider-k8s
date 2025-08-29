output "manifests" {
  value = {
    "example" = data.k8s_ucp_dev_queue_message_v1alpha1_manifest.example.yaml
  }
}
