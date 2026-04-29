// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// TicketMetaData contains all meta data concerning the Ticket contract.
var TicketMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorBurn\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorTransfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"name\":\"_fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	ID:  "Ticket",
	Bin: "0x608060405234801561000f575f5ffd5b50604051612d93380380612d9383398181016040528101906100319190610266565b806040518060400160405280600681526020017f5469636b657400000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f54434b0000000000000000000000000000000000000000000000000000000000815250815f90816100ac91906104ce565b5080600190816100bc91906104ce565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361012f575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161012691906105ac565b60405180910390fd5b61013e8161014560201b60201c565b50506105c5565b5f60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102358261020c565b9050919050565b6102458161022b565b811461024f575f5ffd5b50565b5f815190506102608161023c565b92915050565b5f6020828403121561027b5761027a610208565b5b5f61028884828501610252565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061030c57607f821691505b60208210810361031f5761031e6102c8565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610346565b61038b8683610346565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103cf6103ca6103c5846103a3565b6103ac565b6103a3565b9050919050565b5f819050919050565b6103e8836103b5565b6103fc6103f4826103d6565b848454610352565b825550505050565b5f5f905090565b610413610404565b61041e8184846103df565b505050565b5b81811015610441576104365f8261040b565b600181019050610424565b5050565b601f8211156104865761045781610325565b61046084610337565b8101602085101561046f578190505b61048361047b85610337565b830182610423565b50505b505050565b5f82821c905092915050565b5f6104a65f198460080261048b565b1980831691505092915050565b5f6104be8383610497565b9150826002028217905092915050565b6104d782610291565b67ffffffffffffffff8111156104f0576104ef61029b565b5b6104fa82546102f5565b610505828285610445565b5f60209050601f831160018114610536575f8415610524578287015190505b61052e85826104b3565b865550610595565b601f19841661054486610325565b5f5b8281101561056b57848901518255600182019150602085019450602081019050610546565b868310156105885784890151610584601f891682610497565b8355505b6001600288020188555050505b505050505050565b6105a68161022b565b82525050565b5f6020820190506105bf5f83018461059d565b92915050565b6127c1806105d25f395ff3fe608060405234801561000f575f5ffd5b506004361061012a575f3560e01c8063715018a6116100ab578063c87b56dd1161006f578063c87b56dd14610316578063d3fc986414610346578063e985e9c514610362578063ebce9ccf14610392578063f2fde38b146103ae5761012a565b8063715018a6146102985780638da5cb5b146102a257806395d89b41146102c0578063a22cb465146102de578063b88d4fde146102fa5761012a565b806323b872dd116100f257806323b872dd146101e457806342842e0e1461020057806342966c681461021c5780636352211e1461023857806370a08231146102685761012a565b806301ffc9a71461012e57806306fdde031461015e578063081812fc1461017c578063095ea7b3146101ac5780630d1af103146101c8575b5f5ffd5b61014860048036038101906101439190611cf3565b6103ca565b6040516101559190611d38565b60405180910390f35b6101666103db565b6040516101739190611dc1565b60405180910390f35b61019660048036038101906101919190611e14565b61046a565b6040516101a39190611e7e565b60405180910390f35b6101c660048036038101906101c19190611ec1565b610485565b005b6101e260048036038101906101dd9190611eff565b61049b565b005b6101fe60048036038101906101f99190611eff565b6104b3565b005b61021a60048036038101906102159190611eff565b6105b2565b005b61023660048036038101906102319190611e14565b6105d1565b005b610252600480360381019061024d9190611e14565b6105e7565b60405161025f9190611e7e565b60405180910390f35b610282600480360381019061027d9190611f4f565b6105f8565b60405161028f9190611f89565b60405180910390f35b6102a06106ae565b005b6102aa6106c1565b6040516102b79190611e7e565b60405180910390f35b6102c86106e9565b6040516102d59190611dc1565b60405180910390f35b6102f860048036038101906102f39190611fcc565b610779565b005b610314600480360381019061030f9190612136565b61078f565b005b610330600480360381019061032b9190611e14565b6107b4565b60405161033d9190611dc1565b60405180910390f35b610360600480360381019061035b9190612254565b6107c6565b005b61037c600480360381019061037791906122c0565b6107e7565b6040516103899190611d38565b60405180910390f35b6103ac60048036038101906103a79190611e14565b610875565b005b6103c860048036038101906103c39190611f4f565b610889565b005b5f6103d48261090d565b9050919050565b60605f80546103e99061232b565b80601f01602080910402602001604051908101604052809291908181526020018280546104159061232b565b80156104605780601f1061043757610100808354040283529160200191610460565b820191905f5260205f20905b81548152906001019060200180831161044357829003601f168201915b5050505050905090565b5f6104748261096d565b5061047e826109f3565b9050919050565b6104978282610492610a2c565b610a33565b5050565b6104a3610a45565b6104ae838383610acc565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610523575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161051a9190611e7e565b60405180910390fd5b5f6105368383610531610a2c565b610c34565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146105ac578382826040517f64283d7b0000000000000000000000000000000000000000000000000000000081526004016105a39392919061235b565b60405180910390fd5b50505050565b6105cc83838360405180602001604052805f81525061078f565b505050565b6105e35f826105de610a2c565b610c34565b5050565b5f6105f18261096d565b9050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610669575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016106609190611e7e565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6106b6610a45565b6106bf5f610e3f565b565b5f60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600180546106f89061232b565b80601f01602080910402602001604051908101604052809291908181526020018280546107249061232b565b801561076f5780601f106107465761010080835404028352916020019161076f565b820191905f5260205f20905b81548152906001019060200180831161075257829003601f168201915b5050505050905090565b61078b610784610a2c565b8383610f02565b5050565b61079a8484846104b3565b6107ae6107a5610a2c565b8585858561106b565b50505050565b60606107bf82611217565b9050919050565b6107ce610a45565b6107d88383611322565b6107e2828261133f565b505050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b61087d610a45565b61088681611399565b50565b610891610a45565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610901575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016108f89190611e7e565b60405180910390fd5b61090a81610e3f565b50565b5f634906490660e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061096657506109658261141b565b5b9050919050565b5f5f610978836114fc565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109ea57826040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016109e19190611f89565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610a408383836001611535565b505050565b610a4d610a2c565b73ffffffffffffffffffffffffffffffffffffffff16610a6b6106c1565b73ffffffffffffffffffffffffffffffffffffffff1614610aca57610a8e610a2c565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610ac19190611e7e565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b3c575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610b339190611e7e565b60405180910390fd5b5f610b4883835f610c34565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bba57816040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610bb19190611f89565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c2e578382826040517f64283d7b000000000000000000000000000000000000000000000000000000008152600401610c259392919061235b565b60405180910390fd5b50505050565b5f5f610c3f846114fc565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610c8057610c7f8184866116f4565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610d0b57610cbf5f855f5f611535565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610d8a57600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610f7257816040517f5b08ba18000000000000000000000000000000000000000000000000000000008152600401610f699190611e7e565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405161105e9190611d38565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115611210578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b81526004016110c994939291906123e2565b6020604051808303815f875af192505050801561110457506040513d601f19601f820116820180604052508101906111019190612440565b60015b611185573d805f8114611132576040519150601f19603f3d011682016040523d82523d5f602084013e611137565b606091505b505f81510361117d57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016111749190611e7e565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461120e57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016112059190611e7e565b60405180910390fd5b505b5050505050565b60606112228261096d565b505f60065f8481526020019081526020015f2080546112409061232b565b80601f016020809104026020016040519081016040528092919081815260200182805461126c9061232b565b80156112b75780601f1061128e576101008083540402835291602001916112b7565b820191905f5260205f20905b81548152906001019060200180831161129a57829003601f168201915b505050505090505f6112c76117b7565b90505f8151036112db57819250505061131d565b5f8251111561130f5780826040516020016112f79291906124a5565b6040516020818303038152906040529250505061131d565b611318846117cd565b925050505b919050565b61133b828260405180602001604052805f815250611833565b5050565b8060065f8481526020019081526020015f20908161135d9190612668565b507ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce78260405161138d9190611f89565b60405180910390a15050565b5f6113a55f835f610c34565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361141757816040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161140e9190611f89565b60405180910390fd5b5050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806114e557507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806114f557506114f482611856565b5b9050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061156d57505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1561169f575f61157c8461096d565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156115e657508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156115f957506115f781846107e7565b155b1561163b57826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016116329190611e7e565b60405180910390fd5b811561169d57838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6116ff8383836118bf565b6117b2575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361177357806040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161176a9190611f89565b60405180910390fd5b81816040517f177e802f0000000000000000000000000000000000000000000000000000000081526004016117a9929190612737565b60405180910390fd5b505050565b606060405180602001604052805f815250905090565b60606117d88261096d565b505f6117e26117b7565b90505f8151116118005760405180602001604052805f81525061182b565b8061180a8461197f565b60405160200161181b9291906124a5565b6040516020818303038152906040525b915050919050565b61183d8383611a49565b611851611848610a2c565b5f85858561106b565b505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561197657508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480611937575061193684846107e7565b5b8061197557508273ffffffffffffffffffffffffffffffffffffffff1661195d836109f3565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b60605f600161198d84611b3c565b0190505f8167ffffffffffffffff8111156119ab576119aa612012565b5b6040519080825280601f01601f1916602001820160405280156119dd5781602001600182028036833780820191505090505b5090505f82602001820190505b600115611a3e578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8581611a3357611a3261275e565b5b0494505f85036119ea575b819350505050919050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611ab9575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401611ab09190611e7e565b60405180910390fd5b5f611ac583835f610c34565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611b37575f6040517f73c6ac6e000000000000000000000000000000000000000000000000000000008152600401611b2e9190611e7e565b60405180910390fd5b505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611b98577a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008381611b8e57611b8d61275e565b5b0492506040810190505b6d04ee2d6d415b85acef81000000008310611bd5576d04ee2d6d415b85acef81000000008381611bcb57611bca61275e565b5b0492506020810190505b662386f26fc100008310611c0457662386f26fc100008381611bfa57611bf961275e565b5b0492506010810190505b6305f5e1008310611c2d576305f5e1008381611c2357611c2261275e565b5b0492506008810190505b6127108310611c52576127108381611c4857611c4761275e565b5b0492506004810190505b60648310611c755760648381611c6b57611c6a61275e565b5b0492506002810190505b600a8310611c84576001810190505b80915050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611cd281611c9e565b8114611cdc575f5ffd5b50565b5f81359050611ced81611cc9565b92915050565b5f60208284031215611d0857611d07611c96565b5b5f611d1584828501611cdf565b91505092915050565b5f8115159050919050565b611d3281611d1e565b82525050565b5f602082019050611d4b5f830184611d29565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611d9382611d51565b611d9d8185611d5b565b9350611dad818560208601611d6b565b611db681611d79565b840191505092915050565b5f6020820190508181035f830152611dd98184611d89565b905092915050565b5f819050919050565b611df381611de1565b8114611dfd575f5ffd5b50565b5f81359050611e0e81611dea565b92915050565b5f60208284031215611e2957611e28611c96565b5b5f611e3684828501611e00565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611e6882611e3f565b9050919050565b611e7881611e5e565b82525050565b5f602082019050611e915f830184611e6f565b92915050565b611ea081611e5e565b8114611eaa575f5ffd5b50565b5f81359050611ebb81611e97565b92915050565b5f5f60408385031215611ed757611ed6611c96565b5b5f611ee485828601611ead565b9250506020611ef585828601611e00565b9150509250929050565b5f5f5f60608486031215611f1657611f15611c96565b5b5f611f2386828701611ead565b9350506020611f3486828701611ead565b9250506040611f4586828701611e00565b9150509250925092565b5f60208284031215611f6457611f63611c96565b5b5f611f7184828501611ead565b91505092915050565b611f8381611de1565b82525050565b5f602082019050611f9c5f830184611f7a565b92915050565b611fab81611d1e565b8114611fb5575f5ffd5b50565b5f81359050611fc681611fa2565b92915050565b5f5f60408385031215611fe257611fe1611c96565b5b5f611fef85828601611ead565b925050602061200085828601611fb8565b9150509250929050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61204882611d79565b810181811067ffffffffffffffff8211171561206757612066612012565b5b80604052505050565b5f612079611c8d565b9050612085828261203f565b919050565b5f67ffffffffffffffff8211156120a4576120a3612012565b5b6120ad82611d79565b9050602081019050919050565b828183375f83830152505050565b5f6120da6120d58461208a565b612070565b9050828152602081018484840111156120f6576120f561200e565b5b6121018482856120ba565b509392505050565b5f82601f83011261211d5761211c61200a565b5b813561212d8482602086016120c8565b91505092915050565b5f5f5f5f6080858703121561214e5761214d611c96565b5b5f61215b87828801611ead565b945050602061216c87828801611ead565b935050604061217d87828801611e00565b925050606085013567ffffffffffffffff81111561219e5761219d611c9a565b5b6121aa87828801612109565b91505092959194509250565b5f67ffffffffffffffff8211156121d0576121cf612012565b5b6121d982611d79565b9050602081019050919050565b5f6121f86121f3846121b6565b612070565b9050828152602081018484840111156122145761221361200e565b5b61221f8482856120ba565b509392505050565b5f82601f83011261223b5761223a61200a565b5b813561224b8482602086016121e6565b91505092915050565b5f5f5f6060848603121561226b5761226a611c96565b5b5f61227886828701611ead565b935050602061228986828701611e00565b925050604084013567ffffffffffffffff8111156122aa576122a9611c9a565b5b6122b686828701612227565b9150509250925092565b5f5f604083850312156122d6576122d5611c96565b5b5f6122e385828601611ead565b92505060206122f485828601611ead565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061234257607f821691505b602082108103612355576123546122fe565b5b50919050565b5f60608201905061236e5f830186611e6f565b61237b6020830185611f7a565b6123886040830184611e6f565b949350505050565b5f81519050919050565b5f82825260208201905092915050565b5f6123b482612390565b6123be818561239a565b93506123ce818560208601611d6b565b6123d781611d79565b840191505092915050565b5f6080820190506123f55f830187611e6f565b6124026020830186611e6f565b61240f6040830185611f7a565b818103606083015261242181846123aa565b905095945050505050565b5f8151905061243a81611cc9565b92915050565b5f6020828403121561245557612454611c96565b5b5f6124628482850161242c565b91505092915050565b5f81905092915050565b5f61247f82611d51565b612489818561246b565b9350612499818560208601611d6b565b80840191505092915050565b5f6124b08285612475565b91506124bc8284612475565b91508190509392505050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026125247fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826124e9565b61252e86836124e9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61256961256461255f84611de1565b612546565b611de1565b9050919050565b5f819050919050565b6125828361254f565b61259661258e82612570565b8484546124f5565b825550505050565b5f5f905090565b6125ad61259e565b6125b8818484612579565b505050565b5b818110156125db576125d05f826125a5565b6001810190506125be565b5050565b601f821115612620576125f1816124c8565b6125fa846124da565b81016020851015612609578190505b61261d612615856124da565b8301826125bd565b50505b505050565b5f82821c905092915050565b5f6126405f1984600802612625565b1980831691505092915050565b5f6126588383612631565b9150826002028217905092915050565b61267182611d51565b67ffffffffffffffff81111561268a57612689612012565b5b612694825461232b565b61269f8282856125df565b5f60209050601f8311600181146126d0575f84156126be578287015190505b6126c8858261264d565b86555061272f565b601f1984166126de866124c8565b5f5b82811015612705578489015182556001820191506020850194506020810190506126e0565b86831015612722578489015161271e601f891682612631565b8355505b6001600288020188555050505b505050505050565b5f60408201905061274a5f830185611e6f565b6127576020830184611f7a565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea26469706673582212203e39d564d126e9221b3f992ed795ed8826c814291523006b8140e0e883bfa9dd64736f6c634300081c0033",
}

