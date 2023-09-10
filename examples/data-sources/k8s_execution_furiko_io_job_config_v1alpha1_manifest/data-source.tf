data "k8s_execution_furiko_io_job_config_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    concurrency = {
      policy = "Forbid"
    }
    schedule = {
      cron = {
        expression = "0 */15 * * * * *"
        timezone   = "Asia/Singapore"
      }
      disabled = false
    }
    option = {
      options = [
        {
          type  = "String"
          name  = "username"
          label = "username"
          string = {
            default     = "Example User"
            trim_spaces = true
          }
        }
      ]
    }
    template = {
      metadata = {
        annotations = {
          "annotations.furiko.io/job-group" : "cool-jobs"
        }
      }
      spec = {
        max_attempts                 = 3
        retry_delay_seconds          = 10
        task_pending_timeout_seconds = 1800
        task_template = {
          pod = {
            spec = {
              containers = [
                {
                  name = "job-container"
                  args = ["echo", "Hello world, $${option.username}!"]
                  env = [
                    {
                      name  = "JOBCONFIG_NAME"
                      value = "$${jobconfig.name}"
                    },
                    {
                      name  = "JOB_NAME"
                      value = "$${job.name}"
                    },
                  ]
                  image = "alpine"
                  resources = {
                    limits = {
                      cpu    = "100m"
                      memory = "64Mi"
                    }
                  }
                }
              ]
            }
          }
        }
      }
    }
  }
}
