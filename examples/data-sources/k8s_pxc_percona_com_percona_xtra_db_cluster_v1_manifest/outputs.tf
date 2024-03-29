output "manifests" {
  value = {
    "example" = data.k8s_pxc_percona_com_percona_xtra_db_cluster_v1_manifest.example.yaml
  }
}
