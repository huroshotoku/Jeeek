APP_NAME := jeeekapi
REPO := github.com/tonouchi510/Jeeek
SWAGGER_DIR := ./static/swagger-ui/swagger/


# goa周り
goagen:
	@goa gen ${REPO}/design
	@cp -f ./gen/http/openapi.json ${SWAGGER_DIR}
	@cp -f ./gen/http/openapi.yaml ${SWAGGER_DIR}

example:
	@goa example $(REPO)/design
	@script/fix_goa_example.sh

clean:
	@rm -rf cmd/
	@rm -rf gen/
	@rm *.go

run-local:
	@script/run-local.sh ${ADMIN_PASSWORD}

test:
	@script/test.sh
