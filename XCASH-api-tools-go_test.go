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
