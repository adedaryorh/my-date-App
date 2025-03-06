import 'package:celebut/apps/personal/timeline/presentation/widgets/archived_celebration.dart';
import 'package:celebut/apps/personal/timeline/presentation/widgets/post_widget.dart';
import 'package:celebut/apps/personal/timeline/presentation/widgets/stories_widget.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:go_router/go_router.dart';

class TimelinePersonal extends StatefulWidget {
  const TimelinePersonal({super.key});

  @override
  State<TimelinePersonal> createState() => _TimelinePersonalState();
}

class _TimelinePersonalState extends State<TimelinePersonal> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: TimelineAppBar(
        onTapAvatar: () => context.pushNamed(AppRoute.pProfile.name),
        onTapStore: () {},
        onTapSearch: () {},
      ),
      backgroundColor: context.colorScheme.primary,
      body: Stack(
        children: [
          SafeArea(
            child: Expanded(
              child: ListView(
                padding: const EdgeInsets.symmetric(horizontal: 20),
                shrinkWrap: true,
                physics: const AlwaysScrollableScrollPhysics(),
                children: [
                  ...getPosts(),
                  const StoriesWidget(),
                  const ArchivedCelebration(),
                  ...getPosts(),
                ],
              ),
            ),
          ),
          Align(
            alignment: Alignment.centerRight,
            child: Padding(
              padding: const EdgeInsets.only(right: 16),
              child: FloatingActionButton(
                backgroundColor: Colors.white,
                onPressed: () {
                  context.pushNamed(AppRoute.pNewPost.name);
                },
                child: SvgPicture.asset(
                  AppAssets.speak,
                  width: 26,
                  height: 26,
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
