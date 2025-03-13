import 'package:celebut/features/timeline/presentation/widgets/pic_asset.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class PhotoWidget extends StatefulWidget {
  const PhotoWidget({
    required this.pageController,
    super.key,
  });

  final PageController pageController;

  @override
  State<PhotoWidget> createState() => _PhotoWidgetState();
}

class _PhotoWidgetState extends State<PhotoWidget> {
  int page = 0;
  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        context.push('/pTimeline/pMediaDetail/photo');
      },
      child: SizedBox(
        height: 350,
        width: double.maxFinite,
        child: PageView(
          controller: widget.pageController,
          onPageChanged: (value) {
            page = value;
            setState(() {});
          },
          children: getPictures(widget.pageController),
        ),
      ),
    );
  }
}
