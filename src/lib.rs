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



#[cfg(all(test, feature = "test-env"))]
mod tests {
    use super::Mpt;
    use ink_core::env;

    #[test]
    fn it_deploys() {
        // no events prior to deploy
        assert_eq!(env::test::emitted_events().count(), 0);
        let mut trie = Mpt::deploy_mock();
    }
}
