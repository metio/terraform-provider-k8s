output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_cluster_policy_v1_manifest.example.yaml
    "int_value" = data.k8s_kyverno_io_cluster_policy_v1_manifest.int_value.yaml
    "bool_value" = data.k8s_kyverno_io_cluster_policy_v1_manifest.bool_value.yaml
    "array_value" = data.k8s_kyverno_io_cluster_policy_v1_manifest.array_value.yaml
    "map_value" = data.k8s_kyverno_io_cluster_policy_v1_manifest.map_value.yaml
    "mixed_value" = data.k8s_kyverno_io_cluster_policy_v1_manifest.mixed_value.yaml
  }
}
