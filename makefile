.PHONY:ui
ui:
	@cd ui && pnpm run dev

.PHONY:api
api:
	@cd api && go run .

