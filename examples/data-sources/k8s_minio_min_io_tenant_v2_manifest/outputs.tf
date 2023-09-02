output "manifests" {
  value = {
    "example" = data.k8s_minio_min_io_tenant_v2_manifest.example.yaml
  }
}
