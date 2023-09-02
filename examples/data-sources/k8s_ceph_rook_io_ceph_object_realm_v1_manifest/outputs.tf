output "manifests" {
  value = {
    "example" = data.k8s_ceph_rook_io_ceph_object_realm_v1_manifest.example.yaml
  }
}
