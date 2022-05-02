mkdir -p .freshcloud
cat <<EOF > .freshcloud/harbor-values.yaml
harborAdminPassword: {{.Password}}
service:
  type: ClusterIP
  tls:
    enabled: true
    existingSecret: harbor-tls-prod
    notaryExistingSecret: notary-tls-prod
ingress:
  enabled: true
  hosts:
    core: registry.{{.Domain}}
    notary: notary.{{.Domain}}
  annotations:
    cert-manager.io/cluster-issuer: letsencrypt-prod     # use letsencrypt-prod as the cluster issuer for TLS certs
    ingress.kubernetes.io/force-ssl-redirect: "true"     # force https, even if http is requested
    kubernetes.io/ingress.class: contour                 # using Contour for ingress
    kubernetes.io/tls-acme: "true"                       # using ACME certificates for TLS
externalURL: https://registry.{{.Domain}}
portal:
  tls:
    existingSecret: harbor-tls-prod
EOF
kubectl create namespace harbor
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install harbor bitnami/harbor -f .freshcloud/harbor-values.yaml -n harbor --version 11.2.4
if [ $? != 0 ]; then
  echo "Failed to install Harbor. Bummer"
  exit 1
fi
kubectl wait --for=condition=Ready pods --timeout=900s --all -n harbor
for REPO in {concourse-images,kpack}; do
  echo "Creating: ${REPO} in Harbor."
  curl --user "admin:{{.Password}}" -X POST \
    https://registry.{{.Domain}}/api/v2.0/projects \
    -H "Content-type: application/json" --data \
    '{ "project_name": "'${REPO}'",
    "metadata": {
    "auto_scan": "true",
    "enable_content_trust": "false",
    "prevent_vul": "false",
    "public": "true",
    "reuse_sys_cve_whitelist": "true",
    "severity": "high" }
    }'
done
cat << EOF
echo "Remove harbor by running - kubectl delete ns harbor"
url: https://registry.{{.Domain}}
username: admin
password: {{.Password}}
EOF