#!/bin/sh

#   __        ___    _   _  ____ _           _       ____             
#   \ \      / / \  | \ | |/ ___| |__   __ _(_)_ __ |  _ \  _____   __
#    \ \ /\ / / _ \ |  \| | |   | '_ \ / _` | | '_ \| | | |/ _ \ \ / /
#     \ V  V / ___ \| |\  | |___| | | | (_| | | | | | |_| |  __/\ V / 
#      \_/\_/_/   \_\_| \_|\____|_| |_|\__,_|_|_| |_|____/ \___| \_/  
#                                                                     

echo "run gwan in pluto bootnode testnet"
SCRIPT_DIR=$(dirname "$0")
make && \
rm -rf ~/.wanchain/pluto/gwan && \
echo -n 'Wanchain' > pw.txt && \
mkdir -p  ~/.wanchain/pluto/keystore && \
cp -f ${SCRIPT_DIR}/UTC--2017-05-14T03-13-33.929385593Z--2d0e7c0813a51d3bd1d08246af2a8a7a57d8922e ~/.wanchain/pluto/keystore/ && \
build/bin/gwan --pluto --http --allow-insecure-unlock     --nodiscover --mine --miner.etherbase  "0x2d0e7c0813a51d3bd1d08246af2a8a7a57d8922e"  --unlock "0x2d0e7c0813a51d3bd1d08246af2a8a7a57d8922e" --password ./pw.txt --miner.threads=1 $@

