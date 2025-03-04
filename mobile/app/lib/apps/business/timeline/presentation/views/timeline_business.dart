import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class TimelineBusiness extends StatefulWidget {
  const TimelineBusiness({super.key});

  @override
  State<TimelineBusiness> createState() => _TimelineBusinessState();
}

class _TimelineBusinessState extends State<TimelineBusiness> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: TimelineAppBar(
        onTapAvatar: () => context.pushNamed(AppRoute.bProfile.name),
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
