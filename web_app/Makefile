templ:
	@templ generate -proxy="http://localhost:8080" -open-browser=false -cmd="go run main.go"

tw:
	@npx tailwindcss -i styles.css -o public/static/css/tw.css

docker:
	docker build -t web_app:latest .
	docker run -p 8080:8080 web_app