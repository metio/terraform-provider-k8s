output "manifests" {
  value = {
    "example" = data.k8s_acid_zalan_do_postgres_team_v1_manifest.example.yaml
  }
}
