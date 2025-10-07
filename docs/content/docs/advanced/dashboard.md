---
title: Dashboard
prev: /docs/advanced
---

{{% steps %}}

### Dashboard in docker

#### Tailnet docker compose

Update docker-compose.yml with the following:

```yaml  {filename="/config/tailnet.yaml"}
    labels:
      - tailnet.enable=true
      - tailnet.name=dash
```

#### Restart Tailnet

```bash
docker compose restart
```

### Standalone

#### Configure with a Proxy List provider

Configure a new files provider or configure it in any existing files provider.

```yaml  {filename="/config/tailnet.yaml"}
files:
  proxies:
    filename: /config/proxies.yaml
```

#### Add Dashboard entry on your Proxy List file

```yaml {filename="/config/proxies.yaml"}
dash:
  url: http://127.0.0.1:8080
```

### Test Dashboard access

```bash
curl https://dash.FUNNY-NAME.ts.net
```

> [!NOTE]
> Note that you need to replace `FUNNY-NAME` with the name of your network.

{{% /steps %}}
