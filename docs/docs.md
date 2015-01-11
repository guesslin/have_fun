Find Fun Between Us
===================

Concept
-------
以各個點為圓心
找出點與點間重覆的好玩的指定場所

Flow
----
1.User enter 2 address (at least), and the weight
2.maps API convert address to longitude and latitude. // Address2GPS
example: http://maps.googleapis.com/maps/api/geocode/json?address=<ADDRESS>&sensor=false
3.find the midpoint
http://www.movable-type.co.uk/scripts/latlong.html
4.query nearby place through place API
5.throw result back to maps API. show it!
