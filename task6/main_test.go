package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

type XMLUsers struct {
	Users []XMLUser `xml:"row"`
}

type XMLUser struct {
	ID        int    `xml:"id"`
	FirstName string `xml:"first_name"`
	LastName  string `xml:"last_name"`
	Age       int    `xml:"age"`
	About     string `xml:"about"`
	Gender    string `xml:"gender"`
}

func (xml_user XMLUser) toUser() User {
	return User{
		Id:     xml_user.ID,
		Name:   xml_user.LastName + " " + xml_user.FirstName,
		About:  xml_user.About,
		Age:    xml_user.Age,
		Gender: xml_user.Gender,
	}
}

// --------------------------------------------- USER'S FIELDS ----------------------------------------------

//-------USER_ID-------
type User_ID_asc []User
type User_ID_desc []User

// ASC
func (a User_ID_asc) Len() int           { return len(a) }
func (a User_ID_asc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a User_ID_asc) Less(i, j int) bool { return a[i].Id < a[j].Id }

// DESC
func (a User_ID_desc) Len() int           { return len(a) }
func (a User_ID_desc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a User_ID_desc) Less(i, j int) bool { return a[i].Id > a[j].Id }

//-------USER_ID-------

//-------USER_NAME-------
type User_name_asc []User
type User_name_desc []User

// ASC
func (a User_name_asc) Len() int           { return len(a) }
func (a User_name_asc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a User_name_asc) Less(i, j int) bool { return a[i].Name < a[j].Name }

// DESC
func (a User_name_desc) Len() int           { return len(a) }
func (a User_name_desc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a User_name_desc) Less(i, j int) bool { return a[i].Name > a[j].Name }

//-------USER_NAME-------

//-------USER_AGE-------
type User_age_asc []User
type User_age_desc []User

// ASC
func (a User_age_asc) Len() int           { return len(a) }
func (a User_age_asc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a User_age_asc) Less(i, j int) bool { return a[i].Age < a[j].Age }

// DESC
func (a User_age_desc) Len() int           { return len(a) }
func (a User_age_desc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a User_age_desc) Less(i, j int) bool { return a[i].Age > a[j].Age }

//-------USER_AGE-------

// --------------------------------------------- USER'S FIELDS ----------------------------------------------

func get_IDs(users []User) []int { // ? needed
	var IDs []int

	for _, user := range users {
		IDs = append(IDs, user.Id)
	}

	return IDs
}

func get_Names(users []User) []string { // ? needed
	var Names []string

	for _, user := range users {
		Names = append(Names, user.Name)
	}

	return Names
}

func get_Ages(users []User) []int { // ? needed
	var Ages []int

	for _, user := range users {
		Ages = append(Ages, user.Age)
	}

	return Ages
}

// --------------------------------------------- GET SEARCH PARAMS ------------------------------------------
func get_search_params(r *http.Request) url.Values {
	search_params, _ := url.ParseQuery(r.URL.String()[2:])

	return search_params
}

func get_limit(search_params url.Values) int {
	limit, err := strconv.Atoi(search_params["limit"][0])

	PanicOnErr(err)

	return limit
}

func get_offset(search_params url.Values) int {
	offset, err := strconv.Atoi(search_params["offset"][0])

	PanicOnErr(err)

	return offset
}

func get_query(seach_params url.Values) string {
	return seach_params["query"][0]
}

func get_order_field(search_params url.Values) string {
	return search_params["order_field"][0]
}

func get_order_by(search_params url.Values) string {
	return search_params["order_by"][0]
}

// --------------------------------------------- GET SEARCH PARAMS ------------------------------------------

func SearchServer(w http.ResponseWriter, r *http.Request) {
	filename := "dataset.xml" // ? filename

	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	defer f.Close()

	xml_data, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	var xml_users XMLUsers
	xml.Unmarshal(xml_data, &xml_users)

	search_params := get_search_params(r)

	limit := get_limit(search_params)
	offset := get_offset(search_params)
	query := get_query(search_params)
	order_field := get_order_field(search_params)
	order_by := get_order_by(search_params)

	var result []User

	for _, xml_user := range xml_users.Users {
		user := xml_user.toUser()

		if strings.Contains(user.Name, query) || strings.Contains(user.About, query) {
			result = append(result, user)
		}
	}

	switch order_field { // ? switch order_field
	case "Id":
		if order_by == "0" {
			sort.Sort(User_ID_desc(result))
		} else {
			sort.Sort(User_ID_asc(result))
		}
	case "Name":
		if order_by == "0" {
			sort.Sort(User_name_desc(result))
		} else {
			sort.Sort(User_name_asc(result))
		}
	case "Age":
		if order_by == "0" {
			sort.Sort(User_age_desc(result))
		} else {
			sort.Sort(User_age_asc(result))
		}
	default:
		sort.Sort(User_name_asc(result))
	}

	len := len(result)

	start := min(len, offset)
	fin := min(len, offset+limit)

	xml_data, _ = json.Marshal(result[start:fin])

	w.Write(xml_data)
}

func SearchServerTest(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(XMLUser{})

	w.Write(data)
}

