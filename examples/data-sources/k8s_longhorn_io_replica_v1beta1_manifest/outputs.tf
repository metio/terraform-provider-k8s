output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_replica_v1beta1_manifest.example.yaml
  }
}
