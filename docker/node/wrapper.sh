#!/usr/bin/env sh

##
## Input parameters
##
TENDERMINT_BINARY=/tendermint/tendermint
NSB_BINARY=/usr/bin/NSB
ID=${ID:-0}
LOG=${LOG:-tendermint.log}

echo current ID=${ID}
echo current LOG=${LOG}

if [ ! -f "${NSB_BINARY}" ]; then
	echo "file ${NSB_BINARY} not found"
	exit 1
fi
nohup ${NSB_BINARY} -db=$DB_DIR -server=$TCP_AD -port=$PORT &
##
## Assert linux binary
##
if ! [ -f "${TENDERMINT_BINARY}" ]; then
	echo "The binary $(basename "${TENDERMINT_BINARY}") cannot be found. Please add the binary to the shared folder. Please use the TENDERMINT_BINARY environment variable if the name of the binary is not 'tendermint' E.g.: -e TENDERMINT_BINARY=tendermint_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$TENDERMINT_BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Run binary with all parameters
##
export TMHOME="/tendermint/node${ID}"

if [ -d "`dirname ${TMHOME}/${LOG}`" ]; then
  echo running "$TENDERMINT_BINARY" "$@" | tee "${TMHOME}/${LOG}"
  "$TENDERMINT_BINARY" "$@" | tee "${TMHOME}/${LOG}"
else
  echo running "$TENDERMINT_BINARY" "$@"
  "$TENDERMINT_BINARY" "$@"
fi

chmod 777 -R /tendermint

