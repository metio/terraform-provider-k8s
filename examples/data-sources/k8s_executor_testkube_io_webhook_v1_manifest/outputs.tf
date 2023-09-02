output "manifests" {
  value = {
    "example" = data.k8s_executor_testkube_io_webhook_v1_manifest.example.yaml
  }
}
