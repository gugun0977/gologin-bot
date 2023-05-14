package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Token struct {
	IdToken      string `json:"token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func main() {
	// Menentukan seed yang berbeda setiap kali program dijalankan
	rand.Seed(time.Now().UnixNano())

	// Daftar karakter yang dapat digunakan
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	fmt.Print("Masukkan jumlah iterasi: ")
	var jumlahIterasi int
	fmt.Scanln(&jumlahIterasi)

	// Membuka file baru atau membuka file yang sudah ada dengan nama list.txt
	file, err := os.OpenFile("list.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < jumlahIterasi; i++ {
		// Membuat string acak dengan panjang maksimum 5 karakter
		result := make([]rune, 8)
		for i := range result {
			result[i] = chars[rand.Intn(len(chars))]
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		var data = strings.NewReader(`{"email":"` + string(result) + `@gmail.com","password":"12345678","passwordConfirm":"12345678","googleClientId":"108461031.1683996431","filenameParserError":"","fromApp":false,"fromAppTrue":false,"canvasAndFontsHash":"af08e22b3d622379","affiliate":"","fontsHash":"763ae5c0520834ac","userOs":"lin","canvasHash":"681760097"}`)
		req, err := http.NewRequest("POST", "https://api.gologin.com/user?free-plan=true", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Host", "api.gologin.com")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Language", "en-US,en;q=0.5")
		// req.Header.Set("Accept-Encoding", "gzip, deflate")
		req.Header.Set("Referer", "https://app.gologin.com/")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Gologin-Meta-Header", "site-win-10.0")
		req.Header.Set("Origin", "https://app.gologin.com")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-site")
		req.Header.Set("Te", "trailers")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s\n", bodyText)
		var token Token
		errUnmarshal := json.Unmarshal(bodyText, &token)
		if errUnmarshal != nil {
			log.Fatal(errUnmarshal)
		}
		// fmt.Printf("Token: %v\n", token.IdToken)
		// fmt.Fprintln(file, string(result)+"@gmail.com:"+token.IdToken)
		fmt.Fprintln(file, "Email Kamu : "+string(result)+"@gmail.com"+"\n"+"Password : 12345678\n")
		fmt.Printf("Email Kamu : " + string(result) + "@gmail.com" + "\n" + "Password : 12345678\n")
		// time.Sleep(time.Second * 10) // Menunggu selama 10 detik sebelum looping kembali
	}
}
