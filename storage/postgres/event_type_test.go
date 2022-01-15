package postgres_test

import (
	"Go-Project-With-Postgres/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateArticle(t *testing.T) {
	t.Parallel()

	ts := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.EventType
		want    *storage.EventType
		wantErr bool
	}{
		{
			name: "Success Create EventType",
			in: storage.EventType{
				EventTypeName: "classic",
			},
			want: &storage.EventType{
				EventTypeName: "classic",
			},
		},
		{
			name: "Failed To Create EventType",
			in: storage.EventType{
				EventTypeName: "res",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ts.CreateEventType(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateEventType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// if !tt.wantErr && got.ID == "" {
			// 	t.Error("ID cannot be empty")
			// 	return
			// }

			opt := cmpopts.IgnoreFields(storage.EventType{}, "ID", "Created", "Updated", "Deleted")
			if !cmp.Equal(tt.want, got, opt) {
				t.Errorf("Storage.CreateArticle() = - want, + got\n %+v", cmp.Diff(tt.want, got, opt))
			}
		})
	}
}
