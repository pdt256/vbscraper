#!/bin/bash

curl -s "http://bvbinfo.com/season.asp" -o "all-seasons.html"
curl -s "http://bvbinfo.com/Season.asp?AssocID=1&Year=2017" -o "2017-avp-tournaments.html"
curl -s "http://bvbinfo.com/Tournament.asp?ID=3332&Process=Matches" -o "2017-avp-manhattan-beach-mens-matches.html"