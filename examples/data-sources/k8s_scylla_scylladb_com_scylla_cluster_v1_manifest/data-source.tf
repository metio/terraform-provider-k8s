data "k8s_scylla_scylladb_com_scylla_cluster_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    version                         = "2.3.1"
    repository                      = "scylladb/scylla"
    developer_mode                  = true
    cpuset                          = false
    automatic_orphaned_node_cleanup = true
    repairs = [
      {
        name      = "weekly us-east-1 repair"
        intensity = "2"
        interval  = "7d"
        dc        = ["us-east-1"]
      }
    ]
    backups = [
      {
        name       = "daily users backup"
        rate_limit = ["50"]
        location   = ["s3 =cluster-backups"]
        interval   = "1d"
        keyspace   = ["users"]
      },
      {
        name       = "weekly full cluster backup"
        rate_limit = ["50"]
        location   = ["s3 =cluster-backups"]
        interval   = "7d"
      }
    ]
    datacenter = {
      name = "us-east-1"
      racks = [
        {
          name    = "us-east-1a"
          members = 3
          storage = {
            capacity           = "500G"
            storage_class_name = "local-raid-disks"
          }
          resources = {
            requests = {
              cpu    = 8
              memory = "32Gi"
            }
            limits = {
              cpu    = 8
              memory = "32Gi"
            }
          }
          placement = {
            node_affinity = {
              required_during_scheduling_ignored_during_execution = {
                node_selector_terms = [
                  {
                    match_expressions = [
                      {
                        key      = "failure-domain.beta.kubernetes.io/region"
                        operator = "In"
                        values   = ["us-east-1"]
                      },
                      {
                        key      = "failure-domain.beta.kubernetes.io/zone"
                        operator = "In"
                        values   = ["us-east-1a"]
                      },
                    ]
                  }
                ]
              }
            }
            tolerations = [
              {
                key      = "role"
                operator = "Equal"
                value    = "scylla-clusters"
                effect   = "NoSchedule"
              }
            ]
          }
        }
      ]
    }
  }
}
