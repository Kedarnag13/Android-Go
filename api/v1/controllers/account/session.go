package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Android-Go/api/v1/models"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type sessionController struct{}

var Session sessionController

func (s sessionController) Create(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	var u models.User
	var ses models.Session
	flag := 1

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", "password=password host=localhost dbname=android_go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	users, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL,firstname varchar(255),lastname varchar(255),email varchar(255),password varchar(255),password_confirmation varchar(255),city varchar(255),state varchar(255),country varchar(255),createdat timestamptz,user_thumbnail varchar(2083),mobile_number text,devise_token varchar(2083),status boolean DEFAULT FALSE,status_message varchar(2083), PRIMARY KEY(id, mobile_number))")
	if err != nil || users == nil {
		log.Fatal(err)
	}
	devices, err := db.Exec("CREATE TABLE  IF NOT EXISTS devices (id int,devise_token varchar(320),PRIMARY KEY(devise_token),user_id int, CONSTRAINT user_id_key FOREIGN KEY(user_id) REFERENCES users(id))")
	if err != nil || devices == nil {
		log.Fatal(err)
	}
	sessions, err := db.Exec("CREATE TABLE IF NOT EXISTS sessions (id int,start_time timestamptz,end_time timestamptz,user_id int,CONSTRAINT sessions_id_key FOREIGN KEY(user_id) REFERENCES users(id), devise_token varchar(320), CONSTRAINT sessions_devise_key FOREIGN KEY(devise_token) REFERENCES devices(devise_token))")
	if err != nil || sessions == nil {
		log.Fatal(err)
	}
	mobile_number_validation := `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`
	exp, err := regexp.Compile(mobile_number_validation)
	if err != nil {
		os.Exit(1)
	}
	get_user_id_pass_token, err := db.Query("SELECT id,password,devise_token FROM users where mobile_number=$1", u.Mobile_number)
	if err != nil {
		log.Fatal(err)
	}
	for get_user_id_pass_token.Next() {
		var user_id int
		var password string
		var devise_token string
		err := get_user_id_pass_token.Scan(&user_id, &password, &devise_token)
		get_login_details, err := db.Query("SELECT devise_token FROM sessions")
		if err != nil {
			log.Fatal(err)
		}
		for get_login_details.Next() {
			var devise_token string
			err := get_login_details.Scan(&devise_token)
			if err != nil {
				log.Fatal(err)
			}
			if u.Mobile_number == "" || !exp.MatchString(u.Mobile_number) || u.Devise_token == "" {
				result, err := govalidator.ValidateStruct(u)
				if err != nil || result == false {
					println("error: " + err.Error())
				}
				b, err := json.Marshal(models.LogInErrorMessage{
					Success: "false",
					Error:   err.Error(),
				})
				if err != nil {
					log.Fatal(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				flag = 0
				goto login_end
			} else if u.Devise_token == devise_token {
				b, err := json.Marshal(models.LogInErrorMessage{
					Success: "false",
					Error:   "User already Exists",
				})
				if err != nil {
					log.Fatal(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				flag = 0
				goto login_end
			}
		}
		if flag == 1 {
			var insert_into_devise_query string = "insert into devices (devise_token,user_id) values ($1,$2)"
			insert_into_devise_prepare, err := db.Prepare(insert_into_devise_query)
			if err != nil {
				log.Fatal(err)
			}
			insert_into_devise_result, err := insert_into_devise_prepare.Exec(u.Devise_token, user_id)
			if err != nil || insert_into_devise_result == nil {
				log.Fatal(err)
			}
			var login_query string = "insert into sessions (start_time,user_id,devise_token) values ($1,$2,$3)"
			login_prepare, err := db.Prepare(login_query)
			if err != nil {
				log.Fatal(err)
			}
			login_result, err := login_prepare.Exec(ses.StartTime, user_id, u.Devise_token)
			if err != nil || login_result == nil {
				log.Fatal(err)
			}
			user := models.User{}
			b, err := json.Marshal(models.LogInSuccessMessage{
				Success: "true",
				Message: "User logged in successfully",
				User:    user,
			})

			if err != nil {
				log.Fatal(err)
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			flag = 1
			goto login_end
		}
	}
	if flag == 0 {
		b, err := json.Marshal(models.LogInErrorMessage{
			Success: "false",
			Error:   "User does not exist",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
login_end:
}

func (s sessionController) Destroy(rw http.ResponseWriter, req *http.Request) {

	var u models.User
	flag := 1
	vars := mux.Vars(req)
	devise_token := vars["devise_token"]
	u.Devise_token = devise_token

	db, err := sql.Open("postgres", "password=password host=localhost dbname=android_go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	users, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL,firstname varchar(255),lastname varchar(255),email varchar(255),password varchar(255),password_confirmation varchar(255),city varchar(255),state varchar(255),country varchar(255),createdat timestamptz,user_thumbnail varchar(2083),mobile_number text,devise_token varchar(2083),status boolean DEFAULT FALSE,status_message varchar(2083), PRIMARY KEY(id, mobile_number))")
	if err != nil || users == nil {
		log.Fatal(err)
	}
	devices, err := db.Exec("CREATE TABLE  IF NOT EXISTS devices (id int,devise_token varchar(320),PRIMARY KEY(devise_token),user_id int, CONSTRAINT user_id_key FOREIGN KEY(user_id) REFERENCES users(id))")
	if err != nil || devices == nil {
		log.Fatal(err)
	}
	sessions, err := db.Exec("CREATE TABLE IF NOT EXISTS sessions (id int,start_time timestamptz,end_time timestamptz,user_id int,CONSTRAINT sessions_id_key FOREIGN KEY(user_id) REFERENCES users(id), devise_token varchar(320), CONSTRAINT sessions_devise_key FOREIGN KEY(devise_token) REFERENCES devices(devise_token))")
	if err != nil || sessions == nil {
		log.Fatal(err)
	}
	get_devise_tokens, err := db.Query("SELECT devise_token FROM sessions")
	fmt.Println(get_devise_tokens)
	if err != nil || get_devise_tokens == nil {
		log.Fatal(err)
	}
	if u.Devise_token == "" {
		d, err := json.Marshal(models.LogOutErrorMessage{
			Success: "false",
			Error:   "Devise token cant be empty",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(d)
		goto destroy_end
	}
	for get_devise_tokens.Next() {
		var devise_token string
		err := get_devise_tokens.Scan(&devise_token)
		if err != nil {
			log.Fatal(err)
		}
		if devise_token == u.Devise_token {
			var destroy_from_sessions_query string = "DELETE FROM sessions where devise_token=$1"
			destroy_from_sessions_prepare, err := db.Prepare(destroy_from_sessions_query)
			if err != nil {
				log.Fatal(err)
			}
			destroy_from_sessions_result, err := destroy_from_sessions_prepare.Exec(u.Devise_token)
			if err != nil || destroy_from_sessions_result == nil {
				log.Fatal(err)
			}
			var destroy_from_devices_query string = "DELETE FROM devices where devise_token=$1"
			destroy_from_devices_prepare, err := db.Prepare(destroy_from_devices_query)
			if err != nil {
				log.Fatal(err)
			}
			destroy_from_devices_result, err := destroy_from_devices_prepare.Exec(u.Devise_token)
			if err != nil || destroy_from_devices_result == nil {
				log.Fatal(err)
			}
			user := models.User{}
			b, err := json.Marshal(models.LogOutSuccessMessage{
				User:    user,
				Success: "true",
				Message: "Logged out Successfully",
			})
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			flag = 0
			goto destroy_end
		}
	}
	if flag == 1 {
		d, err := json.Marshal(models.LogOutErrorMessage{
			Success: "false",
			Error:   "Session does not exist",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(d)
	}
destroy_end:
}
