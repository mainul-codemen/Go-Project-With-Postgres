package storage

// import "time"

type (
	User struct {
		ID        int32  `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Username  string `db:"username"`
		Email     string `db:"email"`
		/* IsActive  bool      `db:"is_active"`
		IsAdmin   bool      `db:"is_admin"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"` */
	}
)
