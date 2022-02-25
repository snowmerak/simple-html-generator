package main

import (
	"fmt"

	"gitlab.com/snowmerak/simple-html-generator/e"
)

func main() {
	titleString := "Serve it Yourself!"

	metas := e.Elements{
		e.Text{
			Value: "<meta charset=\"UTF-8\">",
		},
		e.Text{
			Value: "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">",
		},
		e.Text{
			Value: "<meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">",
		},
		e.Text{
			Value: "<title>" + titleString + "</title>",
		},
	}

	title := e.Elements{
		e.Header{
			Level: 1,
			Value: e.Text{
				Value: titleString,
			},
		},
		e.Bold{
			Value: e.Italic{
				Value: e.Text{
					Value: "\"Serve it Yourself!\"는 백엔드 인프라를 직접 구현해보는 프로젝트입니다.",
				},
			},
		},
	}

	summery := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "개요",
			},
		},
	}

	roadmap := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "로드맵",
			},
		},
	}

	contributer := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "기여자",
			},
		},
	}

	community := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "커뮤니티",
			},
		},
	}

	projects := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "프로젝트",
			},
		},
	}

	doc := e.Document{
		Head: e.Head{
			Values: []e.Element{
				metas,
				e.Link{
					Href: "./mvp.css",
				},
			},
		},
		Body: e.Body{
			Values: []e.Element{
				title,
				e.Hr{},
				summery,
				e.Hr{},
				roadmap,
				e.Hr{},
				contributer,
				e.Hr{},
				community,
				e.Hr{},
				projects,
				e.Hr{},
			},
		},
	}
	fmt.Println(doc.Parse())
}
