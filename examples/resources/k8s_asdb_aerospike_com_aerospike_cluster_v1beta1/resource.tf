resource "k8s_asdb_aerospike_com_aerospike_cluster_v1beta1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    aerospike_config = {}
    image = "aerospike/aerospike-server-enterprise:6.1.0.1"
    size = 123
  }
}

resource "k8s_asdb_aerospike_com_aerospike_cluster_v1beta1" "example" {
  metadata = {
    name      = "aerocluster"
    namespace = "aerospike"
  }
  spec = {
    aerospike_config = {
      namespaces = [
        {
          memory-size = 3000000000
          name = "test"
          replication-factor = 2
          storage-engine = {
            type = "memory"
          }
        }
      ]
      network = {
        fabric = {
          port = 3001
        }
        heartbeat = {
          port = 3002
        }
        service = {
          port = 3000
        }
      }
      service = {
        feature-key-file = "/etc/aerospike/secret/features.conf"
      }
    }
    image = "aerospike/aerospike-server-enterprise:6.1.0.1"
    pod_spec = {
      multi_pod_per_host = true
    }
    size = 2
    storage = {
      volumes = [
        {
          name = "aerospike-config-secret"
          aerospike = {
            path = "/etc/aerospike/secret"
          }
          source = {
            secret = {
              secret_name = "aerospike-secret"
            }
          }
        }
      ]
    }
    validation_policy = {
      skip_work_dir_validate = true
      skip_xdr_dlog_file_validate = true
    }
  }
}
