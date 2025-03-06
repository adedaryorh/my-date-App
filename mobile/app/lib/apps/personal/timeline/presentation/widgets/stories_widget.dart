import 'package:celebut/apps/personal/timeline/presentation/widgets/story_widget.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class StoriesWidget extends StatelessWidget {
  const StoriesWidget({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Celebration Around You',
          style: context.textTheme.titleSmall?.copyWith(
            fontWeight: FontWeight.w600,
          ),
        ),
        const Space(20),
        SizedBox(
          height: 200,
          width: double.maxFinite,
          child: ListView.builder(
            shrinkWrap: true,
            scrollDirection: Axis.horizontal,
            physics: const AlwaysScrollableScrollPhysics(),
            itemCount: assetPaths.length,
            itemBuilder: (context, index) {
              final path = assetPaths[index];
              return StoryWidget(assetPath: path);
            },
          ),
        ),
        const Space(20),
      ],
    );
  }
}

final List<String> assetPaths = [
  AppAssets.defaultAvatarGirl,
  AppAssets.girlStory,
  AppAssets.guyStory,
];
