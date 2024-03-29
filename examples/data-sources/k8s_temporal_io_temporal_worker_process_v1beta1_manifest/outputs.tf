output "manifests" {
  value = {
    "example" = data.k8s_temporal_io_temporal_worker_process_v1beta1_manifest.example.yaml
  }
}
