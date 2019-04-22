#[macro_use]
extern crate exonum_derive;
#[macro_use]
extern crate failure;
#[macro_use]
extern crate serde_derive;

pub mod proto;

pub mod schema{
    use super::proto;
    //use exonum_merkledb::{IndexAccess, MapIndex};
    use exonum::crypto::PublicKey;
    use exonum::storage::{Fork, MapIndex, Snapshot};

    #[derive(Serialize, Deserialize, Clone, Debug, ProtobufConvert)]
    #[exonum(pb = "proto::Account")]
    pub struct Account{
        pub student_id: PublicKey,
        pub name: String,
    }

    impl Account{
        pub fn new(&student_id: &PublicKey, name: &str) -> Self{
            Self{
                student_id,
                name: name.to_owned(),
            }
        }
    }

    #[derive(Debug)]
    pub struct AccountSchema<T>{
        view: T,
    }

    impl<T: AsRef<dyn Snapshot>> AccountSchema<T>{
        pub fn new(view: T) -> Self{
            AccountSchema{view}
        }

        pub fn accounts(&self) -> MapIndex<&dyn Snapshot, PublicKey, Account>{
            MapIndex::new("sitcompetence.accounts", self.view.as_ref())
        }

        pub fn account(&self, student_id: &PublicKey) -> Option<Account>{
            self.accounts().get(student_id)
        }
    }

    impl<'a> AccountSchema<&'a mut Fork>{
        pub fn accounts_mut(&mut self) -> MapIndex<&mut Fork, PublicKey, Account>{
            MapIndex::new("sitcompetence.accounts", &mut self.view)
        }
    }

}

pub mod transactions{
    use super::proto;
    use super::service::SERVICE_ID;
    use exonum::{
        crypto::{PublicKey, SecretKey},
        messages::{Message, RawTransaction, Signed},
    };
    use super::schema::{AccountSchema, Account};

    #[derive(Serialize, Deserialize, Clone, Debug, ProtobufConvert)]
    #[exonum(pb = "proto::TxCreateAccount")]
    pub struct TxCreateAccount{
        pub name: String,
    }

    #[derive(Serialize, Deserialize, Clone, Debug, TransactionSet)]
    pub enum SITCompetenceTransactions{
        CreateAccount(TxCreateAccount),
    }

    impl TxCreateAccount{
        pub fn sign(name: &str, pk: &PublicKey, sk: &SecretKey) -> Signed<RawTransaction>{
            let name = name.to_owned();
            Message::sign_transaction(
                Self{
                    name,
                },
                SERVICE_ID,
                *pk,
                sk,
            )
        }
    }
}

pub mod errors{
    #![allow(bare_trait_objects)]
    use exonum::blockchain::ExecutionError;
    
    #[derive(Debug, Fail)]
    #[repr(u8)]
    pub enum Error{
        #[fail(display = "Account already exists")]
        AccountAlreadyExists = 0,
    }

    impl From<Error> for ExecutionError{
        fn from(value: Error) -> ExecutionError{
            let description = format!("{}", value);
            ExecutionError::with_description(value as u8, description)
        }
    }
}

pub mod contracts{
    use exonum::blockchain::{ExecutionResult, Transaction, TransactionContext};

    use crate::{
        errors::Error,
        schema::{AccountSchema, Account},
        transactions::TxCreateAccount,
    };
    
    impl Transaction for TxCreateAccount{
        fn execute(&self, mut context: TransactionContext) -> ExecutionResult{
            let author = context.author();
            let view = context.fork();
            let mut schema = AccountSchema::new(view);
            if schema.account(&author).is_none(){
                let account = Account::new(&author, &self.name);
                println!("Create account: {:?}", account);
                schema.accounts_mut().put(&author, account);
                Ok(())
            }else{
                Err(Error::AccountAlreadyExists)?
            }
        }
    }
}

pub mod api{
    use exonum::{
        api::{self, ServiceApiBuilder, ServiceApiState},
        crypto::PublicKey,
    };

    use crate::schema::{AccountSchema, Account};

    #[derive(Debug, Clone)]
    pub struct CompetenceApi;

    #[derive(Debug, Serialize, Deserialize, Clone, Copy)]
    pub struct AccountQuery{
        pub student_id: PublicKey,
    }

    impl CompetenceApi{
        pub fn get_account(state: &ServiceApiState, query: AccountQuery) -> api::Result<Account>{
            let snapshot = state.snapshot();
            let schema = AccountSchema::new(snapshot);
            schema
                .account(&query.student_id)
                .ok_or_else(|| api::Error::NotFound("\"Account not found\"".to_owned()))
        }

        pub fn get_accounts(state: &ServiceApiState, _query: ()) -> api::Result<Vec<Account>>{
            let snapshot = state.snapshot();
            let schema = AccountSchema::new(snapshot);
            let idx = schema.accounts();
            let accounts = idx.values().collect();
            Ok(accounts)
        }

        pub fn wire(builder: &mut ServiceApiBuilder){
            builder
                .public_scope()
                .endpoint("v1/account", Self::get_account)
                .endpoint("v1/accounts", Self::get_accounts);
            println!("*******************Wire*********************");
        }
    }
}

pub mod service{
    //use exonum_merkledb::snapshot;
    use exonum::{
        api::ServiceApiBuilder,
        blockchain::{Service, Transaction, TransactionSet},
        crypto::Hash,
        messages::RawTransaction,
        storage::Snapshot,
    };

    use crate::{api::CompetenceApi, transactions::SITCompetenceTransactions};

    pub const SERVICE_ID: u16 = 1;

    #[derive(Debug)]
    pub struct CompetenceService;


    impl Service for CompetenceService{
        
        fn service_name(&self) -> &'static str{
            "SITCompetence"
        }

        fn service_id(&self) -> u16{
            SERVICE_ID
        }

        fn tx_from_raw(&self, raw: RawTransaction) -> Result<Box<dyn Transaction>, failure::Error>{
            let tx = SITCompetenceTransactions::tx_from_raw(raw)?;
            Ok(tx.into())
        }

        fn state_hash(&self, _: &dyn Snapshot) -> Vec<Hash>{
            vec![]
        }

        fn wire_api(&self, builder: &mut ServiceApiBuilder){
            CompetenceApi::wire(builder)
        }
    }
}