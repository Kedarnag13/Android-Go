package account

import (
"fmt"
)

type registrationController struct{}

var Registration registrationController

func (r registrationController) Create(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	var u models.User
	flag := 0

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
	users, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL,firstname varchar(255),lastname varchar(255),email varchar(255),password varchar(255),password_confirmation varchar(255),city varchar(255),state varchar(255),
		country varchar(255),
		createdat timestamptz,
		user_thumbnail varchar(2083),
		mobile_number text,
		devise_token varchar(2083),
		status boolean DEFAULT FALSE,
		status_message varchar(2083),
		PRIMARY KEY(id, mobile_number))")
	if err != nil || users == nil {
		log.Fatal(err)
	}
	devices, err := db.Exec("CREATE TABLE  IF NOT EXISTS devices (
		id int,
		devise_token varchar(320),
		PRIMARY KEY(devise_token),
		user_id int,
		CONSTRAINT user_id_key FOREIGN KEY(user_id)
		REFERENCES users(id))")
	if err != nil || devices == nil {
		log.Fatal(err)
	}
	sessions, err := db.Exec("CREATE TABLE IF NOT EXISTS sessions (
		id int,
		start_time timestamptz,
		end_time timestamptz,
		user_id int,
		CONSTRAINT sessions_id_key FOREIGN KEY(user_id)
		REFERENCES users(id),
		devise_token varchar(320),
		CONSTRAINT sessions_devise_key FOREIGN KEY(devise_token)
		REFERENCES devices(devise_token))")
	if err != nil || sessions == nil {
		log.Fatal(err)
	}
	fetch_id, err := db.Query("select coalesce(max(id), 0)from users")
	if err != nil {
		log.Fatal(err)
	}
	mobile_number_validation := `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`
	exp, err := regexp.Compile(mobile_number_validation)
	if err != nil {
		os.Exit(1)
	}
	if u.Firstname == "" || u.Password == "" || u.Password_confirmation == "" || u.Mobile_number == "" || !exp.MatchString(u.Mobile_number) || u.Devise_token == "" {
		result, err := govalidator.ValidateStruct(u)
		if err != nil || result == false {
			println("error: " + err.Error())
		}
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   err.Error(),
			})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		flag = 0
		goto sign_up_end
	} else if u.Password != u.Password_confirmation {
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   "Passwords do not match",
			})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		flag = 0
		goto sign_up_end
	} else {
		var sign_up_query string = "insert into users (id, firstname, lastname, email, password, password_confirmation,city,state,country,user_thumbnail,mobile_number,devise_token,status,status_message) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)"
		sign_up_prepare, err := db.Prepare(sign_up_query)
		if err != nil {
			log.Fatal(err)
		} 
		sign_up_result, err := sign_up_prepare.Exec(id, u.Firstname, u.Lastname, u.Email, u.Password, u.Password_confirmation, u.City, u.State, u.Country, u.User_thumbnail, u.Mobile_number,u.Devise_token, u.Status, u.Status_message)
		if err != nil || sign_up_result == nil {
			log.Fatal(err)
		}
		user := models.User{}
		b, err := json.Marshal(models.SignUpMessage{
			Success: "true",
			Message: "User created Successfully!",
			User:    user,
			})

		if err != nil || res == nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
	get_mobile_number, err := db.Query("SELECT mobile_number FROM users where mobile_number=$1", u.Mobile_number)
	if err != nil || get_mobile_numbers == nil {
		log.Fatal(err)
	}
	get_mobile_number.Next() {
		var mobile_number int
		err := get_mobile_numbers.Scan(&mobile_number)
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   "Number already exist",
			})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		flag = 0
	}
	sign_up_end:
}
