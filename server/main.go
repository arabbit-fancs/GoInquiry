package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started")

	//Mongoセットアップ
	initMongo()

	//hc
	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I'm Alive")
	})

	http.HandleFunc("/inquiry/send", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set( "Access-Control-Allow-Origin", "*" )
		w.Header().Set( "Access-Control-Allow-Headers", "Content-Type" )

		//Postか判断
		if r.Method != http.MethodPost {
			resp := Response{"405", "NotAllowed"}
			json.NewEncoder(w).Encode(resp)
		}

		//jsonをInquiry型に変換
		var inquiry Inquiry
		if json.NewDecoder(r.Body).Decode(&inquiry) != nil {
			resp := Response{"400", "BadRequest"}
			json.NewEncoder(w).Encode(resp)
			return
		}

		//Mongo書き込み処理
		if writeInquiry(inquiry) != nil {
			resp := Response{"500", "InternalServerError"}
			json.NewEncoder(w).Encode(resp)
			return
		}

		//成功を返す
		resp := Response{"200", "データが挿入されました"}
		json.NewEncoder(w).Encode(resp)
		return
	})

	_ = http.ListenAndServe(":8000", nil)
}

//問い合わせ
type Inquiry struct {
	Name string
	Mail string
	Msg  string
}

//レスポンス
type Response struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
