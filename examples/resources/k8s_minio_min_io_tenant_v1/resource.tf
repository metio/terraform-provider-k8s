resource "k8s_minio_min_io_tenant_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    zones = []
  }
}
