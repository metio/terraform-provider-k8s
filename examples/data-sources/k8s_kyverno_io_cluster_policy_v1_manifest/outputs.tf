output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_cluster_policy_v1_manifest.example.yaml
  }
}
