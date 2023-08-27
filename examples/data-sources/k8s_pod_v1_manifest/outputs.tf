output "manifests" {
  value = {
    "example" = data.k8s_pod_v1_manifest.example.yaml
  }
}
