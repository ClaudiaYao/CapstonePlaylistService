package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lithammer/shortuuid"
)

// open address file and store them in a slice of structs
type address struct {
	unit_number   string
	address_line1 string
	address_line2 string
	postal_code   int
}

func main() {
	categoryCodes := generateCategory()
	restaurantIDs := generateRestaurant(categoryCodes)
	// since Dish table has a foreign reference key to RestaurantID, so we
	// pass restaurantIDs to the dish generation function
	dishIDs := generateDish(restaurantIDs)
	playlistIDs := generatePlaylist(categoryCodes)
	generatePlaylistDishRelation(playlistIDs, dishIDs)

}

func generatePlaylist(categoryCodes []string) []string {

	read_f, err := os.Open("Initial/playlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer read_f.Close()

	write_f, err := os.Create("Generated/playlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer write_f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(read_f)
	categoryNumber := len(categoryCodes)

	dietaryInfo := []string{
		"perferendis voluptatibus veniam",
		"veniam ut excepturi nulla conse",
		"fugit, amet quia nulla culpa",
		"nulla consequatur natus tempore officiis",
	}
	dietaryNumber := len(dietaryInfo)

	playlistIDs := []string{}

	// Get each playlist name and then form the table content
	for scanner.Scan() {
		id := "Play" + shortuuid.New()
		// do something with a line
		name := scanner.Text()
		categoryCode := categoryCodes[rand.Intn(categoryNumber)]
		dietary := dietaryInfo[rand.Intn(dietaryNumber)]

		startDateRandom := time.Now().AddDate(0, 0, rand.Intn(60)-30)
		startDate := startDateRandom.Format("2006-01-02")
		endDateRandom := startDateRandom.AddDate(0, 2, 15)
		end_date := endDateRandom.Format("2006-01-02")
		popularity := 1 + rand.Intn(5)

		statusInfo := ""
		if startDateRandom.After(time.Now()) {
			statusInfo = "Pending"
		} else if endDateRandom.Before(time.Now()) {
			statusInfo = "Expired"
		} else if endDateRandom.After(time.Now()) {
			statusInfo = "Active"
		}

		new_text := id + "|" + name + "|" + categoryCode + "|" +
			dietary + "|" + statusInfo + "|" + startDate +
			"|" + end_date + "|" + strconv.Itoa(popularity)

		_, err := write_f.WriteString(new_text + "\n")
		if err != nil {
			log.Fatal("error occurs when writing to file:", err)
		}
		playlistIDs = append(playlistIDs, id)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return playlistIDs

}

func generatePlaylistDishRelation(playlistIDs []string, dishIDs []string) {

	write_f, err := os.Create("Generated/playlist_dish.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer write_f.Close()

	// read the file line by line using scanner

	totalDish := len(dishIDs)
	if totalDish < 8 {
		fmt.Println("too less dishes in the database.")
	}

	// Get each playlist name and then form the table content
	for _, playlistID := range playlistIDs {

		// do something with a line

		dishNum := 2 + rand.Intn(3)
		chosen := 0

		for chosen < dishNum {
			id := "PD" + shortuuid.New()
			dishID := dishIDs[rand.Intn(totalDish)]
			new_text := id + "|" + dishID + "|" + playlistID

			_, err := write_f.WriteString(new_text + "\n")
			if err != nil {
				log.Fatal("error occurs when writing to file:", err)
			}

			chosen += 1

		}

	}
}

func generateDish(restaurantIDs []string) []string {

	read_f, err := os.Open("Initial/dish.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer read_f.Close()

	write_f, err := os.Create("Generated/dish.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer write_f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(read_f)
	restaurantNumber := len(restaurantIDs)
	dishIDs := []string{}

	cuisineStyles := []string{
		"perferendis voluptatibus veniam",
		"veniam ut excepturi nulla conse",
		"fugit, amet quia nulla culpa",
		"nulla consequatur natus tempore officiis",
	}

	dishOptions := []string{
		`[Mentaico Source, Yes, No]`,
		`[Wasabi, Yes, No]`,
	}

	imageUrls := []string{
		"https://images.unsplash.com/photo-1600891964092-4316c288032e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80",
		"https://images.unsplash.com/photo-1597289124948-688c1a35cb48?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=987&q=80",
		"https://images.unsplash.com/photo-1633271332313-04df64c0105b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80",
		"https://images.unsplash.com/photo-1579871494447-9811cf80d66c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80",
	}
	cuisineStyleNumber := len(cuisineStyles)

	ingredients := []string{
		"voluptatibus, veniam",
		"veniam, ut, excepturi",
		"fugit, amet quia, nulla",
		"consequatur, natus, tempore, officiis",
	}
	ingredientNumber := len(ingredients)

	// Get each dish name and then form the table content
	for scanner.Scan() {
		id := "Dish" + shortuuid.New()
		// do something with a line
		name := scanner.Text()
		restaurantID := restaurantIDs[rand.Intn(restaurantNumber)]
		comment := "extra information"
		price := fmt.Sprintf("%.2f", 4.0+rand.Float32()*20)
		if err != nil {
			fmt.Println(err)
		}
		cuisineStyle := cuisineStyles[rand.Intn(cuisineStyleNumber)]
		ingredient := ingredients[rand.Intn(ingredientNumber)]
		dishOption := dishOptions[rand.Intn(len(dishOptions))]
		imageUrl := imageUrls[rand.Intn(len(imageUrls))]

		new_text := id + "|" + name + "|" + restaurantID + "|" +
			price + "|" + cuisineStyle + "|" + ingredient + "|" +
			dishOption + "|" + comment + "|" + imageUrl

		_, err := write_f.WriteString(new_text + "\n")
		if err != nil {
			log.Fatal("error occurs when writing to file:", err)
		}
		dishIDs = append(dishIDs, id)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dishIDs

}
func generateCategory() []string {
	// open file
	read_f, err := os.Open("Initial/category.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer read_f.Close()

	// create file
	features := []string{
		"Temporibus aliquid, obcaecati soluta consequatur a veritatis ad omnis",
		"Rerum sequi, earum delectus quidem tenetur est dicta exercitationem eius labore ipsa",
		"Quibusdam perferendis voluptatibus veniam ut excepturi nulla",
	}

	write_f, err := os.Create("Generated/category.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer write_f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(read_f)
	feature_number := len(features)

	categoryCodes := []string{}
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		code := line[:3]
		new_text := code + "|" + line + "|" + features[rand.Intn(feature_number)]
		_, err := write_f.WriteString(new_text + "\n")
		if err != nil {
			log.Fatal("error occurs when writing to file:", err)
		}

		categoryCodes = append(categoryCodes, code)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return categoryCodes
}

func GenerateAddress() []address {

	var addresses []address
	// read the file line by line using scanner
	read_address, err := os.Open("Initial/address.txt")

	if err != nil {
		log.Fatal("read address file:", err)
	}
	defer read_address.Close()

	scanner := bufio.NewScanner(read_address)

	for scanner.Scan() {
		line := scanner.Text()
		adds := strings.Split(line, ",")
		if len(adds) == 3 {
			postal, err := strconv.Atoi(strings.TrimSpace(adds[2]))
			if err != nil {
				log.Fatal("postal code format is incorrect: ", err)
			}
			new_address := address{unit_number: strings.TrimSpace(adds[0]),
				address_line1: strings.TrimSpace(adds[1]),
				address_line2: "(empty)", postal_code: postal}

			addresses = append(addresses, new_address)
		} else if len(adds) == 4 {
			postal, err := strconv.Atoi(strings.TrimSpace(adds[3]))
			if err != nil {
				log.Fatal("postal code format is incorrect.", err)
			}
			new_address := address{unit_number: strings.TrimSpace(adds[0]),
				address_line1: strings.TrimSpace(adds[1]),
				address_line2: strings.TrimSpace(adds[2]),
				postal_code:   postal}

			addresses = append(addresses, new_address)
		} else {
			log.Fatal("format of the address.txt is incorrect.")
		}

	}
	return addresses
}

func generateRestaurant(categories []string) []string {
	addresses := GenerateAddress()

	read_f, err := os.Open("Initial/restaurant.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer read_f.Close()

	write_f, err := os.Create("Generated/restaurant.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer write_f.Close()

	urls := []string{
		"https://example1.com",
		"https://example2.com",
		"https://example3.com",
	}
	operationStartOptions := []string{"0600", "0700", "0800", "0900", "1100", "1200"}
	operationEndOptions := []string{"1900", "2100", "2300", "2400", "0200"}

	// read the file line by line using scanner
	scanner := bufio.NewScanner(read_f)
	address_number := len(addresses)

	restaurantIDs := []string{}

	for scanner.Scan() {
		id := "RES" + shortuuid.New()
		// do something with a line
		name := scanner.Text()
		restaurant := addresses[rand.Intn(address_number)]
		openStart := operationStartOptions[rand.Intn(len(operationStartOptions))]
		openEnd := operationEndOptions[rand.Intn(len(operationEndOptions))]

		// oh := []string{openStart, openEnd}
		// b, err := json.Marshal(oh)
		// operateHours := string(b)
		operateHours := "[" + openStart + "," + openEnd + "]"

		logoUrl := urls[rand.Intn(len(urls))]
		headerUrl := urls[rand.Intn(len(urls))]
		// headerUrl := ""
		tag := categories[rand.Intn(len(categories))]

		new_text := id + "|" + name + "|" + restaurant.unit_number + "|" +
			restaurant.address_line1 + "|" + restaurant.address_line2 + "|" +
			strconv.Itoa(restaurant.postal_code) + "|" + tag + "|" + operateHours +
			"|" + logoUrl + "|" + headerUrl

		_, err = write_f.WriteString(new_text + "\n")
		if err != nil {
			log.Fatal("error occurs when writing to file:", err)
		}
		restaurantIDs = append(restaurantIDs, id)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return restaurantIDs
}

// first_names := []string{"James", "Robert", "John", "Michael", "David", "William", "Richard",
// 	"Joseph", "Thomas", "Charles", "Christopher", "Daniel", "Matthew", "Anthony", "Mark",
// 	"Donald", "Steven", "Paul", "Andrew", "Joshua", "Kenneth", "Kevin", "Brian", "George",
// 	"Timothy", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan",
// 	"Jessica", "Sarah", "Karen", "Lisa", "Nancy", "Betty", "Margaret", "Sandra",
// 	"Ashley", "Kimberly", "Emily", "Donna", "Michelle", "Carol", "Amanda", "Dorothy",
// 	"Melissa", "Deborah", "Stephanie", "Rebecca"}
// len_first := len(first_names)

// last_names := []string{"Tan", "Lim", "Lee", "Ng", "Ong", "Wong", "Goh",
// 	"Chua", "Chan", "Koh", "Teo", "Ang", "Yeo", "Tay", "Ho", "Low", "Toh", "Sim",
// 	"Chong", "Chia", "Seah"}
// len_last := len(last_names)

// email_dn := []string{"@gmail.com", "@hotmail.com", "@yahoo.com", "@dental.com"}
// len_dn := len(email_dn)

// // all the methods to generate unique user id: https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
// // use a shorter version

// for count := 0; count < 200; count++ {
// 	id := shortuuid.New()

// 	first_name := first_names[rand.Intn(len_first)]
// 	last_name := last_names[rand.Intn(len_last)]
// 	user_name := first_name + last_name + id[:4]
// 	email_add := first_name + "." + last_name + id[:4] + email_dn[rand.Intn(len_dn)]
// 	phone_num := strconv.Itoa(70000000 + rand.Intn(10000000))
// 	bPassword, err := bcrypt.GenerateFromPassword([]byte(user_name), bcrypt.MinCost)
// 	if err != nil {
// 		log.Panic("could not generate test data password successfully.")
// 		return
// 	}
// 	user := User{UserID: id, UserName: user_name, FirstName: first_name,
// 		LastName: last_name, Password: bPassword, EmailAddress: email_add,
// 		ContactPhone: phone_num}
// 	UserTestMap[id] = user

// }
// SaveToTestUserJSON()
