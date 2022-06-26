package xcash

import (
	"testing"
	"time"
)

func TestInitialize(t *testing.T) {
    Initialize("us1.xcash.foundation:18281","us1.xcash.foundation:18289")
    time.Sleep(5 * time.Second)
}

/*func TestBlockchain_get_block_count(t *testing.T) {
    got,error := Blockchain_get_block_count()
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_block_count has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_block_count has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_block_hash(t *testing.T) {
    got,error := Blockchain_get_block_hash(800000)
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_block_hash has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_block_hash has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_block_template(t *testing.T) {
    got,error := Blockchain_get_block_template("XCA1XPzaSeXgwrBrGbh96UD5bk21a4WabcrgtB14A7WGGdcagjVQVV1PMAg5Rj1SM3ca8ZPDvysi78HyZF9imGg48wRK2Ntqov",128)
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_block_template has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_block_template has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_last_block_header(t *testing.T) {
    got,error := Blockchain_get_last_block_header()
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_last_block_header has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_last_block_header has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_block_header_by_hash(t *testing.T) {
    got,error := Blockchain_get_block_header_by_hash("008b8147114089e9577be0644c67a28d5227f4701e345280567c589256fcd4cc")
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_block_header_by_hash has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_block_header_by_hash has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_block_header_by_height(t *testing.T) {
    got,error := Blockchain_get_block_header_by_height(800000)
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_block_header_by_height has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_block_header_by_height has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_block_headers_range(t *testing.T) {
    got,error := Blockchain_get_block_headers_range(800000,800002)
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_block_headers_range has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_block_headers_range has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_block(t *testing.T) {
    got,error := Blockchain_get_block(800000)
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_block has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_block has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_info(t *testing.T) {
    got,error := Blockchain_get_info()
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_info has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_info has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_hard_fork_info(t *testing.T) {
    got,error := Blockchain_hard_fork_info()
    if got == nil && error == nil {                
		t.Errorf("Blockchain_hard_fork_info has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_hard_fork_info has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_version(t *testing.T) {
    got,error := Blockchain_version()
    if got == nil && error == nil {                
		t.Errorf("Blockchain_version has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_version has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_tx(t *testing.T) {
    got,error := Blockchain_get_tx("35f9dccaf21dfe1df0945ebfc8b3ef28977b4ea1a78b3726c9f866facd27f7ad")
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_tx has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_tx has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_balance(t *testing.T) {
    got,error := Blockchain_get_balance()
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_balance has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_balance has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_address(t *testing.T) {
    got,error := Blockchain_get_address()
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_address has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_address has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}

func TestBlockchain_get_height(t *testing.T) {
    got,error := Blockchain_get_height()
    if got == nil && error == nil {                
		t.Errorf("Blockchain_get_height has had an error")
	}
    if got == nil && error != nil {  
        t.Errorf("Blockchain_get_height has had an error. %s",error.Error.Message)
	}
	time.Sleep(5 * time.Second)
}*/

