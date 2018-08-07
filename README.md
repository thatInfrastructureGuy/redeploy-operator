# redeploy-operator
Kubernetes Redeploy Operator to redeploy pods in a deployment in the form of rolling updates. 

This project is built with help of [Operator-SDK](https://github.com/operator-framework/operator-sdk).

### How it works
Redeploy Operator basically adds a label to PodSpec field in deployments. 
This triggers a deployment rollout. Type of rollout depends on `MaxSurge` and `MaxUnavailable` fields in your deployments.

The label you will see added is the datetime field.
Eg:   `redeployed: 2018-08-07-05_28_41`

### Limitations:
1. It only works for deployments. ( Statefulsets, replicasets and daemonsets are not covered. )

### Installation

1. Install CRD, rbac and operator.

```
kubectl apply -f deploy/crd.yaml
kubectl apply -f deploy/rbac.yaml
kubectl apply -f deploy/operator.yaml
```
2. Install a demo app:

```
kubectl create namespace nginx
kubectl run nginx --namespace nginx --image nginx:alpine --replicas=2
```
3. Trigger a redeploy:

```
kubectl apply -f deploy/cr.yaml
```

4. See it work:

```
kubectl -n nginx get pods -w
kubectl -n redeploy-operator logs -l name=redeploy-operator
```


### Building from source:

1. Install [Operator-SDK](https://github.com/operator-framework/operator-sdk)
2. `operator-sdk generate k8s`
3. `operator-sdk build <IMAGE-NAME>`
4. `docker push <IMAGE-NAME>`

Note: operator.yaml is custom built. It is overwritten upon building the project. 
Please make necessary changes to make it work.

