output "resources" {
  value = {
    "minimal" = k8s_ceph_rook_io_ceph_object_store_v1.minimal.yaml
  }
}
