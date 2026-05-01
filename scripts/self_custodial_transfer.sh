#!/bin/bash
set -e

TICKET_ID=$1
TO_ADDRESS=$2
PRIVATE_KEY=$3 # key from the owner

if [ -z "$TICKET_ID" ] || [ -z "$TO_ADDRESS" ] || [ -z "$PRIVATE_KEY" ]; then
  echo "Usage: self_custodial_transfer.sh <ticket_id> <to_address> <private_key>"
  exit 1
fi

API_URL=${API_URL:-http://localhost:8080}
RPC_URL=${RPC_URL:-http://localhost:8545}

echo "Preparing transfer for ticket $TICKET_ID to $TO_ADDRESS..."
PREP=$(curl -sf "$API_URL/tickets/$TICKET_ID/transfer/prepare?to=$TO_ADDRESS")
if [ $? -ne 0 ]; then
  echo "Failed to prepare transfer"
  exit 1
fi

CONTRACT=$(echo "$PREP" | jq -r '.to')
FROM_ADDRESS=$(echo "$PREP" | jq -r '.from')
NONCE=$(echo "$PREP" | jq -r '.nonce')
GAS_LIMIT=$(echo "$PREP" | jq -r '.gasLimit')

echo "Signing transaction..."
SIGNED_TX=$(cast mktx "$CONTRACT" "transferFrom(address,address,uint256)" "$FROM_ADDRESS" "$TO_ADDRESS" "$TICKET_ID" \
  --private-key "$PRIVATE_KEY" \
  --rpc-url "$RPC_URL" \
  --nonce "$NONCE" \
  --gas-limit "$GAS_LIMIT" \
  --gas-price 0 \
  --legacy)

echo "Broadcasting..."
RESULT=$(curl -sf -X POST "$API_URL/tickets/$TICKET_ID/transfer/broadcast" \
  -H "Content-Type: application/json" \
  -d "{\"signedTx\": \"$SIGNED_TX\"}")

echo "Transfer complete:"
echo "$RESULT" | jq .
