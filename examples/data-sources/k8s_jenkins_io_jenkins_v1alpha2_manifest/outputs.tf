output "manifests" {
  value = {
    "example" = data.k8s_jenkins_io_jenkins_v1alpha2_manifest.example.yaml
  }
}
