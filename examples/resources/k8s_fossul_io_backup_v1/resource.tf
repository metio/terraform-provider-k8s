resource "k8s_fossul_io_backup_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_fossul_io_backup_v1" "example" {
  metadata = {
    name = "mariadb-sample"
  }
  spec = {
    deployment_name = "mariadb"
    policy          = "daily"
  }
}
