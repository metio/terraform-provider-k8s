output "manifests" {
  value = {
    "example" = data.k8s_topolvm_cybozu_com_logical_volume_v1_manifest.example.yaml
  }
}
