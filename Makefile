.PHONY: build-css watch-css
setup-css:
	npm init -y
	npm i @tailwindcss/forms @tailwindcss/typography autoprefixer postcss tailwindcss
build-css:
	npx tailwindcss -i src/css/app.postcss -o public/styles.css
watch-css:
	npx tailwindcss --watch -i src/css/app.postcss -o public/styles.css

.PHONY: gen-views
gen-views:
	templ generate internal/view

.PHONY: install-tool
install-tool:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/air-verse/air@latest

.phony: tidy
tidy:
	go mod tidy
	go mod vendor