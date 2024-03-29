output "manifests" {
  value = {
    "example" = data.k8s_actions_github_com_autoscaling_runner_set_v1alpha1_manifest.example.yaml
  }
}
