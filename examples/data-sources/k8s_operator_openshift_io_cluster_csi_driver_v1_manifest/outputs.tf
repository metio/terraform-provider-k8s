output "manifests" {
  value = {
    "example" = data.k8s_operator_openshift_io_cluster_csi_driver_v1_manifest.example.yaml
  }
}
