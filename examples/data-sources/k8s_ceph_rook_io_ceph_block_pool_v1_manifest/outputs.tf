output "manifests" {
  value = {
    "example" = data.k8s_ceph_rook_io_ceph_block_pool_v1_manifest.example.yaml
  }
}
