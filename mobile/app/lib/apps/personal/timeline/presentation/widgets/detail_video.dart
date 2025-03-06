import 'package:celebut/apps/personal/timeline/presentation/widgets/bottom_widget_components.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:video_player/video_player.dart';

class DetailVideo extends StatefulWidget {
  const DetailVideo({super.key});

  @override
  State<DetailVideo> createState() => _DetailVideoState();
}

class _DetailVideoState extends State<DetailVideo> {
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
    return Scaffold(
      body: Stack(
        children: [
          Column(
            children: [
              Container(
                margin: const EdgeInsets.symmetric(horizontal: 5),
                height: MediaQuery.of(context).size.height * 0.82,
                width: double.maxFinite,
                decoration: const BoxDecoration(
                  borderRadius: BorderRadius.all(
                    Radius.circular(20),
                  ),
                  color: Colors.black26,
                ),
                child: _videoPlayerController.value.isInitialized
                    ? ClipRRect(
                        borderRadius: BorderRadius.circular(20),
                        child: SizedBox.expand(
                          child: AspectRatio(
                            aspectRatio:
                                _videoPlayerController.value.aspectRatio,
                            child: VideoPlayer(_videoPlayerController),
                          ),
                        ),
                      )
                    : const Center(
                        child: SizedBox(
                          width: 40,
                          height: 40,
                          child: CircularProgressIndicator.adaptive(
                            valueColor:
                                AlwaysStoppedAnimation<Color>(Colors.white),
                          ),
                        ),
                      ),
              ),
            ],
          ),
          Positioned(
            top: 30,
            left: 0,
            right: 0,
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                IconButton(
                  onPressed: () {
                    context.pop();
                  },
                  icon: const Icon(
                    Icons.arrow_back_ios,
                    size: 15,
                    color: Colors.black,
                  ),
                ),
                IconButton(
                  onPressed: () {},
                  icon: const Icon(
                    Icons.more_vert,
                    size: 19,
                  ),
                ),
              ],
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
                const Space(30),
                VideoProgressIndicator(
                  padding: const EdgeInsets.only(
                    right: 25,
                    left: 25,
                    top: 10,
                  ),
                  colors: const VideoProgressColors(
                    playedColor: Colors.black,
                    bufferedColor: Colors.white24,
                    backgroundColor: Colors.black87,
                  ),
                  _videoPlayerController,
                  allowScrubbing: true,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Row(
                      children: [
                        InkWell(
                          onTap: () {
                            setState(
                              () {
                                _videoPlayerController.value.isPlaying
                                    ? _videoPlayerController.pause()
                                    : _videoPlayerController.play();
                              },
                            );
                          },
                          child: Container(
                            padding: const EdgeInsets.only(left: 10),
                            height: 30,
                            width: 50,
                            child: _videoPlayerController.value.isPlaying
                                ? const Icon(
                                    Icons.pause,
                                    size: 20,
                                    color: Colors.black,
                                  )
                                : const Icon(
                                    Icons.play_arrow,
                                    size: 20,
                                    color: Colors.black,
                                  ),
                          ),
                        ),
                        Text(
                          '158 views',
                          style: context.textTheme.bodySmall
                              ?.copyWith(color: Colors.black),
                        ),
                      ],
                    ),
                    Padding(
                      padding: const EdgeInsets.only(right: 20),
                      child: Text(
                        '${formatDuration(_currentPosition)} / ${formatDuration(_totalDuration)}',
                        style: context.textTheme.bodySmall
                            ?.copyWith(color: Colors.black),
                      ),
                    ),
                  ],
                ),
                const Space(20),
                const Padding(
                  padding: EdgeInsets.symmetric(horizontal: 20),
                  child: BottomWidgetComponents(),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
