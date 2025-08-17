package service

import (
	"fmt"
	"socialai/backend"
	"socialai/constants"
	"socialai/model"

	"github.com/olivere/elastic/v7"
)

// true / false: not nil
// true : nil: successfully
// false: nil: user not existed or wrong password

// 1. search by username, then compare passwords
// 2. search by username + password
//该function的作用是当client sign in时，验证user是否存在，service告诉repository去查，查完返回search result
func CheckUser(username, password string) (bool, error) {//bool 表示存在与否，error处理错误情况
   // 1. construct query
   query := elastic.NewBoolQuery()
   query.Must(elastic.NewTermQuery("username", username))
   query.Must(elastic.NewTermQuery("password", password))


   // 2. call backend
   searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
   if err != nil {
       return false, err
   }


   // 3 construct response
   if searchResult.TotalHits() > 0 {
       fmt.Printf("Login as %s\n", username)
       return true, nil
   }


   return false, nil
}


// false: user existed
// true: successfully added
func AddUser(user *model.User) (bool, error) {
   // check user existed or not
   query := elastic.NewTermQuery("username", user.Username)
   searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
   if err != nil {
       return false, err
   }


   if searchResult.TotalHits() > 0 {
       return false, nil
   }


   // call backend to save
   err = backend.ESBackend.SaveToES(user, constants.USER_INDEX, user.Username)
   if err != nil {
       return false, err
   }


   // construct response
   fmt.Printf("User is added: %s\n", user.Username)
   return true, nil
}