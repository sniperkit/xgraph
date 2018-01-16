package bolt

import (
	"errors"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/schema"
)

// AdminService represents a BoltDB implementation of admin.UserService.
type AdminService struct {
	Store *cayley.Handle
}

// open a db connection and store it's handle inside AdminService instance
func NewAdminService(dbFile string) (*AdminService, error) {
	graph.InitQuadStore("bolt", dbFile, nil)

	// Open and use the database
	db, err := cayley.NewGraph("bolt", dbFile, nil)

	if err != nil {
		log.Fatalln(err)
	}

	adminService := &AdminService{Store: db}

	return adminService, nil
}

func validateEmail(email string) error {
	var ErrBadFormat = errors.New("invalid email format")
	var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func insert(store *cayley.Handle, o interface{}) error {
	qw := graph.NewWriter(store)
	defer qw.Close() // don't forget to close a writer; it has some internal buffering
	_, err := schema.WriteAsQuads(qw, o)
	return err
}
