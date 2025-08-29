output "manifests" {
  value = {
    "example" = data.k8s_csi_ceph_io_driver_v1alpha1_manifest.example.yaml
  }
}
