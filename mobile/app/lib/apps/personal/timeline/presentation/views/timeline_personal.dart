import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
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
      body: SafeArea(
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 20),
          child: Column(
            children: [],
          ),
        ),
      ),
    );
  }
}
