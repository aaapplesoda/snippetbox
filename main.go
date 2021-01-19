package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// กำหนดฟังก์ชันhome handler ซึ่งเขียนด้วยbyte slice ที่มี
// "Hello from Snippetbox" เป็นตัว respond body.
func home(w http.ResponseWriter, r *http.Request) {

	//ตรวจสอบคำขอเส้นทางURLปัจจุบันว่าตรงกับ"/"หรือไม่
	//ถ้าไม่ http.Notfound() function จะส่ง 404 request not found ไปที่ client
	//ที่สำคัญเราต้อง return เพราะถ้าไม่ เราจะยังexecuteแสดงผล "Hello from Snippetbox".
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

//เพิ่ม showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {

	//ดึงค่าของ id parameter จาก query string และ
	//พยายามแปลงเป็น int ไดยใช้ strconv.Atoi() function ถ้า
	//ไม่สามารถแปลง็นintได้ หรือ ค่าน้อยกว่า 1
	//เราจะrespondคืน 404 not found
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Display a specific snippet..."))
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

//เพิ่ม createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {

	//ใช้ r.Method เพื่อเช็คว่า http method เป็น post รึเปล่า
	//ถ้าไม่ จะใช้ w.WriteHeader() method เพื่อส่ง 405 status code
	//w.Write method เพื่อแสดงrespond body "Method Not Allowed"
	//แล้วreturn เพื่อจะได้ไม่ execute โค้ดด้านล่าง
	if r.Method != "POST" {

		//ใช้ Header().Set() method เพื่อเพิ่ม 'Allow: POST' header
		//ที่จะตอบสนอง header map
		//parameter แรกคือ ชื่อheader ตัวสองคือ ค่าheader
		//http.Error function เหมือนรวม w.WriteHeader และ w.Write
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Now Allowed", 405)
		//codeเก่าแบบยาว
		//w.WriteHeader(405)
		//w.Write([]byte("Method Not Allowed"))

		return
	}
	w.Write([]byte("Create net snippet..."))
}

func main() {
	//ใช้ http.NewServeMux function เพื่อเริ่มต้น servemux ใหม่ หลังจากนั้น
	//บันทึก home function เป็น handler สำหรับรูปแบบ URL "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//ใช้ http.ListenAndServe() function เพื่อเริ่มต้น webserver
	//เราส่ง parameter 2 ค่าคือ TCP network address ที่ฟังอยู่บน(ในกรณี้นี้ฟังอยู่ที่พอร์ต":4000") กับ
	//servemux ที่เราพึ่งสร้าง
	//ถ้า http.ListenAndServe() return ค่า error มา เราใช้
	//log.Fatal() function เพื่อ log ค่า error message และ exit.
	log.Println("Start server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	http.ListenAndServe(":4000", mux)

}

//NOTE
//http.HandleFunc("/", home) ไม่ควรใช้ เพราะ defaultservemux เป็น global variable
//และทุกpackage สามารถเข้าถึงเพื่อ register route
// Use your own locally-scoped servemux

//if a user makes a request to /foo/bar/..//baz they will
//automatically be sent a 301 Permanent Redirect to /foo/baz instead
// if you have registered the subtree path /foo/,
//then any request to /foo will be redirected to /foo/

//can use for multiple website different and domain name
//การดำเนินการที่เปลี่ยนสถานะserver หรือdatabase ควรใช้ post method

//you must call w.WriteHeader() before any call to w.Write().
//method ในoop คือfunctionที่อยู่ในclass ส่วนในgolang คือ functionที่มี reciever
