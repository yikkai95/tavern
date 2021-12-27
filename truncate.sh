#!/bin/sh
find ./media/large -type f -name 'IMG*' -execdir convert {} -resize 1280x ./media/large/th_{} \;
