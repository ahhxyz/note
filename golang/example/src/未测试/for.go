	for {
		fInfo, _ := os.Stat(S.file) //该函数不会更改文件的访问时间
		info := fInfo.Sys()
		v := reflect.ValueOf(info).Elem()
		Atim := v.FieldByName("Atim").Field(0).Int()
		fmt.Printf("%s\n%s\n", Atim+int64(s.LifeTime.Seconds()), time.Now().Unix())
		time.Sleep(time.Second * 3)

		<-c
		if Atim+int64(s.LifeTime.Seconds()) < time.Now().Unix() {
			fmt.Printf("del")
			s.lock.Lock()
			defer s.lock.Lock()
			s.Destroy()
		}
	}
for sess, _ := range sessions {
        timeacc, _ := sess.Get("timeAccessed")
        if timeacc.(int64)+int64(sessionTimeout) < time.Now().Unix() {
            sess.RemoveSession()
        }
    }
    time.AfterFunc(time.Duration(60)*time.Second, func() { SessionGC() })