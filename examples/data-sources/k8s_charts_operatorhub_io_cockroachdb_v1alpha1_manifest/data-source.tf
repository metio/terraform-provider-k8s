data "k8s_charts_operatorhub_io_cockroachdb_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    clusterDomain = "cluster.local"
    conf = {
      attrs                             = []
      cache                             = "25%"
      cluster-name                      = ""
      disable-cluster-name-verification = false
      http-port                         = 8080
      join                              = []
      locality                          = ""
      logtostderr                       = "INFO"
      max-sql-memory                    = "25%"
      port                              = 26257
      single-node                       = false
      sql-audit-dir                     = ""
    }
    image = {
      credentials = {}
      pullPolicy  = "IfNotPresent"
      repository  = "cockroachdb/cockroach"
      tag         = "v20.2.4"
    }
    ingress = {
      annotations = {}
      enabled     = false
      hosts       = []
      labels      = {}
      tls         = []
      paths       = ["/"]
    }
    init = {
      affinity     = {}
      annotations  = {}
      nodeSelector = {}
      resources    = {}
      tolerations  = []
      labels = {
        "app.kubernetes.io/component" = "init"
      }
    }
    labels = {}
    networkPolicy = {
      enabled = false
      ingress = {
        grpc = []
        http = []
      }
    }
    service = {
      discovery = {
        annotations = {}
        labels = {
          "app.kubernetes.io/component" = "cockroachdb"
        }
      }
      ports = {
        grpc = {
          external = {
            name = "grpc"
            port = 26257
          }
          internal = {
            name = "grpc-internal"
            port = 26257
          }
        }
        http = {
          name = "http"
          port = 8080
        }
      }
      public = {
        annotations = {}
        labels = {
          "app.kubernetes.io/component" = "cockroachdb"
        }
        type = "ClusterIP"
      }
    }
    serviceMonitor = {
      annotations = {}
      enabled     = false
      interval    = "10s"
      labels      = {}
    }
    statefulset = {
      annotations         = {}
      args                = []
      env                 = []
      nodeAffinity        = {}
      nodeSelector        = {}
      podAffinity         = {}
      podManagementPolicy = "Parallel"
      priorityClassName   = ""
      replicas            = 3
      resources           = {}
      secretMounts        = []
      tolerations         = []
      budget = {
        maxUnavailable = 1
      }
      labels = {
        "app.kubernetes.io/component" = "cockroachdb"
      }
      podAntiAffinity = {
        topologyKey = "kubernetes.io/hostname"
        type        = "soft"
        weight      = 100
      }
      topologySpreadConstraints = {
        maxSkew           = 1
        topologyKey       = "topology.kubernetes.io/zone"
        whenUnsatisfiable = "ScheduleAnyway"
      }
      updateStrategy = {
        type = "RollingUpdate"
      }
    }
    storage = {
      hostPath = ""
      persistentVolume = {
        annotations  = {}
        enabled      = true
        labels       = {}
        size         = "100Gi"
        storageClass = ""
      }
    }
    tls = {
      certs = {
        clientRootSecret = "cockroachdb-root"
        nodeSecret       = "cockroachdb-node"
        provided         = false
        tlsSecret        = false
      }
      enabled = false
      init = {
        image = {
          credentials = {}
          pullPolicy  = "IfNotPresent"
          repository  = "cockroachdb/cockroach-k8s-request-cert"
          tag         = "0.4"
        }
      }
      serviceAccount = {
        create = true
        name   = ""
      }
    }
  }
}
