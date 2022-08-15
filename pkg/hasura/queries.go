package hasura

import (
	"encoding/json"
	
		"log"
		)
		
	 

func Query_blogs() []byte {

	gr := map[string]string{
		"query": ` query MyQuery { 
				nfb_blog(limit: 5) {
					title
					blog_id
					link
					content
				  }
            }

    `,
	}
	log.Println("============>>>gr")
	log.Println(gr)
	json, err := json.Marshal(gr)
	if err != nil {
		log.Println("Error Marshal @blogs")
		log.Println(err)
	}
	log.Println(" Query_Query_blogs")
	log.Println(err) 
	return json
}
func Query_users() []byte {

	gr := map[string]string{
		"query": ` query MyQuery { 
				nfb_users(limit: 5) {
					changed_on
					created_on
					email
					full_name
					id
					password
				  }
            }

    `,
    }
	log.Println("============>>>gr")
	log.Println(gr)
	json, err := json.Marshal(gr)
	if err != nil {
		log.Println("Error Marshal Query_users")
		log.Println(err) 
	}
	log.Println(" Query_users")
	log.Println(" Query_users")
	log.Println(" Query_users")
	return json
	}
func Query_coins() []byte {

	gr := map[string]string{
		"query": ` query MyQuery { 
				nfb_bitcoin(limit: 5) {
					title
					coin_id
					link
					coin_name
					content
					img
					intro
				  }
            }

    `,
	}
	log.Println("============>>>gr")
	log.Println(gr)
	json, err := json.Marshal(gr)
	if err != nil {
		log.Println("Error Marshal Query_ coins")
		log.Println(err) 
	}
	log.Println(" Q coins")
	log.Println(" Query_ coins") 
	return json
}


// func Query_user_asc() []byte {

// 	gr := map[string]string{
// 		"query": ` query MyQuery {
// 			lor_users(order_by: {full_name: asc}) {
// 				changed_on
// 				created_on
// 				email
// 				full_name
// 				id
// 				password
// 			  }
//         }
//     `,
// 	}
// 	log.Println("========>>>gr")
// 	log.Println(gr)
// 	json, err := json.Marshal(gr)
// 	if err != nil {
// 		log.Println("Error Marshal")
// 	}

// 	// log.Println(string(json))
// 	// log.Println("string(json)")
// 	return json
// }

// func Query_gens_byname() []byte {

// 	gr := map[string]string{
// 		"query": ` query MyQuery {
//           lor_users(limit: 10){
// 			changed_on
// 			created_on
// 			email
// 			full_name
// 			id
// 			password
// 		  }
//         }
//     `,
// 	}
// 	log.Println("============>>>gr")
// 	log.Println(gr)
// 	json, err := json.Marshal(gr)
// 	if err != nil {
// 		log.Println("Error Marshal")
// 	}

// 	return json
// }

// func Query_user() []byte {

// 	gr := map[string]string{
// 		"query": ` query MyQuery {
// 			lor_users {
// 				changed_on
// 				created_on
// 				email
// 				full_name
// 				id
// 				password
// 			  }
//         }
//     `,
// 	}
// 	log.Println("============>>>gr")
// 	log.Println(gr)
// 	json, err := json.Marshal(gr)
// 	if err != nil {
// 		log.Println("Error Marshal")
// 	}
// 	log.Println("Query_userjson");log.Println("Query_userjson");log.Println(json)
// 	log.Println("Query_userjson");log.Println("Query_userjson");log.Println(json)
// 	return json
// }

// func Mutation_add_gen(changed_on string, link string, content string,  created_on string , title string) []byte {
// 	log.Println(created_on)
// 	log.Println(changed_on) 
	 
// 	gq := map[string]string{
// 		"query": `
//       mutation MyMutation {
//          insert_lor_gens(objects: {changed_on: "` + changed_on  + `", content: "` + content  + `", created_on: "` + created_on + `", link: "` + link + `", title: "` + title + `"})  
// 	}
    
//     `,
// 	}

// 	json, err := json.Marshal(gq)
// 	if err != nil {
// 		log.Println("Error Marshal")
// 	}
// 	log.Println("json");log.Println("json");log.Println(json)
// 	return json

// }

// func Mutation_signup_user(password string, full_name string, email string,created_on string ,changed_on string  ) []byte {
//     log.Printf("@ Mutation_add_gen @Mutation_add_gen Form in  Mutation %v , %v , %v", password ,email , full_name)
// 	log.Printf("Form in  Mutation %v , %v , %v", password ,email , full_name)
// 	gr := map[string]string{
// 		"query": `mutation MyMutation {
// 			insert_lor_users_one(object: {password: "`+password+`",  full_name: "`+full_name+`", email: "`+email+`", created_on: "`+created_on+`", changed_on: "`+changed_on+`"}) {
//             id }
// 	  }
//     `,
// 	}
// 	log.Printf("qry string %v", gr);	log.Printf("qry string %v", gr);	log.Printf("qry string %v", gr);
// 	json, err := json.Marshal(gr)
// 	if err != nil {
// 		log.Println("Error Marshal")
// 	}
//     log.Println("json");log.Println("json");log.Println(json)
// 	return json
// }