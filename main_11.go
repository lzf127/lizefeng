package main

import (
	"fmt"
	"html/template"
	"net/http"

	"./data"
)

//index 载入界面
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/1.html")
	t.Execute(w, nil)
}

//tiao2 跳到第二界面并加载性格到redis
func tiao2(w http.ResponseWriter, r *http.Request) {
	ss1 := r.URL.Query().Get("01")
	//fmt.Println(ss)
	switch {
	case ss1 == "aa":
		xingge := "性格好，"
		err := data.Add(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	case ss1 == "ab":
		xingge := "性格可以，"
		err := data.Add(xingge)
		if err != nil {
			fmt.Println("ab err shi:", err)
		}
	case ss1 == "ac":
		xingge := "性格一般，"
		_ = data.Add(xingge)
	case ss1 == "ad":
		xingge := "性格暴躁，"
		_ = data.Add(xingge)
	}
	t, _ := template.ParseFiles("templates/2.html")
	t.Execute(w, nil)
}

//tiao3 跳到第三界面并加载爱好到redis
func tiao3(w http.ResponseWriter, r *http.Request) {
	ss2 := r.URL.Query().Get("02")
	switch {
	case ss2 == "ca":
		xingge := "爱好好，"
		err := data.Add2(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	case ss2 == "cb":
		xingge := "爱好一般，"
		err := data.Add2(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	case ss2 == "cc":
		xingge := "爱好广泛，"
		err := data.Add2(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	case ss2 == "cd":
		xingge := "爱好游泳，"
		err := data.Add2(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	}
	t, _ := template.ParseFiles("templates/3.html")
	t.Execute(w, nil)
}

//over 总结
func over(w http.ResponseWriter, r *http.Request) {
	ss2 := r.URL.Query().Get("03")
	switch {
	case ss2 == "da":
		xingge := "没什么特别。"
		err := data.Add3(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	case ss2 == "db":
		xingge := "一般般。"
		err := data.Add3(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	case ss2 == "dc":
		xingge := "就这样。"
		err := data.Add3(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	case ss2 == "dd":
		xingge := "不知道写什么。"
		err := data.Add3(xingge)
		if err != nil {
			fmt.Println("aa err shi:", err)
		}
	}
	text, _ := data.Htext()
	fmt.Println(text)
	t, _ := template.ParseFiles("templates/over.html")
	t.Execute(w, text)
}
