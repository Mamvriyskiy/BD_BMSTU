package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"strconv"
)


func main() {
	createClient()
}

func createClient() {
    name := []string{
        "John", "Luis", "Hiroshi", "Sven",
        "Ahmed", "Isabella", "Santiago", "Pablo", "Chen",
        "Alexei", "Carlos", "Rahul",
        "Diego", "Antonio", "Mikhail", "Hannah", "Javier",
        "Amara", "Gabriel", "Khalid", "Zoe", "Luca", "Leah",
        "Ivan", "Emma", "Andrei", "Grace", "Ravi",
        "Liam", "Ella", "Kazuki", "Lily", "Mateo", "Ava", "Anna", "Elena", 
		"Sophia", "Amina", "Nadia", "Emily", "Eva", "Layla",
        "Olivia", "Amelia", "Emma", "Ava", "Mia", "Lila", "Nina", "Sara", "Lena",
        "Maya", "Aisha", "Tara", "Ella", "Isabella", "Natalia",
    }
	
	lastNames := []string{
        "Smith", "Johnson", "Lee", "Kim", "Garcia", "Martinez", "Hernandez", "Sato",
        "Brown", "Davis", "Williams", "Jones", "Rodriguez", "Gonzalez", "Chen",
        "Nguyen", "Wang", "Lopez", "Gomez", "Wilson", "Johansson", "Ali", "Singh",
        "Johnson", "Harris", "Clark", "Turner", "Scott", "Murphy", "Liu", "Muller",
        "Hall", "Adams", "Green", "Evans", "Patel", "Khan", "Gutierrez", "O'Connor",
        "Zhang", "Martin", "Thomas", "Baker", "Yang", "White", "Taylor", "Wilson",
        "Moore", "Anderson", "Ivanov", "Smirnov", "Kuznetsov", "Popov", "Sokolov", "Lebedev", "Kozlov", "Novikov", 
		"Morozov", "Petrov", "Volkov", "Solovyov", "Vasilyev", "Zaytsev", "Pavlov", "Semyonov", 
		"Golubev", "Vinogradov", "Bogdanov", "Kharitonov", "Andreev", "Kondratiev", 
		"Ilyin", "Sorokin", "Fedorov", "Kuzmin", "Yakovlev", "Sokolov", "Gorbachev", "Konovalov", 
		"Yegorov", "Dmitriev", "Romanov", "Zakharov", "Smirnov", "Ivanov", "Mironov", 
		"Volkov", "Karpov", "Kuznetsov", "Klimov", "Smolov", "Golovin", "Chalov", "Tagirov", 
		"Verzilin", "Pisarev", "Rozhkov", "Gavrilov", "Averin", "Dubov", "Kryuchkov", "Shilov", "Butrimov", 
		"Starovoytov", "Amirov", "Davidov", "Anisimov", "Avdeev", "Bloxin",
		"Voronov", "Bogomolov",
    }

	createServices()

	file, err := os.OpenFile("client.csv", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	fileB, err := os.OpenFile("booking.csv", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer fileB.Close()

	fileIOP, err := os.OpenFile("invoiceforpayment.csv", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer fileIOP.Close()

	fileIS, err := os.OpenFile("invoiceservices.csv", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer fileIS.Close()

	fileR, err := os.OpenFile("room.csv", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	fileBR, err := os.OpenFile("bookingroom.csv", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// _, err = fmt.Fprint(file, "client_id,first_name,last_name,phone_number,date_of_birth,sex,passport\n")
	// _, err = fmt.Fprint(fileB, "booking_id,date_arrive,date_departure,count_room,invoice_id,client_id\n")
	// _, err = fmt.Fprint(fileIOP, "invoice_id,booking_id,date_of_issue,total_price\n")
	// _, err = fmt.Fprint(fileR, "room_id", ",", "room_number", ",", "category", ",", "price", "\n")
	// _, err = fmt.Fprint(fileBR, "booking_id", ",", "room_id", "\n")

	k := 1
	count := 1
	countBR := 1
	for i := 0; i < len(name); i++ {
		for j := 0; j < len(lastNames); j++ {
			s_name := []rune(name[i])
			_, passport := createCountryPsrt()
			if i >= 50 && string(s_name[len(s_name) - 1]) == "a" {
				err = inputFile(file, k, name[i], lastNames[j] + "a", "w", passport)
				//fmt.Println(rand.Intn(100), ",", name[i], surname[j] + "a")
				if err != nil {
					fmt.Println("Ошибка при записи данных в файл:", err)
					return
				}
			} else {
				err = inputFile(file, k, name[i], lastNames[j], "m", passport)
				//fmt.Println(rand.Intn(100), ",", name[i], surname[j])
				if err != nil {
					fmt.Println("Ошибка при записи данных в файл:", err)
					return
				}
			}

			n := rand.Intn(3) + 1

			if n == 2 {
				_, countBR = createBookingInvoice(fileB, fileIS, k, count, fileIOP, fileR, fileBR, countBR)
				count++
				_, countBR = createBookingInvoice(fileB, fileIS, k, count, fileIOP, fileR, fileBR, countBR)
				count++
			} else if n == 1 {
				_, countBR = createBookingInvoice(fileB, fileIS, k, count, fileIOP, fileR, fileBR, countBR)
				count++
			} 
			k++
		}
	}
}

func createRoom(file *os.File, k int) int {
	category := []string{"standard", "superior room", "apartments",
				"suite", "duplex", "presidential"}

	price := []int{2500, 3768, 4600, 5432, 7654, 12345}

	l := rand.Intn(6)

	fmt.Fprint(file, k + 1, ",", rand.Intn(1000) + 100, ",", category[l], ",", price[l], "\n")

	return price[l]
}

func createBookingInvoice(file, fileIS *os.File, client_id, booking_id int, fileIOP, fileR, fileBR *os.File, countBR int) (error, int) {

	// Создаем диапазон Unix timestamp для начала и конца года
    min := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
    max := time.Date(2030, 12, 31, 0, 0, 0, 0, time.UTC).Unix()

    // Генерируем первую случайную дату
    randomUnix1 := rand.Int63n(max-min+1) + min
    randomDate1 := time.Unix(randomUnix1, 0)
	date1 := randomDate1.Format("2006-01-02")

    // Генерируем вторую случайную дату с максимальным промежутком в месяц
    maxInterval := int64(30 * 24 * 60 * 60) // Примерно 30 дней в секундах
    randomUnix2 := randomUnix1 + rand.Int63n(maxInterval + 1)
    randomDate2 := time.Unix(randomUnix2, 0)
	date2 := randomDate2.Format("2006-01-02")

	count := rand.Intn(4) + 1 
	if count == 4 {
		count = 1
	}

	minDays := -150
	maxDays := -15

	randomDays := rand.Intn(maxDays - minDays + 1) + minDays

	newDate := randomDate1.AddDate(0, 0, randomDays)

	 _, err := fmt.Fprint(file, booking_id, ",", date1, ",", date2, ",",
	 						count, ",", booking_id, ",", client_id, "\n")

	price := []int{2000, 1950, 900, 455, 3599, 2599, 3200, 8999, 999, 2500}

	b, _ := time.Parse("2006-01-02", date1)
    a, _ := time.Parse("2006-01-02", date2)

    // Вычисляем разницу между датами
    duration := a.Sub(b)

    // Извлекаем количество дней из разницы
    days := int(duration.Hours() / 24)
	
	if days == 0 {
		days += 1
	}

	var total_price int


	for j := 0; j < count; j++ {
		num := createRoom(fileR, countBR)
		total_price += days * num
		createBookingRoom(fileBR, booking_id, countBR)
		countBR += 1
	}

	count = rand.Intn(10)

	for i := 0; i < count; i++ {
		if i == 0 {
			total_price += price[i] * days
		} else {
			total_price += price[i] * rand.Intn(days)
		}
		createIS(fileIS, booking_id, i)
	}

	_, err = fmt.Fprint(fileIOP, booking_id, ",", booking_id, ",", newDate.Format("2006-01-02"), ",", total_price, "\n")

	return err, countBR
}

func createBookingRoom(file *os.File, booking_id, room_id int) error {
	_, err := fmt.Fprint(file, booking_id, ",", room_id, "\n")

	return err
}

func createIS(fileIS *os.File, id, i int) error {
	_, err := fmt.Fprint(fileIS, id, ",", i + 1,"\n")

	return err
}

func createServices() {
	file, err := os.OpenFile("services.csv", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprint(file, "servies_id", ",", "name_servies", ",", "price", "\n")

	servies := []string{"nutrition", "dry cleaning", "bathhouse", "gym", "luggage storage", 
						"excursions", "translator", "rent of a conference hall", 
						"room service", "car rental"}

	price := []int{2500, 2000, 1950, 900, 455, 2599, 3200, 8999, 999, 3599}

	for i := 1; i < len(price) + 1; i++ {
		_, err = fmt.Fprint(file, i, ",", servies[i - 1], ",", price[i - 1], "\n")		
	}


}

func inputFile(file *os.File, k int, name, surname, sex, passport string) error {

	_, err := fmt.Fprint(file, k, ",", 
						name, ",", surname, ",", createTelefonNUmber(), 
						",", createDate(1940, 2020), ",", sex, ",", passport, "\n")

	return err
}

func createNumberRoom() int {
	return rand.Intn(5500 + 1)
}

func createCountryPsrt() (string, string) {
	countries := []string{
        "Russia", "United States", "Canada", "United Kingdom", "Australia", "Germany", "France", "Spain", "Italy",
        "Japan", "China", "Brazil", "Argentina", "Mexico", "India", "South Africa",
        "Egypt", "Nigeria", "Ghana", "Kenya", "Ethiopia", "Uganda", "Saudi Arabia", "Turkey",
        "Pakistan", "Bangladesh", "Indonesia", "Vietnam", "Thailand", "South Korea", "Philippines",
        "New Zealand", "Netherlands", "Belgium", "Sweden", "Norway", "Denmark", "Finland", "Greece",
    }

	num := rand.Intn(39)
	k := rand.Intn(999999) + 1
	if k < 100001 {
		k += 100000
	}

	psrt := strconv.Itoa(num + 1) + " " + strconv.Itoa(rand.Intn(23) + 1) + " " + strconv.Itoa(k)


	return countries[num], psrt
}

func createDate(a, b int) string {
    min := time.Date(a, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
    max := time.Date(b, 12, 31, 0, 0, 0, 0, time.UTC).Unix()
    randomDate := time.Unix(rand.Int63n(max-min)+min, 0)

	return randomDate.Format("2006-01-02")
}

func createTelefonNUmber() string {
	regionCode := rand.Intn(1000)
    regionCodeStr := fmt.Sprintf("%03d", regionCode)

    // Генерация номера (7 цифр)
    phoneNumber := rand.Intn(10000000)
    phoneNumberStr := fmt.Sprintf("%07d", phoneNumber)

    // Форматирование телефонного номера в стандартный вид
    formattedPhoneNumber := fmt.Sprintf("+7(%s)%s-%s-%s", regionCodeStr, phoneNumberStr[:3], 
											phoneNumberStr[3:5], phoneNumberStr[5:])

	return formattedPhoneNumber
}
