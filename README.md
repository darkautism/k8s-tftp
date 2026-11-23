# k8s-tftp
Tftp server which can running on Kubernetes.
Repository: https://hub.docker.com/r/darkautism/k8s-tftp
Github: https://github.com/darkautism/k8s-tftp

## Image tags

All images are published to [Docker Hub](https://hub.docker.com/r/darkautism/k8s-tftp).

| Tag | Platforms | Description |
| --- | --- | --- |
| `latest` | linux/amd64, linux/arm64 | Latest build from `master`; Docker chooses the right architecture |
| `latest-amd64` | linux/amd64 | Latest AMD64 image |
| `latest-arm64` | linux/arm64 | Latest ARM64 image |
| `sha-<full-commit-sha>` | linux/amd64, linux/arm64 | Immutable build reference for a Git commit |
| `sha-<full-commit-sha>-amd64` | linux/amd64 | AMD64 image for the commit |
| `sha-<full-commit-sha>-arm64` | linux/arm64 | ARM64 image for the commit |

Pushing a Git version tag such as `v1.2.3` also publishes Docker tags `1.2.3`, `1.2`, and `1`, together with `-amd64` and `-arm64` variants. Version aliases point to their latest matching release; use a full version or commit SHA for reproducible deployments. Version-tag builds do not move `latest`, which tracks `master`.

For most users, simply use `darkautism/k8s-tftp:latest`.

## How to use

Create a tftp.yaml like this

``` yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: pxe
  name: pxe-deployment
spec:
  selector:
    matchLabels:
      app: pxe
  template:
    metadata:
      labels:
        app: pxe
    spec:
      containers:
      - name: pxe
        image: darkautism/k8s-tftp
        ports:
        - containerPort: 69
        volumeMounts:
          - name: nfs
            mountPath: /tftpboot
      volumes:
      - name: nfs
        persistentVolumeClaim:
          claimName: YOUR TFTP FILES PVC!!!
---
apiVersion: v1
kind: Service
metadata:
  name: pxe-deployment
  namespace: pxe
spec:
  externalIPs:
  - IP YOU WANT
  ports:
  - port: 69
    protocol: UDP
    targetPort: 69
  selector:
    app: pxe
  sessionAffinity: None
  type: LoadBalancer
```
