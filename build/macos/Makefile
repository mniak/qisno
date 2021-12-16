APPNAME=Pismo
BINNAME=pismo-gui

APPBUNDLE=../../artifacts/$(APPNAME).app
TEMPDIR := $(shell mktemp -d)

TARGET_INFOPLIST=$(APPBUNDLE)/Contents/Info.plist
TARGET_PKGINFO=$(APPBUNDLE)/Contents/PkgInfo
TARGET_ICONS=$(APPBUNDLE)/Contents/Resources/$(APPNAME).icns
TARGET_BINARY=$(APPBUNDLE)/Contents/MacOS/$(BINNAME)


appbundle: $(TARGET_INFOPLIST) $(TARGET_PKGINFO) $(TARGET_ICONS) $(TARGET_BINARY)
	rm -rf $(TEMPDIR)


$(TARGET_INFOPLIST): Info.plist
	mkdir -p $$(dirname $(TARGET_INFOPLIST))
	cp Info.plist $(TARGET_INFOPLIST)


$(TARGET_PKGINFO): PkgInfo
	mkdir -p $$(dirname $(TARGET_PKGINFO))
	cp PkgInfo $(TARGET_PKGINFO)


$(TARGET_ICONS): Icon.png
	mkdir -p $(TEMPDIR)/$(APPNAME).iconset
	sips -z 16 16     Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_16x16.png
	sips -z 32 32     Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_16x16@2x.png
	sips -z 32 32     Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_32x32.png
	sips -z 64 64     Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_32x32@2x.png
	sips -z 128 128   Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_128x128.png
	sips -z 256 256   Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_128x128@2x.png
	sips -z 256 256   Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_256x256.png
	sips -z 512 512   Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_256x256@2x.png
	sips -z 512 512   Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_512x512.png
	sips -z 512 512   Icon.png --out $(TEMPDIR)/$(APPNAME).iconset/icon_512x512@2x.png
	
	mkdir -p $$(dirname $(TARGET_ICONS))
	iconutil -c icns -o $(TARGET_ICONS) $(TEMPDIR)/$(APPNAME).iconset


$(TARGET_BINARY):
	mkdir -p $$(dirname $(TARGET_BINARY))
	go build -o $(TARGET_BINARY) ../../cmd/menubar
