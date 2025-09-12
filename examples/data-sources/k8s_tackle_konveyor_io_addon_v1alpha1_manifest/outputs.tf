output "manifests" {
  value = {
    "example" = data.k8s_tackle_konveyor_io_addon_v1alpha1_manifest.example.yaml
  }
}
