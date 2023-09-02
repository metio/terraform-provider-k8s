output "manifests" {
  value = {
    "example" = data.k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1_manifest.example.yaml
  }
}
