output "manifests" {
  value = {
    "example" = data.k8s_troubleshoot_sh_remote_collector_v1beta2_manifest.example.yaml
  }
}
