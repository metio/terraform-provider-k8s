resource "k8s_minio_min_io_tenant_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
