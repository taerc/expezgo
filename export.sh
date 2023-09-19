#!/bin/sh
VIDEO_FILE=__FILE__
OUT=goal
LIMIT=5
mkdir -p ${OUT}
ffmpeg -i ${VIDEO_FILE} -ss  00:02:38 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_0.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:03:19 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_1.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:04:10 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_2.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:04:29 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_3.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:05:36 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_4.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:05:45 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_5.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:06:18 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_6.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:06:28 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_7.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:07:19 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_8.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:07:37 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_9.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:07:56 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_10.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:16:05 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_11.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:16:40 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_12.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:17:03 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_13.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:18:04 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_14.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:18:27 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_15.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:19:13 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_16.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:25:20 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_17.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:25:47 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_18.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:26:07 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_19.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:26:27 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_20.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:26:50 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_21.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:27:03 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_22.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:27:52 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_23.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:28:11 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_24.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:38:28 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_25.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:39:06 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_26.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:46:29 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_27.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:50:07 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_28.mp4
ffmpeg -i ${VIDEO_FILE} -ss  00:51:25 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_29.mp4
ffmpeg -i ${VIDEO_FILE} -ss  01:47:25 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_30.mp4
ffmpeg -i ${VIDEO_FILE} -ss  01:48:16 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_31.mp4
ffmpeg -i ${VIDEO_FILE} -ss  01:57:37 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_32.mp4
ffmpeg -i ${VIDEO_FILE} -ss  02:05:35 -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_33.mp4
