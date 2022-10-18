resource "k8s_operator_open_cluster_management_io_cluster_manager_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_operator_open_cluster_management_io_cluster_manager_v1" "example" {
  metadata = {
    name = "cluster-manager"
  }
  spec = {
    deploy_option = {
      mode = "Default"
    }
    placement_image_pull_spec = "quay.io/open-cluster-management/placement:v0.8.0"
    registration_configuration = {
      feature_gates = [
        {
          feature = "DefaultClusterSet"
          mode    = "Enable"
        }
      ]
    }
    registration_image_pull_spec = "quay.io/open-cluster-management/registration:v0.8.0"
    work_image_pull_spec         = "quay.io/open-cluster-management/work:v0.8.0"
  }
}
