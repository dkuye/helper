package helper

import (
	"html/template"

	"github.com/kataras/iris"
)

func Paginate(ctx iris.Context, take int) string {
	p := ctx.URLParam("p") // p is page numbter
	skip := 0
	if p == "" || p == "0" || p == "1" {
		skip = 0
	} else {
		skip = (StringToInt(p) - 1) * take
	}
	return "LIMIT " + IntToString(skip) + ", " + IntToString(take) + " "
}

func PaginationNavigator(perPage int, totalInDB int) template.HTML {
	possiblePages := (totalInDB / perPage) + 1
	total := 0
	if possiblePages > 20 {
		total = 20
	} else {
		total = possiblePages
	}

	buttons := `<div class="btn-group" role="group" aria-label="">`
	for i := 0; i < total; i++ {
		var page = IntToString(i + 1)
		buttons += "<button type=\"button\" class=\"btn btn-default pages\" id=\"" + page + "\">" + page + "</button>"
	}
	buttons += "</div>"

	buttons += `
		<script>
			$('.pages').click(function() {
				const urlParams = new URLSearchParams(window.location.search);
				var status = urlParams.get('s');
				var sq = "";
				if (status !== null){
					sq = "&s=" + status
				}

				var url = window.location.href.split('?')[0];
				var pn = $(this).attr('id');
				window.location.replace(url + "?p=" + pn + sq);
			})
		</script>
	`

	return template.HTML(buttons)
}
