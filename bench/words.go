package bench

import "math/rand"

var words = []string{
	"blandit",
	"elit",
	"condimentum",
	"ligula",
	"eros",
	"tellus",
	"consectetur",
	"etiam",
	"sapien",
	"non",
	"dignissim",
	"congue",
	"nibh",
	"tincidunt",
	"bibendum",
	"nisi",
	"consequat",
	"malesuada",
	"integer",
	"odio",
	"faucibus",
	"ipsum",
	"neque",
	"mattis",
	"nisl",
	"viverra",
	"augue",
	"porta",
	"lacinia",
	"orci",
	"ultricies",
	"mi",
	"sagittis",
	"sed",
	"vivamus",
	"eu",
	"et",
	"molestie",
	"risus",
	"nec",
	"quisque",
	"nunc",
	"sem",
	"mollis",
	"suscipit",
	"elementum",
	"velit",
	"leo",
	"euismod",
	"magna",
	"morbi",
	"maecenas",
	"phasellus",
	"amet",
	"aenean",
	"aliquam",
	"vestibulum",
	"justo",
	"libero",
	"tempor",
	"facilisis",
	"nulla",
	"placerat",
	"proin",
	"ultrices",
	"purus",
	"imperdiet",
	"venenatis",
	"enim",
	"ante",
	"dui",
	"praesent",
	"quis",
	"pretium",
	"fringilla",
	"rhoncus",
	"lobortis",
	"eget",
	"tempus",
	"scelerisque",
	"ullamcorper",
	"efficitur",
	"donec",
	"suspendisse",
	"vel",
	"dapibus",
	"tortor",
	"ut",
	"sollicitudin",
	"posuere",
	"egestas",
	"vulputate",
	"maximus",
	"volutpat",
	"felis",
	"ac",
	"dictum",
	"semper",
	"ornare",
	"lectus",
	"duis",
	"at",
	"in",
	"id",
	"urna",
	"mauris",
	"sit",
	"adipiscing",
	"massa",
	"sodales",
	"vitae",
	"pulvinar",
	"arcu",
	"commodo",
	"gravida",
	"turpis",
	"fusce",
	"cursus",
	"lorem",
	"porttitor",
	"a",
	"quam",
	"rutrum",
	"tristique",
	"pellentesque",
	"dolor",
	"finibus",
	"fermentum",
}

func getWords(nBytes int) ([]string, int) {
	l := len(words)
	r := rand.New(rand.NewSource(0))
	size := 0
	ws := make([]string, 0, nBytes)
	for size < nBytes {
		i := r.Intn(l)
		w := words[i]
		ws = append(ws, w)
		size += len(w)
	}
	return ws, size
}
