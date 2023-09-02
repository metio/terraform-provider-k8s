output "manifests" {
  value = {
    "example" = data.k8s_apps_gitlab_com_runner_v1beta2_manifest.example.yaml
  }
}
