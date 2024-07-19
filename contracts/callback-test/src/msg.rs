use cosmwasm_schema::{cw_serde, QueryResponses};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

#[cw_serde]
pub struct InstantiateMsg {
    pub count: i32,
}

#[cw_serde]
pub enum ExecuteMsg {
    Increment {},
    Reset { count: i32 },
}

#[cw_serde]
pub enum SudoMsg {
    Callback {
        job_id: u64,
    },
    Error {
        module_name: String,
        error_code: u32,
        contract_address: String,
        input_payload: String,
        error_message: String,
    },
}

#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {
    // GetCount returns the current count as a json-encoded number
    #[returns(GetCountResponse)]
    GetCount {},
    #[returns(QueryErrorsResponse)]
    QueryCwErrors {},
    #[returns(Option<crate::state::CallbackError>)]
    QueryCallbackError {},
}

#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryErrorsRequest {
    #[prost(string, tag = "1")]
    pub contract_address: ::prost::alloc::string::String,
}

#[derive(Clone, PartialEq, Serialize, Deserialize, JsonSchema, ::prost::Message)]
pub struct SudoError {
    #[prost(string, tag = "1")]
    pub module_name: ::prost::alloc::string::String,
    #[prost(uint32, tag = "2")]
    pub error_code: u32,
    #[prost(string, tag = "3")]
    pub contract_address: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub input_payload: ::prost::alloc::string::String,
    #[prost(string, tag = "5")]
    pub error_message: ::prost::alloc::string::String,
}

#[derive(Clone, PartialEq, Serialize, Deserialize, JsonSchema, ::prost::Message)]
pub struct QueryErrorsResponse {
    #[prost(message, repeated, tag = "1")]
    pub errors: ::prost::alloc::vec::Vec<SudoError>,
}

#[cw_serde]
pub struct GetCountResponse {
    pub count: i32,
}

// pub struct CwErrorsRes
