#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

sqlite3 $DIR/../data/db.sqlite -batch -separator $'\t' 'SELECT Match.id, TAPlayerA.name || "-" || TAPlayerB.name, TBPlayerA.name || "-" || TBPlayerB.name from Match JOIN Team AS TeamA ON TeamA.id = Match.teamA_id Join Team as TeamB ON TeamB.id = Match.teamB_id JOIN Player as TAPlayerA ON TAPlayerA.id = TeamA.playerA_id JOIN Player as TAPlayerB ON TAPlayerB.id = TeamA.playerB_id JOIN Player as TBPlayerA ON TBPlayerA.id = TeamB.playerA_id JOIN Player as TBPlayerB ON TBPlayerB.id = TeamB.playerB_id ORDER BY Match.id DESC LIMIT 20'