resource "k8s_execution_furiko_io_job_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_execution_furiko_io_job_v1alpha1" "from_job_config" {
  metadata = {
    name = "test"
  }
  spec = {
    config_name   = "jobconfig-sample"
    option_values = <<-YAML
username: Example User
image_tag: 3.15
YAML
    start_policy = {
      concurrency_policy = "Enqueue"
      start_after        = "2022-03-06T00:27:00+08:00"
    }
  }
}

resource "k8s_execution_furiko_io_job_v1alpha1" "standalone" {
  metadata = {
    name = "test"
  }
  spec = {
    type = "Adhoc"
    substitutions = {
      "option.username" = "Example User"
    }
    template = {
      parallelism = {
        with_count          = 3
        completion_strategy = "AllSuccessful"
      }
      max_attempts                 = 3
      retry_delay_seconds          = 10
      task_pending_timeout_seconds = 1800
      forbid_task_force_deletion   = true
      task_template = {
        pod = {
          spec = {
            containers = [
              {
                args = ["echo", "Hello world, $${option.username}!"]
                env = [
                  {
                    name  = "JOBCONFIG_NAME"
                    value = "jobconfig-sample"
                  },
                  {
                    name  = "JOB_NAME"
                    value = "$${job.name}"
                  },
                  {
                    name  = "TASK_NAME"
                    value = "$${task.name}"
                  },
                  {
                    name  = "TASK_INDEX"
                    value = "$${task.index_num}"
                  },
                ]
                image = "alpine"
                name  = "job-container"
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
    ttl_seconds_after_finished = 3600
  }
}
