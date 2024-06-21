package services

import (
	"context"
	"time"
	"math/rand"
	"errors"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string 	`json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	City	  string    `json:"city"`
	WeekDays  string    `json:"week_days"`
	UserName  string    `json:"user_name"`
}
type UserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UsersList struct {
	Users []User `json:"users"`
}

func (u *User) GetAllUsers() ([]*User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()
    query := `SELECT id, name, email, created_at, updated_at, city, week_days, user_name FROM users`
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*User
    for rows.Next() {
        var user User
        err := rows.Scan(
            &user.ID,
            &user.Name,
            &user.Email,
            &user.CreatedAt,
            &user.UpdatedAt,
            &user.City,
            &user.WeekDays,
            &user.UserName, // Adicione o campo user_name aqui
        )
        if err != nil {
            return nil, err
        }
        users = append(users, &user)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return users, nil
}


func (u *User) GetUserByID(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	if u.City == "" {
		u.City = mockCity()
	}

	if u.WeekDays == "" {
		u.WeekDays = mockWeekDays()
	}

	return u, nil
}

func mockCity() string {
	// Retorna cidades mockadas
	cities := []string{"Rio de Janeiro", "São Paulo", "João Pessoa", "Salvador", "Brasília"}
	return cities[rand.Intn(len(cities))]
}

func mockWeekDays() string {
	// Retorna dias da semana mockados
	weekdays := []string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"}
	return weekdays[rand.Intn(len(weekdays))]
}

func (u *User) CreateUser(user User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Verificar se o e-mail já está em uso
	var existingUser User
	err := db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", user.Email).Scan(&existingUser.ID)
	if err == nil {
		return nil, errors.New("E-mail já está em uso")
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Inserir usuário com a nova coluna user_name
	query := `INSERT INTO users (name, email, password, created_at, updated_at, city, week_days, user_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, name, email, created_at, updated_at, city, week_days, user_name`
	err = db.QueryRowContext(ctx, query, user.Name, user.Email, hashedPassword, time.Now(), time.Now(), user.City, user.WeekDays, user.UserName).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.City, &user.WeekDays, &user.UserName)
	if err != nil {
		return nil, err
	}
	user.Password = "" // Limpar a senha antes de retornar
	return &user, nil
}



func (u *User) UpdateUser(id string, body User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`
	_, err := db.ExecContext(ctx, query, body.Name, body.Email, time.Now(), id)
	if err != nil {
		return  nil, err
	}
	return &body,nil
}

func (u *User) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserByUsername(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, email, created_at, updated_at, city, week_days, user_name 
              FROM users 
              WHERE user_name = $1`

	log.Printf("Executing query: %s with username: %s", query, username)

	row := db.QueryRowContext(ctx, query, username)
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.City,
		&u.WeekDays,
		&u.UserName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Log para usuário não encontrado
			log.Printf("User with username %s not found", username)
			return nil, nil // Retorne nil se o usuário não for encontrado
		}
		// Log para qualquer outro erro
		log.Printf("Error scanning row for username %s: %v", username, err)
		return nil, err
	}
	// Log para sucesso
	log.Printf("User found: %+v", u)
	return u, nil
}

