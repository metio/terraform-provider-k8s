resource "k8s_camel_apache_org_kamelet_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_camel_apache_org_kamelet_v1alpha1" "example" {
  metadata = {
    name = "telegram-text-source"
    annotations = {
      "camel.apache.org/kamelet.icon" = "data:image/svg+xml;base64,PD94bW..."
    }
    labels = {
      "camel.apache.org/kamelet.type" = "source"
    }
  }
  spec = {
    definition = {
      title       = "Telegram Text Source"
      description = <<-TEXT
Receive all text messages that people send to your telegram bot.

# Instructions
Description can include Markdown and guide the final user to configure the Kamelet parameters.
TEXT
      required    = ["botToken"]
      properties = {
        botToken = {
          title         = "Token"
          description   = "The token to access your bot on Telegram"
          type          = "string"
          x-descriptors = ["urn:alm:descriptor:com.tectonic.ui:password"]
        }
      }
    }
    types = {
      out = {
        media_type = "text/plain"
      }
    }
    template = {
      from = {
        uri = "telegram:bots"
        parameters = {
          authorization_token = "#property:botToken"
        }
        steps = [
          {
            convert-body-to = {
              type       = "java.lang.String"
              type-class = "java.lang.String"
              charset    = "UTF8"
            }
          },
          {
            filter = {
              simple = "$${body} != null"
            }
          },
          {
            log = "$${body}"
          },
          {
            to = "kamelet:sink"
          }
        ]
      }
    }
  }
}
