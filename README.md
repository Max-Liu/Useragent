Useragent
=========

HTTP User Agent parser is wrote by golang

Usage
=========

  export ""useragent""
  
	ua :=useragent.NewUserAgent()
	ua.SetUseragent(r.UserAgent())
	log.Println(ua)
