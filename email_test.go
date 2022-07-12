package email

import "testing"

func TestSendEmail(t *testing.T) {
    SendEmail("jason@localhost", []string{"root@localhost"}, "test subject", "hello world body", "text", []string{})
    /*if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
    */
}
