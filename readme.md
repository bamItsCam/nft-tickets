# nft-tickets

A Go microservice for issuing and managing event tickets as ERC-721 NFTs on a Hyperledger Besu chain. Tickets are minted on purchase, burned on redemption, and transferable — with support for both custodial and self-custodial wallets.

## Prerequisites

> WSL is strongly recommended on Windows.

- Go 1.26+ (try gvm)
- Docker + Docker Compose
- [Foundry](https://book.getfoundry.sh/getting-started/installation) (`forge`, `cast`)
- `abigen` (go-ethereum)
- `jq`
- A running Paladin/Besu kind cluster (see `paladin-kind.yaml`)

Install everything else:
```sh
sudo apt install build-essential jq
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
curl -L https://foundry.paradigm.xyz | bash && foundryup
```

Add to `~/.bashrc`:
```sh
export PATH="$PATH:/${HOME}/.foundry/bin:/usr/local/go/bin:${HOME}/go/bin"
```

## Usage

**1. Start the kind cluster** (Besu RPC available at `localhost:31545`):

[Consult this guide for setting this up](https://lfdt-paladin.github.io/paladin/head/getting-started/installation/)

**2. Deploy the contract**:
```sh
cast wallet new  # save the private key
export OWNER_KEY=<key>

(cd contracts && forge script script/Ticket.s.sol \
  --rpc-url http://localhost:31545 \
  --private-key $OWNER_KEY \
  --broadcast \
  --gas-price 0 \
  --legacy)
```

**3. Start the database and API**:
```sh
cp .env.example .env  # fill in CONTRACT_ADDRESS and OWNER_KEY from the steps above
docker compose up # --build # if you want to rebuild the api each time
```

**3.5 Developing w/out rebuilding docker every time
```sh
docker compose up db
go run .
```

**4. Run migrations**:
```sh
migrate -path db/migrations -database $DATABASE_URL up
```

The API and Swagger UI are available at `http://localhost:8080/openapi.json` and `http://localhost:8080/docs`.

### Self-custodial transfer

To transfer a ticket where the owner holds their own private key:
```sh
chmod +x scripts/self_custodial_transfer.sh
./scripts/self_custodial_transfer.sh <ticket_id> <to_address> <owner_private_key>
```

This calls `/tickets/{id}/transfer/prepare`, signs the transaction locally with `cast`, then broadcasts it via `/tickets/{id}/transfer/broadcast` to keep the DB in sync.

## Development

### Contract changes

After modifying `contracts/src/Ticket.sol`, rebuild and regenerate bindings:
```sh
(cd contracts && forge build)
abigen --v2 \
  --abi <(cat contracts/out/Ticket.sol/Ticket.json | jq .abi) \
  --bin <(cat contracts/out/Ticket.sol/Ticket.json | jq -r '.bytecode.object') \
  --pkg contract --type Ticket --out internal/contract/ticket.go
```

Redeploying the contract means existing on-chain tokens are gone. Wipe the DB to avoid stale references:
```sh
migrate -path db/migrations -database $DATABASE_URL drop -f
migrate -path db/migrations -database $DATABASE_URL up
```

### DB schema changes

Add a migration under `db/migrations/`, run it, then regenerate the query layer:
```sh
migrate -path db/migrations -database $DATABASE_URL up
sqlc generate
```

### Contract dependencies

Foundry deps are gitignored. Restore them on a fresh clone:
```sh
(cd contracts && forge install foundry-rs/forge-std && forge install OpenZeppelin/openzeppelin-contracts)
```

## Future Work Wishlist
There are many items I wish I could have tackled given enough time:
1. Better user auth and session management (beyond just `POST /user`)
2. Adding costs to tickets, not just minting them
3. Hooking in a "resale fee" where every transfer would mean sending some small % to the contract owner's wallet
4. a staging/marketplace pattern where you can post tickets for sale, as well as browse available tickets
5. Clearer distinctions between custodial and non-custodial. Right now only differentiation is on user creation if a wallet is passed into the request

and probably more I'm blanking on!
