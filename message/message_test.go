package message

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMessage = Message{

	Name:      "TestName",
	Timestamp: "2017-12-07",
	Text:      "Test Message",
}

func ExampleMessage_String() {

	fmt.Print(testMessage.String())
	// OUTPUT:
	// 2017-12-07 - message from TestName: Test Message

}

func TestMessage_SetMessage(t *testing.T) {
	type fields struct {
		Name      string
		Timestamp string
		Text      string
	}
	type args struct {
		from string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Chandler",
			fields: fields{
				Name:      "Chandler",
				Timestamp: "2017-12-07",
				Text:      "Hi Chandler\n",
			},
			args: args{
				from: "Nicole",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Name:      tt.fields.Name,
				Timestamp: tt.fields.Timestamp,
				Text:      tt.fields.Text,
			}
			if err := m.SetMessage(tt.args.from, strings.NewReader(m.Text)); (err != nil) != tt.wantErr {
				t.Errorf("Message.SetMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inputMessageText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		wantText string
		wantErr  bool
	}{
		{
			name:     "Jack",
			text:     "Hi Jack\n",
			wantText: "Hi Jack",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotText, err := inputMessageText(strings.NewReader(tt.text))
			if (err != nil) != tt.wantErr {
				t.Errorf("inputMessageText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotText != tt.wantText {
				t.Errorf("inputMessageText() = %v, want %v", gotText, tt.wantText)
			}
		})
	}
}

func TestMessage_String(t *testing.T) {
	type fields struct {
		Name      string
		Timestamp string
		Text      string
	}
	tests := []struct {
		name     string
		fields   fields
		expected string
	}{
		{
			name: "Jack",
			fields: fields{
				Name:      "London",
				Timestamp: "2017-12-07",
				Text:      "How are you?",
			},
			expected: fmt.Sprintf("2017-12-07 - message from London: How are you?\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Name:      tt.fields.Name,
				Timestamp: tt.fields.Timestamp,
				Text:      tt.fields.Text,
			}
			got := m.String()
			assert.Equal(t, tt.expected, got)
		})
	}
}