func TestAPI_Blockchain_stats(t *testing.T) {
    got := API_Blockchain_stats()
    if got == nil {                
		t.Errorf("API_Blockchain_stats has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Blockchain_block_data(t *testing.T) {
    got := API_Blockchain_block_data(800000)
    if got == nil {                
		t.Errorf("API_Blockchain_block_data has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Blockchain_tx_data(t *testing.T) {
    got := API_Blockchain_tx_data("d15005880f5f88b19fc88bdec29faaf57489ba85dd02d41ec87043a5eddf95a9")
    if got == nil {                
		t.Errorf("API_Blockchain_tx_data has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Blockchain_prove_tx(t *testing.T) {
    got := API_Blockchain_prove_tx("d15005880f5f88b19fc88bdec29faaf57489ba85dd02d41ec87043a5eddf95a9","XCA1a9usG2UKajV1Dqzp8fL1BbN3hzuaaJMYjCo7qDoC4C3Vvc5owiLAqKbVw2cRbwRqx3mgrau1Z7LkX6cxR2NC4ZmFBLe2Mf","a15005880f5f88b19fc88bdec29faaf57489ba85dd02d41ec87043a5eddf95a9")
    if got == nil {                
		t.Errorf("API_Blockchain_prove_tx has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Blockchain_prove_balance(t *testing.T) {
    got := API_Blockchain_prove_balance("XCA1a9usG2UKajV1Dqzp8fL1BbN3hzuaaJMYjCo7qDoC4C3Vvc5owiLAqKbVw2cRbwRqx3mgrau1Z7LkX6cxR2NC4ZmFBLe2Mf","ReserveProofV1a15005880f5f88b19fc88bdec29faaf57489ba85dd02d41ec87043a5eddf95a9")
    if got == nil {                
		t.Errorf("API_Blockchain_prove_balance has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Blockchain_tx_history(t *testing.T) {
    got := API_Blockchain_tx_history("sender","XCA1a9usG2UKajV1Dqzp8fL1BbN3hzuaaJMYjCo7qDoC4C3Vvc5owiLAqKbVw2cRbwRqx3mgrau1Z7LkX6cxR2NC4ZmFBLe2Mf")
    if got == nil {                
		t.Errorf("API_Blockchain_tx_history has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Blockchain_validate_address(t *testing.T) {
    got := API_Blockchain_validate_address("XCA1a9usG2UKajV1Dqzp8fL1BbN3hzuaaJMYjCo7qDoC4C3Vvc5owiLAqKbVw2cRbwRqx3mgrau1Z7LkX6cxR2NC4ZmFBLe2Mf")
    if got == nil {                
		t.Errorf("API_Blockchain_validate_address has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Blockchain_create_integrated_address(t *testing.T) {
    got := API_Blockchain_create_integrated_address("XCA1a9usG2UKajV1Dqzp8fL1BbN3hzuaaJMYjCo7qDoC4C3Vvc5owiLAqKbVw2cRbwRqx3mgrau1Z7LkX6cxR2NC4ZmFBLe2Mf","")
    if got == nil {                
		t.Errorf("API_Blockchain_create_integrated_address has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_stats(t *testing.T) {
    got := API_DPOPS_stats()
    if got == nil {                
		t.Errorf("API_DPOPS_stats has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_registered_delegates(t *testing.T) {
    got := API_DPOPS_registered_delegates()
    if got == nil {                
		t.Errorf("API_DPOPS_registered_delegates has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_online_delegates(t *testing.T) {
    got := API_DPOPS_online_delegates()
    if got == nil {                
		t.Errorf("API_DPOPS_online_delegates has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_active_delegates(t *testing.T) {
    got := API_DPOPS_active_delegates()
    if got == nil {                
		t.Errorf("API_DPOPS_active_delegates has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_delegate(t *testing.T) {
    got := API_DPOPS_delegate("us1_xcash_foundation")
    if got == nil {                
		t.Errorf("API_DPOPS_delegate has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_delegate_round_stats(t *testing.T) {
    got := API_DPOPS_delegate_round_stats("us1_xcash_foundation")
    if got == nil {                
		t.Errorf("API_DPOPS_delegates_votes has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_delegates_votes(t *testing.T) {
    got := API_DPOPS_delegates_votes("us1_xcash_foundation",1,2)
    if got == nil {                
		t.Errorf("API_Blockchain_create_integrated_address has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_vote_detail(t *testing.T) {
    got := API_DPOPS_vote_detail("XCA1a9usG2UKajV1Dqzp8fL1BbN3hzuaaJMYjCo7qDoC4C3Vvc5owiLAqKbVw2cRbwRqx3mgrau1Z7LkX6cxR2NC4ZmFBLe2Mf")
    if got == nil {                
		t.Errorf("API_DPOPS_vote_detail has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_round_detail(t *testing.T) {
    got := API_DPOPS_round_detail(810000)
    if got == nil {                
		t.Errorf("API_DPOPS_round_detail has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_DPOPS_last_block_producer(t *testing.T) {
    got := API_DPOPS_last_block_producer()
    if got == nil {                
		t.Errorf("API_DPOPS_last_block_producer has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Namespace_stats(t *testing.T) {
    got := API_Namespace_stats()
    if got == nil {                
		t.Errorf("API_Namespace_stats has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Namespace_registered_delegates(t *testing.T) {
    got := API_Namespace_registered_delegates()
    if got == nil {                
		t.Errorf("API_Namespace_registered_delegates has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Namespace_delegates(t *testing.T) {
    got := API_Namespace_delegates("us1_xcash_foundation")
    if got == nil {                
		t.Errorf("API_Namespace_delegates has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Namespace_names(t *testing.T) {
    got := API_Namespace_names("xcash")
    if got == nil {                
		t.Errorf("API_Namespace_names has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Namespace_name_status(t *testing.T) {
    got := API_Namespace_name_status("xcash")
    if got == nil {                
		t.Errorf("API_Namespace_name_status has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Namespace_address_status(t *testing.T) {
    got := API_Namespace_address_status("XCA1a9usG2UKajV1Dqzp8fL1BbN3hzuaaJMYjCo7qDoC4C3Vvc5owiLAqKbVw2cRbwRqx3mgrau1Z7LkX6cxR2NC4ZmFBLe2Mf")
    if got == nil {                
		t.Errorf("API_Namespace_address_status has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Namespace_name_to_address(t *testing.T) {
    got := API_Namespace_name_to_address("xcash")
    if got == nil {                
		t.Errorf("API_Namespace_name_to_address has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Namespace_address_to_name(t *testing.T) {
    got := API_Namespace_address_to_name("XCA1a9usG2UKajV1Dqzp8fL1BbN3hzuaaJMYjCo7qDoC4C3Vvc5owiLAqKbVw2cRbwRqx3mgrau1Z7LkX6cxR2NC4ZmFBLe2Mf")
    if got == nil {                
		t.Errorf("API_Namespace_address_to_name has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Xpayment_Twitter_stats(t *testing.T) {
    got := API_Xpayment_Twitter_stats()
    if got == nil {                
		t.Errorf("API_Xpayment_Twitter_stats has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Xpayment_Twitter_stats_per_day(t *testing.T) {
    got := API_Xpayment_Twitter_stats_per_day(1,2)
    if got == nil {                
		t.Errorf("API_Xpayment_Twitter_stats_per_day has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Xpayment_Twitter_top_stats(t *testing.T) {
    got := API_Xpayment_Twitter_top_stats(10)
    if got == nil {                
		t.Errorf("API_Xpayment_Twitter_top_stats has had an error")
	}
	time.Sleep(5 * time.Second)
}

func TestAPI_Xpayment_Twitter_recent_tips(t *testing.T) {
    got := API_Xpayment_Twitter_recent_tips(10,"First","All")
    if got == nil {                
		t.Errorf("API_Xpayment_Twitter_recent_tips has had an error")
	}
	time.Sleep(5 * time.Second)
}
