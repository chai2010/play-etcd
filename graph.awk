
BEGIN { printf "digraph G {\n" }
{ printf "\t\"%s\" -> \"%s\"\n", $1, $2 }
END { printf "}\n" }
