package database

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	path string
}

func NewClient(path string) Client {
	return Client{path: path}
}

func (c Client) EnsureDB() error {
	if c.path != "" {
		data, err := os.ReadFile(c.path)
		if err == nil && len(data) < 1 {
			c.createDB()
			return nil
		} else {
			return err
		}
	} else {
		return errors.New("invalid client path")
	}
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	if c.path != "" {
		if email != "" && name != "" && password != "" {
			data, err := c.readDB()
			if err == nil {
				dbUsers := data.Users
				if _, ok := dbUsers[email]; !ok {
					dbUsers[email] = User{
						Email:     email,
						Password:  password,
						Name:      name,
						Age:       age,
						CreatedAt: time.Now().UTC(),
					}
					c.updateDB(data)
					return dbUsers[email], nil
				} else {
					return dbUsers[email], errors.New("user already exists")
				}
			} else {
				return User{}, err
			}
		} else {
			return User{}, errors.New("failed validation")
		}
	} else {
		return User{}, errors.New("invalid client path")
	}
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	if c.path != "" {
		if email != "" && name != "" && password != "" {
			data, err := c.readDB()
			if err == nil {
				dbUsers := data.Users
				if _, ok := dbUsers[email]; ok {
					dbUsers[email] = User{
						Email:     email,
						Password:  password,
						Name:      name,
						Age:       age,
						CreatedAt: time.Now().UTC(),
					}
					c.updateDB(data)
					return dbUsers[email], nil
				} else {
					return dbUsers[email], errors.New("user doesn't exist")
				}
			} else {
				return User{}, err
			}
		} else {
			return User{}, errors.New("failed validation")
		}
	} else {
		return User{}, errors.New("invalid client path")
	}
}

func (c Client) GetUser(email string) (User, error) {
	if c.path != "" {
		if email != "" {
			data, err := c.readDB()
			if err == nil {
				dbUsers := data.Users
				if user, ok := dbUsers[email]; ok {
					return user, nil
				} else {
					return user, errors.New("user doesn't exist")
				}
			} else {
				return User{}, err
			}
		} else {
			return User{}, errors.New("email is empty")
		}
	} else {
		return User{}, errors.New("invalid client path")
	}
}

func (c Client) DeleteUser(email string) error {
	if c.path != "" {
		if email != "" {
			data, err := c.readDB()
			if err == nil {
				dbUsers := data.Users
				delete(dbUsers, email)
				c.updateDB(data)
			} else {
				return err
			}
		} else {
			return errors.New("email is empty")
		}
		return nil
	} else {
		return errors.New("invalid client path")
	}
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	if c.path != "" {
		if userEmail != "" {
			user, err := c.GetUser(userEmail)
			if err == nil && user.Email == userEmail {
				dbData, err := c.readDB()
				if err == nil {
					posts := dbData.Posts
					id := uuid.New().String()
					posts[id] = Post{
						ID:        id,
						CreatedAt: time.Now().UTC(),
						UserEmail: userEmail,
						Text:      text,
					}
					c.updateDB(dbData)
					return posts[id], nil
				} else {
					return Post{}, err
				}
			} else {
				return Post{}, err
			}
		} else {
			return Post{}, errors.New("email is empty")
		}
	} else {
		return Post{}, errors.New("invalid client path")
	}
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	if c.path != "" {
		if userEmail != "" {
			dbData, err := c.readDB()
			if err == nil {
				dbPosts := dbData.Posts
				var posts []Post
				for _, post := range dbPosts {
					if post.UserEmail == userEmail {
						posts = append(posts, post)
					}
				}
				return posts, nil
			} else {
				return nil, err
			}
		} else {
			return nil, errors.New("empty userEmail")
		}
	} else {
		return nil, errors.New("invalid client path")
	}
}

func (c Client) DeletePost(id string) error {
	if c.path != "" {
		dbData, err := c.readDB()
		if err == nil {
			dbPosts := dbData.Posts
			delete(dbPosts, id)
			c.updateDB(dbData)
			return nil
		} else {
			return err
		}
	} else {
		return errors.New("invalid client path")
	}
}

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func (c Client) createDB() error {
	if c.path != "" {
		dbSchema := seedData()
		data, err := json.Marshal(dbSchema)
		if err != nil {
			return err
		}
		write_err := os.WriteFile(c.path, data, os.ModeAppend)
		if write_err != nil {
			return write_err
		}
		return nil

	} else {
		return errors.New("invalid client path")
	}
}

func seedData() databaseSchema {
	return databaseSchema{
		Users: map[string]User{
			"gch@domain.com": {
				Name:      "Giridhara",
				Email:     "gch@domain.com",
				Age:       29,
				Password:  "SimplePass",
				CreatedAt: time.Now().UTC(),
			},
			"siva@domain.com": {
				Name:      "Siva",
				Email:     "siva@domain.com",
				Age:       29,
				Password:  "SamplePass",
				CreatedAt: time.Now().UTC(),
			},
		},
		Posts: map[string]Post{
			"1001": {
				ID:        "1001",
				CreatedAt: time.Now().UTC(),
				UserEmail: "gch@domain.com",
				Text:      "Welcome aboard the new Go BE App!!!",
			},
			"1002": {
				ID:        "1002",
				CreatedAt: time.Now().UTC(),
				UserEmail: "siva@domain.com",
				Text:      "Welcome aboard the new Go BE App!!!",
			},
		},
	}
}

func (c Client) updateDB(db databaseSchema) error {
	if c.path != "" {
		data, err := json.Marshal(db)
		if err == nil {
			write_err := os.WriteFile(c.path, []byte(data), os.ModeAppend)
			if write_err != nil {
				return write_err
			}
			return nil
		} else {
			return err
		}
	} else {
		return errors.New("invalud client path")
	}
}

func (c Client) readDB() (databaseSchema, error) {
	if c.path != "" {
		data, err := os.ReadFile(c.path)
		if err == nil {
			dbData := databaseSchema{}
			unmarshall_err := json.Unmarshal(data, &dbData)
			if unmarshall_err == nil {
				return dbData, nil
			} else {
				return databaseSchema{}, unmarshall_err
			}
		} else {
			return databaseSchema{}, err
		}
	} else {
		return databaseSchema{}, errors.New("invalid client path")
	}
}
