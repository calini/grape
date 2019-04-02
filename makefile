build:
	@go build -o out/grape

run-test-dict:
	@out/grape \
		-url https://github.com/%s \
		-dict junk/dicts/random.txt \
		-query ".p-name .avatar" \
		-output out/output_github.csv

run-test-iter:
	@out/grape \
		-url https://github.com/%d \
		-from 1 \
		-to 10 \
		-query ".p-name .avatar" \
		-output out/output_iter.csv

run-test-json-TODO:
	@out/grape \
		-json
		-url https://github.com/%d \
		-from 1 \
		-to 10 \
		-query ".p-name .avatar" \
		-output out/output_iter.csv


