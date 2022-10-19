# Golang microservice template

Golang microservice template.

Try to create microservice template based on my experience.

Long time I collect different ideas from various services.
All of them has pros and cons.
I just need balanced result to save time when creating a project.

How I see solution anatomy:

1. It is desirable that the service occupies a separate namespace in k8s.
2. One code base is service (microservice).
3. One service can contains multiple components. It allows replicate functionality independently.
4. Each component has separate deployment in k8s.

How I see runtime states of classic service or component of a service:

1. Initialize logger
2. Load default settings
3. Load settings from various sources (file, consul, vault) and override defaults

### Usage

TODO

### Dependencies

* https://github.com/spf13/cobra
* https://github.com/spf13/viper
* github.com/labstack/echo (v4)
* https://github.com/deepmap/oapi-codegen
* https://github.com/ClickHouse/clickhouse-go (v2)