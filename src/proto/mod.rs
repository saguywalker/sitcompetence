#![allow(bare_trait_objects)]
#![allow(renamed_and_removed_lints)]

pub use self::competence::{TxCreateAccount, Account};

include!(concat!(env!("OUT_DIR"), "/protobuf_mod.rs"));

use exonum::proto::schema::*;