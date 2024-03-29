output "manifests" {
  value = {
    "example" = data.k8s_actions_summerwind_dev_runner_replica_set_v1alpha1_manifest.example.yaml
  }
}
