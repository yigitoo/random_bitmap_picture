ifeq ($(OS), Windows_NT)
rm_cmd := rd /S
mv_cmd := move
else
rm_cmd := rm -rf
mv_cmd := mv
endif
build_dir := build
name := random_profile_picture
build:
	mkdir $(build_dir)
	go build
	$(mv_cmd) $(name) $(build_dir)/
run:
	$(build_dir)/$(name)
only-run:
	go run .
clean:
	$(rm_cmd) $(build_dir)
