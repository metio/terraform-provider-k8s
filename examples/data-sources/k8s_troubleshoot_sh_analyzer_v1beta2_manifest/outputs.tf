output "manifests" {
  value = {
    "example" = data.k8s_troubleshoot_sh_analyzer_v1beta2_manifest.example.yaml
  }
}
