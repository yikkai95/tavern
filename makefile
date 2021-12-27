run:
	./tavern

build:
	go build -o tavern main.go

push:
	aws s3 sync ./media/ s3://wine.yykk.xyz/media/ --exclude "*.DS_Store*" --exclude "./media/done/*" --exclude "./media/raw/*"

pull:
	aws s3 sync s3://wine.yykk.xyz/media/ ./media/

deploy:
	rsync -vrP --delete-after --exclude-from=.rsignore . root@nnnyykk.com:/var/www/nnnyykk/tavern
