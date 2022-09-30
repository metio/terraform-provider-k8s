output "resources" {
  value = {
    "minimal" = k8s_minio_min_io_tenant_v2.minimal.yaml
  }
}
