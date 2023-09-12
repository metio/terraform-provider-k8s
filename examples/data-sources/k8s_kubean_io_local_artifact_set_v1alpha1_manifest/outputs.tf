output "manifests" {
  value = {
    "example" = data.k8s_kubean_io_local_artifact_set_v1alpha1_manifest.example.yaml
  }
}
