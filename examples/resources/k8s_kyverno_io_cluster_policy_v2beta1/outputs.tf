output "resources" {
  value = {
    "minimal" = k8s_kyverno_io_cluster_policy_v2beta1.minimal.yaml
  }
}