func TestSearch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	defer ts.Close()

	// ----------TEST: limit < 0-----------------
	res, err := doSearch(ts.URL, 0, 5, "", "Name", 0)
	if err != nil {
		got_error := err.Error()
		expected_error := errors.New("limit must be > 0").Error()

		if got_error != expected_error {
			t.Errorf("expected \n%+v\n, got \n%+v\n", expected_error, got_error)
		}
	}
	// ----------TEST: limit < 0-----------------

	// ----------TEST: offset < 0----------------
	res, err = doSearch(ts.URL, 5, -1, "", "Name", 0)
	if err != nil {
		got_error := err.Error()
		expected_error := errors.New("offset must be > 0").Error()

		if got_error != expected_error {
			t.Errorf("expected \n%+v\n, got \n%+v\n", expected_error, got_error)
		}
	}
	// ----------TEST: offset < 0----------------

	// ----------SIMPLE TEST: limit, age---------
	expected := SearchResponse{Users: []User{User{
		Id:     13,
		Name:   "Davidson Whitley",
		Age:    40,
		About:  "Consectetur dolore anim veniam aliqua deserunt officia eu. Et ullamco commodo ad officia duis ex incididunt proident consequat nostrud proident quis tempor. Sunt magna ad excepteur eu sint aliqua eiusmod deserunt proident. Do labore est dolore voluptate ullamco est dolore excepteur magna duis quis. Quis laborum deserunt ipsum velit occaecat est laborum enim aute. Officia dolore sit voluptate quis mollit veniam. Laborum nisi ullamco nisi sit nulla cillum et id nisi.\n",
		Gender: "male",
	}, User{
		Id:     6,
		Name:   "Mays Jennings",
		Age:    39,
		About:  "Veniam consectetur non non aliquip exercitation quis qui. Aliquip duis ut ad commodo consequat ipsum cupidatat id anim voluptate deserunt enim laboris. Sunt nostrud voluptate do est tempor esse anim pariatur. Ea do amet Lorem in mollit ipsum irure Lorem exercitation. Exercitation deserunt adipisicing nulla aute ex amet sint tempor incididunt magna. Quis et consectetur dolor nulla reprehenderit culpa laboris voluptate ut mollit. Qui ipsum nisi ullamco sit exercitation nisi magna fugiat anim consectetur officia.\n",
		Gender: "male",
	}, User{
		Id:     26,
		Name:   "Cotton Sims",
		Age:    39,
		About:  "Ex cupidatat est velit consequat ad. Tempor non cillum labore non voluptate. Et proident culpa labore deserunt ut aliquip commodo laborum nostrud. Anim minim occaecat est est minim.\n",
		Gender: "male",
	}},
		NextPage: true,
	}

	res, _ = doSearch(ts.URL, 3, 1, "", "Age", 0)
	if !reflect.DeepEqual(*res, expected) {
		t.Errorf("expected \n%+v\n, got \n%+v\n", expected, *res)
	}
	// ----------SIMPLE TEST: limit, age---------

	// ----------TEST: search by about-----------
	expected = SearchResponse{Users: []User{User{
		Id:     16,
		Name:   "Osborn Annie",
		Age:    35,
		About:  "Consequat fugiat veniam commodo nisi nostrud culpa pariatur. Aliquip velit adipisicing dolor et nostrud. Eu nostrud officia velit eiusmod ullamco duis eiusmod ad non do quis.\n",
		Gender: "female",
	},
		User{
			Id:     28,
			Name:   "Hines Cohen",
			Age:    32,
			About:  "Deserunt deserunt dolor ex pariatur dolore sunt labore minim deserunt. Tempor non et officia sint culpa quis consectetur pariatur elit sunt. Anim consequat velit exercitation eiusmod aute elit minim velit. Excepteur nulla excepteur duis eiusmod anim reprehenderit officia est ea aliqua nisi deserunt officia eiusmod. Officia enim adipisicing mollit et enim quis magna ea. Officia velit deserunt minim qui. Commodo culpa pariatur eu aliquip voluptate culpa ullamco sit minim laboris fugiat sit.\n",
			Gender: "male",
		},
	},
		NextPage: false,
	}

	res, err = doSearch(ts.URL, 100, 0, "culpa pariatur", "Name", 0)
	if err != nil || !reflect.DeepEqual(*res, expected) {
		t.Errorf("expected \n%+v\n, got \n%+v\n", expected, *res)
	}
	// ----------TEST: search by about-----------

	res, err = doSearch("https://ru.linkedin.com/", 200, 0, "", "", 0)
	err, ok := err.(net.Error)
	if !ok {
		t.Errorf("expected net.Error, got \n%+v\n", err)
	}
}

func TestServer(t *testing.T) {
	// ----------
	ts := httptest.NewServer(http.HandlerFunc(SearchServerTest))

	res, err := doSearch(ts.URL, 20, 1, "", "Id", 1)

	expected := json.UnmarshalTypeError{Value: "object", Type: reflect.TypeOf([]User{})}

    got_error := err.Error()
    expected_error := expected.Error()
	
    if res != nil || got_error != expected_error {
		t.Errorf("expected err:\n%+v\n, got err:\n%+v\n", expected_error, got_error)
	}

	ts.Close()
	// ----------
	ts = httptest.NewServer(http.HandlerFunc(SearchServerTestTimeout))
	
	res, err = doSearch(ts.URL, 20, 10, "", "Name", 1)

    expected_error = expected.Error()
    
	if res != nil {
		if got_error, ok := err.(net.Error); !ok || got_error.Timeout() {
			t.Errorf("expected err:\n%+v\n, got err:\n%+v\n", expected_error, got_error.Error())
		}
	}

	ts.Close()
	// ----------
}

func SearchServerTestTimeout(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 2)
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
