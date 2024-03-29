output "manifests" {
  value = {
    "example" = data.k8s_operator_openshift_io_csi_snapshot_controller_v1_manifest.example.yaml
  }
}
