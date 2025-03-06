import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:go_router/go_router.dart';
import 'package:video_player/video_player.dart';

class VideoWidget extends StatefulWidget {
  const VideoWidget({
    super.key,
  });

  @override
  State<VideoWidget> createState() => _VideoWidgetState();
}

class _VideoWidgetState extends State<VideoWidget> {
  late VideoPlayerController _videoPlayerController;
  Duration _currentPosition = Duration.zero;
  Duration _totalDuration = Duration.zero;

  Future<void> initializePlayer() async {
    _videoPlayerController = VideoPlayerController.networkUrl(
      Uri.parse(
        'https://flutter.github.io/assets-for-api-docs/assets/videos/butterfly.mp4',
      ),
    );
    await _videoPlayerController.initialize().then((value) {
      _totalDuration = _videoPlayerController.value.duration;
      _videoPlayerController
        ..setLooping(true)
        ..setVolume(1)
        ..pause();
      setState(() {});
    });

    _videoPlayerController.addListener(() {
      setState(() {
        _currentPosition = _videoPlayerController.value.position;
      });
    });
  }

  String formatDuration(Duration duration) {
    String twoDigits(int n) => n.toString().padLeft(2, '0');
    final minutes = twoDigits(duration.inMinutes.remainder(60));
    final seconds = twoDigits(duration.inSeconds.remainder(60));
    return '$minutes:$seconds';
  }

  @override
  void initState() {
    initializePlayer();
    super.initState();
  }

  @override
  void dispose() {
    _videoPlayerController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {},
      child: Container(
        margin: const EdgeInsets.symmetric(
          horizontal: 10,
          vertical: 10,
        ),
        height: 270,
        width: double.maxFinite,
        decoration: const BoxDecoration(
          borderRadius: BorderRadius.all(
            Radius.circular(20),
          ),
          color: Colors.black26,
        ),
        child: _videoPlayerController.value.isInitialized
            ? Stack(
                children: [
                  Align(
                    child: ClipRRect(
                      borderRadius: BorderRadius.circular(20),
                      child: SizedBox.expand(
                        child: AspectRatio(
                          aspectRatio: _videoPlayerController.value.aspectRatio,
                          child: VideoPlayer(_videoPlayerController),
                        ),
                      ),
                    ),
                  ),
                  Align(
                    child: Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        IconButton(
                          onPressed: () {},
                          icon: Image.asset(
                            AppAssets.rewindVideo,
                            height: 30,
                            width: 30,
                          ),
                        ),
                        IconButton(
                          onPressed: () {
                            setState(
                              () {
                                _videoPlayerController.value.isPlaying
                                    ? _videoPlayerController.pause()
                                    : _videoPlayerController.play();
                              },
                            );
                          },
                          icon: _videoPlayerController.value.isPlaying
                              ? SvgPicture.asset(
                                  AppAssets.pauseVideo,
                                )
                              : SvgPicture.asset(
                                  AppAssets.playVideo,
                                  color: Colors.white,
                                  width: 30,
                                  height: 30,
                                ),
                        ),
                        IconButton(
                          onPressed: () {},
                          icon: Image.asset(
                            AppAssets.forwardVideo,
                            height: 30,
                            width: 30,
                          ),
                        ),
                      ],
                    ),
                  ),
                  Positioned(
                    top: 10,
                    bottom: 200,
                    right: 20,
                    child: IconButton(
                      onPressed: () {
                        context.push('/pTimeline/pMediaDetail/video');
                      },
                      icon: Container(
                        height: 24,
                        width: 24,
                        decoration: const BoxDecoration(
                          shape: BoxShape.circle,
                          color: Colors.black26,
                        ),
                        child: SvgPicture.asset(AppAssets.expandVideo),
                      ),
                    ),
                  ),
                  Positioned(
                    right: 0,
                    left: 0,
                    bottom: 30,
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      mainAxisSize: MainAxisSize.min,
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Padding(
                          padding: const EdgeInsets.symmetric(
                            horizontal: 25,
                          ),
                          child: Text(
                            'World War 2',
                            style: context.textTheme.bodyLarge
                                ?.copyWith(color: Colors.white),
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.symmetric(
                            horizontal: 25,
                          ),
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Text(
                                'Documentary',
                                style: context.textTheme.bodySmall
                                    ?.copyWith(color: Colors.white),
                              ),
                              Text(
                                '${formatDuration(_currentPosition)} / ${formatDuration(_totalDuration)}',
                                style: context.textTheme.bodySmall
                                    ?.copyWith(color: Colors.white),
                              ),
                            ],
                          ),
                        ),
                        VideoProgressIndicator(
                          padding: const EdgeInsets.only(
                            right: 25,
                            left: 25,
                            top: 10,
                          ),
                          colors: const VideoProgressColors(
                            playedColor: Colors.white,
                            bufferedColor: Colors.black26,
                            backgroundColor: Colors.white70,
                          ),
                          _videoPlayerController,
                          allowScrubbing: true,
                        ),
                      ],
                    ),
                  ),
                ],
              )
            : const Center(
                child: SizedBox(
                  width: 40,
                  height: 40,
                  child: CircularProgressIndicator.adaptive(
                    valueColor: AlwaysStoppedAnimation<Color>(Colors.white),
                  ),
                ),
              ),
      ),
    );
  }
}
