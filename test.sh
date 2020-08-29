#!/bin/bash

echo "starting berks webserver"

osascript <<EOF
    tell application "iTerm"
        tell current window
            create tab with default profile
            -- split horizontally with default profile
        end tell
        tell current tab of current window
            set _new_session to last item of sessions
        end tell
        tell _new_session
            select
            write text "cd go/src/berksArrUs"
            write text "go run cmd/main.go"
        end tell
    end tell
EOF

echo "waiting a bit for webserver to be ready..."
# TODO: use the root route as a readiness check
sleep 5s


echo "done sleeping, loading test data"
echo
echo "should get book_0"

curl --location --request POST 'http://localhost:8080/books' \
--data-raw '{
	"title":"Book One",
	"author":"Cultured Author",
	"description":"A compelling description",
	"isbn":"a_bunch_of_characters"
}'

echo
echo "should retrieve Book One"

# first book should be stored with key book_0, now let's grab it

curl --location --request GET 'http://localhost:8080/books/book_0'

# now let's update book_0

echo
echo "should update Book One"

curl --location --request PUT 'http://localhost:8080/books/book_0' \
--data-raw '{
	"title":"Book One",
	"author":"Cultured Author",
	"description":"A better description",
	"isbn":"a_bunch_of_characters"
}'

echo "description should be different from before"

curl --location --request GET 'http://localhost:8080/books/book_0'

# now let's delete book_0

echo
echo "now let's delete it"

curl --location --request DELETE 'http://localhost:8080/books/book_0'

echo "Should get 404"

curl --location --request GET 'http://localhost:8080/books/book_0'

echo
echo "Let's put some books back and test bulk retrieval"

curl --location --request POST 'http://localhost:8080/books' \
--data-raw '{
	"title":"Book One",
	"author":"Cultured Author",
	"description":"A compelling description",
	"isbn":"a_bunch_of_characters"
}'

echo

curl --location --request POST 'http://localhost:8080/books' \
--data-raw '{
	"title":"Book Two",
	"author":"Trashy Author",
	"description":"A trashy description",
	"isbn":"a_bundle_of_characters"
}'

echo

curl --location --request POST 'http://localhost:8080/books' \
--data-raw '{
	"title":"Book Three",
	"author":"Pulpy Author",
	"description":"A pulpy description",
	"isbn":"a_bag_of_characters"
}'

echo
echo

curl --location --request GET 'http://localhost:8080/bulk_books'