// Ticket is an auto generated Go binding around an Ethereum contract.
type Ticket struct {
	abi abi.ABI
}

// NewTicket creates a new instance of Ticket.
func NewTicket() *Ticket {
	parsed, err := TicketMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Ticket{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Ticket) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address initialOwner) returns()
func (ticket *Ticket) PackConstructor(initialOwner common.Address) []byte {
	enc, err := ticket.abi.Pack("", initialOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (ticket *Ticket) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (ticket *Ticket) TryPackApprove(to common.Address, tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("approve", to, tokenId)
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (ticket *Ticket) PackBalanceOf(owner common.Address) []byte {
	enc, err := ticket.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (ticket *Ticket) TryPackBalanceOf(owner common.Address) ([]byte, error) {
	return ticket.abi.Pack("balanceOf", owner)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (ticket *Ticket) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := ticket.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burn(uint256 tokenId) returns()
func (ticket *Ticket) PackBurn(tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("burn", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burn(uint256 tokenId) returns()
func (ticket *Ticket) TryPackBurn(tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("burn", tokenId)
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (ticket *Ticket) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (ticket *Ticket) TryPackGetApproved(tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("getApproved", tokenId)
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (ticket *Ticket) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := ticket.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (ticket *Ticket) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := ticket.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (ticket *Ticket) TryPackIsApprovedForAll(owner common.Address, operator common.Address) ([]byte, error) {
	return ticket.abi.Pack("isApprovedForAll", owner, operator)
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (ticket *Ticket) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := ticket.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3fc9864.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mint(address to, uint256 tokenId, string uri) returns()
func (ticket *Ticket) PackMint(to common.Address, tokenId *big.Int, uri string) []byte {
	enc, err := ticket.abi.Pack("mint", to, tokenId, uri)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3fc9864.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mint(address to, uint256 tokenId, string uri) returns()
func (ticket *Ticket) TryPackMint(to common.Address, tokenId *big.Int, uri string) ([]byte, error) {
	return ticket.abi.Pack("mint", to, tokenId, uri)
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (ticket *Ticket) PackName() []byte {
	enc, err := ticket.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (ticket *Ticket) TryPackName() ([]byte, error) {
	return ticket.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (ticket *Ticket) UnpackName(data []byte) (string, error) {
	out, err := ticket.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackOperatorBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xebce9ccf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function operatorBurn(uint256 tokenId) returns()
func (ticket *Ticket) PackOperatorBurn(tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("operatorBurn", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOperatorBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xebce9ccf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function operatorBurn(uint256 tokenId) returns()
func (ticket *Ticket) TryPackOperatorBurn(tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("operatorBurn", tokenId)
}

// PackOperatorTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d1af103.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function operatorTransfer(address from, address to, uint256 tokenId) returns()
func (ticket *Ticket) PackOperatorTransfer(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("operatorTransfer", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOperatorTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d1af103.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function operatorTransfer(address from, address to, uint256 tokenId) returns()
func (ticket *Ticket) TryPackOperatorTransfer(from common.Address, to common.Address, tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("operatorTransfer", from, to, tokenId)
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (ticket *Ticket) PackOwner() []byte {
	enc, err := ticket.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (ticket *Ticket) TryPackOwner() ([]byte, error) {
	return ticket.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (ticket *Ticket) UnpackOwner(data []byte) (common.Address, error) {
	out, err := ticket.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (ticket *Ticket) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (ticket *Ticket) TryPackOwnerOf(tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("ownerOf", tokenId)
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (ticket *Ticket) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := ticket.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (ticket *Ticket) PackRenounceOwnership() []byte {
	enc, err := ticket.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceOwnership() returns()
func (ticket *Ticket) TryPackRenounceOwnership() ([]byte, error) {
	return ticket.abi.Pack("renounceOwnership")
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (ticket *Ticket) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (ticket *Ticket) TryPackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("safeTransferFrom", from, to, tokenId)
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (ticket *Ticket) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := ticket.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (ticket *Ticket) TryPackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) ([]byte, error) {
	return ticket.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (ticket *Ticket) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := ticket.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (ticket *Ticket) TryPackSetApprovalForAll(operator common.Address, approved bool) ([]byte, error) {
	return ticket.abi.Pack("setApprovalForAll", operator, approved)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (ticket *Ticket) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := ticket.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (ticket *Ticket) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return ticket.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (ticket *Ticket) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := ticket.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function symbol() view returns(string)
func (ticket *Ticket) PackSymbol() []byte {
	enc, err := ticket.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function symbol() view returns(string)
func (ticket *Ticket) TryPackSymbol() ([]byte, error) {
	return ticket.abi.Pack("symbol")
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (ticket *Ticket) UnpackSymbol(data []byte) (string, error) {
	out, err := ticket.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (ticket *Ticket) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (ticket *Ticket) TryPackTokenURI(tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("tokenURI", tokenId)
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (ticket *Ticket) UnpackTokenURI(data []byte) (string, error) {
	out, err := ticket.abi.Unpack("tokenURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (ticket *Ticket) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := ticket.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (ticket *Ticket) TryPackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) ([]byte, error) {
	return ticket.abi.Pack("transferFrom", from, to, tokenId)
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (ticket *Ticket) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := ticket.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (ticket *Ticket) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return ticket.abi.Pack("transferOwnership", newOwner)
}

// TicketApproval represents a Approval event raised by the Ticket contract.
type TicketApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const TicketApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (TicketApproval) ContractEventName() string {
	return TicketApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (ticket *Ticket) UnpackApprovalEvent(log *types.Log) (*TicketApproval, error) {
	event := "Approval"
	if len(log.Topics) == 0 || log.Topics[0] != ticket.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TicketApproval)
	if len(log.Data) > 0 {
		if err := ticket.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ticket.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TicketApprovalForAll represents a ApprovalForAll event raised by the Ticket contract.
type TicketApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const TicketApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (TicketApprovalForAll) ContractEventName() string {
	return TicketApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (ticket *Ticket) UnpackApprovalForAllEvent(log *types.Log) (*TicketApprovalForAll, error) {
	event := "ApprovalForAll"
	if len(log.Topics) == 0 || log.Topics[0] != ticket.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TicketApprovalForAll)
	if len(log.Data) > 0 {
		if err := ticket.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ticket.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TicketBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Ticket contract.
type TicketBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const TicketBatchMetadataUpdateEventName = "BatchMetadataUpdate"

// ContractEventName returns the user-defined event name.
func (TicketBatchMetadataUpdate) ContractEventName() string {
	return TicketBatchMetadataUpdateEventName
}

// UnpackBatchMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (ticket *Ticket) UnpackBatchMetadataUpdateEvent(log *types.Log) (*TicketBatchMetadataUpdate, error) {
	event := "BatchMetadataUpdate"
	if len(log.Topics) == 0 || log.Topics[0] != ticket.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TicketBatchMetadataUpdate)
	if len(log.Data) > 0 {
		if err := ticket.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ticket.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TicketMetadataUpdate represents a MetadataUpdate event raised by the Ticket contract.
type TicketMetadataUpdate struct {
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TicketMetadataUpdateEventName = "MetadataUpdate"

// ContractEventName returns the user-defined event name.
func (TicketMetadataUpdate) ContractEventName() string {
	return TicketMetadataUpdateEventName
}

// UnpackMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (ticket *Ticket) UnpackMetadataUpdateEvent(log *types.Log) (*TicketMetadataUpdate, error) {
	event := "MetadataUpdate"
	if len(log.Topics) == 0 || log.Topics[0] != ticket.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TicketMetadataUpdate)
	if len(log.Data) > 0 {
		if err := ticket.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ticket.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TicketOwnershipTransferred represents a OwnershipTransferred event raised by the Ticket contract.
type TicketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const TicketOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (TicketOwnershipTransferred) ContractEventName() string {
	return TicketOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (ticket *Ticket) UnpackOwnershipTransferredEvent(log *types.Log) (*TicketOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != ticket.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TicketOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := ticket.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ticket.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TicketTransfer represents a Transfer event raised by the Ticket contract.
type TicketTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const TicketTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (TicketTransfer) ContractEventName() string {
	return TicketTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (ticket *Ticket) UnpackTransferEvent(log *types.Log) (*TicketTransfer, error) {
	event := "Transfer"
	if len(log.Topics) == 0 || log.Topics[0] != ticket.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TicketTransfer)
	if len(log.Data) > 0 {
		if err := ticket.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ticket.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (ticket *Ticket) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], ticket.abi.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]) {
		return ticket.UnpackERC721IncorrectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]) {
		return ticket.UnpackERC721InsufficientApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]) {
		return ticket.UnpackERC721InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]) {
		return ticket.UnpackERC721InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]) {
		return ticket.UnpackERC721InvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]) {
		return ticket.UnpackERC721InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["ERC721InvalidSender"].ID.Bytes()[:4]) {
		return ticket.UnpackERC721InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]) {
		return ticket.UnpackERC721NonexistentTokenError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return ticket.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ticket.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return ticket.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TicketERC721IncorrectOwner represents a ERC721IncorrectOwner error raised by the Ticket contract.
type TicketERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func TicketERC721IncorrectOwnerErrorID() common.Hash {
	return common.HexToHash("0x64283d7b313c8117c125f736876fa2b4e90ea3831a4716dfdb87d2f540e26289")
}

// UnpackERC721IncorrectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func (ticket *Ticket) UnpackERC721IncorrectOwnerError(raw []byte) (*TicketERC721IncorrectOwner, error) {
	out := new(TicketERC721IncorrectOwner)
	if err := ticket.abi.UnpackIntoInterface(out, "ERC721IncorrectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketERC721InsufficientApproval represents a ERC721InsufficientApproval error raised by the Ticket contract.
type TicketERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func TicketERC721InsufficientApprovalErrorID() common.Hash {
	return common.HexToHash("0x177e802f6f313bc89797ecace66d6d29ab4719cbaaacbb87367264048b1eb861")
}

// UnpackERC721InsufficientApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func (ticket *Ticket) UnpackERC721InsufficientApprovalError(raw []byte) (*TicketERC721InsufficientApproval, error) {
	out := new(TicketERC721InsufficientApproval)
	if err := ticket.abi.UnpackIntoInterface(out, "ERC721InsufficientApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketERC721InvalidApprover represents a ERC721InvalidApprover error raised by the Ticket contract.
type TicketERC721InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidApprover(address approver)
func TicketERC721InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xa9fbf51f86b8e03595d59dc726bb10c329bb24f62589be276d8dd193ca0b69ea")
}

// UnpackERC721InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidApprover(address approver)
func (ticket *Ticket) UnpackERC721InvalidApproverError(raw []byte) (*TicketERC721InvalidApprover, error) {
	out := new(TicketERC721InvalidApprover)
	if err := ticket.abi.UnpackIntoInterface(out, "ERC721InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketERC721InvalidOperator represents a ERC721InvalidOperator error raised by the Ticket contract.
type TicketERC721InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOperator(address operator)
func TicketERC721InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0x5b08ba185e8f577075361f3a3555a6580a227ce22734dcc979c1aeadf894658b")
}

// UnpackERC721InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOperator(address operator)
func (ticket *Ticket) UnpackERC721InvalidOperatorError(raw []byte) (*TicketERC721InvalidOperator, error) {
	out := new(TicketERC721InvalidOperator)
	if err := ticket.abi.UnpackIntoInterface(out, "ERC721InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketERC721InvalidOwner represents a ERC721InvalidOwner error raised by the Ticket contract.
type TicketERC721InvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOwner(address owner)
func TicketERC721InvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x89c62b6479af2e623826dcc39c5133061d35b66d72de92833401dd2fd6567480")
}

// UnpackERC721InvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOwner(address owner)
func (ticket *Ticket) UnpackERC721InvalidOwnerError(raw []byte) (*TicketERC721InvalidOwner, error) {
	out := new(TicketERC721InvalidOwner)
	if err := ticket.abi.UnpackIntoInterface(out, "ERC721InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketERC721InvalidReceiver represents a ERC721InvalidReceiver error raised by the Ticket contract.
type TicketERC721InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func TicketERC721InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x64a0ae9278f805eaf991dcd18ca78756d280b7508b764ef1b255c55845c11df9")
}

// UnpackERC721InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func (ticket *Ticket) UnpackERC721InvalidReceiverError(raw []byte) (*TicketERC721InvalidReceiver, error) {
	out := new(TicketERC721InvalidReceiver)
	if err := ticket.abi.UnpackIntoInterface(out, "ERC721InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketERC721InvalidSender represents a ERC721InvalidSender error raised by the Ticket contract.
type TicketERC721InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidSender(address sender)
func TicketERC721InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x73c6ac6e10798e95d99e1f130d923eb40193ecb8d094ec3dce93292564eb3b17")
}

// UnpackERC721InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidSender(address sender)
func (ticket *Ticket) UnpackERC721InvalidSenderError(raw []byte) (*TicketERC721InvalidSender, error) {
	out := new(TicketERC721InvalidSender)
	if err := ticket.abi.UnpackIntoInterface(out, "ERC721InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketERC721NonexistentToken represents a ERC721NonexistentToken error raised by the Ticket contract.
type TicketERC721NonexistentToken struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func TicketERC721NonexistentTokenErrorID() common.Hash {
	return common.HexToHash("0x7e273289a3a9ef6670f06df7dca227856fc925e956db96980692764a8bc734d7")
}

// UnpackERC721NonexistentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func (ticket *Ticket) UnpackERC721NonexistentTokenError(raw []byte) (*TicketERC721NonexistentToken, error) {
	out := new(TicketERC721NonexistentToken)
	if err := ticket.abi.UnpackIntoInterface(out, "ERC721NonexistentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the Ticket contract.
type TicketOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func TicketOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (ticket *Ticket) UnpackOwnableInvalidOwnerError(raw []byte) (*TicketOwnableInvalidOwner, error) {
	out := new(TicketOwnableInvalidOwner)
	if err := ticket.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TicketOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the Ticket contract.
type TicketOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func TicketOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (ticket *Ticket) UnpackOwnableUnauthorizedAccountError(raw []byte) (*TicketOwnableUnauthorizedAccount, error) {
	out := new(TicketOwnableUnauthorizedAccount)
	if err := ticket.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}
