output "manifests" {
  value = {
    "example" = data.k8s_velero_io_server_status_request_v1_manifest.example.yaml
  }
}
