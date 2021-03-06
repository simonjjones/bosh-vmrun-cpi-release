#!/bin/bash
set -o errexit
set -o pipefail
set -o nounset

cd $(dirname $0)
RELEASE_DIR=../
: ${STATE_DIR:="$PWD/state"}
if ! [ -f $STATE_DIR/env.sh ]; then
  echo "no $STATE_DIR/env.sh file. Create and fill with required fields"
  exit 1
fi

source $STATE_DIR/env.sh
: ${VMRUN_BIN_PATH?"!"}
: ${OVFTOOL_BIN_PATH?"!"}
: ${VDISKMANAGER_BIN_PATH?"!"}
: ${VMRUN_NETWORK:?"!"}
: ${WEB_IP?"!"}
: ${WORKER_IP?"!"}
: ${NETWORK_CIDR:?"!"}
: ${NETWORK_GW:?"!"}
: ${NETWORK_DNS:?"!"}
: ${WINDOWS_STEMCELL:?"!"}
: ${LINUX_STEMCELL:?"!"}

if [ -n ${RESET:-""} ]; then
  RECREATE_RELEASE="y"
  RECREATE_VM="y"
fi

bosh_cli_linux_url="https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-5.1.1-linux-amd64"
bosh_cli_darwin_url="https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-5.1.1-darwin-amd64"
bosh_bin="bin/bosh-$OSTYPE"
if ! [ -f $bosh_bin ]; then
  curl -L $bosh_cli_linux_url > bin/bosh-linux-gnu
  curl -L $bosh_cli_darwin_url > bin/bosh-darwin17
  chmod +x bin/bosh*
fi

concourse_bosh_deployment_url="https://github.com/concourse/concourse-bosh-deployment.git"
if ! [ -d $STATE_DIR/concourse-bosh-deployment ]; then
  git clone $concourse_bosh_deployment_url $STATE_DIR/concourse-bosh-deployment
  pushd $STATE_DIR/concourse-bosh-deployment
    git checkout 14323c4cc2320107dfb2dc622bca0d6861c517d4
  popd
fi

if ! [ -f $LINUX_STEMCELL ]; then
  echo "Error: linux stemcell is required. Downlaod manually"
fi

if ! [ -f $WINDOWS_STEMCELL ]; then
  echo "Error: windows stemcell is required. Downlaod manually"
fi

if [ -n ${RECREATE_RELEASE:-""} -o ! -f $STATE_DIR/cpi.tgz ] ; then
  echo "-----> `date`: Create dev release"
  HOME=$STATE_DIR/bosh_home \
    $bosh_bin create-release --sha2 --force --dir $RELEASE_DIR --tarball $STATE_DIR/cpi.tgz
fi

echo "-----> `date`: Deploy Start"

vm_store_path=$STATE_DIR/vm-store-path
if ! [ -d $vm_store_path ]; then
  mkdir -p $vm_store_path
fi

linux_stemcell_sha1=$(shasum -a1 < $LINUX_STEMCELL | awk '{print $1}')
windows_stemcell_sha1=$(shasum -a1 < $WINDOWS_STEMCELL | awk '{print $1}')

HOME=$STATE_DIR/bosh_home \
$bosh_bin interpolate $STATE_DIR/concourse-bosh-deployment/lite/concourse.yml \
  -o concourse-vars-opsfile.yml \
  -v web_ip="$WEB_IP" \
  -v worker_ip="$WORKER_IP" \
  --vars-store $STATE_DIR/concourse-creds.yml \
> /dev/null;

web_mbus_bootstrap_ssl="$($bosh_bin int $STATE_DIR/concourse-creds.yml --path /web_mbus_bootstrap_ssl)"
worker_mbus_bootstrap_ssl="$($bosh_bin int $STATE_DIR/concourse-creds.yml --path /worker_mbus_bootstrap_ssl)"

cpi_url=file://$STATE_DIR/cpi.tgz
cpi_sha1=$(shasum -a1 < $STATE_DIR/cpi.tgz | awk '{print $1}')

HOME=$STATE_DIR/bosh_home \
$bosh_bin ${BOSH_COMMAND:-"create-env"} $STATE_DIR/concourse-bosh-deployment/lite/concourse.yml \
  -o concourse-vmrun-opsfile.yml \
  -o concourse-vmrun-web-opsfile.yml \
  --vars-file $STATE_DIR/concourse-bosh-deployment/versions.yml \
  --vars-file $STATE_DIR/concourse-creds.yml \
  --state $STATE_DIR/concourse_web_state.json \
  -v cpi_url=$cpi_url \
  -v cpi_sha1=$cpi_sha1 \
  -v public_ip="$WEB_IP" \
  -v internal_ip="$WEB_IP" \
  -v internal_cidr="$NETWORK_CIDR" \
  -v internal_gw="$NETWORK_GW" \
  -v stemcell_url=file://$LINUX_STEMCELL \
  -v stemcell_sha1=$linux_stemcell_sha1 \
  -v network_name="$VMRUN_NETWORK" \
  -v vm_store_path="$vm_store_path" \
  -v vmrun_bin_path="$VMRUN_BIN_PATH" \
  -v ovftool_bin_path="$OVFTOOL_BIN_PATH" \
  -v vdiskmanager_bin_path="$VDISKMANAGER_BIN_PATH" \
  -v mbus_bootstrap_ssl="$web_mbus_bootstrap_ssl" \
  ${RECREATE_VM:+"--recreate"} \
  ;

HOME=$STATE_DIR/bosh_home \
$bosh_bin ${BOSH_COMMAND:-"create-env"} $STATE_DIR/concourse-bosh-deployment/lite/concourse.yml \
  -o concourse-vmrun-opsfile.yml \
  -o concourse-vmrun-windows-worker-opsfile.yml \
  --vars-file $STATE_DIR/concourse-bosh-deployment/versions.yml \
  --vars-file $STATE_DIR/concourse-creds.yml \
  --state $STATE_DIR/concourse_worker_state.json \
  -v cpi_url=file://$STATE_DIR/cpi.tgz \
  -v web_ip="$WEB_IP" \
  -v public_ip="$WORKER_IP" \
  -v internal_ip="$WORKER_IP" \
  -v internal_cidr="$NETWORK_CIDR" \
  -v internal_gw="$NETWORK_GW" \
  -v stemcell_url=file://$WINDOWS_STEMCELL \
  -v stemcell_sha1=$windows_stemcell_sha1 \
  -v network_name="$VMRUN_NETWORK" \
  -v vm_store_path="$vm_store_path" \
  -v vmrun_bin_path="$VMRUN_BIN_PATH" \
  -v ovftool_bin_path="$OVFTOOL_BIN_PATH" \
  -v vdiskmanager_bin_path="$VDISKMANAGER_BIN_PATH" \
  -v mbus_bootstrap_ssl="$worker_mbus_bootstrap_ssl" \
  ${RECREATE_VM:+"--recreate"} \
  ;

echo "-----> `date`: Deploy Done"
