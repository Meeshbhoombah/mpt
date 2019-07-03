#![cfg_attr(not(any(test, feature = "std")), no_std)]

use parity_codec::{
    Decode,
    Encode
};
use ink_core::{
    env::{
        Env,
        EnvTypes,
        Hash,
        AccountId
    },
    memory::string::String,
    storage::{
        self,
        alloc::{
            AllocateUsing,
            BumpAlloc,
            Initialize,
        },
        Flush,
        Key
    },
};
use ink_lang::contract;

#[derive(Debug, Encode, Decode)]
pub struct Node {
    /// collection of children elements
    ///
    /// branch: contains one element
    /// extension: slots for all values - a-z,0-9
    /// leaf: None
    children: storage::Vec<Hash>,
    /// varying key type, assembles key for value 
    ///
    /// branch: prefixstring
    /// extension: shared string
    /// leaf: contains the value
    key: storage::Value<String>,
    /// varying value type
    ///
    /// branch: None 
    /// extension: None 
    /// leaf: contains the value
    value: storage::Value<String>,
}

impl Node {
    pub fn hash<H: EnvTypes>(&self, state: &mut H) {
        
    }
}

contract! {
    struct Mpt {
        owner: storage::Value<AccountId>,
        root: storage::Value<Hash>,
        nodes: storage::HashMap<Hash, Node>,
    }

    impl Deploy for Mpt {
        fn deploy(&mut self) {
            self.owner.set(env.caller());
        }
    }

    impl Mpt {
    }
}

#[cfg(all(test, feature = "test-env"))]
mod tests {
    use super::Mpt;
    use ink_core::env;

    #[test]
    fn it_deploys_owned() {
        // no events prior to deploy
        assert_eq!(env::test::emitted_events().count(), 0);
        let mut trie = Mpt::deploy_mock();
        assert_eq!(trie.owner, env::TestEnvData.caller);
    }
}
