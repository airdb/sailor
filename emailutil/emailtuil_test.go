package emailutil

import "testing"

func TestSendEmail(t *testing.T) {
	type args struct {
		toEmails string
		subject  string
		content  string
	}

	tests := []struct {
		name string
		args args
	}{
		{``, args{"airdb@qq.com", "hello", "test mail"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendEmail(tt.args.toEmails, tt.args.subject, tt.args.content)
		})
	}
}
