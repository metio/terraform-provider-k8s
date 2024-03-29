output "manifests" {
  value = {
    "example" = data.k8s_forklift_konveyor_io_ovirt_volume_populator_v1beta1_manifest.example.yaml
  }
}
