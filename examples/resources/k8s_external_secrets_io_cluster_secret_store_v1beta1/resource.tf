resource "k8s_external_secrets_io_cluster_secret_store_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_external_secrets_io_cluster_secret_store_v1beta1" "issue_110" {
  metadata = {
    name = "gcp-backend"
  }

  spec = {
    provider = {
      gcpsm = {
        project_id = "some-project-id"
        auth = {
          workload_identity = {
            cluster_name     = "some-cluster-name"
            cluster_location = "some-cluster-location"
            service_account_ref = {
              name      = "some-service-account"
              namespace = "some-namespace"
            }
          }
        }
      }
    }
  }
}
