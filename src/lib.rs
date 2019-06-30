#![cfg_attr(not(any(test, feature = "std")), no_std)]

use ink_core::{
    storage,
};
use ink_lang::contract;

contract! {
    struct Mpt {
        owner: storage::Value<AccountId>,
        node: storage::Vec<i32>,
        nodes: storage::HashMap<Hash, self.node>,
    }

    impl Deploy for Mpt {
        fn deploy(&mut self) {
            self.owner.set(env.caller());
        }
    }

    /// Public functions
    impl Mpt {
    }
}

#[cfg(all(test, feature = "test-env"))]
mod tests {
    use super::*;

    #[test]
    fn it_deploys() {
        let mut trie = Mpt::deploy_mock();
    }
}
