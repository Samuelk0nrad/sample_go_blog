package blog

func GetAllPosts() []BlogPost {
	return []BlogPost{
		{ID: 1, Title: "First blog post"},
		{ID: 2, Title: "Second blog post"},
	}
}
