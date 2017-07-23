#!/bin/bash
#oc cluster up --host-config-dir=/Users/kelly/openshift-config/openshift.local.config --service-catalog=true


targetNamespace=mobile
apiserverConfigDir=/tmp/mobile-apiserver/config
masterConfigDir=${HOME}/openshift-config/openshift.local.config/master
mkdir -p ${apiserverConfigDir} || true
#nodeManifestDir=/Users/kelly/openshift-config/openshift.local.config/node-localhost/static-pods


# 1.  Installer creates namespace of its choice
oc create namespace ${targetNamespace}

oc adm policy add-cluster-role-to-user system:auth-delegator -n ${targetNamespace} -z apiserver
oc create policybinding kube-system -n kube-system
oc adm policy add-role-to-user extension-apiserver-authentication-reader -n kube-system --role-namespace=kube-system system:serviceaccount:${targetNamespace}:apiserver
oc new-app -f hack/install-apiserver/openshift/deployment.json -n ${targetNamespace}
oc create -f hack/install-apiserver/openshift/apiservice.json
# 2.  Installer creates `api` service with a fixed selector (app: api) with service serving cert annotation
#oc -n ${targetNamespace} create service clusterip api --tcp=443:3101
#oc -n ${targetNamespace} annotate svc/api service.alpha.openshift.io/serving-cert-secret-name=api-serving-cert
until oc -n ${targetNamespace} get secrets/api-serving-cert; do
	echo "waiting for oc -n ${targetNamespace} get secrets/api-serving-cert"
	sleep 1
done

# 3.  Installer creates `apiserver` service account
#oc  -n ${targetNamespace} create sa apiserver
#until oc -n ${targetNamespace} sa get-token apiserver; do
#	echo "waiting for oc -n ${targetNamespace} get secrets/api-serving-cert"
#	sleep 1
#done

# 4.  Installer harvests serving cert
oc -n ${targetNamespace} extract secret/api-serving-cert --to=${apiserverConfigDir}
mv ${apiserverConfigDir}/tls.crt ${apiserverConfigDir}/serving.crt
mv ${apiserverConfigDir}/tls.key ${apiserverConfigDir}/serving.key

# 5.  Installer harvests SA token
#saToken=$(oc -n ${targetNamespace} sa get-token apiserver)


cp ${masterConfigDir}/front-proxy-ca.crt ${apiserverConfigDir}/

# 6.  Installer creates SA kubeconfig
# TODO do this a LOT better
# start with admin.kubeconfig
#cp ${masterConfigDir}/admin.kubeconfig ${apiserverConfigDir}/kubeconfig
# remove all users
#oc --config=${apiserverConfigDir}/kubeconfig config unset users
# set the service account token
#configContext=$(oc --config=${apiserverConfigDir}/kubeconfig config current-context)
#oc --config=${apiserverConfigDir}/kubeconfig config set-credentials serviceaccount --token=${saToken}
#oc --config=${apiserverConfigDir}/kubeconfig config set-context ${configContext} --user=serviceaccount

# 7.  Installer binds “known” roles to SA user
# TODO remove this bit once we bootstrap these roles
#oc create -f hack/install-apiserver/openshift/prestart.json || true

#oc adm policy add-cluster-role-to-user system:auth-delegator -n ${targetNamespace} -z apiserver
#oc create policybinding kube-system -n kube-system
#oc adm policy add-role-to-user extension-apiserver-authentication-reader -n kube-system --role-namespace=kube-system system:serviceaccount:${targetNamespace}:apiserver


# 8.  Installer provides the following information to all templates.
#cp ${masterConfigDir}/ca.crt ${apiserverConfigDir}/etcd-ca.crt
#cp ${masterConfigDir}/ca-bundle.crt ${apiserverConfigDir}/client-ca.crt
#cp ${masterConfigDir}/master.etcd-client.crt ${apiserverConfigDir}/etcd-write.crt
#cp ${masterConfigDir}/master.etcd-client.key ${apiserverConfigDir}/etcd-write.key
#oc export ClusterRole edit > hack/install-apiserver/openshift/edit-role.yaml



# update the edit role to allow users to create mobileapps (prob beter way to do this)
#oc replace ClusterRole/edit -f hack/install-apiserver/openshift/edit-role.yaml
#oc new-app -f hack/install-apiserver/openshift/deployment.json
# register the new api
#oc create -f hack/install-apiserver/openshift/apiservice.json
#create the api server
#oc create -f hack/install-apiserver/openshift/rc.json