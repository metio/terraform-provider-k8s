output "manifests" {
  value = {
    "example" = data.k8s_actions_github_com_ephemeral_runner_v1alpha1_manifest.example.yaml
  }
}
