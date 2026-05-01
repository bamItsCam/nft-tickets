pragma solidity ^0.8.28;

import {Script} from "forge-std/Script.sol";
import {Ticket} from "../src/Ticket.sol";

contract TicketScript is Script {
    function run() external {
        vm.startBroadcast();
        new Ticket(msg.sender);
        vm.stopBroadcast();
    }
}
