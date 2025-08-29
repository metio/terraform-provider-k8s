output "manifests" {
  value = {
    "example" = data.k8s_csi_ceph_io_client_profile_mapping_v1_manifest.example.yaml
  }
}
