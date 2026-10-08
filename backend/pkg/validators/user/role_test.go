// pkg/validators/user/role_test.go
package user

import "testing"

func TestValidateRole(t *testing.T) {
	tests := []struct {
		name string
		role string
		want bool
	}{
		{
			name: "role owner válida",
			role: "owner",
			want: true,
		},
		{
			name: "role tenant válida",
			role: "tenant",
			want: true,
		},
		{
			name: "role inválida",
			role: "admin",
			want: false,
		},
		{
			name: "role vazia",
			role: "",
			want: false,
		},
		{
			name: "role com maiúscula",
			role: "Owner",
			want: false,
		},
		{
			name: "role com espaços",
			role: " owner ",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateRole(tt.role)

			if got != tt.want {
				t.Errorf(
					"ValidateRole(%q) = %v, want %v",
					tt.role,
					got,
					tt.want,
				)
			}
		})
	}
}
