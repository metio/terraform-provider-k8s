resource "k8s_operator_aquasec_com_aqua_database_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    common = {
      active_active = true
    }
    deploy = {
      replicas = 12
    }
    disk_size = 120
    infra = {
      requirements = true
    }
  }
}
