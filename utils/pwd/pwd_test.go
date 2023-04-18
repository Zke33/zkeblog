package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Printf(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$aafgBzpmL3MRCc3qzuxk4eS2rzEBdv2J9NYlBCFvsyJS0/wjYl/tq", "1234"))
}
