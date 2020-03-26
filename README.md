# k8s-tftp
Tftp server which can running on Kubernetes.
Repository: https://hub.docker.com/r/darkautism/k8s-tftp
Github: https://github.com/darkautism/k8s-tftp

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
