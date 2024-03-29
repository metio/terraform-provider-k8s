output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_update_request_v2_manifest.example.yaml
  }
}
