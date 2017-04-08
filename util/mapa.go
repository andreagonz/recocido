package recocido

import (
	"strconv"
	imp "github.com/andreagonz/recocido/implementacion"
)

func Mapa(ruta string, ciudades *[]imp.Ciudad) string {
	r := CadenaARuta(ruta)
	s := "<!DOCTYPE html>\n<html>\n<head>\n<meta name='viewport'content='initial-scale=1.0, user-scalable=no'>\n<meta charset='utf-8'>\n<title>Mapa</title>\n<style>\n#map {\nheight: 100%;\n}\nhtml, body {\nheight: 100%;\nmargin: 0;\npadding: 0;\n}\n</style>\n</head>\n<body>\n<div id='map'></div>\n<script>\nfunction initMap() {\n var map = new google.maps.Map(document.getElementById('map'), {\nzoom: 3,\ncenter: {lat: 0, lng: -180},\nmapTypeId: 'terrain'\n});\nvar flightPlanCoordinates = [\n"
	for i := 0; i < len(r); i++ {
		s += "{lat: " + strconv.FormatFloat((*ciudades)[r[i]].Latitud, 'f', -1, 64) + ", lng: " + strconv.FormatFloat((*ciudades)[r[i]].Longitud, 'f', -1, 64) + "}"
		if i < len(r) - 1 {
			s += ",\n"
		}
	}
	s += "      ];\n      var flightPath = new google.maps.Polyline({\n      path: flightPlanCoordinates,\n      geodesic: true,\n      strokeColor: '#FF0000',\n      strokeOpacity: 1.0,\n      strokeWeight: 2\n      });      \n      flightPath.setMap(map);\n      }\n    </script>\n    <script async defer\n            src='https://maps.googleapis.com/maps/api/js?key=AIzaSyDdBAKYa4kQqUStHeV39ngfUVZwRAl84bk&callback=initMap'>\n    </script>\n  </body>\n</html>"
	return s
}
