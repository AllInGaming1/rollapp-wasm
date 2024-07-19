use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

use cosmwasm_std::Addr;
use cw_storage_plus::Item;

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq, JsonSchema)]
pub struct State {
    pub count: i32,
    pub owner: Addr,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq, JsonSchema)]
pub struct CallbackError {
    pub message: String,
    pub code: u32,
    pub module_name: String,
}

pub const STATE: Item<State> = Item::new("state");
pub const CALLBACK_ERROR: Item<Option<CallbackError>> = Item::new("callback_error");
