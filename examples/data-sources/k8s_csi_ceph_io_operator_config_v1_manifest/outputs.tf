output "manifests" {
  value = {
    "example" = data.k8s_csi_ceph_io_operator_config_v1_manifest.example.yaml
  }
}
