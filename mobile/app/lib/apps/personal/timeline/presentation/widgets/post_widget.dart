import 'package:celebut/apps/personal/timeline/presentation/widgets/bottom_widget_components.dart';
import 'package:celebut/apps/personal/timeline/presentation/widgets/photo_widget.dart';
import 'package:celebut/apps/personal/timeline/presentation/widgets/top_widget_components.dart';
import 'package:celebut/apps/personal/timeline/presentation/widgets/video_widget.dart';
import 'package:celebut/core/utils/components/app_spacer.dart';
import 'package:flutter/material.dart';

class PostWidget extends StatefulWidget {
  const PostWidget({
    required this.userName,
    required this.time,
    required this.showPhotos,
    required this.showVideo,
    required this.content,
    super.key,
  });

  final String userName;
  final String time;
  final bool showPhotos;
  final bool showVideo;
  final String content;

  @override
  State<PostWidget> createState() => _PostWidgetState();
}

class _PostWidgetState extends State<PostWidget> {
  PageController pageController = PageController();

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.maxFinite,
      margin: const EdgeInsets.only(bottom: 10),
      padding: const EdgeInsets.symmetric(
        vertical: 10,
      ),
      decoration: const BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.all(
          Radius.circular(20),
        ),
      ),
      child: Column(
        children: [
          TopWidgetComponents(
            name: widget.userName,
            time: widget.time,
          ),
          const Space(10),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 10),
            child: Text(widget.content),
          ),
          const Space(10),
          if (widget.showPhotos) ...[
            PhotoWidget(pageController: pageController),
          ],
          if (widget.showVideo) ...[
            const VideoWidget(),
          ],
          const BottomWidgetComponents(),
        ],
      ),
    );
  }
}

List<PostWidget> getPosts() {
  return [
    const PostWidget(
      userName: 'Alice Smith',
      time: '20 April at 4:20 PM',
      showPhotos: false,
      showVideo: false,
      content: '@Samuel_Balogun We’re interested in your ideas '
          'and would be glad to build something bigger out of it. ',
    ),
    const PostWidget(
      userName: 'Jane George',
      time: '20 April at 4:20 PM',
      showPhotos: true,
      showVideo: false,
      content: '@Samuel_Balogun We’re interested in your ideas '
          'and would be glad to build something bigger out of it. '
          'Share your ideas about features/design and '
          'we’ll bring them on to our full case. #Celebut',
    ),
    const PostWidget(
      userName: 'Jane George',
      time: '20 April at 4:20 PM',
      showPhotos: false,
      showVideo: true,
      content: '@Samuel_Balogun We’re interested in your ideas '
          'and would be glad to build something bigger out of it. '
          'Share your ideas about features/design and '
          'we’ll bring them on to our full case. #Celebut',
    ),
    const PostWidget(
      userName: 'Jane Lewis',
      time: '20 April at 4:20 PM',
      showPhotos: false,
      showVideo: false,
      content: 'We’re interested in your ideas #Celebut',
    ),
  ];
}
