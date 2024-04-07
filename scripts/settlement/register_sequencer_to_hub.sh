#!/bin/bash

EXECUTABLE="rollappd"
KEYRING_PATH="$HOME/.rollapp/sequencer_keys"
KEY_NAME_SEQUENCER="sequencer"

#Register Sequencer
DESCRIPTION="{\"Moniker\":\"${ROLLAPP_CHAIN_ID}-sequencer\",\"Identity\":\"\",\"Website\":\"\",\"SecurityContact\":\"\",\"Details\":\"\"}"
SEQ_PUB_KEY="$($EXECUTABLE dymint show-sequencer)"
BOND_AMOUNT="$(dymd q sequencer params -o json | jq -r '.params.min_bond.amount')$(dymd q sequencer params -o json | jq -r '.params.min_bond.denom')"

set -x
dymd tx sequencer create-sequencer "$SEQ_PUB_KEY" "$ROLLAPP_CHAIN_ID" "$DESCRIPTION" "$BOND_AMOUNT" \
  --from "$KEY_NAME_SEQUENCER" \
  --keyring-dir "$KEYRING_PATH" \
  --keyring-backend test \
  --broadcast-mode block \
  --fees 1dym \
  -y

set +x
