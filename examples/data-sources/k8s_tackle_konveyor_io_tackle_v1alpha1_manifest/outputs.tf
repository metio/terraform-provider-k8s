output "manifests" {
  value = {
    "example" = data.k8s_tackle_konveyor_io_tackle_v1alpha1_manifest.example.yaml
  }
}
