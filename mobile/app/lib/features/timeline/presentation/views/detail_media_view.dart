import 'package:celebut/features/timeline/presentation/widgets/detail_photo.dart';
import 'package:celebut/features/timeline/presentation/widgets/detail_video.dart';
import 'package:flutter/material.dart';

class DetailMediaView extends StatefulWidget {
  const DetailMediaView({super.key, this.mediaType});

  final String? mediaType;

  @override
  State<DetailMediaView> createState() => _DetailMediaViewState();
}

class _DetailMediaViewState extends State<DetailMediaView> {
  @override
  Widget build(BuildContext context) {
    if (widget.mediaType == 'video') {
      return const DetailVideo();
    } else {
      return const DetailPhoto();
    }
  }
}
