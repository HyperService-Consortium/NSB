#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/usr/bin/tendermint
ID=${ID:-0}
LOG=${LOG:-tendermint.log}

echo current ID=${ID}
echo current LOG=${LOG}

# cp /usr/bin/NSB /usr/bin/NSB${ID}
if [ ! -f "/usr/bin/NSB" ]; then
	echo "file /usr/bin/NSB not found"
	exit 1
fi
nohup /usr/bin/NSB -db=$DB_DIR -server=$TCP_AD -port=$PORT &
##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'tendermint' E.g.: -e BINARY=tendermint_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Run binary with all parameters
##
export TMHOME="/tendermint/node${ID}"

if [ -d "`dirname ${TMHOME}/${LOG}`" ]; then
  echo running "$BINARY" "$@" | tee "${TMHOME}/${LOG}"
  "$BINARY" "$@" | tee "${TMHOME}/${LOG}"
else
  echo running "$BINARY" "$@"
  "$BINARY" "$@"
fi

chmod 777 -R /tendermint

