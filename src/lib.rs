#![no_std]

// Import to interact with contract storage
use ink_core::storage;
// Import the `contract!` macro
use ink_core::contract;

// The code for your contract will live entirely in the `contract!` macro
contract! {
    struct ContractName {
        // Contract Storage 
    }

    impl Deploy for ContractName {
        fn deploy(&mut self) {
            // Deployment logic that runs once upon contract creation 
        }
    }

    impl ContractName {
        // Public/Private Function Definitions 
    }
}

