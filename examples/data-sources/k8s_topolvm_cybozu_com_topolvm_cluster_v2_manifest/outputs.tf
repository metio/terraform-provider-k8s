output "manifests" {
  value = {
    "example" = data.k8s_topolvm_cybozu_com_topolvm_cluster_v2_manifest.example.yaml
  }
}
