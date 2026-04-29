# Dev setup
just use `wsl` and save yourself the headache

Go
```sh
VERSION=$(curl -s https://go.dev/VERSION?m=text | head -1)
curl -L https://go.dev/dl/${VERSION}.linux-amd64.tar.gz -o go.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go.tar.gz
rm go.tar.gz
```

other stuff
```sh
sudo apt install build-essential
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
curl -L https://foundry.paradigm.xyz | bash
foundryup
sudo apt install jq
```

update your path in `~/.bashrc`
```sh
export PATH="$PATH:/${HOME}/.foundry/bin:/usr/local/go/bin:${HOME}/go/bin"
```

## Development

### contract changes
```sh
mkdir -p internal/contract
(cd contracts/ && forge build)
abigen --v2 --abi <(cat contracts/out/Ticket.sol/Ticket.json | jq .abi) --bin <(cat contracts/out/Ticket.sol/Ticket.json | jq -r '.bytecode.object') --pkg contract --type Ticket --out internal/contract/ticket.go
```

Get a wallet to deploy the contract
```sh
cast wallet new
export OWNER_KEY=<the key>
```

contract deploy
```sh
(cd contracts && forge script script/Ticket.s.sol \
  --rpc-url http://localhost:31545 \
  --private-key $OWNER_KEY \
  --broadcast \
  --gas-price 0 \
  --legacy)
export CONTRACT_ADDRESS=<the contract address>
```

### DB migrations
```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

```sh
migrate -path db/migrations -database $DATABASE_URL up
```

```sh
sqlc generate
```
