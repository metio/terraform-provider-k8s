data "k8s_apps_stateful_set_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
    annotations = {
      SomeAnnotation = "foobar"
    }
    labels = {
      k8s-app                           = "prometheus"
      "kubernetes.io/cluster-service"   = "true"
      "addonmanager.kubernetes.io/mode" = "Reconcile"
      version                           = "v2.2.1"
    }
  }
  spec = {
    pod_management_policy  = "Parallel"
    replicas               = 1
    revision_history_limit = 5
    service_name           = "prometheus"

    selector = {
      match_labels = {
        k8s-app = "prometheus"
      }
    }

    template = {
      metadata = {
        labels = {
          k8s-app = "prometheus"
        }
      }

      spec = {
        service_account_name             = "prometheus"
        termination_grace_period_seconds = 300

        init_containers = [
          {
            name              = "init-chown-data"
            image             = "busybox:latest"
            image_pull_policy = "IfNotPresent"
            command           = ["chown", "-R", "65534:65534", "/data"]

            volume_mounts = [
              {
                name       = "prometheus-data"
                mount_path = "/data"
                sub_path   = ""
              }
            ]
          }
        ]
        containers = [
          {
            name              = "prometheus-server-configmap-reload"
            image             = "jimmidyson/configmap-reload:v0.1"
            image_pull_policy = "IfNotPresent"

            args = [
              "--volume-dir=/etc/config",
              "--webhook-url=http://localhost:9090/-/reload",
            ]

            volume_mounts = [
              {
                name       = "config-volume"
                mount_path = "/etc/config"
                read_only  = true
              }
            ]

            resources = {
              limits = {
                cpu    = "10m"
                memory = "10Mi"
              }

              requests = {
                cpu    = "10m"
                memory = "10Mi"
              }
            }

            readiness_probe = {
              http_get = {
                path = "/-/ready"
                port = "web"
              }

              initial_delay_seconds = 30
              timeout_seconds       = 30
            }

            liveness_probe = {
              http_get = {
                path   = "/-/healthy"
                port   = 9090
                scheme = "HTTPS"
              }

              initial_delay_seconds = 30
              timeout_seconds       = 30
            }
          },
          {
            name              = "prometheus-server"
            image             = "prom/prometheus:v2.2.1"
            image_pull_policy = "IfNotPresent"

            args = [
              "--config.file=/etc/config/prometheus.yml",
              "--storage.tsdb.path=/data",
              "--web.console.libraries=/etc/prometheus/console_libraries",
              "--web.console.templates=/etc/prometheus/consoles",
              "--web.enable-lifecycle",
            ]

            ports = [
              {
                container_port = 9090
              }
            ]

            resources = {
              limits = {
                cpu    = "200m"
                memory = "1000Mi"
              }

              requests = {
                cpu    = "200m"
                memory = "1000Mi"
              }
            }

            volume_mounts = [
              {
                name       = "config-volume"
                mount_path = "/etc/config"
              },
              {
                name       = "prometheus-data"
                mount_path = "/data"
                sub_path   = ""
              },
            ]

            readiness_probe = {
              http_get = {
                path = "/-/ready"
                port = 9090
              }

              initial_delay_seconds = 30
              timeout_seconds       = 30
            }

            liveness_probe = {
              http_get = {
                path   = "/-/healthy"
                port   = 9090
                scheme = "HTTPS"
              }

              initial_delay_seconds = 30
              timeout_seconds       = 30
            }
          },
        ]
        volumes = [
          {
            name = "config-volume"

            config_map = {
              name = "prometheus-config"
            }
          }
        ]
      }
    }

    update_strategy = {
      type = "RollingUpdate"
      rolling_update = {
        partition = 1
      }
    }

    volume_claim_templates = [
      {
        metadata = {
          name = "prometheus-data"
        }
        spec = {
          access_modes       = ["ReadWriteOnce"]
          storage_class_name = "standard"

          resources = {
            requests = {
              storage = "16Gi"
            }
          }
        }
      },
    ]
  }
}
