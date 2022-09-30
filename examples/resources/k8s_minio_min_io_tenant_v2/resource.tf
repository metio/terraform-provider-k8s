resource "k8s_minio_min_io_tenant_v2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    pools = []
  }
}
