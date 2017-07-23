### Mobile Api Server

The mobile api server is an extension to the Kubernetes api in order to represent mobile objects and apis.

### Install locally on oc cluster up
After cloning this repo take the following steps:

*Note* improvements will be made to this over time. It is currently just a POC

- Start your oc cluster with a host config and service catalog enabled. We enabled the service catalog so that Kubernetes and OpenShift are set up to recognise the ```apiregistration.k8s.io``` api.

```
oc cluster up --host-config-dir=${HOME}/openshift-config/openshift.local.config
```
- Login as system:admin

```
oc login -u system:admin
```

- Change the restricted scc to allow hostPath volumes. This is so we can mount the front-proxy-ca.crt which is needed for authentication. Plan to change how this is done in the future.

```
oc edit scc restricted

change: 

allowHostDirVolumePlugin:false 

to 

allowHostDirVolumePlugin:true

```

- Run the install script. Note there will likely be an error about roles already existing. This is fine.

```
./hack/install-apiserver/openshift/install.sh
```

- Check the api is set up. You should be able to create MobileApps and get MobileApps

```
oc login 
developer
developer

oc create -f ./hack/install-apiserver/openshift/MobileApp.json

oc get mobileapps
```