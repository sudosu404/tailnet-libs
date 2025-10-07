# Tailnet - Tailscale Proxy

Tailnet simplifies the process of securely exposing services and Docker containers
to your Tailscale network by automatically creating Tailscale machines for each
tagged container. This allows services to be accessible via unique, secure URLs
without the need for complex configurations or additional Tailscale containers.

## Version 2

Version 2 is already in beta. Please try and open issues if bugs detected.
Some configurations of Version 1.x are deprecated or changed please verify it in
 [Documentation changelog](https://sudosu404.github.io/gnx-cli/docs/changelog/).

Because of some breaking changes, final version 2 arrive will not set as latest
Docker image. We will wait some weeks to give you time to update.

Version 1 will not get new features.

## Full Documentation

- [Official Documentation](https://sudosu404.github.io/gnx-cli/)

## Breaking Changes

Please read the [Documentation changelog](https://sudosu404.github.io/gnx-cli/docs/changelog/)
for details.

## Help needed

Please help with documentation, tests development, new features, bug fixes.
If you don't feel comfortable to this kind of tasks, [sponsor](https://github.com/sponsors/sudosu404)
the project.

## Docker Images

1. sudosu404/gnx-cli:vx.x.x  - Version x.x.x
2. sudosu404/gnx-cli:1       - Latest release of version 1.x.x
3. sudosu404/gnx-cli:2       - Latest release of version 2.x.x (beta)
4. sudosu404/gnx-cli:latest  - Latest stable
5. sudosu404/gnx-cli:dev     - Latest Development Build

## Core Functionality

- **Automatic Tailscale Machine Creation**: For each Docker container tagged
with the appropriate labels, Tailnet creates a new Tailscale machine.
- **Default Serving**: By default, each service is accessible via
`https://{machine-name}.funny-name.ts.net`, where `{machine-name}` is derived
from your container name or custom label.

## Key Features

- **Simplified Networking**: Eliminates the need for a separate Tailscale
container for each service.
- **Label-Based Configuration**: Easy setup using Docker container labels.
- **Automatic HTTPS**: Leverages Tailscale's built-in LetsEncrypt certificate support.
- **Flexible Protocol Support**: Handles HTTP and HTTPS traffic (defaulting to HTTPS).
- **Lightweight Architecture**: Efficient, Docker-based design for minimal overhead.

## How It Works

Tailnet operates by creating a seamless integration between your Docker
containers and Tailscale network:

1. **Container Scanning**: Tailnet continuously monitors your Docker
environment for containers with the `tailnet.enable=true` label.
2. **Tailscale Machine Creation**: When a tagged container is detected, Tailnet
automatically creates a new Tailscale machine for that container.
3. **Hostname Assignment**: The Tailscale machine is assigned a hostname based
on the `tailnet.name` label or the container's name.
4. **Port Mapping**: Tailnet maps the container's internal port to the Tailscale
machine.
5. **Traffic Routing**: Incoming requests to the Tailscale machine are routed to
the appropriate Docker container and port.
6. **Dynamic Management**: As containers start and stop, Tailnet automatically
creates and removes the corresponding Tailscale machines and routing configurations.

## Requirements

Before using this application, make sure you have:

- [Tailscale](https://tailscale.com/) installed and configured on your host machine.
- [Docker](https://www.docker.com/) installed and running.

## License

This project is licensed under the AGPL3  License. See the [LICENSE](LICENSE) file
for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve the app.
