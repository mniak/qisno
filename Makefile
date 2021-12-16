main:
	# Please specify one of the following rules:
	#   app-mac
	#   clean

app-mac:
	cd build/macos && make

clean:
	rm -r artifacts