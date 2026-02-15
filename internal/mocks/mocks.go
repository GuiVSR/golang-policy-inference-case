package mocks

// var MockGraph = gographviz.Graph{
// 	Name: "Policy",
// 	Nodes: map[string]*gographviz.Node{
// 		"start": {
// 			Name: "start",
// 			Attrs: map[string]string{
// 				"result": "",
// 			},
// 		},
// 		"approved": {
// 			Name: "approved",
// 			Attrs: map[string]string{
// 				"result": "approved=true,segment=prime",
// 			},
// 		},
// 		"rejected": {
// 			Name: "rejected",
// 			Attrs: map[string]string{
// 				"result": "approved=false",
// 			},
// 		},
// 		"review": {
// 			Name: "review",
// 			Attrs: map[string]string{
// 				"result": "approved=false,segment=manual",
// 			},
// 		},
// 	},
// 	Edges: []*gographviz.Edge{
// 		{
// 			Src:  "start",
// 			Dst:  "approved",
// 			Label: "",
// 			Attrs: map[string]string{
// 				"cond": "age>=18 && score>700",
// 			},
// 		},
// 		{
// 			Src:  "start",
// 			Dst:  "review",
// 			Label: "",
// 			Attrs: map[string]string{
// 				"cond": "age>=18 && score<=700",
// 			},
// 		},
// 		{
// 			Src:  "start",
// 			Dst:  "rejected",
// 			Label: "",
// 			Attrs: map[string]string{
// 				"cond": "age<18",
// 			},
// 		},
// 	},
// }
