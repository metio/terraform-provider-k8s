output "manifests" {
  value = {
    "example" = data.k8s_csi_ceph_io_ceph_connection_v1_manifest.example.yaml
  }
}
